-- +goose Up
-- +goose StatementBegin
-- ответственные лица учеников
CREATE TABLE responsible_persons (
    id UUID PRIMARY KEY,
    client_id UUID NOT NULL,
    student_id UUID NOT NULL REFERENCES students(id) ON DELETE CASCADE,
    full_name VARCHAR(150) NOT NULL,
    phone VARCHAR(20),
    email VARCHAR(255),
    relationship_type VARCHAR(50) NOT NULL,
    is_primary BOOLEAN DEFAULT FALSE,
    deleted_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_responsible_person_student ON responsible_persons(student_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS responsible_persons;
-- +goose StatementEnd