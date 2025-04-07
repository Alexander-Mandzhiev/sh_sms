-- +goose Up
-- +goose StatementBegin
CREATE TABLE classes (
    id SERIAL PRIMARY KEY,
    grade INT NOT NULL,   -- Номер класса (например, 8)
    letter CHAR(1) NOT NULL, -- Буква класса (например, "А")
    UNIQUE (grade, letter),
    is_active BOOLEAN NOT NULL DEFAULT true

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS classes;
-- +goose StatementEnd
