-- +goose Up
-- +goose StatementBegin
CREATE TABLE tenant_apps (
    tenant_id UUID NOT NULL REFERENCES tenants(id),
    app_id INT NOT NULL REFERENCES apps(id),
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (tenant_id, app_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tenant_apps;
-- +goose StatementEnd
