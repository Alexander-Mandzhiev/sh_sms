-- +goose Up
-- +goose StatementBegin
-- таблица группы
CREATE TABLE groups (
    id UUID PRIMARY KEY,
    client_id UUID NOT NULL,
    name VARCHAR(100) NOT NULL,
    curator_id UUID REFERENCES teachers(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (client_id, name)
);

CREATE INDEX idx_groups_curator ON groups(curator_id);
CREATE INDEX idx_groups_client ON groups(client_id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS groups;
-- +goose StatementEnd