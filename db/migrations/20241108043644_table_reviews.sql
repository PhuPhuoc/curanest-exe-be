-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `reviews` (
    `id` VARCHAR(100) PRIMARY KEY,
    `appointment_id` VARCHAR(100) NOT NULL,
    `patient_name` VARCHAR(255) NOT NULL,
    `rate` SMALLINT,
    `techniques` TEXT,
    `content` TEXT,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` DATETIME,
    CONSTRAINT `fk_reviews_appointment`
    FOREIGN KEY (`appointment_id`) REFERENCES `appointments`(`id`)
    ON DELETE CASCADE ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- ALTER TABLE `reviews` DROP FOREIGN KEY `fk_reviews_appointment`;
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE IF EXISTS `reviews`;
-- +goose StatementEnd
