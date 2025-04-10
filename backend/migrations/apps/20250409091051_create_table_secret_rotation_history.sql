-- +goose Up
-- +goose StatementBegin
-- История ротации секретов
CREATE TABLE secret_rotation_history (
    client_id UUID NOT NULL,
    app_id INT NOT NULL,
    secret_type VARCHAR(10) NOT NULL,
    old_secret VARCHAR(512) NOT NULL,
    new_secret VARCHAR(512) NOT NULL,
    rotated_by UUID,
    rotated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (app_id) REFERENCES apps(id)
);

CREATE INDEX idx_rotation_history_client ON secret_rotation_history(client_id);
CREATE INDEX idx_rotation_history_app ON secret_rotation_history(app_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS secret_rotation_history;
-- +goose StatementEnd
