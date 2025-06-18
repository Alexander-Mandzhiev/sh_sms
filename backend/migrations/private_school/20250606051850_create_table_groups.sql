-- +goose Up
-- +goose StatementBegin
-- таблица группы
CREATE TABLE groups (
    internal_id SERIAL PRIMARY KEY,
    public_id UUID UNIQUE NOT NULL,
    client_id UUID NOT NULL,
    name VARCHAR(100) NOT NULL,
    curator_id UUID REFERENCES teachers(id) ON DELETE SET NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (client_id, name)
);

CREATE INDEX idx_groups_client_internal ON groups(client_id, internal_id);
CREATE INDEX idx_groups_curator ON groups(curator_id) WHERE curator_id IS NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS groups;
-- +goose StatementEnd