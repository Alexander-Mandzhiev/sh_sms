-- +goose Up
-- +goose StatementBegin
-- Таблица прав
CREATE TABLE permissions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    code VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    category VARCHAR(50),
    app_id INT NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);
-- Индексы для прав
CREATE UNIQUE INDEX idx_permissions_code_app ON permissions(code, app_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_permissions_category ON permissions(category);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS permissions;
-- +goose StatementEnd
