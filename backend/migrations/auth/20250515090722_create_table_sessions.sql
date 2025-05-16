-- +goose Up
-- +goose StatementBegin
CREATE TABLE sessions (
    session_id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    client_id UUID NOT NULL,
    app_id INT NOT NULL,
    access_token_hash TEXT NOT NULL,
    refresh_token_hash TEXT,
    ip_address INET,
    user_agent TEXT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    last_activity TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMPTZ NOT NULL,
    revoked_at TIMESTAMPTZ,
    is_active BOOLEAN GENERATED ALWAYS AS (
        revoked_at IS NULL
            AND expires_at > CURRENT_TIMESTAMP
        ) STORED
);

-- Индексы для быстрого поиска
CREATE INDEX idx_sessions_user ON sessions(user_id);
CREATE INDEX idx_sessions_active ON sessions(is_active);
CREATE INDEX idx_sessions_client_app ON sessions(client_id, app_id);
CREATE INDEX idx_sessions_expiration ON sessions(expires_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS sessions;
-- +goose StatementEnd