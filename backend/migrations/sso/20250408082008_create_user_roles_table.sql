-- +goose Up
-- +goose StatementBegin
-- Назначение ролей пользователям
CREATE TABLE user_roles (
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    role_id UUID NOT NULL,
    client_id UUID NOT NULL,
    assigned_by UUID REFERENCES users(id),
    expires_at TIMESTAMPTZ,
    PRIMARY KEY (user_id, role_id, client_id),
    CONSTRAINT fk_role_client FOREIGN KEY (client_id, role_id) REFERENCES roles(client_id, id) ON DELETE CASCADE
);

-- Индексы для назначений ролей
CREATE INDEX idx_user_roles_expires ON user_roles(expires_at) WHERE expires_at IS NOT NULL;
CREATE INDEX idx_user_roles_client_user ON user_roles(client_id, user_id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_roles;
-- +goose StatementEnd
