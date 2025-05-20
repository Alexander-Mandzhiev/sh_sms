-- +goose Up
-- +goose StatementBegin
-- Назначение ролей пользователям
CREATE TABLE user_roles (
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    role_id UUID NOT NULL,
    client_id UUID NOT NULL,
    app_id INT NOT NULL,
    assigned_by UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    expires_at TIMESTAMPTZ,
    assigned_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, role_id, client_id, app_id),
    CONSTRAINT fk_role_client_app
        FOREIGN KEY (client_id, app_id, role_id)
            REFERENCES roles(client_id, app_id, id)
            ON DELETE CASCADE
);

CREATE INDEX idx_user_roles_expires ON user_roles(expires_at) WHERE expires_at IS NOT NULL;
CREATE INDEX idx_user_roles_client_user ON user_roles(client_id, user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_roles;
-- +goose StatementEnd
