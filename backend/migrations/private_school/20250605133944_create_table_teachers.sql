-- +goose Up
-- +goose StatementBegin
-- Таблица преподавателей
CREATE TABLE teachers (
    id UUID PRIMARY KEY,
    client_id UUID NOT NULL,
    full_name VARCHAR(150) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    email VARCHAR(255),
    additional_info TEXT,
    deleted_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_teachers_client ON teachers(client_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS teachers;
-- +goose StatementEnd
