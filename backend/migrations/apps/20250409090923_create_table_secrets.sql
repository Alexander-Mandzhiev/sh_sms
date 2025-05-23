-- +goose Up
-- +goose StatementBegin
-- Секреты
CREATE TABLE secrets (
    client_id UUID NOT NULL,
    app_id INT NOT NULL,
    secret_type VARCHAR(10) NOT NULL CHECK (secret_type IN ('access', 'refresh')),
    current_secret VARCHAR(512) NOT NULL,
    algorithm VARCHAR(20) NOT NULL DEFAULT 'bcrypt',
    secret_version INT DEFAULT 1,
    generated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    revoked_at TIMESTAMP,
    PRIMARY KEY (client_id, app_id, secret_type),
    FOREIGN KEY (app_id) REFERENCES apps(id)
);

CREATE INDEX idx_secrets_generated ON secrets(generated_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS secrets;
-- +goose StatementEnd
