-- +goose Up
-- +goose StatementBegin
-- История ротации секретов (опционально)
CREATE TABLE key_rotation_history (
    tenant_id UUID NOT NULL,
    app_id INT NOT NULL,
    secret_type VARCHAR(10) NOT NULL,
    old_secret VARCHAR(512) NOT NULL,
    new_secret VARCHAR(512) NOT NULL,
    rotated_by UUID,                -- Кто инициировал ротацию
    rotated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (tenant_id, app_id, secret_type) REFERENCES secrets(tenant_id, app_id, secret_type)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS key_rotation_history;
-- +goose StatementEnd
