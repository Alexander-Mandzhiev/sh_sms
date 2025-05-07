# Client Management Service

gRPC-сервис для управления клиентами с поддержкой мягкого удаления, связями с типами клиентов и комплексной валидацией данных.

## 🚀 Особенности

- Полный CRUD для клиентов
- Мягкое удаление с восстановлением
- Пагинация с фильтрацией по типу и активности
- Валидация URL веб-сайтов
- Уникальные UUID идентификаторы
- Автоматическое управление временными метками
- Оптимизированные индексы для быстрого поиска

## 🗃 Структура данных

### Таблица `clients`

Поле | Тип | Ограничения | Описание
-----|-----|-------------|----------
id | UUID | PRIMARY KEY | Уникальный идентификатор (генерируется БД)
name | VARCHAR(255) | NOT NULL | Название клиента
description | TEXT | NOT NULL | Детальное описание
type_id | INT | REFERENCES client_types(id) | Связь с типами клиентов
website | VARCHAR(255) | NOT NULL | URL веб-сайта
is_active | BOOLEAN | DEFAULT TRUE | Флаг активности
created_at | TIMESTAMPTZ | DEFAULT CURRENT_TIMESTAMP | Дата создания
updated_at | TIMESTAMPTZ | DEFAULT CURRENT_TIMESTAMP | Дата обновления

### Индексы
- `idx_clients_type` (type_id)
- `idx_clients_active` (is_active)
- `idx_clients_created_at` (created_at)

## 📡 API Методы

### 1. Создание клиента (Create)
```protobuf
rpc Create(CreateRequest) returns (Client);

message CreateRequest {
  string name = 1;
  string description = 2;
  int32 type_id = 4;
  string website = 5;
}
```
**Пример запроса**
```json
{
  "name": "Test Corp",
  "description": "Пример компании",
  "type_id": 1,
  "website": "https://example.com"
}
```
**Пример успешного ответа**

```json
{
  "id": "fd5e81b6-e7db-45a8-9e62-80fd16db1764",
  "name": "Test Corp",
  "description": "Пример компании",
  "type_id": 1,
  "website": "https://example.com",
  "is_active": true,
  "created_at": {
    "seconds": "1746605299",
    "nanos": 747326000
  },
  "updated_at": {
    "seconds": "1746605299",
    "nanos": 747326000
  }
}
```
**Примеры ошибок**

**INVALID_ARGUMENT (code 3):**

```json
{
  "code": 3,
  "message": "invalid argument: name must be 1-255 characters",
  "details": []
}
```
```json
{
  "code": 3,
  "message": "invalid argument: website must be valid URL",
  "details": []
}
```
**NOT_FOUND (code 5):**
```json
{
  "code": 5,
  "message": "client type not found: 999",
  "details": []
}
```
**Особенности реализации**

**Валидация полей:**
- name: 1-255 символов (обязательное)
- description: 1-10000 символов (обязательное)
- website: валидный URL (обязательное)
- type_id: существующий ID типа клиента

**Автоматические значения:**
- id: генерируется как UUID v4
- is_active: устанавливается в true
- created_at/updated_at: текущее время сервера

**Проверки в БД:**
- Уникальность связки name+type_id
- Существование type_id в client_types
- Валидность URL через прикладную логику

**Безопасность:**
- SQL-инъекции предотвращаются параметризованными запросами
- XSS-фильтрация для текстовых полей

### 2. Получение клиента (Get)
```protobuf
rpc Get(GetRequest) returns (Client);

message GetRequest {
  string id = 1;
}
```
**Пример запроса**

```json
{
  "id": "fd5e81b6-e7db-45a8-9e62-80fd16db1764"
}
```
**Пример успешного ответа**

```json
{
  "id": "fd5e81b6-e7db-45a8-9e62-80fd16db1764",
  "name": "Test Corp",
  "description": "Пример компании",
  "type_id": 1,
  "website": "https://example.com",
  "is_active": true,
  "created_at": {
    "seconds": "1746605299",
    "nanos": 747326000
  },
  "updated_at": {
    "seconds": "1746605299",
    "nanos": 747326000
  }
}
```
**Примеры ошибок**

**INVALID_ARGUMENT (code 3):**

```json
{
  "code": 3,
  "message": "invalid argument: invalid UUID format",
  "details": [
    {
      "@type": "type.googleapis.com/google.rpc.BadRequest",
      "field_violations": [
        {
          "field": "id",
          "description": "invalid UUID length: 10"
        }
      ]
    }
  ]
}
```
**NOT_FOUND (code 5):**

```json
{
  "code": 5,
  "message": "client not found",
  "details": []
}
```

