-- +goose Up
-- +goose StatementBegin
INSERT INTO roles (id, client_id, app_id, name, description, level, is_custom) VALUES
        ('5f43769e-bb58-4120-9ec8-9caaf9409ff3', '8268ec76-d6c2-48b5-a0e4-a9c2538b8f48', 1, 'super_admin', 'System super administrator with full access', 1, FALSE),
        ('9936aed6-680a-4f4b-a093-389a30be4a15', '8268ec76-d6c2-48b5-a0e4-a9c2538b8f48', 1, 'admin', 'Administrator with broad access', 2, FALSE),
        ('898617a3-c23f-4a9a-8247-141b7d723e9a', '8268ec76-d6c2-48b5-a0e4-a9c2538b8f48', 1, 'user', 'Regular user with basic access', 3, FALSE)
    ON CONFLICT (client_id, app_id, name) WHERE deleted_at IS NULL DO NOTHING;


-- Вставляем связи пользователей и ролей
INSERT INTO user_roles (user_id, role_id, client_id, app_id, assigned_by) VALUES
      ('027f7c54-deb3-4210-9fed-71b4f7271fba', '9936aed6-680a-4f4b-a093-389a30be4a15', '8268ec76-d6c2-48b5-a0e4-a9c2538b8f48', 1, '77a94b8e-84b7-4900-9307-83de4f73d6d6'),
      ('875491b6-81e3-41b0-a2de-674a9f5c61e0', '898617a3-c23f-4a9a-8247-141b7d723e9a', '8268ec76-d6c2-48b5-a0e4-a9c2538b8f48', 1, '77a94b8e-84b7-4900-9307-83de4f73d6d6'),
      ('77a94b8e-84b7-4900-9307-83de4f73d6d6', '5f43769e-bb58-4120-9ec8-9caaf9409ff3', '8268ec76-d6c2-48b5-a0e4-a9c2538b8f48', 1,  '77a94b8e-84b7-4900-9307-83de4f73d6d6' )
    ON CONFLICT (user_id, role_id, client_id, app_id) DO NOTHING;

-- +goose StatementEnd

