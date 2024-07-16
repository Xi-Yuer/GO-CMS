/*
 Navicat Premium Dump SQL

 Source Server         : mac
 Source Server Type    : MySQL
 Source Server Version : 80038 (8.0.38)
 Source Host           : localhost:3306
 Source Schema         : micro_cms

 Target Server Type    : MySQL
 Target Server Version : 80038 (8.0.38)
 File Encoding         : 65001

 Date: 15/07/2024 18:01:05
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for roles
-- ----------------------------
DROP TABLE IF EXISTS `roles`;
CREATE TABLE `roles` (
  `id` varchar(256) NOT NULL,
  `role_name` varchar(256) NOT NULL,
  `description` varchar(256) DEFAULT NULL,
  `can_edit` tinyint(1) DEFAULT '1',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_roles_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of roles
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` varchar(256) NOT NULL,
  `account` varchar(256) DEFAULT NULL,
  `password` varchar(256) NOT NULL,
  `nick_name` varchar(256) DEFAULT NULL,
  `avatar` varchar(256) DEFAULT NULL,
  `status` tinyint(1) DEFAULT '1',
  `department_id` varchar(256) DEFAULT NULL,
  `is_admin` tinyint(1) DEFAULT '0',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_users_account` (`account`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO `users` (`id`, `account`, `password`, `nick_name`, `avatar`, `status`, `department_id`, `is_admin`, `created_at`, `updated_at`, `deleted_at`) VALUES ('62462788185886720', '秦娜', '123456', '黄磊', 'http://dummyimage.com/125x125', 1, '', 1, '2024-06-21 16:44:49', '2024-06-25 15:51:39', NULL);
INSERT INTO `users` (`id`, `account`, `password`, `nick_name`, `avatar`, `status`, `department_id`, `is_admin`, `created_at`, `updated_at`, `deleted_at`) VALUES ('62462793915305984', '熊静', '123456', 'Sandra', 'http://dummyimage.com/120x600', 1, '', 1, '2024-06-21 16:44:51', '2024-06-21 16:44:51', NULL);
INSERT INTO `users` (`id`, `account`, `password`, `nick_name`, `avatar`, `status`, `department_id`, `is_admin`, `created_at`, `updated_at`, `deleted_at`) VALUES ('62462796561911808', '史艳', '123456', 'Shirley', 'http://dummyimage.com/120x90', 1, '', 1, '2024-06-21 16:44:51', '2024-06-21 16:44:51', NULL);
INSERT INTO `users` (`id`, `account`, `password`, `nick_name`, `avatar`, `status`, `department_id`, `is_admin`, `created_at`, `updated_at`, `deleted_at`) VALUES ('63888987562971136', '于杰', '123456', 'Jessica', 'http://dummyimage.com/250x250', 1, '', 1, '2024-06-25 15:12:02', '2024-06-25 15:12:02', NULL);
INSERT INTO `users` (`id`, `account`, `password`, `nick_name`, `avatar`, `status`, `department_id`, `is_admin`, `created_at`, `updated_at`, `deleted_at`) VALUES ('63889859118370816', '龚丽', '123456', 'Donna', 'http://dummyimage.com/468x60', 1, '', 1, '2024-06-25 15:15:29', '2024-06-25 15:15:29', NULL);
INSERT INTO `users` (`id`, `account`, `password`, `nick_name`, `avatar`, `status`, `department_id`, `is_admin`, `created_at`, `updated_at`, `deleted_at`) VALUES ('63898827001696256', '孙秀英', '123456', 'Laura', 'http://dummyimage.com/88x31', 1, '', 1, '2024-06-25 15:51:08', '2024-06-25 15:51:08', NULL);
INSERT INTO `users` (`id`, `account`, `password`, `nick_name`, `avatar`, `status`, `department_id`, `is_admin`, `created_at`, `updated_at`, `deleted_at`) VALUES ('67082467554627584', '魏娜', '123456', 'Melissa', 'http://dummyimage.com/120x60', 1, '', 1, '2024-07-04 10:41:47', '2024-07-04 10:41:47', NULL);
INSERT INTO `users` (`id`, `account`, `password`, `nick_name`, `avatar`, `status`, `department_id`, `is_admin`, `created_at`, `updated_at`, `deleted_at`) VALUES ('71177962640117760', '龚刚', '123456', 'Jessica', 'http://dummyimage.com/120x90', 1, '', 1, '2024-07-15 09:55:49', '2024-07-15 09:55:49', NULL);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
