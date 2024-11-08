-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `appointments` (
    `id` VARCHAR(100) PRIMARY KEY
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `appointments`;
-- +goose StatementEnd
