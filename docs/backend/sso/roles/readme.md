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
    string name = 2;
    string description = 3;
    int32 level = 4;
    bool is_custom = 5;
    optional string created_by = 6;
}
```
**Пример запроса**
```json
{
  "client_id": "8268ec76-d6c2-48b5-a0e4-a9c2538b8f48",
  "name": "super admin",
  "description": "Super admin role",
  "level": 0,
  "is_custom": true
}
```
**Пример ответа**
```json
{
	"permission_ids": [],
	"id": "6d08281f-7ee6-4019-a9f3-f6004a5d6acd",
	"client_id": "8268ec76-d6c2-48b5-a0e4-a9c2538b8f48",
	"name": "super admin",
	"description": "Super admin role",
	"level": 0,
	"is_active": true,
	"is_custom": true,
	"created_at": {
		"seconds": "1745473148",
		"nanos": 32631200
	},
	"updated_at": {
		"seconds": "1745473148",
		"nanos": 32631200
	}
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
}
```
**Пример запроса**
```json
{
	"id": "6d08281f-7ee6-4019-a9f3-f6004a5d6acd",
	"client_id": "8268ec76-d6c2-48b5-a0e4-a9c2538b8f48"
}
```
**Пример ответа**
```json
{
  "permission_ids": [],
  "id": "6d08281f-7ee6-4019-a9f3-f6004a5d6acd",
  "client_id": "8268ec76-d6c2-48b5-a0e4-a9c2538b8f48",
  "name": "super admin",
  "description": "Super admin role",
  "level": 0,
  "is_active": true,
  "is_custom": true,
  "created_at": {
    "seconds": "1745473148",
    "nanos": 32631000
  },
  "updated_at": {
    "seconds": "1745473148",
    "nanos": 32631000
  }
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
    optional string name = 3;
    optional string description = 4;
    optional int32 level = 5;
    optional bool is_active = 6;
}
```
**Пример запроса**
```json
{
  "id": "6d08281f-7ee6-4019-a9f3-f6004a5d6acd",
  "client_id": "8268ec76-d6c2-48b5-a0e4-a9c2538b8f48",
  "name": "admin",
  "description": "Admin role",
  "level": 1,
  "is_active": true,
  "is_custom": true
}
```
**Пример ответа**
```json
{
  "permission_ids": [],
  "id": "6d08281f-7ee6-4019-a9f3-f6004a5d6acd",
  "client_id": "8268ec76-d6c2-48b5-a0e4-a9c2538b8f48",
  "name": "admin",
  "description": "Admin role",
  "level": 1,
  "is_active": true,
  "is_custom": true,
  "created_at": {
    "seconds": "1745473148",
    "nanos": 32631000
  },
  "updated_at": {
    "seconds": "1745476512",
    "nanos": 246523000
  }
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
    bool permanent = 3;
}
```
**Особенности**
- Мягкое удаление по умолчанию (is_active=false + deleted_at)
- Запрет удаления системных ролей
- Каскадное удаление зависимостей при permanent=true
  **Ответ**
```json
{
    "success": true
}
```
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
      "permission_ids": [],
      "id": "6d08281f-7ee6-4019-a9f3-f6004a5d6acd",
      "client_id": "8268ec76-d6c2-48b5-a0e4-a9c2538b8f48",
      "name": "admin",
      "description": "Admin role",
      "level": 1,
      "is_active": true,
      "is_custom": true,
      "created_at": {
        "seconds": "1745473148",
        "nanos": 32631000
      },
      "updated_at": {
        "seconds": "1745476512",
        "nanos": 246523000
      }
    }
  ],
  "total_count": 1,
  "current_page": 1
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