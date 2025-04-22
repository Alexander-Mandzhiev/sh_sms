# User Management Service

gRPC-сервис для управления пользователями с поддержкой мультитенантности, безопасной аутентификацией и гибкой системой фильтрации.

👥 **Особенности**
- Полный CRUD для пользователей с мягким/жестким удалением
- Безопасное хранение паролей (bcrypt)
- Поддержка мультитенантности через client_id
- Валидация email и телефона при создании/обновлении
- Пагинация и фильтрация по email/телефону/активности
- Автоматическое управление временными метками
- Оптимизированные индексы для быстрого поиска

🗃 **Структура данных**

### Таблица `users`
Поле | Тип | Ограничения | Описание
-----|-----|-------------|----------
id | UUID | PRIMARY KEY, DEFAULT gen_random_uuid() | Уникальный идентификатор
client_id | UUID | NOT NULL | Идентификатор клиента (тенанта)
email | VARCHAR(255) | NOT NULL, CHECK(LENGTH >= 5) | Уникальный email пользователя
password_hash | VARCHAR(60) | NOT NULL, CHECK(LENGTH = 60) | Хеш пароля (bcrypt)
full_name | TEXT |  | Полное имя
phone | VARCHAR(20) |  | Номер телефона
is_active | BOOLEAN | DEFAULT TRUE | Флаг активности
created_at | TIMESTAMPTZ | DEFAULT CURRENT_TIMESTAMP | Дата создания
updated_at | TIMESTAMPTZ |  | Дата последнего обновления
deleted_at | TIMESTAMPTZ |  | Мягкое удаление

### Индексы
- `idx_users_client_active` (client_id) WHERE deleted_at IS NULL
- `idx_users_email_client` UNIQUE (client_id, email) WHERE deleted_at IS NULL
- `idx_users_active` (is_active) WHERE is_active = TRUE

# User Management Service

## 📡 API Методы

### 1. Создание пользователя (Create)
**gRPC Contract**
```protobuf
rpc Create(CreateRequest) returns (User);

message CreateRequest {
    string client_id = 1;
    string email = 2;
    string password = 3;
    string full_name = 4;
    string phone = 5;
}
```
Пример запроса

```json
{
    "client_id": "a3d8f7c2-45b1-4e90-8c6a-12b8e7f4d9c0",
    "email": "user@example.com",
    "password": "SecurePass123!",
    "full_name": "John Doe",
    "phone": "+1234567890"
}
```
**Пример ответа**
```json
{
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "client_id": "a3d8f7c2-45b1-4e90-8c6a-12b8e7f4d9c0",
    "email": "user@example.com",
    "full_name": "John Doe",
    "phone": "+1234567890",
    "is_active": true,
    "created_at": "2024-04-05T12:34:56Z"
}
```
**Особенности**
- Автоматическая генерация UUID
- Хеширование пароля с salt (bcrypt)
- Проверка уникальности email в рамках client_id

**Валидация**:
- email: RFC 5322
- password: минимум 8 символов, 1 цифра, 1 спецсимвол
- phone: E.164 формат

**Ошибки**

```json
{
    "error": {
        "code": 6,
        "message": "User with this email already exists"
    }
}
```
### 2. Получение пользователя (Get)
gRPC Contract

```protobuf
rpc Get(GetRequest) returns (User);

message GetRequest {
    string client_id = 1;
    string id = 2;
}
```
**Пример запроса**

```json
{
    "client_id": "a3d8f7c2-45b1-4e90-8c6a-12b8e7f4d9c0",
    "id": "550e8400-e29b-41d4-a716-446655440000"
}
```
**Пример ответа**

```json
{
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "client_id": "a3d8f7c2-45b1-4e90-8c6a-12b8e7f4d9c0",
    "email": "user@example.com",
    "full_name": "John Doe",
    "phone": "+1234567890",
    "is_active": true,
    "created_at": "2024-04-05T12:34:56Z",
    "updated_at": "2024-04-05T12:34:56Z"
}
```
**Особенности**
- Возвращает 404 если пользователь удален
- Не возвращает password_hash

**Ошибки**
```json
{
    "error": {
        "code": 5,
        "message": "User not found"
    }
}
```
### 3. Обновление пользователя (Update)
gRPC Contract

```protobuf
rpc Update(UpdateRequest) returns (User);

message UpdateRequest {
    string id = 1;
    string client_id = 2;
    optional string email = 3;
    optional string full_name = 4;
    optional string phone = 5;
    optional bool is_active = 6;
}
```
**Пример запроса**

