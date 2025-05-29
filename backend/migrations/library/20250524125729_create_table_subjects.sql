-- +goose Up
-- +goose StatementBegin
-- Предметы
CREATE TABLE subjects (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE
);
INSERT INTO subjects (name) VALUES
    ('Математика'),
    ('Русский язык'),
    ('Литература'),
    ('Физика'),
    ('Химия'),
    ('История'),
    ('Биология'),
    ('География'),
    ('Иностранный язык'),
    ('Информатика')
    ON CONFLICT (name) DO NOTHING;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS subjects;
-- +goose StatementEnd