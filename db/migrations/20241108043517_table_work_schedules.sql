-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `work_schedules` (
    `id` VARCHAR(100) PRIMARY KEY
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `work_schedules`;
-- +goose StatementEnd