**Особенности реализации**

**Валидация входных данных:**
- Проверка формата UUID (версия 4)
- ID должен соответствовать RFC 4122

**Безопасность:**
- Отсутствие утечки информации о несуществующих записях
- Логирование попыток доступа

**Производительность:**
- Использование индекса по полю id
- Оптимизированный запрос с выборкой всех полей

**Обработка ошибок:**
- 3 попытки повтора при transient errors
- Кеширование результатов на 5 минут

### 3. Обновление клиента (Update)
```protobuf
rpc Update(UpdateRequest) returns (Client);

message UpdateRequest {
  string id = 1;
  optional string name = 2;
  optional string description = 3;
  optional int32 type_id = 4;
  optional string website = 5;
}
```
**Пример запроса**

```json
{
  "id": "fd5e81b6-e7db-45a8-9e62-80fd16db1764",
  "name": "Test Corp",
  "description": "Пример компании upd",
  "type_id": 1,
  "website": "https://example.com"
}
```
**Пример успешного ответа**

```json
{
  "id": "fd5e81b6-e7db-45a8-9e62-80fd16db1764",
  "name": "Test Corp",
  "description": "Пример компании upd",
  "type_id": 1,
  "website": "https://example.com",
  "is_active": true,
  "created_at": {
    "seconds": "1746605299",
    "nanos": 747326000
  },
  "updated_at": {
    "seconds": "1746609903",
    "nanos": 103921000
  }
}
```
**Примеры ошибок**

**INVALID_ARGUMENT (code 3):**

```json
{
  "code": 3,
  "message": "invalid argument: at least one field must be provided",
  "details": []
}
```

```json
{
  "code": 3,
  "message": "invalid argument: no fields to update",
  "details": []
}
```
**NOT_FOUND (code 5):**
    
```json
{
  "code": 5,
  "message": "client type not found: 999",
  "details": []
}
```
```json
{
  "code": 5, 
  "message": "client not found: a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11", 
  "details": []
}
```
**Особенности реализации**

**Правила обновления:**
- Обновляются только переданные не-null поля
- Пустые строки считаются валидными значениями
- При изменении type_id проверяется существование записи в client_types

**Валидация:**
- name: 1-255 символов (если передано)
- website: валидный URL (если передано)
- type_id: положительное целое число (если передано)

**Безопасность:**
- Проверка прав на модификацию
- Оптимистическая блокировка при конфликтах версий
- SQL-инъекции предотвращаются через параметризацию

**Производительность:**
- Использование частичных индексов
- Кеширование предыдущих версий на 1 минуту
- Пакетное обновление связанных сущностей

**Атомарность:**
- Все проверки и обновления в одной транзакции
- Автоматический откат при ошибках
- Гарантия согласованности данных

### 4. Удаление клиента (Delete)
```protobuf
rpc Delete(DeleteRequest) returns (DeleteResponse);

message DeleteRequest {
  string id = 1;
  optional bool permanent = 2;
}

message DeleteResponse {
  bool success = 1;
}
```
**Примеры запросов**

**Мягкое удаление (по умолчанию):**

```json
{
  "id": "fd5e81b6-e7db-45a8-9e62-80fd16db1764"
}
```
Физическое удаление:
```json
{
  "id": "fd5e81b6-e7db-45a8-9e62-80fd16db1764",
  "permanent": true
}
```

**Пример успешного ответа**

```json
{
  "success": true
}
```
**Примеры ошибок**

**NOT_FOUND (code 5):**
```json
{
  "code": 5,
  "message": "client not found: a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11",
  "details": []
}
```
**FAILED_PRECONDITION (code 9):**
```json
{
  "code": 9,
  "message": "cannot delete client with active contracts",
  "details": []
}
```
**INVALID_ARGUMENT (code 3):**

```json
{
  "code": 3,
  "message": "invalid argument: invalid UUID format",
  "details": []
}
```
**Особенности реализации**

**Режимы удаления:**
- Мягкое: UPDATE clients SET is_active = FALSE
- Физическое: DELETE FROM clients (только если нет связанных записей)

**Проверки:**
- Существование клиента перед удалением
- Наличие зависимых записей (заказы, контракты) при физическом удалении
- Корректность UUID

**Безопасность:**
- Каскадное удаление запрещено
- Требуется подтверждение прав администратора для физического удаления
- Аудит операции в системном журнале

**Производительность:**
- Использование составных индексов для проверки зависимостей
- Оптимизация блокировок таблиц
- Пакетная обработка связанных сущностей

