-- +goose Up
-- +goose StatementBegin

-- Права для работы с пользователями
INSERT INTO permissions (id, code, description, category, app_id, is_active) VALUES
      ('041a6cff-5d0d-4537-9f8a-2f58a2b3d429', 'user.create', 'Create a new user', 'users', 1, TRUE),
      ('1bc0e2b7-eab5-4a07-9a27-fbd6f2d0a570', 'user.delete', 'Delete a user permanently', 'users', 1, TRUE),
      ('ddc8ba1a-81ec-47c4-8e2c-c86818b40d4c', 'user.read', 'Retrieve details of a specific user', 'users', 1, TRUE),
      ('854869bf-4455-4c97-a19e-53a5a531dc77', 'user.list', 'List all users', 'users', 1, TRUE),
      ('147f5bc5-f071-47a3-b23b-b6d076f51676', 'user.update', 'Update user information', 'users', 1, TRUE),
      ('da3ca285-8001-4dca-8486-ad667bd5c4a8', 'user.restore', 'Restore a deleted user', 'users', 1, TRUE)
    ON CONFLICT (code, app_id) WHERE deleted_at IS NULL DO NOTHING;

-- Права для работы с ролями
INSERT INTO permissions (id, code, description, category, app_id, is_active) VALUES
      ('a6feac36-3d67-45f8-b787-9e1e6c322096', 'role.create', 'Create a new role', 'roles', 1, TRUE),
      ('adaa977a-6573-4437-a15c-2b9511193022', 'role.delete', 'Delete a role permanently', 'roles', 1, TRUE),
      ('2e62e6de-1764-4d74-be50-d08be673bc66', 'role.read', 'Retrieve details of a specific role', 'roles', 1, TRUE),
      ('ff71d389-e0f7-4cba-9df8-0ea95fb18e1f', 'role.list', 'List all roles', 'roles', 1, TRUE),
      ('d7fdfacd-55b5-4c9e-bb39-93240ad55546', 'role.update', 'Update role information', 'roles', 1, TRUE),
      ('11da4c13-8252-439f-a5d8-f85c167b2141', 'role.restore', 'Restore a deleted role', 'roles', 1, TRUE)
    ON CONFLICT (code, app_id) WHERE deleted_at IS NULL DO NOTHING;

-- Права для работы с правами (permissions)
INSERT INTO permissions (id, code, description, category, app_id, is_active) VALUES
      ('ca24487a-dfd6-4ca4-82b7-a86dbc0a8992', 'permission.create', 'Create a new permission', 'permissions', 1, TRUE),
      ('fc257e28-14db-4030-8bc6-741287aab3f3', 'permission.delete', 'Delete a permission permanently', 'permissions', 1, TRUE),
      ('fc57a6b0-9d96-47da-b9df-7745394365f5', 'permission.read', 'Retrieve details of a specific permission', 'permissions', 1, TRUE),
      ('d1f0d48b-32bf-4087-9b50-443a80c0bc06', 'permission.list', 'List all permissions', 'permissions', 1, TRUE),
      ('2fbdf044-18e9-478b-b3e1-eb9f5670ca7c', 'permission.update', 'Update permission information', 'permissions', 1, TRUE),
      ('c7d12711-6d51-4156-b53e-539c517c9c97', 'permission.restore', 'Restore a deleted permission', 'permissions', 1, TRUE)
    ON CONFLICT (code, app_id) WHERE deleted_at IS NULL DO NOTHING;
-- +goose StatementEnd

