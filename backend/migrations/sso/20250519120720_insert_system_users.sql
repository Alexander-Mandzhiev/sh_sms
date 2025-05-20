-- +goose Up
-- +goose StatementBegin
-- Генерируем хэш для "Password123!" с помощью bcrypt
INSERT INTO users (id, client_id, email, password_hash, full_name, phone, is_active) VALUES
      ('027f7c54-deb3-4210-9fed-71b4f7271fba','8268ec76-d6c2-48b5-a0e4-a9c2538b8f48', 'admin@example.com', '$2a$10$9qF7lXzO4YxjSHRzKtMjCeGvLdGz6B6iZsI48K0W7ZyN3mP7uU75C', 'Admin User', '+9998884455', TRUE),
      ('875491b6-81e3-41b0-a2de-674a9f5c61e0','8268ec76-d6c2-48b5-a0e4-a9c2538b8f48', 'user@example.com', '$2a$10$9qF7lXzO4YxjSHRzKtMjCeGvLdGz6B6iZsI48K0W7ZyN3mP7uU75C', 'Regular User', '+9998884456', TRUE),
      ('77a94b8e-84b7-4900-9307-83de4f73d6d6','8268ec76-d6c2-48b5-a0e4-a9c2538b8f48', 'superadmin@example.com', '$2a$10$9qF7lXzO4YxjSHRzKtMjCeGvLdGz6B6iZsI48K0W7ZyN3mP7uU75C', 'Super Admin', '+9998884457', TRUE)
    ON CONFLICT (client_id, email) WHERE deleted_at IS NULL DO NOTHING;
-- +goose StatementEnd
