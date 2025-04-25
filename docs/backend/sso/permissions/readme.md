# Permission Management Service

gRPC-сервис для управления правами доступа с мультитенантной поддержкой и гибкой системой жизненного цикла.

🔑 **Особенности**
- Полный CRUD для прав с мягким/жестким удалением
- Группировка прав по категориям
- Привязка к конкретному приложению через app_id
- Пагинация и фильтрация по коду/категории/активности
- Контроль уникальности в рамках приложения
- Оптимизированные индексы для быстрого поиска

🗃 **Структура данных**

### Таблица `permissions`
Поле | Тип | Ограничения | Описание
-----|-----|-------------|----------
id | UUID | PRIMARY KEY, DEFAULT gen_random_uuid() | Уникальный идентификатор
code | VARCHAR(100) | NOT NULL | Уникальный код права
description | TEXT | NOT NULL | Описание назначения права
category | VARCHAR(50) |  | Категория для группировки
app_id | INT | NOT NULL | ID приложения-владельца
is_active | BOOLEAN | DEFAULT TRUE | Флаг активности
created_at | TIMESTAMPTZ | DEFAULT CURRENT_TIMESTAMP | Дата создания
updated_at | TIMESTAMPTZ | DEFAULT CURRENT_TIMESTAMP | Дата обновления
deleted_at | TIMESTAMPTZ |  | Мягкое удаление

### Индексы
- `idx_permissions_code_app` UNIQUE (code, app_id) WHERE deleted_at IS NULL
- `idx_permissions_category` (category)
- `idx_permissions_created` (created_at DESC)

## 📡 API Методы

### 1. Создание права (Create)
```protobuf
rpc Create(CreateRequest) returns (Permission);

message CreateRequest {
    string code = 1;
    string description = 2;
    string category = 3;
    int32 app_id = 4;
}
```
**Пример запроса**
```json
{
  "code": "user.create",
  "description": "Create new users",
  "category": "users",
  "app_id": 1
}
```
**Пример ответа**
```json
{
	"id": "37a2339d-4778-42ab-aadd-be47de79a1e9",
	"code": "user.create",
	"description": "Create new users",
	"category": "users",
	"app_id": 1,
	"is_active": true,
	"created_at": {
		"seconds": "1745562453",
		"nanos": 727466000
	},
	"updated_at": {
		"seconds": "1745562453",
		"nanos": 727466000
	}
}
```
**Особенности**
- Автоматическая генерация UUID
- Проверка уникальности code в рамках app_id
- Валидация длины code (≤100) и category (≤50)
- Значение is_active=true по умолчанию

**Ошибки**

```json
{
    "code": 6,
    "message": "Permission code already exists for this app"
}
```
### 2. Получение права (Get)
```protobuf
rpc Get(GetRequest) returns (Permission);

message GetRequest {
    string id = 1;
    int32 app_id = 2;
}
```
**Пример запроса**

```json
{
  "id": "37a2339d-4778-42ab-aadd-be47de79a1e9",
  "app_id": 1
}
```
**Пример ответа**

```json
{
  "id": "37a2339d-4778-42ab-aadd-be47de79a1e9",
  "code": "user.create",
  "description": "Create new users",
  "category": "users",
  "app_id": 1,
  "is_active": true,
  "created_at": {
    "seconds": "1745562453",
    "nanos": 727466000
  },
  "updated_at": {
    "seconds": "1745562453",
    "nanos": 727466000
  }
}
```
**Особенности**

- Проверка принадлежности права приложению

- Не возвращает мягко удаленные права

### 3. Обновление права (Update)
```protobuf
rpc Update(UpdateRequest) returns (Permission);

message UpdateRequest {
    string id = 1;
    int32 app_id = 2;
    optional string code = 3;
    optional string description = 4;
    optional string category = 5;
    optional bool is_active = 6;
}
```
**Пример запроса**
```json
{
  "id": "37a2339d-4778-42ab-aadd-be47de79a1e9",
  "code": "user.create",
  "description": "Create new users - permissions",
  "category": "users",
  "app_id": 1,
  "is_active": true
}
```
**Пример ответа**
```json
{
  "id": "37a2339d-4778-42ab-aadd-be47de79a1e9",
  "code": "user.create",
  "description": "Create new users - permissions",
  "category": "users",
  "app_id": 1,
  "is_active": true,
  "created_at": {
    "seconds": "1745562453",
    "nanos": 727466000
  },
  "updated_at": {
    "seconds": "1745562651",
    "nanos": 635528000
  }
}
```
**Особенности**
- Частичное обновление полей
- Проверка уникальности code при изменении
- Автоматическое обновление updated_at

