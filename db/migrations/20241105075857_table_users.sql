-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `users` (
    `id` VARCHAR(100) PRIMARY KEY,
    `email` VARCHAR(30) NOT NULL,
    `password` VARCHAR(30) NOT NULL,
    `name` VARCHAR(255),
    `avatar` VARCHAR(255),
    `role` ENUM('admin', 'user', 'customer', 'nurse'),
    `wallet_amount` DECIMAL(15,2) NOT NULL,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` DATETIME,
    UNIQUE(`email`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `users`;
-- +goose StatementEnd
