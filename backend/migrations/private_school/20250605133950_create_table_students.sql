-- +goose Up
-- +goose StatementBegin
-- таблица учеников
CREATE TABLE students (
    id UUID PRIMARY KEY,
    client_id UUID NOT NULL,
    full_name VARCHAR(150) NOT NULL,
    contract_number VARCHAR(50) NOT NULL,
    phone VARCHAR(20),
    email VARCHAR(255),
    additional_info TEXT,
    deleted_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_students_client ON students(client_id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS students;
-- +goose StatementEnd