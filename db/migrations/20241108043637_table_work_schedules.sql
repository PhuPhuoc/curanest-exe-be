-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `work_schedules` (
    `id` VARCHAR(100) PRIMARY KEY,
    `nurse_id` VARCHAR(100) NOT NULL,
    `appointment_id` VARCHAR(100),
    `shift_date` DATE NOT NULL,
    `shift_from` TIME NOT NULL,
    `shift_to` TIME NOT NULL,
    `status` ENUM('available', 'not-available'),
    CONSTRAINT `fk_work_schedules_appointment`
    FOREIGN KEY (`appointment_id`) REFERENCES `appointments`(`id`)
    ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `fk_work_schedules_nurse`
    FOREIGN KEY (`nurse_id`) REFERENCES `nurses`(`user_id`)
    ON DELETE CASCADE ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE `work_schedules` DROP FOREIGN KEY `fk_work_schedules_appointment`;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE `work_schedules` DROP FOREIGN KEY `fk_work_schedules_nurse`;
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE IF EXISTS `work_schedules`;
-- +goose StatementEnd