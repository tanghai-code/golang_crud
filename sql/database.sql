DROP DATABASE IF EXISTS `golang_curd`;
CREATE DATABASE `golang_curd`;
USE `golang_curd`;

CREATE TABLE `user` (
    `user_id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT 'User Id',
    `username` VARCHAR(100) NOT NULL COMMENT 'Username',
    `real_name` VARCHAR(100) NOT NULL COMMENT 'User real name',
    `phone_number` VARCHAR(50) NOT NULL COMMENT 'User mobile phone number',
    `email` VARCHAR(50) COMMENT 'User email',
    `password` VARCHAR(100) NOT NULL COMMENT 'User login password',
    `birthday` DATE COMMENT 'User birthday',
    `gender` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'User gender; 1=Unknown, 2=male, 3=female',
    `del_status` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'User delete status; 0=normal, 1=deleted'
) COMMENT 'User Table';
