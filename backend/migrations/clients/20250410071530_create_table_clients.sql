-- +goose Up
-- +goose StatementBegin
-- Основная таблица клиентов
CREATE TABLE clients (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    address JSONB,
    type_id INT NOT NULL REFERENCES client_types(id) ON DELETE CASCADE,
    email VARCHAR(320),
    phone VARCHAR(20),
    website VARCHAR(255),
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_clients_type ON clients(type_id);
CREATE INDEX idx_clients_active ON clients(is_active);
CREATE INDEX idx_clients_created_at ON clients(created_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS clients;
-- +goose StatementEnd