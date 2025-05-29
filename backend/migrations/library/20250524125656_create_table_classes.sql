-- +goose Up
-- +goose StatementBegin
-- Учебные классы (добавлен UNIQUE для grade)
CREATE TABLE classes (
    id SERIAL PRIMARY KEY,
    grade INT NOT NULL UNIQUE CHECK (grade BETWEEN 1 AND 11)
);

INSERT INTO classes (grade) VALUES (1), (2), (3), (4), (5), (6), (7), (8), (9), (10), (11)
    ON CONFLICT (grade) DO NOTHING;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS classes;
-- +goose StatementEnd