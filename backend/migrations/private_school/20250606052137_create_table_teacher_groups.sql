-- +goose Up
-- +goose StatementBegin
-- талица связи учителей с группами
CREATE TABLE teacher_groups (
    teacher_id UUID NOT NULL REFERENCES teachers(id) ON DELETE CASCADE,
    group_id UUID NOT NULL REFERENCES groups(id) ON DELETE CASCADE,
    PRIMARY KEY (teacher_id, group_id)
);

CREATE INDEX idx_teacher_groups_teacher_id ON teacher_groups(teacher_id);
CREATE INDEX idx_teacher_groups_group ON teacher_groups(group_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS teacher_groups;
-- +goose StatementEnd