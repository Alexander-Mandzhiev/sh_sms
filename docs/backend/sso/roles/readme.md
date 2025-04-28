# Role Management Service

gRPC-сервис для управления ролями с поддержкой мультитенантности, иерархии доступа и гибкой системой управления жизненным циклом.

👑 **Особенности**
- Полный CRUD для ролей с мягким/жестким удалением
- Иерархическая система уровней доступа (level)
- Поддержка системных и кастомных ролей
- Межклиентская изоляция через client_id
- Пагинация и фильтрация по имени/уровню/активности
- Контроль целостности при удалении
- Оптимизированные индексы для быстрого поиска

🗃 **Структура данных**

### Таблица `roles`
Поле | Тип | Ограничения | Описание
-----|-----|-------------|----------
id | UUID | PRIMARY KEY, DEFAULT gen_random_uuid() | Уникальный идентификатор
client_id | UUID | NOT NULL | Идентификатор клиента (тенанта)
name | VARCHAR(150) | NOT NULL | Уникальное имя роли
description | TEXT |  | Описание роли
level | INT | DEFAULT 0, CHECK (>= 0) | Уровень доступа (0 - минимальный)
is_custom | BOOLEAN | DEFAULT FALSE | Флаг кастомной роли
is_active | BOOLEAN | DEFAULT TRUE | Флаг активности
created_by | UUID |  | Создатель роли (опционально)
created_at | TIMESTAMPTZ | DEFAULT CURRENT_TIMESTAMP | Дата создания
updated_at | TIMESTAMPTZ |  | Дата последнего обновления
deleted_at | TIMESTAMPTZ |  | Мягкое удаление

### Индексы
- `idx_roles_client_active` (client_id) WHERE deleted_at IS NULL
- `idx_roles_name_client` UNIQUE (client_id, name) WHERE deleted_at IS NULL
- `idx_roles_level` (level)
- `idx_roles_created` (created_at DESC)

## 📡 API Методы

### 1. Создание роли (Create)
```protobuf
rpc Create(CreateRequest) returns (Role);

message CreateRequest {
  string client_id = 1;
  int32 app_id = 2;
  string name = 3;
  string description = 4;
  int32 level = 5;
  optional bool is_custom =6;
  optional string created_by = 7;
}
```
**Пример запроса**
```json
{
  "client_id": "8268ec76-d6c2-48b5-a0e4-a9c2538b8f48",
  "app_id": 1,
  "name": "super admin",
  "description": "Super admin role",
  "level": 0,
  "is_custom": true
}
```
**Пример ответа**
```json
{
  "id": "bc870624-ea19-4cc9-951e-9cb19b740812",
  "client_id": "8268ec76-d6c2-48b5-a0e4-a9c2538b8f48",
  "app_id": 1,
  "name": "super admin",
  "level": 0,
  "is_active": true,
  "is_custom": true,
  "created_at": {
    "seconds": "1745829837",
    "nanos": 490219000
  },
  "updated_at": {
    "seconds": "1745830788",
    "nanos": 765930000
  },
  "description": "Super admin",
  "_description": "description"
}
```
**Особенности**
- Автоматическая генерация UUID
- Проверка уникальности имени в рамках client_id
- Валидация уровня (level ≥ 0)
- Автоматическое проставление временных меток

**Ошибки**

```json
{
    "code": 6,
    "message": "Role name already exists"
}
```
### 2. Получение роли (Get)
```protobuf
rpc Get(GetRequest) returns (Role);

message GetRequest {
    string client_id = 1;
    string id = 2;
    int32 app_id = 3;
}
```
**Пример запроса**
```json
{
  "id": "bc870624-ea19-4cc9-951e-9cb19b740812", 
  "client_id": "8268ec76-d6c2-48b5-a0e4-a9c2538b8f48", 
  "app_id": 1
}

```
**Пример ответа**
```json
{
  "id": "bc870624-ea19-4cc9-951e-9cb19b740812",
  "client_id": "8268ec76-d6c2-48b5-a0e4-a9c2538b8f48",
  "app_id": 1,
  "name": "super admin",
  "level": 0,
  "is_active": true,
  "is_custom": true,
  "created_at": {
    "seconds": "1745829837",
    "nanos": 490219000
  },
  "updated_at": {
    "seconds": "1745830788",
    "nanos": 765930000
  },
  "description": "Super admin",
  "_description": "description"
}
```
**Особенности**
- Проверка принадлежности роли клиенту
- Не возвращает мягко удаленные роли

