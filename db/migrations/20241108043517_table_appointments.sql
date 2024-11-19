-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `appointments` (
    `id` VARCHAR(100) PRIMARY KEY,
    `patient_id` VARCHAR(100) NOT NULL,
    `nurse_id` VARCHAR(100) NOT NULL,
    `appointment_date` DATE NOT NULL,
    `time_from_to` VARCHAR(50) NOT NULL,
    `status` ENUM('confirmed', 'reconfirm', 'cancel', 'completed'),
    `techniques` TEXT,
    `total_fee` INT,
    CONSTRAINT `fk_appointments_patient`
    FOREIGN KEY (`patient_id`) REFERENCES `patients`(`id`)
    ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `fk_appointments_nurse`
    FOREIGN KEY (`nurse_id`) REFERENCES `nurses`(`user_id`)
    ON DELETE CASCADE ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE `appointments` DROP FOREIGN KEY `fk_appointments_patient`;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE `appointments` DROP FOREIGN KEY `fk_appointments_nurse`;
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE IF EXISTS `appointments`;
-- +goose StatementEnd