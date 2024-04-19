/*
 Navicat MySQL Data Transfer

 Source Server         : mac
 Source Server Type    : MySQL
 Source Server Version : 80300 (8.3.0)
 Source Host           : localhost:3306
 Source Schema         : cms

 Target Server Type    : MySQL
 Target Server Version : 80300 (8.3.0)
 File Encoding         : 65001

 Date: 19/04/2024 15:34:09
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for area
-- ----------------------------
DROP TABLE IF EXISTS `area`;
CREATE TABLE `area` (
  `area_id` smallint unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` smallint unsigned NOT NULL DEFAULT '0',
  `name` varchar(120) NOT NULL DEFAULT '',
  PRIMARY KEY (`area_id`) USING BTREE,
  KEY `parent_id` (`parent_id`) USING BTREE
) ENGINE=MyISAM AUTO_INCREMENT=4069 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Table structure for commits
-- ----------------------------
DROP TABLE IF EXISTS `commits`;
CREATE TABLE `commits` (
  `commit_id` varchar(40) DEFAULT NULL,
  `author` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `commit_date` timestamp NULL DEFAULT NULL,
  `message` text
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for department
-- ----------------------------
DROP TABLE IF EXISTS `department`;
CREATE TABLE `department` (
  `id` varchar(36) NOT NULL,
  `department_name` varchar(36) NOT NULL,
  `parent_department` varchar(36) DEFAULT NULL,
  `create_time` timestamp NULL DEFAULT (now()),
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `delete_time` timestamp NULL DEFAULT NULL,
  `department_order` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `department_pk_2` (`id`),
  UNIQUE KEY `department_pk_3` (`department_name`),
  KEY `department_department_id_fk` (`parent_department`),
  CONSTRAINT `department_department_id_fk` FOREIGN KEY (`parent_department`) REFERENCES `department` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for file
-- ----------------------------
DROP TABLE IF EXISTS `file`;
CREATE TABLE `file` (
  `id` varchar(255) NOT NULL,
  `file_name` varchar(255) DEFAULT NULL,
  `file_size` int DEFAULT NULL,
  `upload_user` varchar(255) DEFAULT NULL,
  `upload_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `file_pk_2` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for interfaces
-- ----------------------------
DROP TABLE IF EXISTS `interfaces`;
CREATE TABLE `interfaces` (
  `interface_id` varchar(36) NOT NULL,
  `interface_name` varchar(255) DEFAULT NULL,
  `interface_method` varchar(10) NOT NULL,
  `interface_path` varchar(255) NOT NULL,
  `interface_page_id` varchar(36) NOT NULL,
  `interface_dic` varchar(36) NOT NULL,
  `interface_desc` varchar(255) DEFAULT NULL,
  `create_time` timestamp NULL DEFAULT (now()),
  `update_time` timestamp NULL DEFAULT (now()) ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`interface_id`),
  UNIQUE KEY `interface_pk` (`interface_id`),
  KEY `interface_pages_page_id_fk` (`interface_page_id`),
  CONSTRAINT `interface_pages_page_id_fk` FOREIGN KEY (`interface_page_id`) REFERENCES `pages` (`page_id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='接口/资源表';

-- ----------------------------
-- Table structure for logs
-- ----------------------------
DROP TABLE IF EXISTS `logs`;
CREATE TABLE `logs` (
  `id` varchar(36) NOT NULL,
  `user_id` varchar(36) DEFAULT NULL,
  `user_account` varchar(255) DEFAULT NULL,
  `ip` varchar(255) DEFAULT NULL,
  `method` varchar(10) DEFAULT NULL,
  `path` varchar(255) DEFAULT NULL,
  `status` int DEFAULT NULL,
  `duration` varchar(255) DEFAULT NULL,
  `create_time` timestamp NOT NULL DEFAULT (now()),
  PRIMARY KEY (`id`),
  UNIQUE KEY `logs_pk_2` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='日志';

-- ----------------------------
-- Table structure for pages
-- ----------------------------
DROP TABLE IF EXISTS `pages`;
CREATE TABLE `pages` (
  `page_id` varchar(36) NOT NULL,
  `parent_page` varchar(36) DEFAULT NULL,
  `page_name` varchar(255) NOT NULL,
  `page_path` varchar(255) NOT NULL,
  `page_icon` varchar(255) DEFAULT NULL,
  `page_component` varchar(255) DEFAULT NULL,
  `page_order` int NOT NULL,
  `can_edit` int DEFAULT '1',
  `is_out_site` int DEFAULT '0' COMMENT '是否外链',
  `out_site_link` varchar(255) DEFAULT NULL COMMENT '外链地址',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `delete_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`page_id`),
  UNIQUE KEY `pages_pk_2` (`page_id`),
  KEY `pages_parent_page_fk` (`parent_page`),
  CONSTRAINT `pages_parent_page_fk` FOREIGN KEY (`parent_page`) REFERENCES `pages` (`page_id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for roles
-- ----------------------------
DROP TABLE IF EXISTS `roles`;
CREATE TABLE `roles` (
  `role_id` varchar(36) NOT NULL,
  `role_name` varchar(255) NOT NULL,
  `description` text,
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `delete_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`role_id`),
  UNIQUE KEY `roles_pk` (`role_name`),
  UNIQUE KEY `roles_pk_2` (`role_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for roles_interfaces
-- ----------------------------
DROP TABLE IF EXISTS `roles_interfaces`;
CREATE TABLE `roles_interfaces` (
  `role_id` varchar(36) NOT NULL,
  `interface_id` varchar(36) NOT NULL,
  KEY `roles_interfaces_interfaces_interface_id_fk` (`interface_id`),
  KEY `roles_interfaces_roles_role_id_fk` (`role_id`),
  CONSTRAINT `roles_interfaces_interfaces_interface_id_fk` FOREIGN KEY (`interface_id`) REFERENCES `interfaces` (`interface_id`) ON DELETE CASCADE,
  CONSTRAINT `roles_interfaces_roles_role_id_fk` FOREIGN KEY (`role_id`) REFERENCES `roles` (`role_id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='角色接口中间表';

-- ----------------------------
-- Table structure for roles_pages
-- ----------------------------
DROP TABLE IF EXISTS `roles_pages`;
CREATE TABLE `roles_pages` (
  `role_id` varchar(36) DEFAULT NULL,
  `page_id` varchar(36) DEFAULT NULL,
  KEY `roles_pages_pages_page_id_fk` (`page_id`),
  KEY `roles_pages_roles_role_id_fk` (`role_id`),
  CONSTRAINT `roles_pages_pages_page_id_fk` FOREIGN KEY (`page_id`) REFERENCES `pages` (`page_id`) ON DELETE CASCADE,
  CONSTRAINT `roles_pages_roles_role_id_fk` FOREIGN KEY (`role_id`) REFERENCES `roles` (`role_id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` varchar(36) NOT NULL,
  `account` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `nickname` varchar(255) NOT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `create_time` timestamp NULL DEFAULT (now()) ON UPDATE CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT (now()) ON UPDATE CURRENT_TIMESTAMP,
  `delete_time` timestamp NULL DEFAULT NULL,
  `status` int DEFAULT '1',
  `department_id` varchar(36) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `account` (`account`),
  KEY `users_department_id_fk` (`department_id`),
  CONSTRAINT `users_department_id_fk` FOREIGN KEY (`department_id`) REFERENCES `department` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for users_roles
-- ----------------------------
DROP TABLE IF EXISTS `users_roles`;
CREATE TABLE `users_roles` (
  `user_id` varchar(36) NOT NULL,
  `role_id` varchar(36) NOT NULL,
  PRIMARY KEY (`user_id`,`role_id`),
  KEY `role_id` (`role_id`),
  CONSTRAINT `users_roles_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  CONSTRAINT `users_roles_ibfk_2` FOREIGN KEY (`role_id`) REFERENCES `roles` (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

SET FOREIGN_KEY_CHECKS = 1;
