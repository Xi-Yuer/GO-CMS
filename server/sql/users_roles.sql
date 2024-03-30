/*
 Navicat Premium Data Transfer

 Source Server         : mac
 Source Server Type    : MySQL
 Source Server Version : 80300
 Source Host           : localhost:3306
 Source Schema         : cms

 Target Server Type    : MySQL
 Target Server Version : 80300
 File Encoding         : 65001

 Date: 05/03/2024 16:43:33
*/
SET NAMES utf8mb4;

SET FOREIGN_KEY_CHECKS = 0;
-- ----------------------------
-- Table structure for users_roles
-- ----------------------------
DROP TABLE
    IF
        EXISTS `users_roles`;
CREATE TABLE `users_roles`
(
    `user_id` VARCHAR(36) NOT NULL,
    `role_id` VARCHAR(36) NOT NULL,
    PRIMARY KEY (`user_id`, `role_id`),
    KEY `role_id` (`role_id`),
    CONSTRAINT `users_roles_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
    CONSTRAINT `users_roles_ibfk_2` FOREIGN KEY (`role_id`) REFERENCES `roles` (`role_id`)
) ENGINE = INNODB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;

SET FOREIGN_KEY_CHECKS = 1;