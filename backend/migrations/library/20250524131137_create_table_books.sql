-- +goose Up
-- +goose StatementBegin
-- Таблица книг
CREATE TABLE books (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    description TEXT,
    subject_id INT NOT NULL REFERENCES subjects(id) ON DELETE RESTRICT,
    class_id INT NOT NULL REFERENCES classes(id) ON DELETE RESTRICT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS books;
-- +goose StatementEnd