### 4. Удаление права (Delete)
```protobuf
rpc Delete(DeleteRequest) returns (DeleteResponse);

message DeleteRequest {
    string id = 1;
    int32 app_id = 2;
    bool permanent = 3;
}

```
**Пример запроса**
```json
{
	"id": "37a2339d-4778-42ab-aadd-be47de79a1e9",
  "app_id": 1
}
```
**Пример ответа**
```json
{
	"success": true
}
```
Особенности
- Мягкое удаление по умолчанию (deleted_at)
- Полное удаление истории при permanent=true
- Проверка использования права в ролях
**Ответ**
```json
{
  "success": true
}
```
5. Список прав (List)
```protobuf
rpc List(ListRequest) returns (ListResponse);

message ListRequest {
    int32 app_id = 1;
    optional string code_filter = 2;
    optional string category = 3;
    optional bool active_only = 4;
    int32 page = 5;
    int32 count = 6;
}
```
**Пример запроса**

```json
{
  "app_id": 1,
  "page": 1,
  "count": 20,
  "active_only": true
}
```
**Пример ответа**

```json
{
  "permissions": [
    {
      "id": "37a2339d-4778-42ab-aadd-be47de79a1e9",
      "code": "user.create",
      "description": "Create new users - permissions",
      "category": "users",
      "app_id": 1,
      "is_active": true,
      "created_at": {
        "seconds": "1745562453",
        "nanos": 727466000
      },
      "updated_at": {
        "seconds": "1745562651",
        "nanos": 635528000
      }
    }
  ],
  "total_count": 1,
  "current_page": 1
}
```
**Особенности**

- Поиск по маске code (LIKE)
- Фильтрация по категории и активности
- Сортировка по дате создания
- Максимальный размер страницы 1000 записей

### 6. Восстановление права (Restore)
```protobuf
rpc Restore(RestoreRequest) returns (Permission);

message RestoreRequest {
string id = 1;
int32 app_id = 2;
}
```
**Пример запроса**
```json
{
  "id": "37a2339d-4778-42ab-aadd-be47de79a1e9",
  "app_id": 1
}
```
**Пример ответа**
```json
{
  "id": "37a2339d-4778-42ab-aadd-be47de79a1e9",
  "code": "user.create",
  "description": "Create new users - permissions",
  "category": "users",
  "app_id": 1,
  "is_active": true,
  "created_at": {
    "seconds": "1745562453",
    "nanos": 727466000
  },
  "updated_at": {
    "seconds": "1745562651",
    "nanos": 635528000
  }
}
```
**Особенности**
- Восстанавливает мягко удаленные права (сбрасывает deleted_at в NULL)
- Проверяет принадлежность права к указанному приложению (app_id)
- Возвращает полный объект с актуальным состоянием
- Автоматически обновляет updated_at

**Ошибки**

```json
{
"code": 5,
"message": "Permission not found or already active"
}
```
```json
{
"code": 3,
"message": "Permission is permanently deleted"
}
```

**Логика работы**
- Проверка существования права по id и app_id
- Валидация состояния:
- Право должно быть мягко удалено (deleted_at IS NOT NULL)
- Не должно быть полностью удалено (нет записи в БД)
- Сброс deleted_at и обновление updated_at
- Возврат восстановленного объекта

### 🛡️ Политики безопасности
- Валидация app_id для межприложенной изоляции
- Шифрование чувствительных метаданных
- Контроль доступа на уровне приложения
- JWT-авторизация с проверкой прав
- Подробное логирование операций

### 🔐 Особенности безопасности
- Ограничение на изменение системных прав
- Проверка зависимостей перед удалением
- Межприложенная изоляция через app_id
- Регулярный аудит назначений прав

### 📦 Зависимости
- PostgreSQL 14+ с расширением pgcrypto
- Go 1.21+ с поддержкой gRPC
- Библиотеки: pgx/v5, google.golang.org/protobuf