-- +goose Up
-- +goose StatementBegin
CREATE TABLE attachments (
    book_id INT NOT NULL REFERENCES books(id) ON DELETE CASCADE,
    format VARCHAR(10) NOT NULL REFERENCES file_formats(format),
    file_url TEXT NOT NULL,
    PRIMARY KEY (book_id, format)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS attachments;
-- +goose StatementEnd