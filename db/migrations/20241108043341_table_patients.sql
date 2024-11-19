-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `patients` (
    `id` VARCHAR(100) PRIMARY KEY,
    `avatar` VARCHAR(255),
    `full_name` VARCHAR(255),
    `phone_number` VARCHAR(20),
    `old` TINYINT,
    `dob` VARCHAR(20),
    `citizen_id` VARCHAR(30) NOT NULL,
    `ward` VARCHAR(50),
    `district` VARCHAR(50),
    `city` VARCHAR(50),
    `address` TEXT,
    `medical_description` TEXT,
    `note_for_nurses` TEXT,
    UNIQUE(`citizen_id`)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `customer_patient` (
    `user_id` VARCHAR(100),
    `patient_id` VARCHAR(100),
    PRIMARY KEY (`user_id`, `patient_id`),
    CONSTRAINT `fk_customer_patient_user`
        FOREIGN KEY (`user_id`) REFERENCES `customers`(`user_id`)
        ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `fk_customer_patient_patient`
        FOREIGN KEY (`patient_id`) REFERENCES `patients`(`id`)
        ON DELETE CASCADE ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE `customer_patient` DROP FOREIGN KEY `fk_customer_patient_user`;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE `customer_patient` DROP FOREIGN KEY `fk_customer_patient_patient`;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS `customer_patient`;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS `patients`;
-- +goose StatementEnd