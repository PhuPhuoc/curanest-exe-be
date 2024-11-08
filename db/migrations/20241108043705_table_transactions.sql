-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `transactions` (
    `id` VARCHAR(100) PRIMARY KEY
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `transactions`;
-- +goose StatementEnd
