# Client Type Management Service

gRPC-сервис для управления типами клиентов с поддержкой мягкого удаления, уникальных идентификаторов и управления жизненным циклом.

## 🚀 Особенности

- Полный CRUD для типов клиентов
- Мягкое удаление с возможностью восстановления
- Пагинация и фильтрация по активности
- Уникальные коды типов (case-sensitive)
- Автоматическое управление временными метками
- Оптимизированные индексы для быстрого поиска

## 🗃 Структура данных

### Таблица `client_types`

Поле | Тип | Ограничения | Описание
-----|-----|-------------|----------
id | SERIAL | PRIMARY KEY | Автоинкрементный идентификатор
code | VARCHAR(50) | UNIQUE NOT NULL | Уникальный код типа
name | VARCHAR(100) | NOT NULL | Название типа
description | TEXT |  | Описание типа
is_active | BOOLEAN | DEFAULT TRUE | Флаг активности
created_at | TIMESTAMPTZ | DEFAULT CURRENT_TIMESTAMP | Дата создания
updated_at | TIMESTAMPTZ | DEFAULT CURRENT_TIMESTAMP | Дата обновления

### Индексы
- `idx_client_types_active` (is_active)
- Уникальный индекс на поле code

## 📡 API Методы

### 1. Создание типа (Create)
```protobuf
rpc Create(CreateRequest) returns (ClientType);

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
  "code": "VIP",
  "name": "VIP Client",
  "description": "Very Important Client Type",
  "is_active": true
}
```
**Пример успешного ответа**
```json
{
  "id": 42,
  "code": "VIP",
  "name": "VIP Client",
  "description": "Very Important Client Type",
  "is_active": true,
  "created_at": {
    "seconds": 1718901234,
    "nanos": 567890000
  },
  "updated_at": {
    "seconds": 1718901234,
    "nanos": 567890000
  }
}
```
**Пример ошибки (ALREADY_EXISTS)**
```json
{
  "code": 6,
  "message": "code already exists: VIP",
  "details": []
}
```
**Пример ошибки (INVALID_ARGUMENT)**
```json
{
  "code": 3,
  "message": "invalid argument: code is required",
  "details": [
    {
      "@type": "type.googleapis.com/google.rpc.BadRequest",
      "field_violations": [
        {
          "field": "code",
          "description": "field is required"
        }
      ]
    }
  ]
}
```
**Особенности реализации:**
1. Валидация длины полей:
- Code: 1-50 символов
- Name: 1-100 символов
2. Автоматические значения:
- id генерируется последовательностью БД
- created_at и updated_at устанавливаются сервером
- is_active по умолчанию true
3. Проверки в базе данных:
- Уникальность кода через UNIQUE constraint
- Автоматическая проверка длин через прикладную логику

### 2. Получение типа (Get)
```protobuf
rpc Get(GetRequest) returns (ClientType);

message GetRequest {
  int32 id = 1;
}
```
**Пример запроса**
```json
{
  "id": 1
}
```
**Пример успешного ответа**
```json
{
  "id": 1,
  "code": "VIP",
  "name": "VIP Client",
  "description": "Very Important Client Type",
  "is_active": true,
  "created_at": {
    "seconds": 1718901234,
    "nanos": 567890000
  },
  "updated_at": {
    "seconds": 1718901234,
    "nanos": 567890000
  }
}
```
**Пример ошибки (NOT_FOUND)**
```json
{
  "code": 5,
  "message": "client type not found: 42",
  "details": []
}
```
**Пример ошибки (INVALID_ARGUMENT)**
```json
{
  "code": 3,
  "message": "invalid argument: invalid ID format",
  "details": [
    {
      "@type": "type.googleapis.com/google.rpc.BadRequest",
      "field_violations": [
        {
          "field": "id",
          "description": "must be positive integer"
        }
      ]
    }
  ]
}
```
**Особенности реализации:**

**Валидация входных данных:**
- ID должен быть положительным целым числом
**Проверки в базе данных:**
- Поиск только активных записей (is_active = TRUE)
- Полная выборка всех полей записи
**Логирование:**
- Запись попыток доступа к несуществующим ID
- Отслеживание времени выполнения запроса
### 3. Обновление типа (Update)
```protobuf
rpc Update(UpdateRequest) returns (ClientType);

message UpdateRequest {
  int32 id = 1;
  string code = 2;
  string name = 3;
  string description = 4;
}
```
**Пример запроса**
```json
{
  "id": 42,
  "code": "VIP_2",
  "name": "Updated VIP",
  "description": "New description"
}
```

**Пример успешного ответа**
```json
{
  "id": 42,
  "code": "VIP_2",
  "name": "Updated VIP",
  "description": "New description",
  "is_active": true,
  "created_at": {
    "seconds": 1718901234,
    "nanos": 567890000
  },
  "updated_at": {
    "seconds": 1718905678,
    "nanos": 901234000
  }
}
```

