```protobuf
service UserRoleService {
// Назначение роли пользователю
rpc Assign(AssignRequest) returns (UserRole);

// Отзыв роли у пользователя
rpc Revoke(RevokeRequest) returns (RevokeResponse);

// Получение всех назначений для пользователя
rpc ListForUser(UserRequest) returns (UserRolesResponse);

// Получение всех пользователей с указанной ролью
rpc ListForRole(RoleRequest) returns (UserRolesResponse);

// Массовое назначение/отзыв (опционально)
rpc BulkAssign(BulkAssignRequest) returns (BulkResponse);
rpc BulkRevoke(BulkRevokeRequest) returns (BulkResponse);
}

```
