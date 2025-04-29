# Role-Permission Management Service

gRPC-сервис для управления связями между ролями и правами с поддержкой мультитенантности и контроля доступа.

🔗 **Особенности**
- Управление связями "многие-ко-многим" между ролями и правами
- Автоматическая проверка принадлежности к клиенту и приложению
- Проверка активности связанных сущностей
- Пакетные операции добавления/удаления прав
- Оптимизированные запросы с использованием индексов
- Межклиентская изоляция через комбинацию client_id + app_id

🗃 **Структура данных**

### Таблица `role_permissions`
Поле | Тип | Ограничения | Описание
-----|-----|-------------|----------
role_id | UUID | PRIMARY KEY, REFERENCES roles(id) ON DELETE CASCADE | Идентификатор роли
permission_id | UUID | PRIMARY KEY, REFERENCES permissions(id) ON DELETE CASCADE | Идентификатор права

### Индексы
- `role_permissions_permission_idx` (permission_id)
- `idx_role_permissions_role` (role_id)
- `idx_role_permissions_composite` (role_id, permission_id)

## 📡 API Методы

### 1. Добавление прав к роли (AddPermissionsToRole)
```protobuf
rpc AddPermissionsToRole(PermissionsRequest) returns (OperationStatus);

message PermissionsRequest {
  string role_id = 1;
  string client_id = 2;
  int32 app_id = 3;
  repeated string permission_ids = 4;
}
```
**Пример запроса**

```json
{
  "role_id": "bc870624-ea19-4cc9-951e-9cb19b740812",
  "client_id": "8268ec76-d6c2-48b5-a0e4-a9c2538b8f48",
  "app_id": 1,
  "permission_ids": ["30a7b096-2248-48c9-9427-197cd00bce1f"]
}
```
**Пример ответа**
```json
{
  "success": true,
  "message": "Successfully added 1 permissions",
  "timestamp": {
    "seconds": "1745906417",
    "nanos": 95267800
  }
}
```
**Особенности**
- Пакетное добавление прав
- Проверка принадлежности роли и прав одному приложению
- Автоматическое игнорирование дубликатов
- Проверка активности роли и прав

**Ошибки**
```json
{
  "code": 3,
  "message": "Permission not found"
}
```
### 2. Удаление прав из роли (RemovePermissionsFromRole)
```protobuf
rpc RemovePermissionsFromRole(PermissionsRequest) returns (OperationStatus);

message PermissionsRequest {
  string role_id = 1;
  string client_id = 2;
  int32 app_id = 3;
  repeated string permission_ids = 4;
}
```
**Пример запроса**

```json
{
  "role_id": "bc870624-ea19-4cc9-951e-9cb19b740812",
  "client_id": "8268ec76-d6c2-48b5-a0e4-a9c2538b8f48",
  "app_id": 1,
  "permission_ids": ["30a7b096-2248-48c9-9427-197cd00bce1f"]
}
```
**Пример ответа**
```json
{
  "success": true,
  "message": "Removed 1/1 permissions",
  "timestamp": {
    "seconds": "1745906868",
    "nanos": 938520200
  }
}
```
**Особенности**
- Атомарное удаление нескольких прав
- Возвращает количество фактически удаленных связей
- Проверка прав доступа перед удалением

### 3. Проверка связи (HasPermission)
```protobuf
rpc HasPermission(HasPermissionRequest) returns (HasPermissionResponse);

message HasPermissionRequest {
  string role_id = 1;
  string client_id = 2;
  string permission_id = 3;
  int32 app_id = 4;
}
```
**Пример запроса**
```json
{
  "role_id": "bc870624-ea19-4cc9-951e-9cb19b740812",
  "client_id": "8268ec76-d6c2-48b5-a0e4-a9c2538b8f48",
  "app_id": 1,
  "permission_id": "30a7b096-2248-48c9-9427-197cd00bce1f"
}
```
**Пример ответа**
```json
{
  "has_permission": true,
  "checked_at": {
    "seconds": "1745906707",
    "nanos": 42958700
  }
}
```
**Особенности**
- Проверка активности связанных сущностей
- Оптимизированный запрос с использованием индексов
- Межклиентская изоляция

### 4. Список прав для роли (ListPermissionsForRole)
```protobuf
rpc ListPermissionsForRole(ListPermissionsRequest) returns (ListPermissionsResponse);

message ListPermissionsRequest {
  string role_id = 1;
  string client_id = 2;
  int32 app_id = 3;
}
```
**Пример запроса**
```json
{
  "role_id": "bc870624-ea19-4cc9-951e-9cb19b740812",
  "client_id": "8268ec76-d6c2-48b5-a0e4-a9c2538b8f48",
  "app_id": 1
}

```
**Пример ответа**
```json
{
  "permission_ids": [
    "30a7b096-2248-48c9-9427-197cd00bce1f"
  ]
}
```
**Особенности**
- Возвращает только активные права
- Проверка принадлежности роли клиенту
- Сортировка по времени добавления

### 5. Список ролей для права (ListRolesForPermission)
```protobuf
rpc ListRolesForPermission(ListRolesRequest) returns (ListRolesResponse);

message ListRolesRequest {
  string permission_id = 1;
  string client_id = 2;
  int32 app_id = 3;
}
```
**Пример запроса**
```json
{
  "permission_id": "30a7b096-2248-48c9-9427-197cd00bce1f",
  "client_id": "8268ec76-d6c2-48b5-a0e4-a9c2538b8f48",
  "app_id": 1
}
```
**Пример ответа**
```json
{
  "role_ids": [
    "bc870624-ea19-4cc9-951e-9cb19b740812"
  ]
}
```
**Особенности**
- Фильтрация по активности ролей
- Пагинация результатов
- Проверка принадлежности права приложению

### 🛡️ Политики безопасности
- Строгая проверка client_id и app_id для всех операций
- Валидация UUID на уровне API
- Автоматическая отмена невалидных операций
- Логирование всех изменений связей
- Ограничение максимального количества прав на операцию (1000)

### 🔐 Особенности безопасности
- Транзакционное выполнение пакетных операций
- Запрет модификации системных ролей
- Проверка уровня доступа перед изменением связей
- Шифрование идентификаторов в логах

### 📦 Зависимости
- Таблица roles с полями: id, client_id, app_id, is_active
- Таблица permissions с полями: id, app_id, is_active
- PostgreSQL 14+ с расширением pgcrypto
- Библиотеки: pgx/v5, google.golang.org/protobuf