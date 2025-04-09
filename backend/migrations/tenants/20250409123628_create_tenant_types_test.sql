-- +goose Up
-- +goose StatementBegin
-- Таблица типов арендаторов (справочник)
CREATE TABLE tenant_types (
    id SERIAL PRIMARY KEY,
    code VARCHAR(50) UNIQUE NOT NULL, -- school, university, college
    name VARCHAR(100) NOT NULL, -- Название клиента
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Индексы для tenant_types
CREATE INDEX idx_tenant_types_name ON tenant_types(name);
CREATE INDEX idx_tenant_types_deleted ON tenant_types(deleted_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tenant_types;
-- +goose StatementEnd
