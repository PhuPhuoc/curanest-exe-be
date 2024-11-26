-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `transactions` (
    `order_id` VARCHAR(100) PRIMARY KEY,
    `user_id` VARCHAR(100) NOT NULL,
    `amount` DECIMAL(15,2) NOT NULL,
    `order_info` TEXT,
    `transaction_no` VARCHAR(50),
    `response_code` VARCHAR(2),
    `payment_status` VARCHAR(20) NOT NULL,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `last_updated` DATETIME,
    CONSTRAINT `fk_transaction_user`
    FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
    ON DELETE CASCADE ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE `transactions` DROP FOREIGN KEY `fk_transaction_user`;
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE IF EXISTS `transactions`;
-- +goose StatementEnd