**Пример ошибки (NOT_FOUND)**
```json
{
  "code": 5,
  "message": "client type not found: 42",
  "details": []
}
```
**Пример ошибки (ALREADY_EXISTS)**
```json
{
  "code": 6,
  "message": "code conflict: VIP_2",
  "details": []
}
```
**Пример ошибки (INVALID_ARGUMENT)**
```json
{
  "code": 3,
  "message": "invalid argument: code must be 1-50 chars",
  "details": [
    {
      "@type": "type.googleapis.com/google.rpc.BadRequest",
      "field_violations": [
        {
          "field": "code",
          "description": "exceeds 50 characters"
        }
      ]
    }
  ]
}
```
**Особенности реализации:**
1. **Валидация полей:**
- Обязательные поля: code, name
- Длина code: 1-50 символов
- Длина name: 1-100 символов
2. **Проверки:**
- Существование записи по ID
- Уникальность нового кода (игнорируя текущий ID)
3. **Автоматическое обновление:**
- Поле updated_at устанавливается в текущее время
- Сохраняет оригинальные created_at и is_active
4. **Транзакционность:**
- Все проверки и обновления выполняются в одной транзакции
- Атомарное выполнение операции
### 4. Список типов (List)
```protobuf
rpc List(ListRequest) returns (ListResponse);

message ListRequest {
  int32 page = 1;
  int32 count = 2;
  optional string search = 3;
  optional bool active_only = 4;
}
```
**Пример запроса**
```json
{
  "page": 2,
  "count": 20,
  "search": "vip",
  "active_only": true
}
```

**Пример успешного ответа**
```json
{
  "client_types": [
    {
      "id": 42,
      "code": "VIP",
      "name": "VIP Client",
      "description": "Very Important Client Type",
      "is_active": true,
      "created_at": {
        "seconds": 1718901234,
        "nanos": 567890000
      },
      "updated_at": {
        "seconds": 1718905678,
        "nanos": 901234000
      }
    }
  ],
  "total_count": 35,
  "current_page": 2
}
```

**Пример ошибки (INVALID_ARGUMENT)**
```json
{
  "code": 3,
  "message": "invalid argument: page must be ≥ 1",
  "details": [
    {
      "@type": "type.googleapis.com/google.rpc.BadRequest",
      "field_violations": [
        {
          "field": "page",
          "description": "must be positive integer"
        }
      ]
    }
  ]
}
```

**Особенности реализации:**
1. **Пагинация:**
- Максимальный размер страницы: 100 записей
- Нумерация страниц начинается с 1
- Автоматический расчет смещения: OFFSET = (page-1)*count
2. **Поиск:**
- Ищет по полям code (ILIKE), name (ILIKE) и description (ILIKE)
- Регистронезависимый поиск
- Поддерживает частичное совпадение (добавляет % вокруг запроса)
3. **Фильтрация:**
- active_only=true: только записи с is_active=true
- active_only=false или не указан: все записи
4. **Сортировка:**
- По умолчанию: created_at DESC
- Гарантирует стабильный порядок при одинаковых датах

### 5. Удаление типа (Delete)
```protobuf
rpc Delete(DeleteRequest) returns (google.protobuf.Empty);

message DeleteRequest {
  int32 id = 1;
  optional bool permanent = 2;
}
```
**Пример запроса (мягкое удаление)**
```json
{
  "id": 42,
  "permanent": false
}
```
**Пример запроса (физическое удаление)**
```json
{
  "id": 42,
  "permanent": true
}
```

**Пример успешного ответа**
```json
{}
```

**Пример ошибки (CONFLICT)**
```json
{
  "code": 10,
  "message": "cannot delete used client type: existing clients",
  "details": []
}
```

**Пример ошибки (NOT_FOUND)**
```json
{
  "code": 5,
  "message": "client type not found: 42",
  "details": []
}
```

**Особенности реализации:**
1. **Режимы удаления:**
- Мягкое (по умолчанию): Устанавливает is_active=false
- Физическое: Полное удаление записи из БД (только если нет зависимостей)
2. **Проверки:**
- Существование записи перед удалением
- Наличие связанных клиентов при физическом удалении
- Гарантия атомарности операций через транзакции
3. **Безопасность:**
- Каскадное удаление запрещено
- Автоматический откат транзакции при ошибках
- Проверка прав доступа перед выполнением

### 6. Восстановление типа (Restore)
```protobuf
rpc Restore(RestoreRequest) returns (ClientType);

message RestoreRequest {
  int32 id = 1;
}
```
**Пример успешного ответа**
```json
{
  "id": 42,
  "code": "VIP",
  "name": "VIP Client",
  "description": "Very Important Client Type",
  "is_active": true,
  "created_at": {
    "seconds": 1718901234,
    "nanos": 567890000
  },
  "updated_at": {
    "seconds": 1718905678,
    "nanos": 901234000
  }
}
```
**Пример ошибки (NOT_FOUND)**
```json
{
  "code": 5,
  "message": "client type not found: 42",
  "details": []
}
```
**Пример ошибки (FAILED_PRECONDITION)**
```json
{
  "code": 9,
  "message": "client type already active: 42",
  "details": []
}
```
**Пример ошибки (CONFLICT)**
```json
{
  "code": 10,
  "message": "code conflict: VIP",
  "details": []
}
```
**Особенности реализации:**
1. **Логика восстановления:**
- Активирует is_active=true
- Обновляет updated_at
- Проверяет уникальность кода среди всех записей
2. **Проверки:**
- Существование записи
- Текущий статус is_active
- Уникальность кода после активации
3. **Транзакционность:**
- Все операции выполняются в одной транзакции
- Автоматический откат при конфликте

### 🛡️ Политики безопасности
- Валидация входных параметров
- Проверка прав доступа перед модификацией
- Шифрование чувствительных данных
- Логирование всех операций изменения
- Ограничение частоты запросов (rate limiting)

### 📦 Зависимости
- PostgreSQL 12+
- Go 1.21+

### Библиотеки:
- pgx/v5
- google.golang.org/protobuf
- google.golang.org/grpc