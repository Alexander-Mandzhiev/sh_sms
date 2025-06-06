-- +goose Up
-- +goose StatementBegin
-- таблица связи учеников с группами
CREATE TABLE student_groups (
    student_id UUID NOT NULL REFERENCES students(id) ON DELETE CASCADE,
    group_id UUID NOT NULL REFERENCES groups(id) ON DELETE CASCADE,
    PRIMARY KEY (student_id, group_id)
);

CREATE INDEX idx_student_groups_group ON student_groups(group_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS student_groups;
-- +goose StatementEnd