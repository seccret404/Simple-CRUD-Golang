CREATE TABLE menus (
     `id` INT AUTO_INCREMENT PRIMARY KEY,
     `name_product` VARCHAR(255) NOT NULL,
     `price` BIGINT NOT NULL,
     `description_product` VARCHAR(255) NOT NULL,
     `image_product` VARCHAR(255) NOT NULL,
     `stock_product` INT NOT NULL,
     `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
     `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE users (
     `id` INT AUTO_INCREMENT PRIMARY KEY,
     `username` VARCHAR(255) NOT NULL,
     `password` VARCHAR(255) NOT NULL,
     `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
     `updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP


);