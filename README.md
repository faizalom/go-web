# go-web

# Rename `.env.example` to `.env`

Please complete all required field in .env

# Please Run this SQL Commands
```sh
CREATE TABLE `users` (
	`id` INT NOT NULL AUTO_INCREMENT, PRIMARY KEY (`id`),
	`google_id` VARCHAR(50) NULL DEFAULT NULL, UNIQUE (`google_id`),
	`email` VARCHAR(50) NOT NULL DEFAULT '', UNIQUE (`email`),
	`first_name` VARCHAR(50) NOT NULL DEFAULT '',
	`last_name` VARCHAR(50) NOT NULL DEFAULT '',
	`password` VARCHAR(255) NOT NULL DEFAULT '',
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` TIMESTAMP on update CURRENT_TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE = InnoDB;

CREATE TABLE `access_token` (
	`id` INT NOT NULL AUTO_INCREMENT, PRIMARY KEY (`id`),
	`user_id` INT NOT NULL, INDEX (`user_id`),
	`token` VARCHAR(255) NOT NULL,
	`expired_at` DATETIME NOT NULL , INDEX (`expired_at`),
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` TIMESTAMP on update CURRENT_TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE = InnoDB;
```