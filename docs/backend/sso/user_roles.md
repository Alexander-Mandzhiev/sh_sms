# User-Role Management Service

gRPC-сервис для управления назначением ролей пользователям с поддержкой мультитенантности и контроля сроков доступа.

🔗 **Особенности**
- Назначение и отзыв ролей с учетом принадлежности к клиенту и приложению
- Управление сроками действия назначений
- Пагинированные списки назначений
- Межклиентская изоляция через комбинацию client_id + app_id
- Проверка активности связанных сущностей (пользователей и ролей)

🗃 **Структура данных**

### Таблица `user_roles`
Поле | Тип | Ограничения | Описание
-----|-----|-------------|----------
user_id | UUID | PRIMARY KEY, REFERENCES users(id) ON DELETE CASCADE | Идентификатор пользователя
role_id | UUID | PRIMARY KEY, REFERENCES roles(id) | Идентификатор роли
client_id | UUID | PRIMARY KEY, FOREIGN KEY (client_id, role_id) REFERENCES roles(client_id, id) | Идентификатор клиента
app_id | INT | NOT NULL | Идентификатор приложения
assigned_by | UUID | REFERENCES users(id) | Идентификатор назначившего
expires_at | TIMESTAMPTZ |  | Срок действия назначения
assigned_at | TIMESTAMPTZ | NOT NULL DEFAULT NOW() | Время назначения

### Индексы
- `idx_user_roles_expires` (expires_at) - для быстрого поиска активных назначений
- `idx_user_roles_client_user` (client_id, user_id) - оптимизация выборок по клиенту и пользователю

## 📡 API Методы

### 1. Назначение роли пользователю (Assign)
```protobuf
rpc Assign(AssignRequest) returns (UserRole);

message AssignRequest {
  string user_id = 1;
  string role_id = 2;
  string client_id = 3;
  int32 app_id = 4;
  optional google.protobuf.Timestamp expires_at = 5;
  string assigned_by = 6;
}
```
**Пример запроса**
```json
{
  "user_id": "3e3f9902-ae96-4be8-94f0-762197619e06",
  "role_id": "f35c40f9-a321-425c-9e50-8371ee305c85",
  "client_id": "8268ec76-d6c2-48b5-a0e4-a9c2538b8f48",
  "app_id": 1,
  "assigned_by": "3e3f9902-ae96-4be8-94f0-762197619e06"
}
```
**Пример ответа**
```json
{
  "user_id": "3e3f9902-ae96-4be8-94f0-762197619e06",
  "role_id": "f35c40f9-a321-425c-9e50-8371ee305c85",
  "client_id": "8268ec76-d6c2-48b5-a0e4-a9c2538b8f48",
  "app_id": 1,
  "assigned_by": "3e3f9902-ae96-4be8-94f0-762197619e06",
  "expires_at": null,
  "assigned_at": {
    "seconds": "1745990516",
    "nanos": 172467000
  }
}
```
**Особенности**
- Автоматическая проверка существования пользователя и роли
- Валидация срока действия (не может быть в прошлом)
- Проверка уникальности назначения
- Каскадное удаление при удалении пользователя или роли

**Ошибки**

```json
{
  "code": 6,
  "message": "role assignment exists"
}
```
### 2. Отзыв роли у пользователя (Revoke)
```protobuf
rpc Revoke(RevokeRequest) returns (RevokeResponse);

message RevokeRequest {
  string user_id = 1;
  string role_id = 2;
  string client_id = 3;
  int32 app_id = 4;
}
```
**Пример запроса**

```json
{
  "user_id": "3e3f9902-ae96-4be8-94f0-762197619e06",
  "role_id": "f35c40f9-a321-425c-9e50-8371ee305c85",
  "client_id": "8268ec76-d6c2-48b5-a0e4-a9c2538b8f48",
  "app_id": 1
}
```
**Пример ответа**

