-- +goose Up
-- +goose StatementBegin

-- Вставка приложений
INSERT INTO apps (code, name, description, is_active, version) VALUES
    ('CEMS', 'ЦСУО', 'Централизованная система управления образованием/Centralized education management system', TRUE, 1)
    ON CONFLICT (code) DO NOTHING;

-- Назначаем клиенту приложение "mobile_sales"
INSERT INTO client_apps (client_id, app_id, is_active) VALUES
    ('8268ec76-d6c2-48b5-a0e4-a9c2538b8f48', 1, TRUE)
    ON CONFLICT (client_id, app_id) DO NOTHING;

-- Вставляем секреты
INSERT INTO secrets (client_id, app_id, secret_type, current_secret, algorithm, secret_version) VALUES
    ('8268ec76-d6c2-48b5-a0e4-a9c2538b8f48', 1, 'refresh', 'eRiz-rG8gu9aYRazN-27azRAVvZBf3w0mCpqAVv39og=', 'bcrypt', 1),
    ('8268ec76-d6c2-48b5-a0e4-a9c2538b8f48', 1, 'access', '0q3ZgHiVVQOG8xn1I4D_uGy-3yM4dwH9b1zpq-XAhDA=', 'bcrypt', 1)
    ON CONFLICT (client_id, app_id, secret_type) DO NOTHING;

-- +goose StatementEnd