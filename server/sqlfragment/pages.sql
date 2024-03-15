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

 Date: 05/03/2024 16:15:15
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for pages
-- ----------------------------
DROP TABLE IF EXISTS `pages`;
CREATE TABLE `pages`
(
    `page_id`        int          NOT NULL AUTO_INCREMENT,
    `page_name`      varchar(255) NOT NULL,
    `page_path`      varchar(255) NOT NULL,
    `page_icon`      varchar(255)      DEFAULT NULL,
    `page_component` varchar(255) NOT NULL,
    `parent_page`    int               DEFAULT NULL,
    `all_delete`     tinyint(1)        DEFAULT '1',
    `create_time`    timestamp    NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time`    timestamp    NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `delete_time`    timestamp    NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`page_id`),
    KEY `parent_page` (`parent_page`),
    CONSTRAINT `pages_ibfk_1` FOREIGN KEY (`parent_page`) REFERENCES `pages` (`page_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;

SET FOREIGN_KEY_CHECKS = 1;
