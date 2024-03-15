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

 Date: 08/03/2024 09:20:26
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for roles_pages
-- ----------------------------
DROP TABLE IF EXISTS `roles_pages`;
CREATE TABLE `roles_pages`
(
    `role_id` int NOT NULL,
    `page_id` int NOT NULL,
    PRIMARY KEY (`page_id`, `role_id`),
    KEY `role_id` (`role_id`),
    CONSTRAINT `roles_pages_ibfk_1` FOREIGN KEY (`page_id`) REFERENCES `pages` (`page_id`),
    CONSTRAINT `roles_pages_ibfk_2` FOREIGN KEY (`role_id`) REFERENCES `roles` (`role_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;

SET FOREIGN_KEY_CHECKS = 1;
