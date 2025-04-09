-- +goose Up
-- +goose StatementBegin
INSERT INTO apps (code, name, description, is_active)
VALUES (
           'school_test',
           'Школа Тест',
           'Тестовое приложение для демонстрации работы школы',
           TRUE
       );
-- +goose StatementEnd