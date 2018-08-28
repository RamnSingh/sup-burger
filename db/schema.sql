USE sup_burger;

DROP TABLE IF EXISTS `stuff`;
CREATE TABLE `stuff` (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(255) NOT NULL,
    `description` TEXT  NULL
)  ENGINE=INNODB;


DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(255) NOT NULL,
    `description` TEXT NULL
)  ENGINE=INNODB;

DROP TABLE IF EXISTS `city`;
CREATE TABLE `city` (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(255) NOT NULL
)  ENGINE=INNODB;


DROP TABLE IF EXISTS `burger`;
CREATE TABLE `burger` (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(255) NOT NULL,
    `description` TEXT NOT NULL,
    `img_path` VARCHAR(1000) NOT NULL,
    `price` DECIMAL NOT NULL,
    `stock` INT NOT NULL,
    `stuff_id` INT NOT NULL,
    FOREIGN KEY (stuff_id)
        REFERENCES stuff (id)
)  ENGINE=INNODB;

DROP TABLE IF EXISTS `burger_stuff`;
CREATE TABLE `burger_stuff` (
    `burger_id` INT NOT NULL,
    `stuff_id` INT NOT NULL,
    FOREIGN KEY (burger_id)
        REFERENCES burger (id),
    FOREIGN KEY (stuff_id)
        REFERENCES stuff (id)
)  ENGINE=INNODB;


DROP TABLE IF EXISTS `city_burger`;
CREATE TABLE `city_burger` (
    `burger_id` INT NOT NULL,
    `city_id` INT NOT NULL,
    FOREIGN KEY (burger_id)
        REFERENCES burger (id),
    FOREIGN KEY (city_id)
        REFERENCES city (id)
)  ENGINE=INNODB;

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `username` VARCHAR(255) NOT NULL,
    `email` VARCHAR(255) NOT NULL,
    `blocked` BOOLEAN NOT NULL,
    `img_path` VARCHAR(1000)  NULL,
    `password` VARCHAR(1000) NOT NULL,
    `street` VARCHAR(1000) NOT NULL,
    `city_id` INT NOT NULL,
    `role_id` INT NOT NULL,
    FOREIGN KEY (city_id)
        REFERENCES city (id),
    FOREIGN KEY (role_id)
        REFERENCES role (id),
    UNIQUE (username),
    UNIQUE (email)
)  ENGINE=INNODB;


DROP TABLE IF EXISTS `order`;
CREATE TABLE `order` (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `date` DATETIME NOT NULL,
    `pdf_path` VARCHAR(1000) NOT NULL,
    `total_price` DECIMAL NOT NULL,
    `user_id` INT NOT NULL,
    FOREIGN KEY (user_id)
        REFERENCES user (id)
)  ENGINE=INNODB;
