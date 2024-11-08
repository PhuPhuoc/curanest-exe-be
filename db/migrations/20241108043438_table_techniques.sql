-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `techniques` (
    `id` VARCHAR(100) PRIMARY KEY,
    `name` VARCHAR(255),
    `esttimated_time` TIME,
    `fee` INT
);

CREATE TABLE IF NOT EXISTS `nurse_technique` (
    `nurse_id` VARCHAR(100),
    `technique_id` VARCHAR(100),
    PRIMARY KEY (`nurse_id`, `technique_id`),
    CONSTRAINT `fk_nurse_technique_nurse`
        FOREIGN KEY (`nurse_id`) REFERENCES `nurses`(`user_id`)
        ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `fk_nurse_technique_technique`
        FOREIGN KEY (`technique_id`) REFERENCES `techniques`(`id`)
        ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS `patient_technique` (
    `patient_id` VARCHAR(100),
    `technique_id` VARCHAR(100),
    PRIMARY KEY (`patient_id`, `technique_id`),
    CONSTRAINT `fk_patient_technique_patient`
        FOREIGN KEY (`patient_id`) REFERENCES `patients`(`id`)
        ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `fk_patient_technique_technique`
        FOREIGN KEY (`technique_id`) REFERENCES `techniques`(`id`)
        ON DELETE CASCADE ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE `patient_technique` DROP FOREIGN KEY `fk_patient_technique_patient`;
ALTER TABLE `patient_technique` DROP FOREIGN KEY `fk_patient_technique_technique`;
DROP TABLE IF EXISTS `patient_technique`;

ALTER TABLE `nurse_technique` DROP FOREIGN KEY `fk_nurse_technique_nurse`;
ALTER TABLE `nurse_technique` DROP FOREIGN KEY `fk_nurse_technique_technique`;
DROP TABLE IF EXISTS `nurse_technique`;

DROP TABLE IF EXISTS `techniques`;
-- +goose StatementEnd
