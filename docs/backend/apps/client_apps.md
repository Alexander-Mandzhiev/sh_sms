# Client Apps Service

gRPC сервис для управления связями клиентов и приложений. Обеспечивает создание, управление статусом и получение информации о связях.

## Сущность ClientApp

| Поле          | Тип                  | Обязательность | Ограничения               | Описание                          |
|---------------|----------------------|----------------|---------------------------|-----------------------------------|
| client_id     | string (UUID)        | Обязательно    | Формат UUID v4            | Идентификатор клиента             |
| app_id        | int32                | Обязательно    | > 0, REFERENCES apps(id)  | Идентификатор приложения          |
| is_active     | bool                 | Опционально    | default=true              | Флаг активности связи             |
| created_at    | google.protobuf.Timestamp | Авто        | -                         | Дата создания связи               |
| updated_at    | google.protobuf.Timestamp | Авто        | -                         | Дата последнего обновления        |

## Примеры запросов

### 1. Создание связи (Create)
**Запрос**:
```json
{
  "app_id": 1,
  "client_id": "8268ec76-d6c2-48b5-a0e4-a9c2538b8f48",
  "is_active": true
}
```
Ответ:

```json
{
  "client_id": "8268ec76-d6c2-48b5-a0e4-a9c2538b8f48",
  "app_id": 1,
  "is_active": true,
  "created_at": {
    "seconds": "1744794472",
    "nanos": 649169000
  },
  "updated_at": {
    "seconds": "1744794472",
    "nanos": 649169000
  }
}
```
### 2. Получение связи (Get)
Запрос:

```json
{
  "app_id": 1,
  "client_id": "8268ec76-d6c2-48b5-a0e4-a9c2538b8f48"
}
```
Ответ:

``` json
{
	"client_id": "8268ec76-d6c2-48b5-a0e4-a9c2538b8f48",
	"app_id": 1,
	"is_active": true,
	"created_at": {
		"seconds": "1744794472",
		"nanos": 649169000
	},
	"updated_at": {
		"seconds": "1744794472",
		"nanos": 649169000
	}
}
```
### 3. Обновление связи (Update)
Запрос:

```json

{
  "app_id": 1,
  "client_id": "8268ec76-d6c2-48b5-a0e4-a9c2538b8f48",
  "is_active": false
}
```
Ответ:

```json
{
  "client_id": "8268ec76-d6c2-48b5-a0e4-a9c2538b8f48",
  "app_id": 1,
  "is_active": false,
  "created_at": {
    "seconds": "1744794472",
    "nanos": 649169000
  },
  "updated_at": {
    "seconds": "1744794547",
    "nanos": 721660000
  }
}
```
### 4. Удаление связи (Delete)
Запрос:

```json
{
  "app_id": 1,
  "client_id": "8268ec76-d6c2-48b5-a0e4-a9c2538b8f48"
}
```
Ответ:

```json
{
  "success": true
}
```
### 5. Список связей (List)

Запрос:
```json
{
  "page": 1,
  "count": 20,
  "is_active": true // опциональное поле
}
```
Ответ:

```json
{
  "client_apps": [
    {
      "client_id": "8268ec76-d6c2-48b5-a0e4-a9c2538b8f48",
      "app_id": 1,
      "is_active": true,
      "created_at": {
        "seconds": "1744794472",
        "nanos": 649169000
      },
      "updated_at": {
        "seconds": "1744794585",
        "nanos": 276850000
      }
    }
  ],
  "total_count": 1,
  "page": 1,
  "count": 20
}
```
Обработка ошибок
```json
{
  "code": 6,
  "message": "client_app relation already exists"
}
```
```json

{
  "code": 5,
  "message": "client_app relation not found"
}
```
```json
{
  "code": 3,
  "message": "invalid client_id format"
}
```
Запуск и тестирование
Запуск сервиса:

```bash
go run cmd/apps/main.go -config-path=config/apps/development.yaml
```
Требования
PostgreSQL 13+

Go 1.21+

Настроенные миграции БД

### Примечание
Полная спецификация API доступна в файле client_apps.proto. Для работы сервиса требуется предварительно развернутый сервис управления приложениями (App Manager Service).