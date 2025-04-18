-- +goose Up
-- +goose StatementBegin
-- Таблица пользователей
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    client_id UUID NOT NULL,
    email VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255) NOT NULL CHECK (LENGTH(password_hash) = 60),
    full_name TEXT,
    phone VARCHAR(20),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

-- Индексы
CREATE INDEX idx_users_client_active ON users(client_id) WHERE deleted_at IS NULL;
CREATE UNIQUE INDEX idx_users_email_client ON users(client_id, email) WHERE deleted_at IS NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
