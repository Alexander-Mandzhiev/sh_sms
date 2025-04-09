-- +goose Up
-- +goose StatementBegin
-- Секреты (JWT ключи)
CREATE TABLE secrets (
    tenant_id UUID NOT NULL,
    app_id INT NOT NULL,
    secret_type VARCHAR(10) NOT NULL CHECK (secret_type IN ('access', 'refresh')),
    current_secret VARCHAR(512) NOT NULL,
    algorithm VARCHAR(20) NOT NULL DEFAULT 'HS256',
    secret_version INT DEFAULT 1,
    generated_at TIMESTAMP NOT NULL,
    revoked_at TIMESTAMP,           -- Для отзыва токенов
    PRIMARY KEY (tenant_id, app_id, secret_type),
    FOREIGN KEY (tenant_id, app_id) REFERENCES tenant_apps(tenant_id, app_id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS secrets;
-- +goose StatementEnd
