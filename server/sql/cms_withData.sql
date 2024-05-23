/*
 Navicat Premium Dump SQL

 Source Server         : Xi-Yuer
 Source Server Type    : MySQL
 Source Server Version : 80032 (8.0.32)
 Source Host           : 112.124.28.77:3306
 Source Schema         : cms

 Target Server Type    : MySQL
 Target Server Version : 80032 (8.0.32)
 File Encoding         : 65001

 Date: 23/05/2024 17:18:01
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

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
-- Records of commits
-- ----------------------------
BEGIN;
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('ef96c3394bbadfe446e3020eeff88b40769dc004', 'Xi-Yuer', '2214380963@qq.com', '2024-05-17 13:15:00', 'feat: 添加代码生成器页面，调整部分页面路径');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('7da1492158595e3ea30d59ad4e9f4e68ab0d14c7', 'Xi-Yuer', '2214380963@qq.com', '2024-05-17 09:18:55', 'fix:解决全局在 TS 文件中调用Antd Message报错警告⚠️问题');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('3f911f052325882dd6b8f7d4655b9e28139b83e6', 'Xi-Yuer', '2214380963@qq.com', '2024-05-16 09:56:04', 'feat:添加定时任务功能：支持创建、编辑、启动、停止及查看定时任务，优化日志清理逻辑。');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('0bee6cb417b3fb8365233d04d3df5bb95974e006', 'Xi-Yuer', '2214380963@qq.com', '2024-05-16 03:31:52', 'feat:大文件上传，支持断点续传，文件秒传，进度查看，大文件下载');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('d10141712d619cf16ef25189a772908021a1aa25', 'Xi-Yuer', '2214380963@qq.com', '2024-05-15 09:01:49', 'feat: 添加文件上传功能');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('21a03030a949cb5b7dced8104bf6aeea84b3adea', 'Xi-Yuer', '2214380963@qq.com', '2024-05-13 03:25:15', 'workspace:文件目录结构修改');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('0e1c1da6fa45afcde8474dbcb7dcb7caeb9017e7', 'Xi-Yuer', '2214380963@qq.com', '2024-05-13 01:29:19', 'feat:文件管理');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('d69504a9b4a9364f19ff71ff7471f1542b14b47e', 'Xi-Yuer', '2214380963@qq.com', '2024-05-11 14:13:15', 'feat:实现对用户操作的权限控制。');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('63f924030dd8d38b2c707744e3037ce08b85c613', 'Xi-Yuer', '2214380963@qq.com', '2024-05-11 09:15:07', 'feat:菜单管理');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('2c8cea302cd60f6961c5c8709e3dee1bbd670467', 'Xi-Yuer', '2214380963@qq.com', '2024-05-10 09:13:10', 'feat:部门管理');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('80f87517f69fc753c566c9989b506ae100240014', 'Xi-Yuer', '2214380963@qq.com', '2024-05-09 09:50:57', '优化类型定义，使 operation 方法变为可选； 调整 SystemUser 组件的 props 定义； 重新排列了 Main 组件的导入语句。');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('a8a397a7aaf7642e0fe76d826c1360c8cf743d19', 'Xi-Yuer', '2214380963@qq.com', '2024-05-09 08:32:08', 'docs:新增README文件');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('9a81ce7db5bc27757d7126d687b006694594f933', 'Xi-Yuer', '2214380963@qq.com', '2024-05-09 07:50:35', 'feat:新增接口字典项');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('bbc3f70fa5e17b59bc919ea3a2538e68780ef426', 'Xi-Yuer', '2214380963@qq.com', '2024-05-09 07:10:31', 'feat: 更新依赖包到最新版本');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('550737870aee2557c86eec2a20ec909bbf5443a4', 'Xi-Yuer', '2214380963@qq.com', '2024-05-09 06:56:51', 'feat: 添加角色用户授权功能');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('1edb4d92afb34eba3aa66193f2c84a75aca77ca8', 'Xi-Yuer', '2214380963@qq.com', '2024-05-08 04:08:24', '角色管理：添加角色权限设置，优化用户分页，看板完善。');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('971b386faccd567991d97dfa99784443592c2a29', 'Xi-Yuer', '2214380963@qq.com', '2024-05-07 14:48:44', '角色管理：添加角色权限设置，优化用户管理分页功能，看板优化完善。');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('c0838880063537327ca6b08668c6e813e16a1384', 'Xi-Yuer', '2214380963@qq.com', '2024-05-06 10:13:31', 'feat:角色管理，角色权限设置');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('73448be4b4988c42ee6ab9b073544cf67bdfc01f', 'Xi-Yuer', '2214380963@qq.com', '2024-05-02 17:43:09', 'feat:角色管理');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('b10e8983062a187ac3f7e82a7530b681520cd022', 'Xi-Yuer', '2214380963@qq.com', '2024-05-02 16:45:45', 'feat:Excel数据导出');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('1db69a92fc3e018967172d6ba09a1fb1c5e4ac50', 'Xi-Yuer', '2214380963@qq.com', '2024-05-02 07:55:59', 'feat:角色管理/搜索表单封装');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('3988b5ccfaff29b82927629a860c1c3819d52a71', 'Xi-Yuer', '2214380963@qq.com', '2024-04-28 01:18:45', 'sql:sql文件备份');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('e25faa12e80918b1262a4b2e1f19139e244a48ab', 'Xi-Yuer', '2214380963@qq.com', '2024-04-26 10:01:23', 'feat:用户管理:分页');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('30c275d4950a41df295f709475f8f4516039113e', 'Xi-Yuer', '2214380963@qq.com', '2024-04-25 15:31:40', 'feat:用户管理');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('b812681ad10c426f3704680066e2394ab48985a8', 'Xi-Yuer', '2214380963@qq.com', '2024-04-25 08:09:07', 'feat:DashBoard看板优化完善');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('b01aef2bc0e44f196fa75091e794b33216360ff9', 'Xi-Yuer', '2214380963@qq.com', '2024-04-25 03:17:25', 'feat:Git提交日志返回数据结构修改');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('7e30c48814cdcce43ecc15cd6188b024c1d8fbd7', 'Xi-Yuer', '2214380963@qq.com', '2024-04-24 15:22:58', 'feat:Git提交日志返回数据结构修改');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('edd54e30fc4df81969ccd5d1f1021663c6bc7516', 'Xi-Yuer', '2214380963@qq.com', '2024-04-24 10:04:27', 'feat:系统运行信息');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('58ae8cc1faa680fc4266dbd5d31c639a9a630793', 'Xi-Yuer', '2214380963@qq.com', '2024-04-24 09:24:24', 'feat:系统运行信息');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('547697445af9c8c9d299b06d332d6fd495ffe6af', 'Xi-Yuer', '2214380963@qq.com', '2024-04-24 06:34:07', 'style:仪表盘');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('0d3dfffce0e9f67428a18c554c0abbfa527a2c02', 'Xi-Yuer', '2214380963@qq.com', '2024-04-24 05:01:19', 'feat:仪表盘');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('914af363737e39d6c010c4f0ec67897af739eb84', 'Xi-Yuer', '2214380963@qq.com', '2024-04-23 07:41:12', 'feat:ts-md5密码加密');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('63c0556de6ba4f9f470d35a1cb095eccd3545e72', 'Xi-Yuer', '2214380963@qq.com', '2024-04-23 07:29:05', 'feat:tab国际化');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('82becd52ee3279f5109dfaaeec79a212e3c201a7', 'Xi-Yuer', '2214380963@qq.com', '2024-04-23 07:08:24', 'fix:菜单以及面包屑跳转bug');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('cfba37c70e9606802fbd973706a4e7bc9abec26f', 'Xi-Yuer', '2214380963@qq.com', '2024-04-22 04:46:29', 'feat:动态路由搭建');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('54a834d26c43fe4b31adff30be5ecb9630dff736', 'Xi-Yuer', '2214380963@qq.com', '2024-04-21 15:47:25', 'feat:菜单搭建');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('1326bfd1c723612535bef19bdcc09728272a5ecf', 'Xi-Yuer', '2214380963@qq.com', '2024-04-21 14:05:15', 'feat:登录页面');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('c38f1f9d394db801c8b8f0c97b7fdfeffb12ba1d', 'Xi-Yuer', '2214380963@qq.com', '2024-04-20 16:19:09', 'feat:前端搭建');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('de346b8317ed5c47c671daa3dbb2588948ec5b9c', 'Xi-Yuer', '2214380963@qq.com', '2024-04-19 07:35:40', 'sql:sql文件');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('741a0c72cf8b730452e2fdfbe6f1b23fb26cc41e', 'Xi-Yuer', '2214380963@qq.com', '2024-04-19 07:18:54', 'feat:定时任务');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('ee62b7b6342fe93168d5a7f7ad618569737410f6', 'Xi-Yuer', '2214380963@qq.com', '2024-04-19 04:23:01', 'feat:文件下载，添加通过a链接直接下载文件');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('541d17c368fd62f8d2ebc23d991531e8d649ed4e', 'Xi-Yuer', '2214380963@qq.com', '2024-04-18 10:02:10', 'feat:角色管理，添加和删除用户');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('389e4068dbedf9ca9ebe9734fc8a41d11787289d', 'Xi-Yuer', '2214380963@qq.com', '2024-04-18 09:54:53', 'feat:角色管理，添加和删除用户');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('41f438b9d3ddcec073340360e227b00b0674c0d1', 'Xi-Yuer', '2214380963@qq.com', '2024-04-17 13:23:53', 'feat:大文件上传');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('6af37d7ac2ada927b57ad11beeffe0b7e541464b', 'Xi-Yuer', '2214380963@qq.com', '2024-04-17 04:03:22', 'git:workflow');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('7dd1f2d39890034297cd11b41fde43c9d542e94d', 'Xi-Yuer', '2214380963@qq.com', '2024-04-17 04:03:22', 'git:workflow');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('4f9f715721280ccf3642783c4a9540152859d959', 'Xi-Yuer', '2214380963@qq.com', '2024-04-17 03:32:28', 'git:workflow');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('e9e2c22294934d8476e0f96992694bfc8d553ac5', 'Xi-Yuer', '2214380963@qq.com', '2024-04-17 03:28:28', 'git:workflow');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('9c1f3479f5405951faff5777c81380864f007039', 'Xi-Yuer', '2214380963@qq.com', '2024-04-17 03:22:43', 'git:workflow');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('e72c63f4804d68a179c5b350906df4bdfc738153', 'Xi-Yuer', '2214380963@qq.com', '2024-04-17 03:22:21', 'git:workflow');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('73b5ffba0cc9d44401159433ba78bc4f960eebeb', 'Xi-Yuer', '2214380963@qq.com', '2024-04-17 03:20:07', 'git:workflow');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('a7e7e027e3b9f4fe05c8b39390e35eac204c72db', 'Xi-Yuer', '2214380963@qq.com', '2024-04-17 03:19:14', 'git:workflow');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('84476e448f07ba7a04eb48ffe2b1eb5da3bd367f', 'Xi-Yuer', '2214380963@qq.com', '2024-04-17 03:17:45', 'git:workflow');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('90668ab7236ed8231a06eae34513be0a90716ee3', 'Xi-Yuer', '2214380963@qq.com', '2024-04-17 03:16:03', 'git:workflow');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('dcb86451590215b1a11e6645512d04caa804ed36', 'Xi-Yuer', '2214380963@qq.com', '2024-04-17 03:14:46', 'git:workflow');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('5a49940523f208cd934c68010a18b1373bf35f2d', 'Xi-Yuer', '2214380963@qq.com', '2024-04-17 03:09:58', 'git:workflow');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('8a683d8654fc0535712f88122bbcb143223fc566', 'Xi-Yuer', '2214380963@qq.com', '2024-04-17 03:08:12', 'git:workflow');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('9b9cf01b41a1137ac62dbee6376ab9591d023cd5', 'Xi-Yuer', '2214380963@qq.com', '2024-04-17 03:07:17', 'git:workflow');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('b9d3da4b6690c41a7749c3460f0648fb69fbd5fe', 'Xi-Yuer', '2214380963@qq.com', '2024-04-17 02:52:02', 'git:workflow');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('f56337536aaae86a1cec7eb3665711ea7add89b7', 'Xi-Yuer', '2214380963@qq.com', '2024-04-17 02:49:12', 'git:workflow');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('29d16aa53225080a679841d3ca2030e995a591f4', 'Xi-Yuer', '2214380963@qq.com', '2024-04-17 02:44:33', 'git:workflow');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('9e0ecb39d2b782795e4b9bb90dc18a74aa71b33e', 'Xi-Yuer', '2214380963@qq.com', '2024-04-17 02:36:46', 'git:workflow');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('9e04643a01dcf261d8be41685ff487db9a8bbe01', 'Xi-Yuer', '2214380963@qq.com', '2024-04-17 02:29:50', 'git:workflow');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('98e9566f2223b8c77e16d0ed00c9d516137b29d1', 'Xi-Yuer', '2214380963@qq.com', '2024-04-17 02:23:49', 'git:workflow');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('0746a5e8c3aee0a1bac2acb05f30f78306a8c219', 'Xi-Yuer', '2214380963@qq.com', '2024-04-17 02:18:07', 'git:workflow');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('5a28e241f11de20ed456265ac7a4b800c2b96ef9', 'Xi-Yuer', '2214380963@qq.com', '2024-04-17 02:13:20', 'git:workflow');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('64021742bf37ea3a6b7d39518d29a0f536c5cc33', 'Xi-Yuer', '2214380963@qq.com', '2024-04-17 02:10:39', 'git:workflow');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('c38a4ea1b75bf368212606559f5bab8407565d6f', 'Xi-Yuer', '2214380963@qq.com', '2024-04-17 02:08:55', 'workflow');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('a53de6e27976ca80f9b68333700abad35456f84a', 'Xi-Yuer', '2214380963@qq.com', '2024-04-17 01:57:55', 'git:文件管理');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('293e69f5e02eac7535067a2ce3f5e32b60a9b74b', 'Xi-Yuer', '2214380963@qq.com', '2024-04-17 01:57:05', 'git:文件管理');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('03c885c63d3a38b2f186d9efbd83c9b9e63aa9a2', 'Xi-Yuer', '2214380963@qq.com', '2024-04-17 01:47:20', 'feat:大文件分片上传');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('ce42bf19a85bdd4026fb2c4debc4ec658de74db5', '小鱼儿', '2214380963@qq.com', '2024-04-17 01:42:31', 'Update go.yml');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('afcd6529310b8f9645038aa4d546ce9a955d6494', '小鱼儿', '2214380963@qq.com', '2024-04-17 01:25:20', 'Create go.yml');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('0b5b2058e1f1af7f884158ef026a3f940efd5482', 'Xi-Yuer', '2214380963@qq.com', '2024-04-16 10:04:30', 'feat:大文件分片上传');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('823f9b28c732ae77f7a15af5cd8d8e05cb93d025', 'Xi-Yuer', '2214380963@qq.com', '2024-04-15 16:17:37', 'feat:定时任务');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('902731fcebd5e551fb2e225643b761a58ade2d6c', 'Xi-Yuer', '2214380963@qq.com', '2024-04-15 08:59:50', 'feat:系统监控');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('62c952df18a119979cde6d3467a157fe537dc6e4', 'Xi-Yuer', '2214380963@qq.com', '2024-04-15 08:56:41', 'feat:系统监控');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('8d34d831a1263ef96e65e698bd1fc8f493454f19', 'Xi-Yuer', '2214380963@qq.com', '2024-04-15 07:19:25', 'docs:Excel数据导出新增权限配置表');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('753d441a2b147ae059b84b4f65b1c8a61beb084b', 'Xi-Yuer', '2214380963@qq.com', '2024-04-15 07:17:05', 'feat:Excel数据导出新增权限配置表');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('705b6afac4aa86a59664bed9acf9740d4d4f20f1', 'Xi-Yuer', '2214380963@qq.com', '2024-04-15 07:08:10', 'feat:Excel数据导出');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('f3f6055af835a1ad29398dffb914bb27409f3805', 'Xi-Yuer', '2214380963@qq.com', '2024-04-14 14:39:27', 'feat:初始化超级管理员用户');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('61e2ef74e52a0e432b7d0852afd90a4793fb8539', 'Xi-Yuer', '2214380963@qq.com', '2024-04-14 11:54:05', 'git:文件管理');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('78c310c81c0fc26efd10df9aaee2e807a8b0fcb0', 'Xi-Yuer', '2214380963@qq.com', '2024-04-14 11:42:38', 'git:文件管理');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('8c6eb61fefe8fbaf66db2689ff006740b474cb2b', 'Xi-Yuer', '2214380963@qq.com', '2024-04-14 11:27:02', 'feat: 移除依赖更新其他依赖版本');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('5c7c991efacafa7d96eedb8e421830242317ed0f', 'Xi-Yuer', '2214380963@qq.com', '2024-04-13 15:13:38', 'feat: 移除了`github.com/deckarep/golang-set`依赖，更新了其他依赖版本。');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('3ddc0499b7b041b2646cc892797287fc5d694165', 'Xi-Yuer', '2214380963@qq.com', '2024-04-13 15:08:16', 'feat: 添加接口权限验证中间件和路由白名单');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('db04c3aa0b42f51f295132fc3e27fd4b0b017639', 'Xi-Yuer', '2214380963@qq.com', '2024-04-13 07:57:55', 'refactor:系统日志/Git提交日志');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('4e606e11508b27b62ab384bb75d8762fe21b3dcc', 'Xi-Yuer', '2214380963@qq.com', '2024-04-13 07:56:47', 'refactor:系统日志/Git提交日志');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('2b68340f161e2cfe211c938ad0a9345fc62de438', 'Xi-Yuer', '2214380963@qq.com', '2024-04-13 07:51:15', 'feat:添加提交记录功能');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('fbcb9eaf63dd7ca833744c8f751b0c70f7114a1c', 'Xi-Yuer', '2214380963@qq.com', '2024-04-13 07:42:50', 'feat:添加提交记录功能');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('4adc5968884a6a7a265f01518b6b283d81fc325d', 'Xi-Yuer', '2214380963@qq.com', '2024-04-13 06:56:23', '\"修复了constant.go中的变量命名，使其符合驼峰命名规则。\"');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('3e8779cdf015b861836254b1e744a144b6666392', 'Xi-Yuer', '2214380963@qq.com', '2024-04-13 06:46:30', 'feat: 添加页面外链');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('774473306b0c542e5056a8ef0322eeaee10f7652', 'Xi-Yuer', '2214380963@qq.com', '2024-04-12 09:03:08', 'feat:资源不存在提示修改');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('c2b6f21a4a11e7e85584caa183fac7b8ca277289', 'Xi-Yuer', '2214380963@qq.com', '2024-04-12 08:53:40', 'feat:接口及权限管理');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('c2ed5e2c5e1228377903399c6382ccb53bb30b26', 'Xi-Yuer', '2214380963@qq.com', '2024-04-12 02:58:03', 'feat:用户Token使用Account');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('80a49ab8d6595c14330d36662b4a466626d256ff', 'Xi-Yuer', '2214380963@qq.com', '2024-04-12 06:21:14', 'feat:用户Token使用Account');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('b177c9b134672100fa4c1121ed364149156a4c34', 'Xi-Yuer', '2214380963@qq.com', '2024-04-12 02:58:03', 'feat:日志管理');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('c28e1490d8ebe067a03072cebf51b33496f57e5e', 'Xi-Yuer', '2214380963@qq.com', '2024-04-12 02:58:03', 'feat:日志管理');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('2358cd3b3236429a451cf281fb9ac96400c78a9c', 'Xi-Yuer', '2214380963@qq.com', '2024-04-11 09:51:48', 'sql:区域数据库');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('d3119236d6c588d10407a2f0beb28957623725fb', 'Xi-Yuer', '2214380963@qq.com', '2024-04-10 15:48:14', 'sql:接口管理');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('50bd384d70d761dee15a493f5876198ff14c8b88', 'Xi-Yuer', '2214380963@qq.com', '2024-04-10 15:46:35', 'feat:接口管理');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('d4b7385eb779eecfb60473eb9167434b0be63109', 'Xi-Yuer', '2214380963@qq.com', '2024-04-10 14:08:49', 'feat:通过角色ID查询用户');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('4553873b793a5518ecfb5b20e7ffa85854b72001', 'Xi-Yuer', '2214380963@qq.com', '2024-04-07 13:38:54', 'feat:角色用户管理');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('1dbde5d7901bdffcad356c8e3941dd7a9f8670cd', 'Xi-Yuer', '2214380963@qq.com', '2024-04-07 11:22:34', 'feat:更新角色');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('0a84ed8dac0000c2a35e436d924212ddc3a545bb', 'Xi-Yuer', '2214380963@qq.com', '2024-04-07 06:52:49', 'feat:更新菜单');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('e4f0b2a71f9330bc2333391890d305cab4252293', 'Xi-Yuer', '2214380963@qq.com', '2024-04-07 06:32:30', 'docs:swagger');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('e2bad77880c2e1d69764bdd8d4dda348dc6e8cc9', 'Xi-Yuer', '2214380963@qq.com', '2024-04-07 06:32:02', 'refactor:创建用户和角色的时候直接分配权限信息');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('48b08086c672b38d626ce6e92a72a7b4a94fed19', 'Xi-Yuer', '2214380963@qq.com', '2024-04-07 06:04:37', 'refactor:获取用户角色信息');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('a445f01de41e750a0f6b965151cf939f6fbaaca9', 'Xi-Yuer', '2214380963@qq.com', '2024-04-07 04:26:52', 'feat:禁用账号不允许登录');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('1a8b79b084fd510f25b408cbff795026773f4cf6', 'Xi-Yuer', '2214380963@qq.com', '2024-04-07 04:21:24', 'refactor:路由模块重构');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('172c54da1575259fd6762aed231d89f65e0ae809', 'Xi-Yuer', '2214380963@qq.com', '2024-04-06 14:10:24', 'feat:删除用户特殊处理');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('557c7860442b517e17d5f256b31ff2dcecc57326', 'Xi-Yuer', '2214380963@qq.com', '2024-04-06 14:02:42', 'docs:查询用户信息');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('f68ca1119d1d49bffc9c7ad0e42f43708350ac8d', 'Xi-Yuer', '2214380963@qq.com', '2024-04-06 13:27:26', 'feat:用户权限信息');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('fcfc1b2856589869c3b17d0c4f9b29648d065072', 'Xi-Yuer', '2214380963@qq.com', '2024-04-06 11:47:46', 'docs:部门管理');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('1fa7337bc7b54615f1f0ad64af7f2b50d58b13f3', 'Xi-Yuer', '2214380963@qq.com', '2024-04-06 11:46:19', 'feat:部门管理');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('db90b34125c48284a096d6c44145fb64b9715f71', 'Xi-Yuer', '2214380963@qq.com', '2024-04-06 08:23:23', 'feat:部门管理');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('4e15df8ffae12865309a0d1cdb5bd2df014472bb', 'Xi-Yuer', '2214380963@qq.com', '2024-04-06 07:03:56', 'feat:获取菜单');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('c697ba33bec4fa8aa98a6317f20d1632d58d98c3', 'Xi-Yuer', '2214380963@qq.com', '2024-04-05 17:57:57', 'docs:swagger');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('f6622c48e6237fa84a33c2387fcff987577d1ae4', 'Xi-Yuer', '2214380963@qq.com', '2024-04-05 17:45:20', 'feat:构建用户菜单');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('b9c93dd6841eaf058c4628421301cf26cf4694d2', 'Xi-Yuer', '2214380963@qq.com', '2024-04-04 17:55:32', 'docs:swagger');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('7eba3f7987195934ffc991a777f8f477c5b3fb6c', 'Xi-Yuer', '2214380963@qq.com', '2024-04-04 17:42:59', 'docs:swagger');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('61ead725b6e316da8ec60555d580b663a20a9a3f', 'Xi-Yuer', '2214380963@qq.com', '2024-04-04 17:40:06', 'feat:获取用户菜单');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('4a29d1fb5f943404fac0e192bc246ae78e36c774', 'Xi-Yuer', '2214380963@qq.com', '2024-04-04 16:43:25', 'feat:给角色分配页面权限');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('064c698b49f40e32b2177d1b15db08acee6de1ee', 'Xi-Yuer', '2214380963@qq.com', '2024-04-03 18:39:35', 'refactor:dto');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('39d26bbe25e3594bf771693f10cc841beba48f2a', 'Xi-Yuer', '2214380963@qq.com', '2024-04-03 18:36:56', 'feat:角色权限');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('e7c603231440d541b11a56d9ef65ada273da9621', 'Xi-Yuer', '2214380963@qq.com', '2024-04-03 18:02:50', 'feat:页面管理（新建删除页面）');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('219714b514701ea4bdd5af16dc64b1965dc92b8e', 'Xi-Yuer', '2214380963@qq.com', '2024-03-31 13:59:39', 'refactor:swagger');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('35960289a1d4d7325bf36fcda543f701076674d7', 'Xi-Yuer', '2214380963@qq.com', '2024-03-31 13:58:23', 'refactor:swagger');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('ee592d852fde27310469eba879255d41865339d2', 'Xi-Yuer', '2214380963@qq.com', '2024-03-31 13:57:48', 'refactor:sql建表语句');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('41a7a946ddba6ea097dd15cd20e0513e0efabd4a', 'Xi-Yuer', '2214380963@qq.com', '2024-03-30 13:06:00', 'feat:页面管理');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('18fd36b7f64bd14cc2b999a7d27267fe94259e3b', 'Xi-Yuer', '2214380963@qq.com', '2024-03-30 13:06:00', 'refactor:用户角色权限');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('703949bcbf35daa92a75d5c155035a069aa51504', 'Xi-Yuer', '2214380963@qq.com', '2024-03-30 13:06:00', 'refactor:用户角色权限');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('2c25f67b85375780b4c2467068fd9ec1f009e14e', 'Xi-Yuer', '2214380963@qq.com', '2024-03-30 13:06:00', 'refactor:sql文件');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('c1e356bc285b671b323164e860c57fdff75bdeec', 'Xi-Yuer', '2214380963@qq.com', '2024-03-30 13:06:00', 'refactor:sql文件');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('dc70cd42cc9e02fc9b9588a1f91ad6e0df3356e4', 'Xi-Yuer', '2214380963@qq.com', '2024-03-29 17:25:57', 'refactor:代码组织结构以查询逻辑修改');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('06f9803626ebf0b6a4d2704807ed23e823280ec1', 'Xi-Yuer', '2214380963@qq.com', '2024-03-29 17:25:57', 'feat:角色模块');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('1b65d96dbc6a1e23eac6af238194b8d8e97b0735', 'Xi-Yuer', '2214380963@qq.com', '2024-03-29 16:50:33', 'feat:角色模块');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('0995d403ba639ecd1e6382f69dfdf6aeee68c0e1', 'Xi-Yuer', '2214380963@qq.com', '2024-03-27 16:03:37', 'feat:角色模块');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('f301454d72ace3fe7b2794eb2c762bf3e907ef9e', 'Xi-Yuer', '2214380963@qq.com', '2024-03-27 10:03:41', 'feat:验证码/jwt权限认证');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('faf140371c4311b78cdb9d4b21df5a33029901ed', 'Xi-Yuer', '2214380963@qq.com', '2024-03-27 06:35:13', 'feat:验证码/jwt权限认证');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('7889004403821ab78ae3cad5b46ad91d9f921486', 'Xi-Yuer', '2214380963@qq.com', '2024-03-23 17:41:44', 'feat:用户模块-新增、删除，更新用户');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('39f730cdd306595f5a6ff6fcd2ced862155f9e70', 'Xi-Yuer', '2214380963@qq.com', '2024-03-15 07:57:24', 'feat:用户模块-查询用户');
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('bffd62d5731e3f25a32a9b5f6b6bd63c83969b64', 'Xi-Yuer', '2214380963@qq.com', '2024-03-04 09:44:50', 'feat:后端基础框架搭建');
COMMIT;

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
  `description` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `department_pk_2` (`id`),
  UNIQUE KEY `department_pk_3` (`department_name`),
  KEY `department_department_id_fk` (`parent_department`),
  CONSTRAINT `department_department_id_fk` FOREIGN KEY (`parent_department`) REFERENCES `department` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of department
-- ----------------------------
BEGIN;
INSERT INTO `department` (`id`, `department_name`, `parent_department`, `create_time`, `update_time`, `delete_time`, `department_order`, `description`) VALUES ('47230539865788416', '旭升科技有限责任公司', NULL, '2024-05-10 15:57:18', '2024-05-10 21:29:05', NULL, 0, '旭升科技有限责任公司');
INSERT INTO `department` (`id`, `department_name`, `parent_department`, `create_time`, `update_time`, `delete_time`, `department_order`, `description`) VALUES ('47230730614345728', '旭升科技有限责任公司（成都分公司）', NULL, '2024-05-10 15:58:03', '2024-05-10 21:28:57', NULL, 1, '旭升科技有限责任公司（成都分公司）');
INSERT INTO `department` (`id`, `department_name`, `parent_department`, `create_time`, `update_time`, `delete_time`, `department_order`, `description`) VALUES ('47230853654253568', '软件开发部', '47230539865788416', '2024-05-10 15:58:33', '2024-05-10 15:58:33', NULL, 0, '软件开发部');
INSERT INTO `department` (`id`, `department_name`, `parent_department`, `create_time`, `update_time`, `delete_time`, `department_order`, `description`) VALUES ('47230946730053632', '前端部门', '47230853654253568', '2024-05-10 15:58:55', '2024-05-10 17:10:40', NULL, 0, '负责软件界面开发');
INSERT INTO `department` (`id`, `department_name`, `parent_department`, `create_time`, `update_time`, `delete_time`, `department_order`, `description`) VALUES ('47231640862199808', '市场营销部', '47230539865788416', '2024-05-10 16:01:40', '2024-05-10 16:01:40', NULL, 1, '市场营销');
INSERT INTO `department` (`id`, `department_name`, `parent_department`, `create_time`, `update_time`, `delete_time`, `department_order`, `description`) VALUES ('47231703478964224', '人事部', '47230539865788416', '2024-05-10 16:01:55', '2024-05-10 16:02:00', NULL, 2, '人事部');
INSERT INTO `department` (`id`, `department_name`, `parent_department`, `create_time`, `update_time`, `delete_time`, `department_order`, `description`) VALUES ('47231806839197696', '采购部', '47230539865788416', '2024-05-10 16:02:20', '2024-05-10 16:02:20', NULL, 3, '采购');
INSERT INTO `department` (`id`, `department_name`, `parent_department`, `create_time`, `update_time`, `delete_time`, `department_order`, `description`) VALUES ('47231951035174912', '国际贸易部', '47230730614345728', '2024-05-10 16:02:54', '2024-05-10 16:02:54', NULL, 0, '国际贸易部');
INSERT INTO `department` (`id`, `department_name`, `parent_department`, `create_time`, `update_time`, `delete_time`, `department_order`, `description`) VALUES ('47248253208498176', 'Java开发部', '47230853654253568', '2024-05-10 17:07:41', '2024-05-10 17:07:41', NULL, 1, 'Java开发部');
COMMIT;

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
-- Records of file
-- ----------------------------
BEGIN;
COMMIT;

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
  `can_edit` tinyint(1) DEFAULT '1',
  `create_time` timestamp NULL DEFAULT (now()),
  `update_time` timestamp NULL DEFAULT (now()) ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`interface_id`),
  UNIQUE KEY `interface_pk` (`interface_id`),
  KEY `interface_pages_page_id_fk` (`interface_page_id`),
  CONSTRAINT `interface_pages_page_id_fk` FOREIGN KEY (`interface_page_id`) REFERENCES `pages` (`page_id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='接口/资源表';

-- ----------------------------
-- Records of interfaces
-- ----------------------------
BEGIN;
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `can_edit`, `create_time`, `update_time`) VALUES ('37899586299236352', '获取菜单', 'GET', '/pages', '40708567220621312', 'GET:/pages', '获取菜单', 0, '2024-04-14 21:59:25', '2024-05-11 16:04:13');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `can_edit`, `create_time`, `update_time`) VALUES ('37899744348999680', '获取用户菜单', 'GET', '/pages/menus', '40708567220621312', 'GET:/pages/menus', '获取用户菜单', 0, '2024-04-14 22:00:03', '2024-05-11 16:04:13');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `can_edit`, `create_time`, `update_time`) VALUES ('40696216958275584', '新建菜单', 'POST', '/pages', '40708567220621312', 'POST:/pages', '新建菜单', 0, '2024-04-22 15:12:14', '2024-05-11 16:04:13');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `can_edit`, `create_time`, `update_time`) VALUES ('41065294818447360', '创建用户', 'POST', '/users', '40351711566499840', 'POST:/users', '创建用户', 0, '2024-04-23 15:38:49', '2024-05-11 16:04:13');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `can_edit`, `create_time`, `update_time`) VALUES ('41071752721207296', '获取系统运行信息', 'GET', '/system', '40350365404631040', 'GET:/system', '获取系统运行信息', 0, '2024-04-23 16:04:28', '2024-05-11 16:04:13');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `can_edit`, `create_time`, `update_time`) VALUES ('41430491085148160', '获取Git提交日志', 'GET', '/log/commits', '40350365404631040', 'GET:/log/commits', '获取Git提交日志', 0, '2024-04-24 15:49:59', '2024-05-11 16:04:13');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `can_edit`, `create_time`, `update_time`) VALUES ('41438779612860416', '获取Git提交次数', 'GET', '/log/commits/count', '40350365404631040', 'GET:/log/commits/count', '获取Git提交次数', 0, '2024-04-24 16:22:54', '2024-05-11 16:04:13');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `can_edit`, `create_time`, `update_time`) VALUES ('41789444067430400', '获取用户', 'POST', '/users/query', '40351711566499840', 'POST:/users/query', '获取用户', 0, '2024-04-25 17:20:52', '2024-05-11 16:04:13');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `can_edit`, `create_time`, `update_time`) VALUES ('41903454221766656', '获取角色', 'POST', '/roles/query', '40352343601975296', 'POST:/roles/query', '获取角色', 0, '2024-04-25 23:09:21', '2024-05-11 16:04:13');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `can_edit`, `create_time`, `update_time`) VALUES ('41906647861301248', '获取部门', 'GET', '/department', '40352749044371456', 'GET:/department', '获取部门', 0, '2024-04-25 23:22:03', '2024-05-11 16:04:13');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `can_edit`, `create_time`, `update_time`) VALUES ('42063400649363456', '更新用户', 'PATCH', '/users/:id', '40351711566499840', 'PATCH:/users/:id', '更新用户', 0, '2024-04-26 09:44:56', '2024-05-11 16:04:13');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `can_edit`, `create_time`, `update_time`) VALUES ('42071947340681216', '删除用户', 'DELETE', '/users/:id', '40351711566499840', 'DELETE:/users/:id', '删除用户', 0, '2024-04-26 10:18:53', '2024-05-11 16:04:13');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `can_edit`, `create_time`, `update_time`) VALUES ('42079060825739264', '获取用户详情', 'GET', '/users/:id', '40351711566499840', 'GET:/users/:id', '获取用户详情', 0, '2024-04-26 10:47:09', '2024-05-11 17:24:40');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `can_edit`, `create_time`, `update_time`) VALUES ('44326032676753408', '删除角色', 'DELETE', '/roles/:id', '40352343601975296', 'DELETE:/roles/:id', '获取Git提交次数', 0, '2024-05-02 15:35:49', '2024-05-11 16:04:13');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `can_edit`, `create_time`, `update_time`) VALUES ('44332995393359872', '导出角色', 'POST', '/roles/export', '40352343601975296', 'POST:/roles/export', '导出角色', 0, '2024-05-02 16:03:29', '2024-05-11 16:04:13');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `can_edit`, `create_time`, `update_time`) VALUES ('44333062707744768', '导出用户', 'POST', '/users/export', '40351711566499840', 'POST:/users/export', '导出用户', 0, '2024-05-02 16:03:45', '2024-05-11 16:04:13');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `can_edit`, `create_time`, `update_time`) VALUES ('44474438040686592', '更新角色', 'PATCH', '/roles/:id', '40352343601975296', 'PATCH:/roles/:id', '更新角色', 0, '2024-05-03 01:25:32', '2024-05-11 16:04:13');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `can_edit`, `create_time`, `update_time`) VALUES ('45864670036234240', '新增角色', 'POST', '/roles', '40352343601975296', 'POST:/roles', '新增角色', 0, '2024-05-06 21:29:49', '2024-05-11 16:04:13');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `can_edit`, `create_time`, `update_time`) VALUES ('46239031234662400', '通过角色ID查询用户', 'GET', '/users/role/:id', '40352343601975296', 'GET:/users/role/:id', '通过角色ID查询用户', 0, '2024-05-07 22:17:24', '2024-05-11 16:04:13');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `can_edit`, `create_time`, `update_time`) VALUES ('46858901232029696', '绑定用户', 'POST', '/roles/bindUser', '40352343601975296', 'POST:/roles/bindUser', '绑定用户', 0, '2024-05-09 15:20:32', '2024-05-11 16:04:13');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `can_edit`, `create_time`, `update_time`) VALUES ('46859015694585856', '解绑用户', 'POST', '/roles/deBindUser', '40352343601975296', 'POST:/roles/deBindUser', '解绑用户', 0, '2024-05-09 15:20:59', '2024-05-11 16:04:13');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `can_edit`, `create_time`, `update_time`) VALUES ('46859424727306240', '查询角色之外的用户', 'POST', '/users/query/role/:id', '40352343601975296', 'POST:/users/query/role/:id', '查询角色之外的用户', 0, '2024-05-09 15:22:37', '2024-05-11 16:04:13');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `can_edit`, `create_time`, `update_time`) VALUES ('46862379094380544', '创建菜单', 'POST', '/pages', '40708567220621312', 'POST:/pages', '创建菜单', 0, '2024-05-09 15:34:21', '2024-05-11 16:04:13');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `can_edit`, `create_time`, `update_time`) VALUES ('46862511944765440', '删除菜单', 'DELETE', '/pages/:id', '40708567220621312', 'DELETE:/pages/:id', '创建菜单', 0, '2024-05-09 15:34:53', '2024-05-11 16:04:13');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `can_edit`, `create_time`, `update_time`) VALUES ('46862830229524480', '获取菜单（All）', 'GET', '/pages', '40708567220621312', 'GET:/pages', '获取菜单（All）', 0, '2024-05-09 15:36:09', '2024-05-11 16:04:13');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `can_edit`, `create_time`, `update_time`) VALUES ('46863099453509632', '获取菜单（User）', 'GET', '/pages/user', '40708567220621312', 'GET:/pages/user', '获取菜单（User）', 0, '2024-05-09 15:37:13', '2024-05-11 16:04:13');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `can_edit`, `create_time`, `update_time`) VALUES ('46863346367991808', '更新菜单', 'PATCH', '/pages/:id', '40708567220621312', 'PATCH:/pages/:id', 'PATCH', 0, '2024-05-09 15:38:12', '2024-05-11 16:04:13');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `can_edit`, `create_time`, `update_time`) VALUES ('46863589419520000', '获取菜单（Role）', 'GET', '/pages/role/:id', '40708567220621312', 'GET:/pages/role/:id', 'GET', 0, '2024-05-09 15:39:10', '2024-05-11 16:04:13');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `can_edit`, `create_time`, `update_time`) VALUES ('47156442767036416', '创建部门', 'POST', '/department', '40352749044371456', 'POST:/department', '创建部门', 0, '2024-05-10 11:02:52', '2024-05-11 16:04:13');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `can_edit`, `create_time`, `update_time`) VALUES ('47216551878725632', '删除部门', 'DELETE', '/department/:id', '40352749044371456', 'DELETE:/department/:id', '删除部门', 0, '2024-05-10 15:01:43', '2024-05-11 16:04:13');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `can_edit`, `create_time`, `update_time`) VALUES ('47216745617821696', '更新部门', 'PATCH', '/department/:id', '40352749044371456', 'PATCH:/department/:id', '更新部门', 0, '2024-05-10 15:02:29', '2024-05-11 16:04:13');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `can_edit`, `create_time`, `update_time`) VALUES ('47584267768696832', '获取接口（Page）', 'GET', '/interface/page/:id', '40708567220621312', 'GET:/interface/page/:id', '获取接口（Page）', 0, '2024-05-11 15:22:53', '2024-05-11 16:04:13');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `can_edit`, `create_time`, `update_time`) VALUES ('47593342061514752', '删除接口', 'DELETE', '/interface/:id', '40708567220621312', 'DELETE:/interface/:id', '删除接口', 0, '2024-05-11 15:58:57', '2024-05-11 16:04:13');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `can_edit`, `create_time`, `update_time`) VALUES ('47608622313639936', '新增接口', 'POST', '/interface', '40708567220621312', 'POST:/interface', '新增接口', 0, '2024-05-11 16:59:40', '2024-05-11 17:00:53');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `can_edit`, `create_time`, `update_time`) VALUES ('47608882222075904', '更新接口', 'PATCH', '/interface/:id', '40708567220621312', 'PATCH:/interface/:id', '更新接口', 0, '2024-05-11 17:00:42', '2024-05-11 17:06:39');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `can_edit`, `create_time`, `update_time`) VALUES ('47973984368594944', '获取文件', 'GET', '/upload', '47973248603787264', 'GET:/upload', '获取文件列表', 0, '2024-05-12 17:11:29', '2024-05-12 17:21:04');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `can_edit`, `create_time`, `update_time`) VALUES ('47974191147782144', '删除文件', 'DELETE', '/upload/del/:id', '47973248603787264', 'DELETE:/upload/del/:id', '删除文件', 0, '2024-05-12 17:12:18', '2024-05-12 17:21:04');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `can_edit`, `create_time`, `update_time`) VALUES ('47974424665657344', '查看文件状态', 'POST', '/upload/check', '47973248603787264', 'POST:/upload/check', '上传文件是需要检查文件是否上传以及上传了多少', 0, '2024-05-12 17:13:14', '2024-05-12 17:21:04');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `can_edit`, `create_time`, `update_time`) VALUES ('47974555385335808', '上传文件', 'POST', '/upload', '47973248603787264', 'POST:/upload', '上传文件', 0, '2024-05-12 17:13:45', '2024-05-12 17:21:04');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `can_edit`, `create_time`, `update_time`) VALUES ('47974732070391808', '完成上传', 'POST', '/upload/finish', '47973248603787264', 'POST:/upload/finish', '告诉服务器文件已全部上传完毕', 0, '2024-05-12 17:14:27', '2024-05-12 17:21:04');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `can_edit`, `create_time`, `update_time`) VALUES ('47974990527598592', '文件下载（POST）', 'POST', '/upload/download/:id', '47973248603787264', 'POST:/upload/download/:id', '通过Ajax下载文件', 0, '2024-05-12 17:15:29', '2024-05-12 17:21:04');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `can_edit`, `create_time`, `update_time`) VALUES ('47975191568977920', '文件下载（GET）', 'GET', '/upload/aHref/download/:id', '47973248603787264', 'GET:/upload/aHref/download/:id', '通过 a 标签下载文件', 0, '2024-05-12 17:16:17', '2024-05-12 17:21:04');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `can_edit`, `create_time`, `update_time`) VALUES ('47975597174951936', '获取Cookie', 'GET', '/auth/cookie', '47973248603787264', 'GET:/auth/cookie', 'a标签下载文件需要先获取Cookie,该Ckooiie只能使用一次就会过期', 0, '2024-05-12 17:17:53', '2024-05-12 17:21:04');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `can_edit`, `create_time`, `update_time`) VALUES ('47976243613667328', '获取系统日志', 'GET', '/log/system', '47975943540576256', 'GET:/log/system', '获取系统日志', 0, '2024-05-12 17:20:27', '2024-05-12 17:21:04');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `can_edit`, `create_time`, `update_time`) VALUES ('49391522268844032', '获取定时任务列表', 'GET', '/timeTask', '49387228375289856', 'GET:/timeTask', '获取定时任务列表', 1, '2024-05-16 15:04:16', '2024-05-16 15:08:02');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `can_edit`, `create_time`, `update_time`) VALUES ('49391714045005824', '开始定时任务', 'POST', '/timeTask/start/:id', '49387228375289856', 'POST:/timeTask/start/:id', '开始定时任务', 1, '2024-05-16 15:05:02', '2024-05-16 15:05:02');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `can_edit`, `create_time`, `update_time`) VALUES ('49391851559456768', '停止定时任务', 'POST', '/timeTask/stop/:id', '49387228375289856', 'POST:/timeTask/stop/:id', '停止定时任务', 1, '2024-05-16 15:05:35', '2024-05-16 15:05:35');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `can_edit`, `create_time`, `update_time`) VALUES ('49392003502313472', '更新定时任务', 'PATCH', '/timeTask/update/:id', '49387228375289856', 'PATCH:/timeTask/update/:id', '更新定时任务', 1, '2024-05-16 15:06:11', '2024-05-16 15:06:11');
COMMIT;

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
-- Records of logs
-- ----------------------------
BEGIN;
COMMIT;

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
  `is_out_site` tinyint(1) DEFAULT '0' COMMENT '是否外链',
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
-- Records of pages
-- ----------------------------
BEGIN;
INSERT INTO `pages` (`page_id`, `parent_page`, `page_name`, `page_path`, `page_icon`, `page_component`, `page_order`, `can_edit`, `is_out_site`, `out_site_link`, `create_time`, `update_time`, `delete_time`) VALUES ('40350365404631040', NULL, '仪表盘', '/dashboard', 'DashboardOutlined', 'dashboard', 1, 0, 0, '', '2024-04-21 17:41:43', '2024-05-22 16:24:47', NULL);
INSERT INTO `pages` (`page_id`, `parent_page`, `page_name`, `page_path`, `page_icon`, `page_component`, `page_order`, `can_edit`, `is_out_site`, `out_site_link`, `create_time`, `update_time`, `delete_time`) VALUES ('40351462613585920', NULL, '系统管理', '/system', 'SettingOutlined', 'system', 2, 0, 0, '', '2024-04-21 17:46:05', '2024-05-10 22:11:01', NULL);
INSERT INTO `pages` (`page_id`, `parent_page`, `page_name`, `page_path`, `page_icon`, `page_component`, `page_order`, `can_edit`, `is_out_site`, `out_site_link`, `create_time`, `update_time`, `delete_time`) VALUES ('40351711566499840', '40351462613585920', '用户管理', '/system/user', 'UserOutlined', 'user', 1, 0, 0, '', '2024-04-21 17:47:04', '2024-05-10 22:11:01', NULL);
INSERT INTO `pages` (`page_id`, `parent_page`, `page_name`, `page_path`, `page_icon`, `page_component`, `page_order`, `can_edit`, `is_out_site`, `out_site_link`, `create_time`, `update_time`, `delete_time`) VALUES ('40352343601975296', '40351462613585920', '角色管理', '/system/role', 'MergeOutlined', 'role', 2, 0, 0, '', '2024-04-21 17:49:35', '2024-05-10 22:11:01', NULL);
INSERT INTO `pages` (`page_id`, `parent_page`, `page_name`, `page_path`, `page_icon`, `page_component`, `page_order`, `can_edit`, `is_out_site`, `out_site_link`, `create_time`, `update_time`, `delete_time`) VALUES ('40352749044371456', '40351462613585920', '部门管理', '/system/department', 'ApartmentOutlined', 'department', 3, 0, 0, '', '2024-04-21 17:51:12', '2024-05-10 22:11:01', NULL);
INSERT INTO `pages` (`page_id`, `parent_page`, `page_name`, `page_path`, `page_icon`, `page_component`, `page_order`, `can_edit`, `is_out_site`, `out_site_link`, `create_time`, `update_time`, `delete_time`) VALUES ('40708567220621312', '40351462613585920', '菜单管理', '/system/menu', 'FolderOpenOutlined', 'menu', 4, 0, 0, '', '2024-04-22 16:01:18', '2024-05-10 22:11:01', NULL);
INSERT INTO `pages` (`page_id`, `parent_page`, `page_name`, `page_path`, `page_icon`, `page_component`, `page_order`, `can_edit`, `is_out_site`, `out_site_link`, `create_time`, `update_time`, `delete_time`) VALUES ('47513502385967104', NULL, '聊一聊', '', 'CommentOutlined', '', 6, 0, 1, 'https://xiyuer.club', '2024-05-11 10:41:41', '2024-05-17 17:33:54', NULL);
INSERT INTO `pages` (`page_id`, `parent_page`, `page_name`, `page_path`, `page_icon`, `page_component`, `page_order`, `can_edit`, `is_out_site`, `out_site_link`, `create_time`, `update_time`, `delete_time`) VALUES ('47973248603787264', NULL, '文件管理', '/file', 'CloudUploadOutlined', 'File', 3, 0, 0, NULL, '2024-05-12 17:08:33', '2024-05-12 17:21:17', NULL);
INSERT INTO `pages` (`page_id`, `parent_page`, `page_name`, `page_path`, `page_icon`, `page_component`, `page_order`, `can_edit`, `is_out_site`, `out_site_link`, `create_time`, `update_time`, `delete_time`) VALUES ('47975943540576256', NULL, '系统监控', '/monitor', 'CodeOutlined', 'Monitor', 4, 0, 0, NULL, '2024-05-12 17:19:16', '2024-05-13 09:26:00', NULL);
INSERT INTO `pages` (`page_id`, `parent_page`, `page_name`, `page_path`, `page_icon`, `page_component`, `page_order`, `can_edit`, `is_out_site`, `out_site_link`, `create_time`, `update_time`, `delete_time`) VALUES ('47989353607073792', '47975943540576256', '操作日志', '/monitor/logs', 'FileDoneOutlined', 'Logs', 1, 0, 0, NULL, '2024-05-13 09:20:50', '2024-05-17 17:40:20', NULL);
INSERT INTO `pages` (`page_id`, `parent_page`, `page_name`, `page_path`, `page_icon`, `page_component`, `page_order`, `can_edit`, `is_out_site`, `out_site_link`, `create_time`, `update_time`, `delete_time`) VALUES ('47989749788446720', '47973248603787264', '大文件上传', '/file/upload', 'CloudUploadOutlined', 'UploadFile', 0, 0, 0, NULL, '2024-05-13 09:22:25', '2024-05-17 17:40:20', NULL);
INSERT INTO `pages` (`page_id`, `parent_page`, `page_name`, `page_path`, `page_icon`, `page_component`, `page_order`, `can_edit`, `is_out_site`, `out_site_link`, `create_time`, `update_time`, `delete_time`) VALUES ('49387228375289856', '47975943540576256', '定时任务', '/monitor/timeTask', 'ClockCircleOutlined', 'TimeTask', 0, 0, 0, NULL, '2024-05-16 14:47:12', '2024-05-17 17:40:20', NULL);
INSERT INTO `pages` (`page_id`, `parent_page`, `page_name`, `page_path`, `page_icon`, `page_component`, `page_order`, `can_edit`, `is_out_site`, `out_site_link`, `create_time`, `update_time`, `delete_time`) VALUES ('50765767066521600', NULL, '系统工具', '/utils', 'ExperimentOutlined', 'Utils', 5, 0, 0, NULL, '2024-05-20 10:05:02', '2024-05-20 10:33:32', NULL);
INSERT INTO `pages` (`page_id`, `parent_page`, `page_name`, `page_path`, `page_icon`, `page_component`, `page_order`, `can_edit`, `is_out_site`, `out_site_link`, `create_time`, `update_time`, `delete_time`) VALUES ('50766109451751424', '50765767066521600', '代码生成器', '/utils/codeGenerator', 'CoffeeOutlined', 'CodeGenerator', 0, 0, 0, NULL, '2024-05-20 10:06:23', '2024-05-20 10:33:32', NULL);
COMMIT;

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
  `can_edit` int DEFAULT '1' COMMENT '是否可编辑删除',
  PRIMARY KEY (`role_id`),
  UNIQUE KEY `roles_pk` (`role_name`),
  UNIQUE KEY `roles_pk_2` (`role_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of roles
-- ----------------------------
BEGIN;
INSERT INTO `roles` (`role_id`, `role_name`, `description`, `create_time`, `update_time`, `delete_time`, `can_edit`) VALUES ('37904208560656384', '超级管理员', '超级管理员，拥有所有权限，不可编辑和删除', '2024-04-14 22:17:47', '2024-05-09 15:29:08', NULL, 1);
INSERT INTO `roles` (`role_id`, `role_name`, `description`, `create_time`, `update_time`, `delete_time`, `can_edit`) VALUES ('45864950110883840', '测试', '测试', '2024-05-06 21:30:56', '2024-05-07 21:25:35', NULL, 1);
INSERT INTO `roles` (`role_id`, `role_name`, `description`, `create_time`, `update_time`, `delete_time`, `can_edit`) VALUES ('46164130582761472', '前端开发', '前端开发', '2024-05-07 17:19:46', '2024-05-07 17:19:46', NULL, 1);
INSERT INTO `roles` (`role_id`, `role_name`, `description`, `create_time`, `update_time`, `delete_time`, `can_edit`) VALUES ('46170564095643648', '后端开发', '后端开发', '2024-05-07 17:45:20', '2024-05-07 17:45:20', NULL, 1);
INSERT INTO `roles` (`role_id`, `role_name`, `description`, `create_time`, `update_time`, `delete_time`, `can_edit`) VALUES ('47615060066963456', 'test', 'test', '2024-05-11 17:25:15', '2024-05-11 17:25:15', NULL, 1);
COMMIT;

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
-- Records of roles_interfaces
-- ----------------------------
BEGIN;
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('46170564095643648', '41071752721207296');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('46170564095643648', '41430491085148160');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('46170564095643648', '41438779612860416');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('46170564095643648', '41065294818447360');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('46170564095643648', '41789444067430400');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('46170564095643648', '42063400649363456');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('46170564095643648', '42071947340681216');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('46170564095643648', '42079060825739264');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('46170564095643648', '44333062707744768');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('46170564095643648', '37899586299236352');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('46170564095643648', '37899744348999680');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('46170564095643648', '40696216958275584');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('46170564095643648', '41903454221766656');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('46170564095643648', '44326032676753408');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('46170564095643648', '44332995393359872');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('46170564095643648', '44474438040686592');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('46170564095643648', '45864670036234240');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('46170564095643648', '41906647861301248');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('45864950110883840', '37899586299236352');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('45864950110883840', '37899744348999680');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('45864950110883840', '40696216958275584');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('45864950110883840', '41065294818447360');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('45864950110883840', '41071752721207296');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('45864950110883840', '41430491085148160');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('45864950110883840', '41438779612860416');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('45864950110883840', '41789444067430400');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('45864950110883840', '41903454221766656');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('45864950110883840', '41906647861301248');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('45864950110883840', '42063400649363456');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('45864950110883840', '42071947340681216');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('45864950110883840', '42079060825739264');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('45864950110883840', '44326032676753408');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('45864950110883840', '44332995393359872');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('45864950110883840', '44333062707744768');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('45864950110883840', '44474438040686592');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('45864950110883840', '45864670036234240');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('45864950110883840', '46862379094380544');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('45864950110883840', '46862511944765440');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('45864950110883840', '46862830229524480');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('45864950110883840', '46863099453509632');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('45864950110883840', '46863346367991808');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('45864950110883840', '46863589419520000');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('45864950110883840', '46239031234662400');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('45864950110883840', '46858901232029696');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('45864950110883840', '46859015694585856');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('45864950110883840', '46859424727306240');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('46164130582761472', '37899586299236352');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('46164130582761472', '37899744348999680');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('46164130582761472', '40696216958275584');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('46164130582761472', '41065294818447360');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('46164130582761472', '41071752721207296');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('46164130582761472', '41430491085148160');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('46164130582761472', '41438779612860416');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('46164130582761472', '41789444067430400');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('46164130582761472', '41903454221766656');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('46164130582761472', '41906647861301248');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('46164130582761472', '42063400649363456');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('46164130582761472', '42071947340681216');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('46164130582761472', '42079060825739264');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('46164130582761472', '44326032676753408');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('46164130582761472', '44332995393359872');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('46164130582761472', '44333062707744768');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('46164130582761472', '44474438040686592');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('46164130582761472', '45864670036234240');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('46164130582761472', '46239031234662400');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('46164130582761472', '46862379094380544');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('46164130582761472', '46862511944765440');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('46164130582761472', '46862830229524480');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('46164130582761472', '46863099453509632');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('46164130582761472', '46863346367991808');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('46164130582761472', '46863589419520000');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('46164130582761472', '46858901232029696');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('46164130582761472', '46859015694585856');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('46164130582761472', '46859424727306240');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('46164130582761472', '47156442767036416');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('46164130582761472', '47216551878725632');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('46164130582761472', '47216745617821696');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('47615060066963456', '37899586299236352');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('47615060066963456', '41071752721207296');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('47615060066963456', '41430491085148160');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('47615060066963456', '41438779612860416');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('47615060066963456', '41789444067430400');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('47615060066963456', '41903454221766656');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('47615060066963456', '41906647861301248');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('47615060066963456', '46862830229524480');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('47615060066963456', '46863099453509632');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('47615060066963456', '46863589419520000');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('47615060066963456', '47973984368594944');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('47615060066963456', '47976243613667328');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '37899586299236352');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '37899744348999680');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '40696216958275584');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '41065294818447360');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '41071752721207296');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '41430491085148160');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '41438779612860416');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '41789444067430400');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '41903454221766656');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '41906647861301248');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '42063400649363456');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '42071947340681216');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '42079060825739264');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '44326032676753408');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '44332995393359872');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '44333062707744768');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '44474438040686592');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '45864670036234240');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '46239031234662400');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '46858901232029696');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '46859015694585856');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '46859424727306240');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '46862379094380544');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '46862511944765440');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '46862830229524480');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '46863099453509632');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '46863346367991808');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '46863589419520000');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '47156442767036416');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '47216551878725632');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '47216745617821696');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '47584267768696832');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '47593342061514752');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '47608622313639936');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '47608882222075904');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '47973984368594944');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '47974191147782144');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '47974424665657344');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '47974555385335808');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '47974732070391808');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '47974990527598592');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '47975191568977920');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '47975597174951936');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '47976243613667328');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '49391522268844032');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '49391714045005824');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '49391851559456768');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '49392003502313472');
COMMIT;

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
-- Records of roles_pages
-- ----------------------------
BEGIN;
INSERT INTO `roles_pages` (`role_id`, `page_id`) VALUES ('46170564095643648', '40350365404631040');
INSERT INTO `roles_pages` (`role_id`, `page_id`) VALUES ('46170564095643648', '40351462613585920');
INSERT INTO `roles_pages` (`role_id`, `page_id`) VALUES ('46170564095643648', '40351711566499840');
INSERT INTO `roles_pages` (`role_id`, `page_id`) VALUES ('46170564095643648', '40352343601975296');
INSERT INTO `roles_pages` (`role_id`, `page_id`) VALUES ('46170564095643648', '40352749044371456');
INSERT INTO `roles_pages` (`role_id`, `page_id`) VALUES ('46170564095643648', '40708567220621312');
INSERT INTO `roles_pages` (`role_id`, `page_id`) VALUES ('46164130582761472', '40350365404631040');
INSERT INTO `roles_pages` (`role_id`, `page_id`) VALUES ('46164130582761472', '40351462613585920');
INSERT INTO `roles_pages` (`role_id`, `page_id`) VALUES ('46164130582761472', '40351711566499840');
INSERT INTO `roles_pages` (`role_id`, `page_id`) VALUES ('46164130582761472', '40352343601975296');
INSERT INTO `roles_pages` (`role_id`, `page_id`) VALUES ('46164130582761472', '40352749044371456');
INSERT INTO `roles_pages` (`role_id`, `page_id`) VALUES ('46164130582761472', '40708567220621312');
INSERT INTO `roles_pages` (`role_id`, `page_id`) VALUES ('47615060066963456', '40350365404631040');
INSERT INTO `roles_pages` (`role_id`, `page_id`) VALUES ('47615060066963456', '40351462613585920');
INSERT INTO `roles_pages` (`role_id`, `page_id`) VALUES ('47615060066963456', '40351711566499840');
INSERT INTO `roles_pages` (`role_id`, `page_id`) VALUES ('47615060066963456', '40352343601975296');
INSERT INTO `roles_pages` (`role_id`, `page_id`) VALUES ('47615060066963456', '40352749044371456');
INSERT INTO `roles_pages` (`role_id`, `page_id`) VALUES ('47615060066963456', '40708567220621312');
INSERT INTO `roles_pages` (`role_id`, `page_id`) VALUES ('47615060066963456', '47513502385967104');
INSERT INTO `roles_pages` (`role_id`, `page_id`) VALUES ('47615060066963456', '47973248603787264');
INSERT INTO `roles_pages` (`role_id`, `page_id`) VALUES ('47615060066963456', '47975943540576256');
INSERT INTO `roles_pages` (`role_id`, `page_id`) VALUES ('47615060066963456', '47989353607073792');
INSERT INTO `roles_pages` (`role_id`, `page_id`) VALUES ('47615060066963456', '47989749788446720');
INSERT INTO `roles_pages` (`role_id`, `page_id`) VALUES ('37904208560656384', '40350365404631040');
INSERT INTO `roles_pages` (`role_id`, `page_id`) VALUES ('37904208560656384', '40351711566499840');
INSERT INTO `roles_pages` (`role_id`, `page_id`) VALUES ('37904208560656384', '40352343601975296');
INSERT INTO `roles_pages` (`role_id`, `page_id`) VALUES ('37904208560656384', '40352749044371456');
INSERT INTO `roles_pages` (`role_id`, `page_id`) VALUES ('37904208560656384', '40708567220621312');
INSERT INTO `roles_pages` (`role_id`, `page_id`) VALUES ('37904208560656384', '47513502385967104');
INSERT INTO `roles_pages` (`role_id`, `page_id`) VALUES ('37904208560656384', '47989353607073792');
INSERT INTO `roles_pages` (`role_id`, `page_id`) VALUES ('37904208560656384', '47989749788446720');
INSERT INTO `roles_pages` (`role_id`, `page_id`) VALUES ('37904208560656384', '49387228375289856');
INSERT INTO `roles_pages` (`role_id`, `page_id`) VALUES ('37904208560656384', '40351462613585920');
INSERT INTO `roles_pages` (`role_id`, `page_id`) VALUES ('37904208560656384', '47973248603787264');
INSERT INTO `roles_pages` (`role_id`, `page_id`) VALUES ('37904208560656384', '47975943540576256');
INSERT INTO `roles_pages` (`role_id`, `page_id`) VALUES ('37904208560656384', '50765767066521600');
INSERT INTO `roles_pages` (`role_id`, `page_id`) VALUES ('37904208560656384', '50766109451751424');
COMMIT;

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
  `create_time` timestamp NULL DEFAULT (now()),
  `update_time` timestamp NULL DEFAULT (now()) ON UPDATE CURRENT_TIMESTAMP,
  `delete_time` timestamp NULL DEFAULT NULL,
  `status` int DEFAULT '1',
  `department_id` varchar(36) DEFAULT NULL,
  `is_admin` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `account` (`account`),
  KEY `users_department_id_fk` (`department_id`),
  CONSTRAINT `users_department_id_fk` FOREIGN KEY (`department_id`) REFERENCES `department` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('34980631960096768', 'admin', '$2a$10$kPjaoLdiPnZitTo.elv8KuzJGIQGChpQuMJfVyLs7yHtVMkjE1qcG', '超级管理员', NULL, '2024-04-23 15:40:14', '2024-05-09 15:29:01', NULL, 1, NULL, 1);
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('46430662030266368', '邹娟', '$2a$10$khkuQh9cpslVtWj2WEAsF.8NH/MxHYAvGENpYorrH./8woUci5ruG', '明', NULL, '2024-05-08 10:58:52', '2024-05-11 22:12:09', NULL, 1, '47230946730053632', 0);
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('46430672482471936', '姚静', '$2a$10$ytNU44lTSdrlePs4wu4X5e90geWc8Sh/Ga5WY7cGse34iVgruMj/O', '磊', NULL, '2024-05-08 10:58:54', '2024-05-10 18:16:21', NULL, 1, '47230539865788416', 0);
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('46430675846303744', '秦勇', '$2a$10$LnKevbGHMjpaR1ctPAyEGOjCwaIubAgd070NbarOz3wvquKa4pXLK', '勇', NULL, '2024-05-08 10:58:55', '2024-05-08 10:58:55', NULL, 1, NULL, 0);
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('46430679004614656', '黎秀兰', '$2a$10$J0dSPRpI62/y.A5Zkyc57.NDPmpZny3U5nLzE0z2J5IkIhbGRwNN2', '秀兰', NULL, '2024-05-08 10:58:56', '2024-05-08 10:58:56', NULL, 1, NULL, 0);
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('46430681936433152', '张静', '$2a$10$OvLG0A.GfWE4maYJ5caXdOSJ9otIjL0d3O.w5rS4rhwxTU1YdOKlS', '杰', NULL, '2024-05-08 10:58:57', '2024-05-11 15:00:30', NULL, 1, NULL, 0);
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('46430684994080768', '杜涛', '$2a$10$Drn9J3a5SUxNqVDHYLwGhuMKn/.G0XXTQ1mx22IGON7hJas3s6V9u', '静', NULL, '2024-05-08 10:58:57', '2024-05-08 10:58:57', NULL, 1, NULL, 0);
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('46431220808028160', '汪军', '$2a$10$VWUAcCqPmrwj.qVPGtAtvO9w5bSE/SHw7ZfbKbKm8Lj.NGvEpNPT2', '磊', NULL, '2024-05-08 11:01:05', '2024-05-08 11:01:05', NULL, 1, NULL, 0);
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('46431225157521408', '郝涛', '$2a$10$PCAH3/DCYw5X/qyugA5QierbJWC3qtUjcSQCFpwqgDXdhvr5.hN3C', '明', NULL, '2024-05-08 11:01:06', '2024-05-08 11:01:06', NULL, 1, NULL, 0);
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('46431228647182336', '曾磊', '$2a$10$HhyOiCGirFBXbIvaPBl4XuFo.sxT1pFfol./bXKLxVL38ftPQgZna', '平', NULL, '2024-05-08 11:01:07', '2024-05-08 11:01:07', NULL, 1, NULL, 0);
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('46431231260233728', '汤桂英', '$2a$10$pSLjqAEzaMj3g5Tt8GYfN.msOTmVnEBYDIamWDPgBVo3LAuf0Zq2S', '伟', NULL, '2024-05-08 11:01:08', '2024-05-17 17:06:40', NULL, 1, NULL, 0);
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('46431233982337024', '金超', '$2a$10$enPFyDmVT/DUqsyxJPATPeoHSNcThqzFG5Fx863MIO4BiQyTfN3fy', '勇', NULL, '2024-05-08 11:01:08', '2024-05-08 11:01:08', NULL, 1, NULL, 0);
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('46431236440199168', '雷杰', '$2a$10$SmYT4qE1zgbEWHRR.TA.q.EKYpq3VwacGxwjIRJeuUVJWSf5z4BSq', '丽', NULL, '2024-05-08 11:01:09', '2024-05-08 11:01:09', NULL, 1, NULL, 0);
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('46431239200051200', '梁娜', '$2a$10$6TGiCA046jZDEXJDeSW5MeNsywv/P2Tu3mP9k1nmVV.a/k5BuFPFa', '洋', NULL, '2024-05-08 11:01:10', '2024-05-08 11:01:10', NULL, 1, NULL, 0);
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('47615361893273600', 'test', '$2a$10$aRXCev.upsYNXfCrREKZcu/eHSQi6gx2M79slcF5GpyEpfhUrVoqi', 'test', NULL, '2024-05-11 17:26:26', '2024-05-11 17:26:26', NULL, 1, '47230539865788416', 0);
COMMIT;

-- ----------------------------
-- Table structure for users_roles
-- ----------------------------
DROP TABLE IF EXISTS `users_roles`;
CREATE TABLE `users_roles` (
  `user_id` varchar(36) NOT NULL,
  `role_id` varchar(36) NOT NULL,
  PRIMARY KEY (`user_id`,`role_id`),
  KEY `role_id` (`role_id`),
  CONSTRAINT `users_roles_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
  CONSTRAINT `users_roles_ibfk_2` FOREIGN KEY (`role_id`) REFERENCES `roles` (`role_id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of users_roles
-- ----------------------------
BEGIN;
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('34980631960096768', '37904208560656384');
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('34980631960096768', '45864950110883840');
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('46430662030266368', '45864950110883840');
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('46430672482471936', '45864950110883840');
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('46430675846303744', '45864950110883840');
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('46430679004614656', '45864950110883840');
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('46430662030266368', '46164130582761472');
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('46430672482471936', '46164130582761472');
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('46430675846303744', '46164130582761472');
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('46430679004614656', '46164130582761472');
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('46430681936433152', '46164130582761472');
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('46430684994080768', '46164130582761472');
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('46431225157521408', '46164130582761472');
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('46431228647182336', '46164130582761472');
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('46431231260233728', '46164130582761472');
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('46431233982337024', '46164130582761472');
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('46431236440199168', '46164130582761472');
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('46431239200051200', '46164130582761472');
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('46430662030266368', '46170564095643648');
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('46430672482471936', '46170564095643648');
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('46430675846303744', '46170564095643648');
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('46430681936433152', '46170564095643648');
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('46430684994080768', '46170564095643648');
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('46431220808028160', '46170564095643648');
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('46431225157521408', '46170564095643648');
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('46431228647182336', '46170564095643648');
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('46431231260233728', '46170564095643648');
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('46431233982337024', '46170564095643648');
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('46431236440199168', '46170564095643648');
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('47615361893273600', '47615060066963456');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
