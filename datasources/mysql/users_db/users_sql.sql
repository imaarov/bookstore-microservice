CREATE TABLE `usersgomicro`.`users` (
	`id` BIGINT(20) NOT NULL AUTO_INCREMENT,
	`first_name` VARCHAR(45) NULL,
	`last_name` VARCHAR(45) NULL,
	`email` VARCHAR(45) NOT NULL,
	`date_created` VARCHAR(45) NULL,
	PRIMARY KEY (`id`),
	UNIQUE INDEX `email_UNIQUE` (`email` ASC)
)