-- +goose Up
-- +goose StatementBegin
CREATE TABLE contacts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    client_id UUID NOT NULL REFERENCES clients(id) ON DELETE CASCADE,
    full_name VARCHAR(255) NOT NULL,
    position VARCHAR(255),
    email VARCHAR(320),
    phone VARCHAR(30),
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- Правильные индексы для таблицы contacts
CREATE INDEX idx_contacts_client_id ON contacts(client_id);
CREATE INDEX idx_contacts_email ON contacts(email);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS contacts;
-- +goose StatementEnd