-- +goose Up
-- +goose StatementBegin
-- таблица учеников
CREATE TABLE students (
    id UUID PRIMARY KEY,
    client_id UUID NOT NULL,
    full_name VARCHAR(150) NOT NULL,
    contract_number VARCHAR(50) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    email VARCHAR(255) NOT NULL,
    additional_info TEXT NOT NULL DEFAULT '',
    deleted_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_students_client ON students(client_id);
CREATE INDEX idx_students_deleted ON students(deleted_at) WHERE deleted_at IS NULL;
CREATE UNIQUE INDEX uniq_student_contract_active ON students (client_id, contract_number) WHERE deleted_at IS NULL;
CREATE INDEX idx_students_pagination ON students (client_id, created_at, id) WHERE deleted_at IS NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS students;
-- +goose StatementEnd