-- +goose Up
-- +goose StatementBegin

-- Основная таблица приложений
CREATE TABLE apps (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL UNIQUE,       -- Уникальное имя приложения
    description TEXT,                        -- Описание назначения приложения
    secret_key BYTEA NOT NULL,               -- Зашифрованный ключ (AES-256)
    is_active BOOLEAN DEFAULT TRUE,          -- Флаг активности приложения
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by UUID NOT NULL,                -- ID из сервиса users
    metadata JSONB DEFAULT '{}'::jsonb       -- Гибкие метаданные (версия, теги и т.д.)
);

-- Индексы для ускорения запросов
CREATE INDEX idx_apps_name ON apps(name);
CREATE INDEX idx_apps_created_by ON apps(created_by);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS apps;
-- +goose StatementEnd
