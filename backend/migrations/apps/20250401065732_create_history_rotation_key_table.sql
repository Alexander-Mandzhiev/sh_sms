-- +goose Up
-- +goose StatementBegin

-- Таблица истории ротации ключей
CREATE TABLE key_rotation_history (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    app_id UUID NOT NULL REFERENCES apps(id) ON DELETE CASCADE,
    old_key BYTEA,                          -- Старый ключ (зашифрованный)
    new_key BYTEA NOT NULL,                 -- Новый ключ (зашифрованный)
    rotated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    rotated_by UUID NOT NULL                -- ID из сервиса users

-- Индексы для ускорения запросов
CREATE INDEX idx_key_rotation_app_id ON key_rotation_history(app_id);
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS key_rotation_history;
-- +goose StatementEnd
