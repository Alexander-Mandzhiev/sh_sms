-- +goose Up
-- +goose StatementBegin
CREATE TABLE clients (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    address JSONB,
    type_id INT NOT NULL REFERENCES client_types(id),
    email VARCHAR(320) UNIQUE,
    phone VARCHAR(20),
    website VARCHAR(255),
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE INDEX idx_clients_type ON clients(type_id);
CREATE INDEX idx_clients_created_at ON clients(created_at);
CREATE INDEX idx_clients_email_trgm ON clients USING gin(email gin_trgm_ops);
CREATE INDEX idx_clients_phone_trgm ON clients USING gin(phone gin_trgm_ops);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS clients;
-- +goose StatementEnd
