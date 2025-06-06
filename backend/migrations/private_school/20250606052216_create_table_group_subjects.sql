-- +goose Up
-- +goose StatementBegin
-- таблица связи групп с предметами
CREATE TABLE group_subjects (
    group_id UUID NOT NULL REFERENCES groups(id) ON DELETE CASCADE,
    subject_id INT NOT NULL REFERENCES subjects(id) ON DELETE CASCADE,
    PRIMARY KEY (group_id, subject_id)
);

CREATE INDEX idx_group_subjects_subject ON group_subjects(subject_id);
CREATE INDEX idx_group_subjects_group_id ON group_subjects(group_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS group_subjects;
-- +goose StatementEnd