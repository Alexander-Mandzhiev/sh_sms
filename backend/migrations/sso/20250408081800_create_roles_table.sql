-- +goose Up
-- +goose StatementBegin
-- Роли с поддержкой иерархии и кастомных прав
CREATE TABLE roles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(50) NOT NULL,
    description TEXT,
    level INT DEFAULT 0, -- Уровень иерархии (0 - системная роль)
    is_custom BOOLEAN DEFAULT FALSE, -- Пользовательская/системная роль
    school_id UUID, -- Привязка к школе (опционально)
    created_by UUID REFERENCES users(id), -- Кто создал роль
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS roles;
-- +goose StatementEnd
