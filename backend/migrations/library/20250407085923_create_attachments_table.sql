-- +goose Up
-- +goose StatementBegin
CREATE TABLE attachments (
    id SERIAL PRIMARY KEY,
    attachment_type_id INT NOT NULL REFERENCES attachment_types(id) ON DELETE RESTRICT,
    subject_id INT REFERENCES subjects(id) ON DELETE SET NULL,
    class_id INT REFERENCES classes(id) ON DELETE SET NULL,
    file_name VARCHAR(255) NOT NULL,
    file_path TEXT NOT NULL,
    mime_type VARCHAR(100),
    file_size BIGINT,
    is_active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Индексы для фильтрации и производительности
CREATE INDEX idx_attachments_active ON attachments(attachment_type_id, subject_id, class_id) WHERE is_active = true;
CREATE INDEX idx_attachments_type ON attachments(attachment_type_id);
CREATE INDEX idx_attachments_subject ON attachments(subject_id);
CREATE INDEX idx_attachments_class ON attachments(class_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS attachments;
-- +goose StatementEnd
