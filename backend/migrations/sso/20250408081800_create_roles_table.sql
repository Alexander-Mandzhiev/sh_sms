-- +goose Up
-- +goose StatementBegin
-- Таблица ролей
CREATE TABLE roles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    client_id UUID NOT NULL,
    app_id INT NOT NULL,
    name VARCHAR(150) NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    level INT DEFAULT 0 CHECK (level >= 0),
    is_custom BOOLEAN NOT NULL DEFAULT TRUE,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_by UUID REFERENCES users(id) ON DELETE SET NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ);


CREATE INDEX roles_client_idx ON roles(client_id);
CREATE INDEX roles_app_idx ON roles(app_id);
CREATE INDEX roles_level_idx ON roles(level);
CREATE UNIQUE INDEX idx_roles_client_id ON roles(client_id, id);
CREATE UNIQUE INDEX roles_client_app_name_unique_idx ON roles(client_id, app_id, name) WHERE deleted_at IS NULL;
CREATE INDEX roles_client_app_id_idx ON roles(client_id, app_id, id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS roles;
-- +goose StatementEnd