**Восстановление:**
- Мягко удаленные клиенты могут быть восстановлены через метод Restore
- Физически удаленные клиенты не подлежат восстановлению
- 
### 5. Список клиентов (List)
```protobuf
rpc List(ListRequest) returns (ListResponse);

message ListRequest {
  int32 page = 1;
  int32 count = 2;
  optional string search = 3;
  optional int32 type_id = 4;
  optional bool active_only = 5;
}

message ListResponse {
  repeated Client clients = 1;
  int32 total_count = 2;
}
```
**Пример запроса**

```json
{
  "page": 1,
  "count": 20,
  "type_id": 1,
  "search": "Test"
}
```
Пример успешного ответа

```json
{
  "clients": [
    {
      "id": "fd5e81b6-e7db-45a8-9e62-80fd16db1764",
      "name": "Test Corp",
      "description": "Пример компании upd",
      "type_id": 1,
      "website": "https://example.com",
      "is_active": true,
      "created_at": {
        "seconds": "1746605299",
        "nanos": 747326000
      },
      "updated_at": {
        "seconds": "1746612916",
        "nanos": 918682000
      }
    }
  ],
  "total_count": 1
}
```
**Примеры ошибок**

**INVALID_ARGUMENT (code 3):**

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
        "description": "value must be positive"
      }
    ]}
  ]
}
```
```json
{
  "code": 3,
  "message": "invalid argument: count exceeds maximum limit (100)",
  "details": []
}
```
**NOT_FOUND (code 5):**

```json
{
  "code": 5,
  "message": "client type not found: 999",
  "details": []
}
```
**Особенности реализации**

**Пагинация:**
- Минимальная страница: 1
- Размер страницы: 1-100 записей
- Формула смещения: offset = (page-1) * count

**Поиск:**
- Ищет одновременно в полях name и description
- Регистронезависимый (ILIKE '%query%')
- Поддерживает спецсимволы: % (любые символы), _ (один символ)

**Фильтрация:**
- active_only=true: только записи с is_active=true (по умолчанию)
- type_id: строгая проверка существования типа

**Производительность:**
- Использует GIN-индексы для полнотекстового поиска
- Кеширование результатов на 1 минуту

**Параллельный подсчет общего количества**

**Сортировка:**
- По умолчанию: updated_at DESC
- Гарантирует порядок при одинаковых значениях: id ASC

**Безопасность:**
- Экранирование специальных символов в поисковом запросе
- Лимит запросов: 10 в секунду на пользователя

### 6. Восстановление клиента (Restore)
```protobuf
rpc Restore(RestoreRequest) returns (Client);

message RestoreRequest {
  string id = 1;
}
```
**Пример запроса**

```json
{
  "id": "fd5e81b6-e7db-45a8-9e62-80fd16db1764"
}
```
**Пример успешного ответа**

```json
{
  "id": "fd5e81b6-e7db-45a8-9e62-80fd16db1764",
  "name": "Test Corp",
  "description": "Пример компании upd",
  "type_id": 1,
  "website": "https://example.com",
  "is_active": true,
  "created_at": {
    "seconds": "1746605299",
    "nanos": 747326000
  },
  "updated_at": {
    "seconds": "1746612780",
    "nanos": 973352000
  }
}
```
**Примеры ошибок**
**INVALID_ARGUMENT (code 3):**

```json
{
  "code": 5,
  "message": "invalid argument: id",
  "details": []
}
```

**NOT_FOUND (code 5):**
```json
{
  "code": 5,
  "message": "NOT_FOUND: client not found",
  "details": []
}
```
**Особенности реализации**

**Логика восстановления:**
- Активирует is_active = TRUE
- Обновляет updated_at до текущего времени
- Проверяет существование связанного type_id

**Валидация:**
- Формат UUID клиента
- Состояние клиента перед восстановлением (is_active = FALSE)
- Наличие конфликтов уникальности (name + type_id)

**Безопасность:**
- Требуется роль администратора для восстановления
- Аудит операции в журнале безопасности
- Проверка ACL перед выполнением

**Транзакционность:**
- Все проверки и обновления в одной транзакции
- Оптимистическая блокировка через версию записи
- Автоматический откат при конфликтах

**Восстановление зависимостей:**
- Автоматическая реактивация связанных сущностей
- Проверка квот и лимитов системы
- Нотификация заинтересованных служб

### 🛡️ Политики безопасности
- Валидация UUID и URL
- Проверка существования связанных сущностей
- Шифрование передаваемых данных
- Аудит изменений через логи
- Ограничение частоты запросов

### 📦 Зависимости
- PostgreSQL 14+
- Go 1.21+
- pgx/v5
- google.golang.org/protobuf
- google.golang.org/grpc