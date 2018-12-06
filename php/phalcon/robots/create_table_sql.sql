CREATE DATABASE IF NOT EXISTS `phalconTutorial`;
use `phalconTutorial`;

CREATE TABLE IF NOT EXISTS `robots` (
	`id` INT(11) NOT NULL AUTO_INCREMENT,
	`name` VARCHAR(50) NOT NULL,
	`type` VARCHAR(50) NOT NULL,
	`year` INT(11) NOT NULL,
	PRIMARY KEY (`id`)
)
COLLATE='latin1_swedish_ci'
ENGINE=InnoDB
;