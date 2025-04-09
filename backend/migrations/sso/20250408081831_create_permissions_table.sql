-- +goose Up
-- +goose StatementBegin
-- Права с категориями для группировки
CREATE TABLE permissions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    code VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    category VARCHAR(50),
    app_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);
-- Индекс для связки код + приложение
CREATE UNIQUE INDEX idx_permissions_code_app ON permissions(code, app_id)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS permissions;
-- +goose StatementEnd
