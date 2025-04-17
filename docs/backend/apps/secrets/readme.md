# Secrets Management Service

gRPC-сервис для безопасного управления секретами приложений с поддержкой криптографической ротации, версионированием и полным аудитом изменений.

🔐 **Особенности**
- Транзакционная генерация и ротация секретов
- Двухфакторная валидация секретов (access/refresh)
- Полная история ротаций с метаданными изменений
- Автоматическая инвалидация предыдущих версий
- Мягкое удаление с возможностью восстановления
- Оптимизированные запросы с использованием составных индексов

🗃 **Структура данных**

### Таблица `secrets`
Поле | Тип | Ограничения | Описание
-----|-----|-------------|----------
client_id | UUID | PRIMARY KEY, NOT NULL | Идентификатор клиента
app_id | INT | PRIMARY KEY, REFERENCES apps(id) | Идентификатор приложения
secret_type | VARCHAR(10) | PRIMARY KEY, CHECK(access/refresh) | Тип секрета
current_secret | VARCHAR(512) | NOT NULL | Текущий активный секрет
algorithm | VARCHAR(20) | DEFAULT 'bcrypt' | Алгоритм хеширования
secret_version | INT | DEFAULT 1 | Версия секрета
generated_at | TIMESTAMP | NOT NULL | Время последней генерации
revoked_at | TIMESTAMP |  | Время отзыва секрета

### Таблица `secret_rotation_history`
Поле | Тип | Ограничения | Описание
-----|-----|-------------|----------
client_id | UUID | NOT NULL | Идентификатор клиента
app_id | INT | NOT NULL, REFERENCES apps(id) | Идентификатор приложения
secret_type | VARCHAR(10) | NOT NULL | Тип секрета
old_secret | VARCHAR(512) | NOT NULL | Предыдущее значение
new_secret | VARCHAR(512) | NOT NULL | Новое значение
rotated_by | UUID |  | Инициатор изменения
rotated_at | TIMESTAMP | DEFAULT CURRENT_TIMESTAMP | Время ротации

🔍 **Индексы**
- `idx_secrets_composite` (client_id, app_id, secret_type) - первичный ключ
- `idx_secrets_generated` (generated_at) - для временных запросов
- `idx_rotation_history` (client_id, rotated_at) - для аудита
- `idx_rotation_timestamp` (rotated_at DESC) - для сортировки по времени

# Secrets Management Service

## 📡 API Методы

### 1. Генерация секрета (Generate)
**gRPC Contract**
```protobuf
rpc Generate(CreateRequest) returns (Secret);

message CreateRequest {
    string client_id = 1;
    int32 app_id = 2;
    string secret_type = 3;
    optional string algorithm = 4;
}
```
**Пример запроса**
```json
{
    "client_id": "a3d8f7c2-45b1-4e90-8c6a-12b8e7f4d9c0",
    "app_id": 42,
    "secret_type": "access"
}
```
**Пример ответа**
```json
{
    "client_id": "a3d8f7c2-45b1-4e90-8c6a-12b8e7f4d9c0",
    "app_id": 42,
    "secret_type": "access",
    "current_secret": "ejJk...WEs=",
    "algorithm": "bcrypt",
    "secret_version": 1,
    "generated_at": {
        "seconds": 1712345678,
        "nanos": 123456789
    }
}
```
**Особенности**
- Автоматическая генерация криптостойкого секрета
- Значение по умолчанию для algorithm: "bcrypt"

**Валидация**:
- client_id: формат UUID v4
- secret_type: "access" или "refresh"
- app_id: > 0

**Ошибки**
```json
{
    "error": {
        "code": 6,
        "message": "Secret already exists for client-app-type combination"
    }
}
```
### 2. Получение секрета (Get)
gRPC Contract
```protobuf

rpc Get(GetRequest) returns (Secret);

message GetRequest {
    string client_id = 1;
    int32 app_id = 2;
    string secret_type = 3;
}
```
**Пример запроса**
```json
{
    "client_id": "a3d8f7c2-45b1-4e90-8c6a-12b8e7f4d9c0",
    "app_id": 1,
    "secret_type": "access"
}
```
**Пример ответа**
```json
{
    "client_id": "a3d8f7c2-45b1-4e90-8c6a-12b8e7f4d9c0",
    "app_id": 42,
    "secret_type": "access",
    "current_secret": "ejJk...WEs=",
    "algorithm": "bcrypt",
    "secret_version": 1,
    "generated_at": {
        "seconds": 1712345678,
        "nanos": 123456789
    },
    "revoked_at": null
}
```
**Особенности**
- Возвращает только активные секреты (revoked_at = null)
- Строгая проверка регистра для secret_type
- Время ответа: < 50ms для активных секретов