```json
{
  "success": true,
  "revoked_at": {
    "seconds": "1745990475",
    "nanos": 333580300
  }
}
```
**Особенности**
- Атомарная операция удаления
- Проверка существования назначения перед удалением
- Возвращает точное время отзыва

### 3. Список ролей пользователя (ListForUser)
```protobuf
rpc ListForUser(UserRequest) returns (UserRolesResponse);

message UserRequest {
  string user_id = 1;
  string client_id = 2;
  int32 app_id = 3;
  int32 page = 4;
  int32 count = 5;
}
```
**Пример запроса**
```json
{
  "user_id": "3e3f9902-ae96-4be8-94f0-762197619e06",
  "client_id": "8268ec76-d6c2-48b5-a0e4-a9c2538b8f48",
  "app_id": 1,
  "page": 1,
  "count": 10
}
```
**Пример ответа**
```json
{
  "assignments": [
    {
      "user_id": "3e3f9902-ae96-4be8-94f0-762197619e06",
      "role_id": "f35c40f9-a321-425c-9e50-8371ee305c85",
      "client_id": "8268ec76-d6c2-48b5-a0e4-a9c2538b8f48",
      "app_id": 1,
      "assigned_by": "3e3f9902-ae96-4be8-94f0-762197619e06",
      "expires_at": null,
      "assigned_at": {
        "seconds": "1745990516",
        "nanos": 172467000
      }
    }
  ],
  "total_count": 1,
  "current_page": 1,
  "app_id": 1
}
```
**Особенности**
- Пагинация с сортировкой по времени назначения
- Фильтрация по активности (expires_at > NOW())
- Проверка прав доступа к данным пользователя

### 4. Список пользователей с ролью (ListForRole)
```protobuf
rpc ListForRole(RoleRequest) returns (UserRolesResponse);

message RoleRequest {
  string role_id = 1;
  string client_id = 2;
  int32 app_id = 3;
  int32 page = 4;
  int32 count = 5;
}
```
**Пример запроса**

```json
{
  "role_id": "f35c40f9-a321-425c-9e50-8371ee305c85",
  "client_id": "8268ec76-d6c2-48b5-a0e4-a9c2538b8f48",
  "app_id": 1,
  "page": 1,
  "count": 10
}
```
**Пример ответа**
```json
{
  "assignments": [
    {
      "user_id": "3e3f9902-ae96-4be8-94f0-762197619e06",
      "role_id": "f35c40f9-a321-425c-9e50-8371ee305c85",
      "client_id": "8268ec76-d6c2-48b5-a0e4-a9c2538b8f48",
      "app_id": 1,
      "assigned_by": "3e3f9902-ae96-4be8-94f0-762197619e06",
      "expires_at": null,
      "assigned_at": {
        "seconds": "1745990516",
        "nanos": 172467000
      }
    }
  ],
  "total_count": 1,
  "current_page": 1,
  "app_id": 1
}
```
**Особенности**
- Постраничная выборка
- Проверка принадлежности роли к указанному приложению
- Фильтрация неактивных назначений

### 🛡️ Политики безопасности
- Строгая валидация UUID для всех идентификаторов
- Проверка принадлежности сущностей к одному клиенту и приложению
- Ограничение максимального количества элементов на страницу (1000)
- Логирование всех операций изменения
- Шифрование чувствительных данных в логах

### 🔐 Особенности безопасности
- Транзакционное выполнение критичных операций
- Запрет модификации системных ролей
- Валидация уровня доступа инициирующего пользователя
- Автоматическая очистка истекших назначений (через задачи Cron)

### 📦 Зависимости
- Таблица users с полями: id (UUID), client_id (UUID), is_active (BOOL)
- Таблица roles с полями: id (UUID), client_id (UUID), app_id (INT), is_active (BOOL)
- PostgreSQL 14+ с расширением pgcrypto
- Библиотеки: pgx/v5, google.golang.org/protobuf