-- +goose Up
-- +goose StatementBegin

-- Вставка типов клиентов
INSERT INTO client_types (code, name, description, is_active) VALUES
    ('internal', 'Internal', 'Внутренние клиенты (например, сотрудники компании)', TRUE),
    ('external', 'External', 'Внешние клиенты (партнёры, пользователи)', TRUE),
    ('partner', 'Partner', 'Клиенты-партнёры', TRUE)
    ON CONFLICT (code) DO NOTHING;

INSERT INTO clients (id, name, description, type_id, website, is_active) VALUES
    ('8268ec76-d6c2-48b5-a0e4-a9c2538b8f48', 'Internal Client', 'Системный клиент для внутренних сервисов', (SELECT id FROM client_types WHERE code = 'internal'), 'https://internal.example.com ', TRUE)
    ON CONFLICT (id) DO NOTHING;
-- +goose StatementEnd
