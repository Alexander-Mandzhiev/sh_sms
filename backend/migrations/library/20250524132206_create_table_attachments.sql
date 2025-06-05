-- +goose Up
-- +goose StatementBegin
CREATE TABLE attachments (
    book_id INT NOT NULL REFERENCES books(id) ON DELETE CASCADE,
    format VARCHAR(10) NOT NULL REFERENCES file_formats(format),
    file_id TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- Индексы для ускорения операций
CREATE INDEX idx_attachments_book_id ON attachments(book_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS attachments;
-- +goose StatementEnd