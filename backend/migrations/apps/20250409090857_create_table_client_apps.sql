-- +goose Up
-- +goose StatementBegin
CREATE TABLE client_apps (
    client_id UUID NOT NULL,
    app_id INT NOT NULL REFERENCES apps(id),
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (client_id, app_id),
    FOREIGN KEY (app_id) REFERENCES apps(id)
);

CREATE INDEX idx_client_apps_client ON client_apps(client_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS client_apps;
-- +goose StatementEnd
