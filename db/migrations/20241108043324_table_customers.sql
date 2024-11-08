-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `customers` (
    `user_id` VARCHAR(100) PRIMARY KEY,
    `citizen_id_num` VARCHAR(30) NOT NULL,
    `full_name` VARCHAR(255),
    `phone_number` VARCHAR(20),
    `dob` VARCHAR(20),
    `address` TEXT,
    UNIQUE(`citizen_id_num`, `phone_number`),
    CONSTRAINT `fk_user_customer`
    FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
    ON DELETE CASCADE ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE `customers` DROP FOREIGN KEY `fk_user_customer`;
DROP TABLE IF EXISTS `customers`;
-- +goose StatementEnd
