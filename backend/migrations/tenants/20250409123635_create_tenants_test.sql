-- +goose Up
-- +goose StatementBegin
-- Основная таблица арендаторов
CREATE TABLE tenants (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    address JSONB,
    type_id INT NOT NULL REFERENCES tenant_types(id),
    email VARCHAR(320) UNIQUE,
    phone VARCHAR(20),
    website VARCHAR(2048),
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);
-- Индексы для tenants
CREATE INDEX idx_tenants_type ON tenants(type_id);
CREATE INDEX idx_tenants_email_active ON tenants(email) WHERE deleted_at IS NULL;
CREATE INDEX idx_tenants_active ON tenants(is_active);
CREATE INDEX idx_tenants_deleted ON tenants(deleted_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tenants;
-- +goose StatementEnd
