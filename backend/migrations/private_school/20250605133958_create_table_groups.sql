-- +goose Up
CREATE TABLE groups (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL,
    curator_id UUID REFERENCES teachers(id) ON DELETE SET NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- Индекс для поиска групп по куратору
CREATE INDEX idx_group_curator ON groups(curator_id);

-- +goose Down
DROP TABLE IF EXISTS groups;