```json
{
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "client_id": "a3d8f7c2-45b1-4e90-8c6a-12b8e7f4d9c0",
    "phone": "+1987654321",
    "is_active": false
}
```
**Пример ответа**

```json
{
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "client_id": "a3d8f7c2-45b1-4e90-8c6a-12b8e7f4d9c0",
    "email": "user@example.com",
    "full_name": "John Doe",
    "phone": "+1987654321",
    "is_active": false,
    "created_at": "2024-04-05T12:34:56Z",
    "updated_at": "2024-04-05T12:35:07Z"
}
```
**Особенности**
- Частичное обновление полей
- Проверка прав доступа (client_id должен совпадать)
- Автоматическое обновление updated_at

**Ошибки**
```json
{
    "error": {
        "code": 7,
        "message": "Permission denied"
    }
}
```
### 4. Удаление пользователя (Delete)
gRPC Contract

```protobuf
rpc Delete(DeleteRequest) returns (SuccessResponse);

message DeleteRequest {
    string id = 1;
    string client_id = 2;
    bool permanent = 3;
}
```
**Пример запроса**
```json
{
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "client_id": "a3d8f7c2-45b1-4e90-8c6a-12b8e7f4d9c0",
    "permanent": false
}
```
**Пример ответа**

```json
{
    "success": true
}
```
**Особенности**
- Мягкое удаление (is_active=false + deleted_at) по умолчанию
- Жесткое удаление при permanent=true
- Каскадное удаление зависимых сущностей

**Ошибки**

```json
{
    "error": {
        "code": 5,
        "message": "User not found"
    }
}
```
### 5. Список пользователей (List)
gRPC Contract

```protobuf
rpc List(ListRequest) returns (ListResponse);

message ListRequest {
    string client_id = 1;
    optional string email_filter = 2;
    optional string phone_filter = 3;
    optional bool active_only = 4;
    int32 page = 5;
    int32 count = 6;
}
```
**Пример запроса**

```json
{
    "client_id": "a3d8f7c2-45b1-4e90-8c6a-12b8e7f4d9c0",
    "email_filter": "@example.com",
    "active_only": true,
    "page": 1,
    "count": 20
}
```
**Пример ответа**

```json
{
    "users": [
        {
            "id": "550e8400-e29b-41d4-a716-446655440000",
            "client_id": "a3d8f7c2-45b1-4e90-8c6a-12b8e7f4d9c0",
            "email": "user@example.com",
            "full_name": "John Doe",
            "phone": "+1987654321",
            "is_active": true,
            "created_at": "2024-04-05T12:34:56Z"
        }
    ],
    "total_count": 15,
    "current_page": 1
}
```
**Особенности**
- Поиск по частичному совпадению email/телефона
- Пагинация с максимальным размером страницы 100
- Сортировка по дате создания (DESC)

**Ошибки**

```json
{
    "error": {
        "code": 3,
        "message": "Invalid pagination parameters"
    }
}
```
### 6. Смена пароля (SetPassword)
gRPC Contract

```protobuf
rpc SetPassword(SetPasswordRequest) returns (SuccessResponse);

message SetPasswordRequest {
    string id = 1;
    string client_id = 2;
    string new_password = 3;
}
```
**Пример запроса**

```json
{
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "client_id": "a3d8f7c2-45b1-4e90-8c6a-12b8e7f4d9c0",
    "new_password": "NewSecurePass123!"
}
```
**Пример ответа**

```json
{
    "success": true
}
```
**Особенности**
- Обязательная проверка текущего пароля в сервисном слое
- Автоматическая инвалидация старых сессий
- История изменений паролей (реализуется отдельно)

**Ошибки**

```json
{
    "error": {
        "code": 9,
        "message": "Password does not meet complexity requirements"
    }
}
```
🛡️ Политики безопасности
- Двухфакторная аутентификация для операций с паролями
- Ограничение 5 попыток входа в минуту
- Шифрование чувствительных данных (телефон, email) в БД
- Автоматическая блокировка после 10 неудачных попыток
- JWT-токены с TTL 15 минут для доступа

### 🔐 Особенности безопасности
- Хранение паролей: bcrypt с cost=12
- Валидация ввода на всех уровнях
- SQL-инъекции: 100% параметризованные запросы
- Логирование: маскировка конфиденциальных данных

### 📦 Зависимости
- PostgreSQL 14+ с расширением pgcrypto
- Go 1.21+ с поддержкой gRPC
- Библиотеки: pgx/v5, google.golang.org/protobuf, golang.org/x/crypto