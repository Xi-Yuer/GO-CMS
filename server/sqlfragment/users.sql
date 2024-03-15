/*
 Navicat Premium Data Transfer

 Source Server         : Xi-Yuer
 Source Server Type    : MySQL
 Source Server Version : 80032
 Source Host           : 112.124.28.77:3306
 Source Schema         : cms

 Target Server Type    : MySQL
 Target Server Version : 80032
 File Encoding         : 65001

 Date: 02/03/2024 16:20:29
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`
(
    `id`          varchar(36)  NOT NULL,
    `account`     varchar(255) NOT NULL,
    `password`    varchar(255) NOT NULL,
    `nickname`    varchar(255) NOT NULL,
    `avatar`      varchar(255)      DEFAULT NULL,
    `create_time` timestamp    NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
    `update_time` timestamp    NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
    `delete_time` timestamp    NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
    `status`      int               DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;

SET FOREIGN_KEY_CHECKS = 1;
