-- +goose Up
-- +goose StatementBegin
-- Роли с поддержкой иерархии и кастомных прав
CREATE TABLE roles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    client_id UUID NOT NULL, -- Ссылка на внешнюю БД
    name VARCHAR(50) NOT NULL,
    description TEXT,
    level INT DEFAULT 0,
    is_custom BOOLEAN DEFAULT FALSE,
    created_by UUID, -- Без FK, т.к. users в той же БД
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);
-- Индекс для поиска по tenant
CREATE INDEX idx_roles_client ON roles(client_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS roles;
-- +goose StatementEnd