**Ошибки**

```json
{
    "error": {
        "code": 5,
        "message": "Secret not found or revoked"
    }
}
```
### 3. Ротация секрета (Rotate)
gRPC Contract
```protobuf
rpc Rotate(RotateRequest) returns (Secret);

message RotateRequest {
    string client_id = 1;
    int32 app_id = 2;
    string secret_type = 3;
    string rotated_by = 4;
}
```
**Пример запроса**
```json
{
    "client_id": "a3d8f7c2-45b1-4e90-8c6a-12b8e7f4d9c0",
    "app_id": 42,
    "secret_type": "access",
    "rotated_by": "d9c0a3d8-45b1-4e90-8c6a-12b8e7f4a3d8"
}
```
**Пример ответа**
```json
{
    "client_id": "a3d8f7c2-45b1-4e90-8c6a-12b8e7f4d9c0",
    "app_id": 42,
    "secret_type": "access",
    "current_secret": "nEwS...Kr3=",
    "algorithm": "bcrypt",
    "secret_version": 2,
    "generated_at": {
        "seconds": 1712345690,
        "nanos": 987654321
    }
}
```
**Особенности**
- Атомарная операция (генерация + обновление + запись в историю)
- Автоматическое увеличение secret_version
- Старый секрет сохраняется в истории ротаций
- Обязательная идентификация инициатора (rotated_by)

**Ошибки**

```json
{
    "error": {
        "code": 9,
        "message": "Secret already revoked, cannot rotate"
    }
}
```

### 4. Отзыв секрета (Revoke)
**gRPC Contract**
```protobuf
rpc Revoke(RevokeRequest) returns (Secret);

message RevokeRequest {
    string client_id = 1;
    int32 app_id = 2;
    string secret_type = 3;
}
```
**Пример запроса**

```json
{
    "client_id": "a3d8f7c2-45b1-4e90-8c6a-12b8e7f4d9c0",
    "app_id": 42,
    "secret_type": "access"
}
```
**Пример ответа**

```json
{
    "client_id": "a3d8f7c2-45b1-4e90-8c6a-12b8e7f4d9c0",
    "app_id": 42,
    "secret_type": "access",
    "current_secret": "ejJk...WEs=",
    "algorithm": "bcrypt",
    "secret_version": 1,
    "generated_at": {
        "seconds": 1712345678,
        "nanos": 123456789
    },
    "revoked_at": {
        "seconds": 1712346000,
        "nanos": 123456789
    }
}
```
**Особенности**
- Мягкий отзыв (отметка времени вместо удаления)
- Последующие запросы Get возвращают ошибку 403
- Возможность восстановления через Rotate

**Ошибки**

```json
{
    "error": {
        "code": 9,
        "message": "Secret already revoked"
    }
}
```
### 5. Удаление секрета (Delete)
gRPC Contract

```protobuf
rpc Delete(DeleteRequest) returns (DeleteResponse);

message DeleteRequest {
    string client_id = 1;
    int32 app_id = 2;
    string secret_type = 3;
}
```
**Пример запроса**
```json
{
    "client_id": "a3d8f7c2-45b1-4e90-8c6a-12b8e7f4d9c0",
    "app_id": 42,
    "secret_type": "access"
}
```
**Пример ответа**

```json
{
    "success": true
}
```
**Особенности**
- Физическое удаление записи
- Каскадное удаление истории ротаций
- Необратимая операция

**Ошибки**

```json
{
    "error": {
        "code": 5,
        "message": "Secret not found"
    }
}
```
### 6. Список секретов (List)
gRPC Contract

```protobuf
rpc List(ListRequest) returns (ListResponse);

message ListRequest {
    message Filter {
        optional string client_id = 1;
        optional int32 app_id = 2;
        optional string secret_type = 3;
        optional bool is_active = 4;
    }
    int64 page = 1;
    int64 count = 2;
    Filter filter = 3;
}
```
**Пример запроса**

```json
{
    "page": 1,
    "count": 10,
    "filter": {
        "app_id": 42,
        "is_active": true
    }
}
```
**Пример ответа**

