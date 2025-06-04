-- +goose Up
-- +goose StatementBegin
CREATE TABLE attachments (
    book_id INT NOT NULL REFERENCES books(id) ON DELETE CASCADE,
    format VARCHAR(10) NOT NULL REFERENCES file_formats(format),
    file_url TEXT NOT NULL,
    deleted_at TIMESTAMP,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- Уникальный индекс для активных записей
CREATE UNIQUE INDEX idx_attachments_active ON attachments (book_id, format) WHERE deleted_at IS NULL;

-- Индексы для ускорения операций
CREATE INDEX idx_attachments_book_id ON attachments(book_id);
CREATE INDEX idx_attachments_deleted_at ON attachments(deleted_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS attachments;
-- +goose StatementEnd