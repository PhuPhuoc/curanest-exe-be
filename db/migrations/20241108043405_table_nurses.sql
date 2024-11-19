-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `nurses` (
    `user_id` VARCHAR(100) PRIMARY KEY,
    `citizen_id` VARCHAR(30) NOT NULL,
    `full_name` VARCHAR(255) NOT NULL,
    `phone_number` VARCHAR(20),
    `current_workplace` TEXT,
    `expertise` TEXT,
    `certificate` TEXT,
    `education_level` TEXT,
    `work_experience` TEXT,
    `slogan` TEXT,
    UNIQUE(`citizen_id`, `phone_number`),
    CONSTRAINT `fk_user_nurse`
    FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
    ON DELETE CASCADE ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE `nurses` DROP FOREIGN KEY `fk_user_nurse`;
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE IF EXISTS `nurses`;
-- +goose StatementEnd