# App Manager Service

gRPC сервис для управления приложениями. Обеспечивает базовые CRUD операции с метаданными приложений.

## Сущность App

| Поле          | Тип                  | Обязательность | Ограничения               | Описание                          |
|---------------|----------------------|----------------|---------------------------|-----------------------------------|
| id            | int32                | Автоинкремент  | > 0                       | Уникальный идентификатор          |
| code          | string               | Обязательно    | UNIQUE, длина ≤ 50        | Человеко-читаемый идентификатор   |
| name          | string               | Обязательно    | длина ≤ 250               | Название приложения               |
| description   | string               | Опционально    | -                         | Описание функционала              |
| is_active     | bool                 | Опционально    | default=true              | Флаг активности                   |
| created_at    | google.protobuf.Timestamp | Авто        | -                         | Дата создания                     |
| updated_at    | google.protobuf.Timestamp | Авто        | -                         | Дата последнего обновления        |

## Примеры запросов

### 1. Создание приложения (Create)
**Запрос**:
```json
{
  "code": "school_crm",
  "name": "School Management",
  "description": "Система управления учебным процессом",
  "is_active": true
}
```
Ответ:
```json
{
  "id": 1,
  "code": "school_crm",
  "name": "School Management",
  "description": "Система управления учебным процессом",
  "is_active": true,
  "created_at": {
    "seconds": "1744618602",
    "nanos": 798227000
  },
  "updated_at": {
    "seconds": "1744618776",
    "nanos": 439292000
  }
}
```
2. Получение приложения (Get)
Запрос по ID:

```json
{
  "id": 1
}
```
Запрос по Code:
```json
{
  "code": "school_crm"
}
```
Ответ:

```json
{
  "id": 1,
  "code": "school_crm",
  "name": "School Management",
  "is_active": true,
  "created_at": {
    "seconds": "1744618602",
    "nanos": 798227000
  },
  "updated_at": {
    "seconds": "1744618776",
    "nanos": 439292000
  }
}
```
3. Обновление приложения (Update)
Запрос:
```json

{
  "id": 1,
  "name": "School Management Pro",
  "code": "school_crm_v2"
}
```
Ответ:

```json
{
  "id": 1,
  "code": "school_crm_v2",
  "name": "School Management Pro",
  "is_active": true,
  "created_at": {
    "seconds": "1744618602",
    "nanos": 798227000
  },
  "updated_at": {
    "seconds": "1744618776",
    "nanos": 439292000
  }
}
```
4. Удаление приложения (Delete)
Запрос:
```json

{
  "id": 1
}
```
Ответ:

```json
{
  "success": true
}
```
5. Список приложений (List)
Запрос:

```json
{
  "page": 1,
  "count": 2,
  "filter_is_active": true
}
```
Ответ:

```json
{
  "apps": [
    {
      "id": 1,
      "code": "school_crm",
      "name": "School Management",
      "is_active": true,
      "created_at": {
        "seconds": "1744618602",
        "nanos": 798227000
      },
      "updated_at": {
        "seconds": "1744618776",
        "nanos": 439292000
      }
    },
    {
      "id": 2,
      "code": "hospital_crm",
      "name": "Hospital Management",
      "is_active": true,
      "created_at": {
        "seconds": "1744618602",
        "nanos": 798227000
      },
      "updated_at": {
        "seconds": "1744618776",
        "nanos": 439292000
      }
    }
  ],
  "total_count": 5,
  "page": 1,
  "count": 2
}
```
Обработка ошибок
```json
{
  "code": 6,
  "message": "code already exists"
}
```
```json
{
  "code": 5,
  "message": "application not found"
}
```
```json
{
  "code": 3,
  "message": "invalid pagination parameters"
}
```
Запуск и тестирование
Запуск сервиса:

```bash
go run cmd/apps/main.go -config-path=config/apps/development.yaml
```


## Требования

PostgreSQL 13+

Go 1.21+

Настроенные миграции БД

## Note
Полная спецификация API доступна в app_manager.proto файле.