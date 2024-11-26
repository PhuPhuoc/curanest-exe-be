-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `wallet_transactions` (
    `id` VARCHAR(100) PRIMARY KEY,
    `user_id` VARCHAR(100) NOT NULL COMMENT 'ID của điều dưỡng',
    `amount` DECIMAL(15,2) NOT NULL COMMENT 'Số tiền cộng hoặc trừ trong ví',
    `type` ENUM('deposit', 'deduction') NOT NULL COMMENT 'Loại giao dịch: nạp tiền hoặc khấu trừ',
    `related_transaction_id` VARCHAR(100) DEFAULT NULL COMMENT 'ID giao dịch ngân hàng (nếu có)',
    `appointment_id` VARCHAR(100) DEFAULT NULL COMMENT 'ID lịch làm việc (nếu có)',
    `detail` TEXT COMMENT 'Ghi chú chi tiết lý do giao dịch',
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT `fk_wallet_user`
    FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
    ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `fk_wallet_related_transaction`
    FOREIGN KEY (`related_transaction_id`) REFERENCES `transactions`(`order_id`)
    ON DELETE SET NULL ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE `wallet_transactions` DROP FOREIGN KEY `fk_wallet_user`;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE `wallet_transactions` DROP FOREIGN KEY `fk_wallet_related_transaction`;
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE IF EXISTS `wallet_transactions`;
-- +goose StatementEnd
