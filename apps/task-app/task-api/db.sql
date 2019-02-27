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

-- Insert tasks into taskstore

INSERT INTO `taskstore`.`task` (name, type, content) VALUES ('demo1', 'url', 'http://www.google.com')
INSERT INTO `taskstore`.`task` (name, type, content) VALUES ('demo2', 'url', 'http://www.google.com')
INSERT INTO `taskstore`.`task` (name, type, content) VALUES ('demo3', 'url', 'http://www.google.com')
