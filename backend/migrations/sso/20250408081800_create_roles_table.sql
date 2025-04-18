-- +goose Up
-- +goose StatementBegin
-- Таблица ролей
CREATE TABLE roles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    client_id UUID NOT NULL,
    name VARCHAR(150) NOT NULL,
    description TEXT,
    level INT DEFAULT 0 CHECK (level >= 0),
    parent_role_id UUID REFERENCES roles(id),
    is_custom BOOLEAN DEFAULT FALSE,
    created_by UUID REFERENCES users(id) ON DELETE SET NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

-- Индексы для ролей
CREATE INDEX idx_roles_client ON roles(client_id);
CREATE INDEX idx_roles_level ON roles(level);
CREATE UNIQUE INDEX idx_roles_name_client ON roles(client_id, name) WHERE deleted_at IS NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS roles;
-- +goose StatementEnd