```json
{
    "secrets": [
        {
            "client_id": "a3d8f7c2-45b1-4e90-8c6a-12b8e7f4d9c0",
            "app_id": 42,
            "secret_type": "access",
            "current_secret": "ejJk...WEs=",
            "algorithm": "bcrypt",
            "secret_version": 2,
            "generated_at": {
                "seconds": 1712345690,
                "nanos": 987654321
            }
        }
    ],
    "total_count": 1,
    "page": 1,
    "count": 10
}
```
**Особенности**
- Максимальный размер страницы: 1000 элементов
- Фильтры комбинируются через AND
- Сортировка по generated_at DESC

**Ошибки**

```json
{
    "error": {
        "code": 3,
        "message": "Invalid pagination parameters"
    }
}
```
### 7. Получение ротации (GetRotation)
gRPC Contract

```protobuf
rpc GetRotation(GetRotationHistoryRequest) returns (RotationHistory);

message GetRotationHistoryRequest {
    string client_id = 1;
    int32 app_id = 2;
    string secret_type = 3;
    google.protobuf.Timestamp rotated_at = 4;
}
```
**Пример запроса**

```json
{
    "client_id": "a3d8f7c2-45b1-4e90-8c6a-12b8e7f4d9c0",
    "app_id": 42,
    "secret_type": "access",
    "rotated_at": {
        "seconds": 1712345690,
        "nanos": 987654321
    }
}
```
**Пример ответа**

```json
{
    "client_id": "a3d8f7c2-45b1-4e90-8c6a-12b8e7f4d9c0",
    "app_id": 42,
    "secret_type": "access",
    "old_secret": "ejJk...WEs=",
    "new_secret": "nEwS...Kr3=",
    "rotated_by": "d9c0a3d8-45b1-4e90-8c6a-12b8e7f4a3d8",
    "rotated_at": {
        "seconds": 1712345690,
        "nanos": 987654321
    }
}
```
**Особенности**
- Точное совпадение времени до наносекунд
- Возвращает полные значения секретов (только для аудита)

**Ошибки**
```json
{
    "error": {
        "code": 5,
        "message": "Rotation record not found"
    }
}
```
### 8. Список ротаций (ListRotations)
gRPC Contract

```protobuf
rpc ListRotations(ListRequest) returns (ListRotationHistoryResponse);

message ListRotationHistoryResponse {
    repeated RotationHistory rotations = 1;
    int32 total_count = 2;
    int64 page = 3;
    int64 count = 4;
}
```
**Пример запроса**

```json
{
    "page": 1,
    "count": 5,
    "filter": {
        "rotated_by": "d9c0a3d8-45b1-4e90-8c6a-12b8e7f4a3d8",
        "rotated_after": {
            "seconds": 1712345000
        }
    }
}
```
**Пример ответа**
```json
{
    "rotations": [
        {
            "client_id": "a3d8f7c2-45b1-4e90-8c6a-12b8e7f4d9c0",
            "app_id": 42,
            "secret_type": "access",
            "old_secret": "ejJk...WEs=",
            "new_secret": "nEwS...Kr3=",
            "rotated_by": "d9c0a3d8-45b1-4e90-8c6a-12b8e7f4a3d8",
            "rotated_at": {
                "seconds": 1712345690,
                "nanos": 987654321
            }
        }
    ],
    "total_count": 1,
    "page": 1,
    "count": 5
}
```
**Особенности**
- Поддержка временных диапазонов (rotated_after/rotated_before)
- Максимальная глубина истории: 90 дней
- Экспорт в CSV через дополнительный параметр

**Ошибки**
```json

{
    "error": {
        "code": 3,
        "message": "Time range exceeds 90 days limit"
    }
}
```
### 🛡️ Политики безопасности
Доступ к истории ротаций только для роли Auditor
- Ограничение 1000 запросов/мин на endpoint ListRotations
- Обязательная аутентификация для операций записи
- Шифрование секретов в базе данных с использованием AES-256-GCM

### 🔐 Особенности безопасности
- Все секреты хешируются перед сохранением
- Маскировка значений в логах (первые 2 и последние 2 символа)
- Ограничение частоты ротаций: 1 раз/мин для access, 1 раз/час для refresh
- Транзакционный уровень изоляции: Repeatable Read

### 📦 **Зависимости**
- PostgreSQL 13+ с расширением pgcrypto
- Go 1.21+ с поддержкой gRPC-Gateway
- Библиотеки: pgx/v5, google/protobuf, opentelemetry