### 3. Обновление роли (Update)
```protobuf
rpc Update(UpdateRequest) returns (Role);

message UpdateRequest {
  string id = 1;
  string client_id = 2;
  int32 app_id = 3;
  optional string name = 4;
  optional string description = 5;
  optional int32 level = 6;
}
```
**Пример запроса**
```json
{
  "id": "bc870624-ea19-4cc9-951e-9cb19b740812",
  "client_id": "8268ec76-d6c2-48b5-a0e4-a9c2538b8f48",
  "app_id": 1,
  "name": "super admin - role",
  "level": 0
}
```
**Пример ответа**
```json
{
  "id": "bc870624-ea19-4cc9-951e-9cb19b740812",
  "client_id": "8268ec76-d6c2-48b5-a0e4-a9c2538b8f48",
  "app_id": 0,
  "name": "super admin - role",
  "level": 0,
  "is_active": true,
  "is_custom": true,
  "created_at": {
    "seconds": "1745829837",
    "nanos": 490219000
  },
  "updated_at": {
    "seconds": "1745835216",
    "nanos": 903339000
  },
  "description": "Super admin",
  "_description": "description"
}
```

**Особенности**
- Частичное обновление полей
- Запрет изменения системных ролей (is_custom=false)
- Валидация уровня доступа
- Автоматическое обновление updated_at
  **Ошибки**
```json
{
    "code": 7,
    "message": "System roles cannot be modified"
}
```
### 4. Удаление роли (Delete)
```protobuf
rpc Delete(DeleteRequest) returns (DeleteResponse);

message DeleteRequest {
    string id = 1;
    string client_id = 2;
    int32  app_id = 3,
    bool permanent = 4;
}
```
```json
{
  "id": "bc870624-ea19-4cc9-951e-9cb19b740812",
  "client_id": "8268ec76-d6c2-48b5-a0e4-a9c2538b8f48",
  "app_id": 1,
  "permanent": false
}
```
**Ответ**
```json
{
    "success": true
}
```
**Особенности**
- Мягкое удаление по умолчанию (is_active=false + deleted_at)
- Запрет удаления системных ролей
- Каскадное удаление зависимостей при permanent=true

### 5. Список ролей (List)
```protobuf
rpc List(ListRequest) returns (ListResponse);

message ListRequest {
    string client_id = 1;
    optional string name_filter = 2;
    optional int32 level_filter = 3;
    optional bool active_only = 4;
    int32 page = 5;
    int32 count = 6;
}
```
**Пример запроса**
```json
{
	"client_id": "8268ec76-d6c2-48b5-a0e4-a9c2538b8f48",
    "app_id": 1,
	"page":1,
    "count":20,
	"active_only": true,
	"name_filter": "admin"
}
```
**Пример ответа**

```json
{
  "roles": [
    {
      "id": "bc870624-ea19-4cc9-951e-9cb19b740812",
      "client_id": "8268ec76-d6c2-48b5-a0e4-a9c2538b8f48",
      "app_id": 1,
      "name": "super admin",
      "level": 0,
      "is_active": true,
      "is_custom": true,
      "created_at": {
        "seconds": "1745829837",
        "nanos": 490219000
      },
      "updated_at": {
        "seconds": "1745830788",
        "nanos": 765930000
      },
      "description": "Super admin",
      "_description": "description"
    }
  ],
  "total_count": 1,
  "current_page": 1
}
```

### 6. Восстановление роли (Restore)
```protobuf
rpc Restore(RestoreRequest) returns (Role);

message RestoreRequest {
    string id = 1;
    string client_id = 2;
    int32 app_id = 3;
}
```
**Пример запроса**

```json
{
  "id": "bc870624-ea19-4cc9-951e-9cb19b740812",
  "client_id": "8268ec76-d6c2-48b5-a0e4-a9c2538b8f48",
  "app_id": 1
}
```
**Пример ответа**
```json
{
  "id": "bc870624-ea19-4cc9-951e-9cb19b740812",
  "client_id": "8268ec76-d6c2-48b5-a0e4-a9c2538b8f48",
  "app_id": 1,
  "name": "super admin",
  "level": 0,
  "is_active": true,
  "is_custom": true,
  "created_at": "2024-05-10T15:23:57Z",
  "updated_at": "2024-05-10T15:39:48Z",
  "description": "Super admin role"
}
```
**Особенности**
- Восстанавливает мягко удаленные роли
- Проверяет уникальность имени после восстановления
- Обновляет поля deleted_at и is_active

**Ошибки**
```json
{
    "code": 5,
    "message": "Role name 'super admin' already exists"
}
```

**Особенности**
- Поиск по частичному совпадению имени
- Фильтрация по уровню и активности
- Сортировка по уровню и дате создания
- Максимальный размер страницы 100 записей

### 🛡️ Политики безопасности
- Валидация уровня доступа при изменении ролей
- Шифрование чувствительных метаданных
- Ограничение на изменение вышестоящих ролей
- JWT-авторизация с проверкой уровня доступа
- Логирование всех операций изменения

### 🔐 Особенности безопасности
- Хранение журнала изменений ролей
- Проверка прав доступа перед модификацией
- Ограничение частоты запросов (rate limiting)
- Межклиентская изоляция на уровне БД

### 📦 Зависимости
- PostgreSQL 14+ с расширением pgcrypto
- Go 1.21+ с поддержкой gRPC
- Библиотеки: pgx/v5, google.golang.org/protobuf