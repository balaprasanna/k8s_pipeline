-- Create DB
CREATE SCHEMA `taskstore` DEFAULT CHARACTER SET utf8 ;

-- Create Table

CREATE TABLE `taskstore`.`task` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(1000) NOT NULL,
  `type` VARCHAR(1000) NULL,
  `content` LONGTEXT NOT NULL,
  `status` INT NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`));
