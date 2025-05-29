-- +goose Up
-- +goose StatementBegin
CREATE TABLE file_formats (format VARCHAR(10) PRIMARY KEY);

INSERT INTO file_formats (format) VALUES
    ('PDF'), ('EPUB'), ('FB2'), ('TXT'), ('MOBI'), ('AZW3'), ('DjVu'), ('RTF')
    ON CONFLICT (format) DO NOTHING;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS file_formats;
-- +goose StatementEnd