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
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_by UUID REFERENCES users(id) ON DELETE SET NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

-- Уникальный составной индекс для ссылок между таблицами
CREATE UNIQUE INDEX roles_client_id_idx ON roles(client_id, id);

-- Оптимизация частых запросов
CREATE INDEX roles_client_idx ON roles(client_id);
CREATE INDEX roles_level_idx ON roles(level);
CREATE UNIQUE INDEX roles_name_client_unique_idx ON roles(client_id, name) WHERE deleted_at IS NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS roles;
-- +goose StatementEnd
