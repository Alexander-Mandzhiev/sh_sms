# App Manager Service

gRPC-сервис для управления приложениями с поддержкой CRUD операций, версионированием и мягким удалением.

## 📌 Особенности
- **Оптимистичная блокировка** через поле версии (version)
- **Мягкое удаление** (деактивация через флаг `is_active`)
- Валидация входных данных
- Пагинация с фильтрацией по активности
- Полная история изменений через версионирование

## 🗄 Структура данных

### Таблица `apps`
| Поле          | Тип          | Ограничения               | Описание                     |
|---------------|--------------|---------------------------|------------------------------|
| id            | SERIAL       | PRIMARY KEY               | Уникальный идентификатор     |
| code          | VARCHAR(50)  | UNIQUE, NOT NULL          | Уникальный символьный идентификатор |
| name          | VARCHAR(250) | NOT NULL                  | Название приложения          |
| description   | TEXT         |                           | Детальное описание           |
| is_active     | BOOLEAN      | DEFAULT TRUE              | Флаг активности              |
| version       | INT          | DEFAULT 1, NOT NULL       | Версия для контроля изменений|
| created_at    | TIMESTAMP    | DEFAULT CURRENT_TIMESTAMP | Дата создания                |
| updated_at    | TIMESTAMP    | DEFAULT CURRENT_TIMESTAMP | Дата обновления              |

**Индексы**:
- `idx_apps_active` (is_active) - для фильтрации активных записей
- Автоматический уникальный индекс на поле `code` (создан через UNIQUE constraint)

## 📡 API Методы
### 1. Создание приложения
#### gRPC Contract

```protobuf
rpc Create(CreateRequest) returns (App);
```
```protobuf
message CreateRequest {
    string code = 1;
    string name = 2;
    optional string description = 3;
    optional bool is_active = 4;
}
```
**Пример запроса**

```json
{
    "code": "my_app",
    "name": "My Application",
    "description": "Описание приложения",
    "is_active": true
}
```
**Пример ответа**

```json
{
   "id": 1,
   "code": "school_crm",
   "name": "School Management",
   "description": "Система управления учебным процессом",
   "is_active": true,
   "created_at": {
      "seconds": "1744720153",
      "nanos": 774665000
   },
   "updated_at": {
      "seconds": "1744720153",
      "nanos": 774665000
   }
}
```
**Особенности**
- Автоматически генерирует: id, version, created_at, updated_at
**Валидация**:
- code: 1-50 символов
- name: 1-250 символов

**Ошибки**

```json
{
    "error": {
      "code": 409,
      "message": "Приложение с code='my_app' уже существует"
    }
}
```
### 2. Получение приложения
#### gRPC Contract
```protobuf
rpc Get(AppIdentifier) returns (App);
```
```protobuf
message AppIdentifier {
    oneof identifier {
        int32 id = 1;
        string code = 2;
    }
}
```
**Пример запроса (ID)**

```json
{
  "id": 1
}
```
**Пример запроса (Code)**
```json
{
  "code": "school" 
}
```
**Пример ответа**

```json
{
   "id": 1,
   "code": "school",
   "name": "School Management",
   "description": "Система управления учебным процессом",
   "is_active": true,
   "created_at": {
      "seconds": "1744720153",
      "nanos": 774665000
   },
   "updated_at": {
      "seconds": "1744720196",
      "nanos": 260388000
   }
}
```

**Особенности**
- Регистрозависимый поиск по коду
- Возвращает 404 для деактивированных записей

**Ошибки**

```json
    {
    "error": {
        "code": 404,
        "message": "Приложение не найдено"
    }
}
```
### 3. Обновление приложения
#### gRPC Contract

```protobuf
rpc Update(UpdateRequest) returns (App);
```
```protobuf
message UpdateRequest {
    int32 id = 1;
    optional string code = 2;
    optional string name = 3;
    optional string description = 4;
    optional bool is_active = 5;
}
```
**Пример запроса**
```json
{
   "id": 1,
   "code": "school",
   "name": "School Management",
   "description": "Система управления учебным процессом"
}
```
**Пример ответа**

```json
{
   "id": 1,
   "code": "school",
   "name": "School Management",
   "description": "Система управления учебным процессом",
   "is_active": true,
   "created_at": {
      "seconds": "1744720153",
      "nanos": 774665000
   },
   "updated_at": {
      "seconds": "1744720196",
      "nanos": 260388000
   }
}
```
**Особенности**
- Частичное обновление (только указанные поля)
- Автоматически увеличивает version

**Валидация**:
- code: уникальный, 1-50 символов
- name: 1-250 символов

**Ошибки**

```json
{
    "error": {
        "code": 409,
        "message": "Конфликт версий (текущая версия: 4)"
    }
}
```

### 4. Удаление приложения
#### gRPC Contract

```protobuf
rpc Delete(AppIdentifier) returns (DeleteResponse);
```
```protobuf
message DeleteResponse {
  bool success = 1;
}
```
**Пример запроса**

```json
{
    "code": "school"
}
```
**Пример ответа**

```json
{
  "success": true
}
```
**Особенности**
- Мягкое удаление (is_active=false)
- Повторные вызовы возвращают success: true

**Ошибки**

```json
{
    "error": {
        "code": 404,
        "message": "Приложение не найдено"
    }
}
```

### 5. Список приложений
#### gRPC Contract

```protobuf
rpc List(ListRequest) returns (ListResponse);
```
```protobuf
message ListRequest {
   int64 page = 1;
   int64 count = 2;
   optional bool filter_is_active = 3;
  }
```
```protobuf
message ListResponse {
repeated App apps = 1;
int32 total_count = 2;
int64 page = 3;
int64 count = 4;
}
```
**Пример запроса**
```json
{
   "page": 1,
   "count": 10,
   "filter_is_active": true
}
```
**Пример ответа**

```json
{
   "apps": [
      {
         "id": 1,
         "code": "school",
         "name": "School Management",
         "description": "Система управления учебным процессом",
         "is_active": true,
         "created_at": {
            "seconds": "1744720153",
            "nanos": 774665000
         },
         "updated_at": {
            "seconds": "1744720196",
            "nanos": 260388000
         }
      }
   ],
   "total_count": 1,
   "page": "1",
   "count": "10"
}
```
**Особенности**

- Сортировка по created_at (DESC)
- Максимум 100 элементов на страницу
- При count=0 отключает пагинацию

**Ошибки**

```json
{
    "error": {
        "code": 400,
        "message": "Недопустимое значение count: 150"
    }
}
```
### Запуск и требования
**Требования**
- PostgreSQL 13+
- Go 1.21+
- Настроенные миграции БД

#### Запуск

```bash
go run cmd/apps/main.go -config-path=config/apps/development.yaml
```