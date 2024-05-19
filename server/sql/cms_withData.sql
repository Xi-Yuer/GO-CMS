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

 Date: 27/04/2024 15:37:44
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
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('0aa70d224bbf8863e4974de8307025cc3b638e8f', 'Xi-Yuer', '2214380963@qq.com', '2024-04-25 15:31:40', 'feat:用户管理');
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
INSERT INTO `commits` (`commit_id`, `author`, `email`, `commit_date`, `message`) VALUES ('064c698b49f40e32b2177d1b15db08acee6de1ee', 'Xi-Yuer', '2214380963@qq.com', '2024-04-03 18:39:35', 'refactor:dto.tmpl');
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
INSERT INTO `department` (`id`, `department_name`, `parent_department`, `create_time`, `update_time`, `delete_time`, `department_order`) VALUES ('34915012879323136', '部门（程桂英）', NULL, '2024-04-06 16:19:47', '2024-04-06 16:24:51', NULL, 0);
INSERT INTO `department` (`id`, `department_name`, `parent_department`, `create_time`, `update_time`, `delete_time`, `department_order`) VALUES ('34915061705216000', '部门（邵军）', '34915012879323136', '2024-04-06 16:19:59', '2024-04-06 16:19:59', NULL, 0);
INSERT INTO `department` (`id`, `department_name`, `parent_department`, `create_time`, `update_time`, `delete_time`, `department_order`) VALUES ('34915090499112960', '部门（吴秀兰）', NULL, '2024-04-06 16:20:06', '2024-04-06 16:20:06', NULL, 0);
INSERT INTO `department` (`id`, `department_name`, `parent_department`, `create_time`, `update_time`, `delete_time`, `department_order`) VALUES ('34915120370946048', '嘿嘿11', '34915012879323136', '2024-04-06 16:20:13', '2024-04-06 19:43:00', NULL, 0);
INSERT INTO `department` (`id`, `department_name`, `parent_department`, `create_time`, `update_time`, `delete_time`, `department_order`) VALUES ('37552846962429952', '部门（程霞）', NULL, '2024-04-13 23:01:36', '2024-04-13 23:01:36', NULL, 0);
INSERT INTO `department` (`id`, `department_name`, `parent_department`, `create_time`, `update_time`, `delete_time`, `department_order`) VALUES ('37905912664428544', '部门（康丽）', NULL, '2024-04-14 22:24:33', '2024-04-14 22:24:33', NULL, 0);
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
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `create_time`, `update_time`) VALUES ('37899586299236352', '获取菜单', 'GET', '/pages', '40350365404631040', 'GET:/pages', '获取菜单', '2024-04-14 21:59:25', '2024-04-14 21:59:25');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `create_time`, `update_time`) VALUES ('37899744348423240', '新建接口', 'POST', '/interface', '40350365404631040', 'POST:/interface', '新建接口', '2024-04-22 15:07:31', '2024-04-22 15:11:31');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `create_time`, `update_time`) VALUES ('37899744348999680', '获取用户菜单', 'GET', '/pages/menus', '40350365404631040', 'GET:/pages/menus', '获取用户菜单', '2024-04-14 22:00:03', '2024-04-14 22:26:10');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `create_time`, `update_time`) VALUES ('40696216958275584', '新建菜单', 'POST', '/pages', '40350365404631040', 'POST:/pages', '新建菜单', '2024-04-22 15:12:14', '2024-04-22 15:12:14');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `create_time`, `update_time`) VALUES ('41065294818447360', '创建用户', 'POST', '/users', '40350365404631040', 'POST:/users', '创建用户', '2024-04-23 15:38:49', '2024-04-23 15:38:49');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `create_time`, `update_time`) VALUES ('41071752721207296', '获取系统运行信息', 'GET', '/system', '40350365404631040', 'GET:/system', '获取系统运行信息', '2024-04-23 16:04:28', '2024-04-23 16:04:28');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `create_time`, `update_time`) VALUES ('41430491085148160', '获取Git提交日志', 'GET', '/log/commits', '40350365404631040', 'GET:/log/commits', '获取Git提交日志', '2024-04-24 15:49:59', '2024-04-24 15:49:59');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `create_time`, `update_time`) VALUES ('41438779612860416', '获取Git提交次数', 'GET', '/log/commits/count', '40350365404631040', 'GET:/log/commits/count', '获取Git提交次数', '2024-04-24 16:22:54', '2024-04-24 16:22:54');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `create_time`, `update_time`) VALUES ('41789444067430400', '获取用户', 'GET', '/users', '40350365404631040', 'GET:/users', '获取用户', '2024-04-25 17:20:52', '2024-04-25 17:20:52');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `create_time`, `update_time`) VALUES ('41903454221766656', '获取角色', 'GET', '/roles', '40350365404631040', 'GET:/roles', '获取角色', '2024-04-25 23:09:21', '2024-04-25 23:09:21');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `create_time`, `update_time`) VALUES ('41906647861301248', '获取部门', 'GET', '/department', '40350365404631040', 'GET:/department', '获取部门', '2024-04-25 23:22:03', '2024-04-25 23:22:03');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `create_time`, `update_time`) VALUES ('42063400649363456', '更新用户', 'PATCH', '/users/:id', '40350365404631040', 'PATCH:/users/:id', '更新用户', '2024-04-26 09:44:56', '2024-04-26 09:48:59');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `create_time`, `update_time`) VALUES ('42071947340681216', '删除用户', 'DELETE', '/users/:id', '40350365404631040', 'DELETE:/users/:id', '删除用户', '2024-04-26 10:18:53', '2024-04-26 10:18:53');
INSERT INTO `interfaces` (`interface_id`, `interface_name`, `interface_method`, `interface_path`, `interface_page_id`, `interface_dic`, `interface_desc`, `create_time`, `update_time`) VALUES ('42079060825739264', '获取用户详细', 'GET', '/users/:id', '40350365404631040', 'GET:/users/:id', '获取用户详细', '2024-04-26 10:47:09', '2024-04-26 10:47:09');
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
-- Records of pages
-- ----------------------------
BEGIN;
INSERT INTO `pages` (`page_id`, `parent_page`, `page_name`, `page_path`, `page_icon`, `page_component`, `page_order`, `can_edit`, `is_out_site`, `out_site_link`, `create_time`, `update_time`, `delete_time`) VALUES ('40350365404631040', NULL, '仪表盘', '/dashboard', 'DashboardOutlined', 'dashboard', 1, 1, 0, '', '2024-04-21 17:41:43', '2024-04-21 17:41:43', NULL);
INSERT INTO `pages` (`page_id`, `parent_page`, `page_name`, `page_path`, `page_icon`, `page_component`, `page_order`, `can_edit`, `is_out_site`, `out_site_link`, `create_time`, `update_time`, `delete_time`) VALUES ('40351462613585920', NULL, '系统管理', '/system', 'SettingOutlined', 'system', 2, 1, 0, '', '2024-04-21 17:46:05', '2024-04-21 17:58:42', NULL);
INSERT INTO `pages` (`page_id`, `parent_page`, `page_name`, `page_path`, `page_icon`, `page_component`, `page_order`, `can_edit`, `is_out_site`, `out_site_link`, `create_time`, `update_time`, `delete_time`) VALUES ('40351711566499840', '40351462613585920', '用户管理', '/system/user', 'UserOutlined', 'user', 1, 1, 0, '', '2024-04-21 17:47:04', '2024-04-21 17:47:04', NULL);
INSERT INTO `pages` (`page_id`, `parent_page`, `page_name`, `page_path`, `page_icon`, `page_component`, `page_order`, `can_edit`, `is_out_site`, `out_site_link`, `create_time`, `update_time`, `delete_time`) VALUES ('40352343601975296', '40351462613585920', '角色管理', '/system/role', 'MergeOutlined', 'role', 2, 1, 0, '', '2024-04-21 17:49:35', '2024-04-21 17:51:37', NULL);
INSERT INTO `pages` (`page_id`, `parent_page`, `page_name`, `page_path`, `page_icon`, `page_component`, `page_order`, `can_edit`, `is_out_site`, `out_site_link`, `create_time`, `update_time`, `delete_time`) VALUES ('40352749044371456', '40351462613585920', '部门管理', '/system/department', 'ApartmentOutlined', 'department', 3, 1, 0, '', '2024-04-21 17:51:12', '2024-04-21 17:51:37', NULL);
INSERT INTO `pages` (`page_id`, `parent_page`, `page_name`, `page_path`, `page_icon`, `page_component`, `page_order`, `can_edit`, `is_out_site`, `out_site_link`, `create_time`, `update_time`, `delete_time`) VALUES ('40696722552262656', NULL, '测试页面', '/test', 'FilterOutlined', 'test', 3, 1, 0, '', '2024-04-22 15:14:14', '2024-04-22 15:26:16', NULL);
INSERT INTO `pages` (`page_id`, `parent_page`, `page_name`, `page_path`, `page_icon`, `page_component`, `page_order`, `can_edit`, `is_out_site`, `out_site_link`, `create_time`, `update_time`, `delete_time`) VALUES ('40696900508192768', '40696722552262656', '页面一', '/test/test1', 'FilterOutlined', 'test', 1, 1, 0, '', '2024-04-22 15:14:57', '2024-04-22 16:13:55', NULL);
INSERT INTO `pages` (`page_id`, `parent_page`, `page_name`, `page_path`, `page_icon`, `page_component`, `page_order`, `can_edit`, `is_out_site`, `out_site_link`, `create_time`, `update_time`, `delete_time`) VALUES ('40704268730109952', '40696900508192768', '页面二', '/test/test2', 'FilterOutlined', 'test', 1, 1, 0, '', '2024-04-22 15:44:13', '2024-04-22 16:13:55', NULL);
INSERT INTO `pages` (`page_id`, `parent_page`, `page_name`, `page_path`, `page_icon`, `page_component`, `page_order`, `can_edit`, `is_out_site`, `out_site_link`, `create_time`, `update_time`, `delete_time`) VALUES ('40708567220621312', '40351462613585920', '菜单管理', '/system/menu', 'FolderOpenOutlined', 'menu', 4, 1, 0, '', '2024-04-22 16:01:18', '2024-04-22 16:01:18', NULL);
INSERT INTO `pages` (`page_id`, `parent_page`, `page_name`, `page_path`, `page_icon`, `page_component`, `page_order`, `can_edit`, `is_out_site`, `out_site_link`, `create_time`, `update_time`, `delete_time`) VALUES ('40710393135370240', '40696900508192768', '测试页面三', '/test/test3', 'FilterOutlined', '', 2, 1, 0, '', '2024-04-22 16:08:34', '2024-04-22 16:08:34', NULL);
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
  PRIMARY KEY (`role_id`),
  UNIQUE KEY `roles_pk` (`role_name`),
  UNIQUE KEY `roles_pk_2` (`role_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of roles
-- ----------------------------
BEGIN;
INSERT INTO `roles` (`role_id`, `role_name`, `description`, `create_time`, `update_time`, `delete_time`) VALUES ('32448651817127936', '角色(薛明)', '角色描述', '2024-03-30 20:59:21', '2024-03-30 20:59:21', NULL);
INSERT INTO `roles` (`role_id`, `role_name`, `description`, `create_time`, `update_time`, `delete_time`) VALUES ('32448655961100288', '角色(谭霞)', '角色描述', '2024-03-30 20:59:22', '2024-03-30 20:59:22', NULL);
INSERT INTO `roles` (`role_id`, `role_name`, `description`, `create_time`, `update_time`, `delete_time`) VALUES ('32506178231603200', '角色(钱洋)', '角色描述', '2024-03-31 00:47:56', '2024-03-31 00:47:56', NULL);
INSERT INTO `roles` (`role_id`, `role_name`, `description`, `create_time`, `update_time`, `delete_time`) VALUES ('32506184078462976', '角色(熊超)', '角色描述', '2024-03-31 00:47:58', '2024-03-31 00:47:58', NULL);
INSERT INTO `roles` (`role_id`, `role_name`, `description`, `create_time`, `update_time`, `delete_time`) VALUES ('32506187438100480', '角色(黄娜)', '角色描述', '2024-03-31 00:47:59', '2024-03-31 00:47:59', NULL);
INSERT INTO `roles` (`role_id`, `role_name`, `description`, `create_time`, `update_time`, `delete_time`) VALUES ('34696529495199744', '角色(唐超)', '角色描述', '2024-04-06 01:51:37', '2024-04-06 01:51:37', NULL);
INSERT INTO `roles` (`role_id`, `role_name`, `description`, `create_time`, `update_time`, `delete_time`) VALUES ('34974688459362304', '角色(江丽)', '角色描述', '2024-04-06 20:16:55', '2024-04-06 20:16:55', NULL);
INSERT INTO `roles` (`role_id`, `role_name`, `description`, `create_time`, `update_time`, `delete_time`) VALUES ('34974697074462720', '角色(邱丽)', '角色描述', '2024-04-06 20:16:57', '2024-04-06 20:16:57', NULL);
INSERT INTO `roles` (`role_id`, `role_name`, `description`, `create_time`, `update_time`, `delete_time`) VALUES ('34974700945805312', '角色(夏勇)', '角色描述', '2024-04-06 20:16:58', '2024-04-06 20:16:58', NULL);
INSERT INTO `roles` (`role_id`, `role_name`, `description`, `create_time`, `update_time`, `delete_time`) VALUES ('35247440273608704', '角色(黎伟)', '角色描述', '2024-04-07 14:20:44', '2024-04-07 14:20:44', NULL);
INSERT INTO `roles` (`role_id`, `role_name`, `description`, `create_time`, `update_time`, `delete_time`) VALUES ('35247574801715200', '角色(萧涛)', '角色描述', '2024-04-07 14:21:16', '2024-04-07 14:21:16', NULL);
INSERT INTO `roles` (`role_id`, `role_name`, `description`, `create_time`, `update_time`, `delete_time`) VALUES ('35247623136874496', '角色(秦静)', '角色描述', '2024-04-07 14:21:28', '2024-04-07 14:21:28', NULL);
INSERT INTO `roles` (`role_id`, `role_name`, `description`, `create_time`, `update_time`, `delete_time`) VALUES ('35248170237693952', '角色(侯刚)', '角色描述', '2024-04-07 14:23:38', '2024-04-07 14:23:38', NULL);
INSERT INTO `roles` (`role_id`, `role_name`, `description`, `create_time`, `update_time`, `delete_time`) VALUES ('35248274852024320', '角色(汪平)', '角色描述', '2024-04-07 14:24:03', '2024-04-07 14:24:03', NULL);
INSERT INTO `roles` (`role_id`, `role_name`, `description`, `create_time`, `update_time`, `delete_time`) VALUES ('35248410621644800', '角色(段洋)', '角色描述', '2024-04-07 14:24:35', '2024-04-07 14:24:35', NULL);
INSERT INTO `roles` (`role_id`, `role_name`, `description`, `create_time`, `update_time`, `delete_time`) VALUES ('35248847487766528', '角色(万霞)', '角色描述', '2024-04-07 14:26:20', '2024-04-07 14:26:20', NULL);
INSERT INTO `roles` (`role_id`, `role_name`, `description`, `create_time`, `update_time`, `delete_time`) VALUES ('35248898666663936', '角色(文磊)', '角色描述', '2024-04-07 14:26:32', '2024-04-07 14:26:32', NULL);
INSERT INTO `roles` (`role_id`, `role_name`, `description`, `create_time`, `update_time`, `delete_time`) VALUES ('35249103713603584', '角色(陆娜)', '角色描述', '2024-04-07 14:27:21', '2024-04-07 14:27:21', NULL);
INSERT INTO `roles` (`role_id`, `role_name`, `description`, `create_time`, `update_time`, `delete_time`) VALUES ('35249264233811968', '角色2', '角色描述111', '2024-04-07 14:27:59', '2024-04-07 19:20:50', NULL);
INSERT INTO `roles` (`role_id`, `role_name`, `description`, `create_time`, `update_time`, `delete_time`) VALUES ('35353039908900864', '1', '1', '2024-04-07 21:20:21', '2024-04-07 21:20:21', NULL);
INSERT INTO `roles` (`role_id`, `role_name`, `description`, `create_time`, `update_time`, `delete_time`) VALUES ('37070072052715520', '角色(毛秀兰)', '角色描述', '2024-04-12 15:03:13', '2024-04-12 15:03:13', NULL);
INSERT INTO `roles` (`role_id`, `role_name`, `description`, `create_time`, `update_time`, `delete_time`) VALUES ('37070107242926080', '角色(龙强)', '角色描述', '2024-04-12 15:03:22', '2024-04-12 15:03:22', NULL);
INSERT INTO `roles` (`role_id`, `role_name`, `description`, `create_time`, `update_time`, `delete_time`) VALUES ('37070141069987840', '角色(余明)', '角色描述', '2024-04-12 15:03:30', '2024-04-12 15:03:30', NULL);
INSERT INTO `roles` (`role_id`, `role_name`, `description`, `create_time`, `update_time`, `delete_time`) VALUES ('37070190629883904', '角色(郭磊)', '角色描述', '2024-04-12 15:03:42', '2024-04-12 15:03:42', NULL);
INSERT INTO `roles` (`role_id`, `role_name`, `description`, `create_time`, `update_time`, `delete_time`) VALUES ('37070267247235072', '角色(许杰)', '角色描述', '2024-04-12 15:04:00', '2024-04-12 15:04:00', NULL);
INSERT INTO `roles` (`role_id`, `role_name`, `description`, `create_time`, `update_time`, `delete_time`) VALUES ('37070514094608384', '角色(董秀英)', '角色描述', '2024-04-12 15:04:59', '2024-04-12 15:04:59', NULL);
INSERT INTO `roles` (`role_id`, `role_name`, `description`, `create_time`, `update_time`, `delete_time`) VALUES ('37070682684657664', '角色(姚超)', '角色描述', '2024-04-12 15:05:39', '2024-04-12 15:05:39', NULL);
INSERT INTO `roles` (`role_id`, `role_name`, `description`, `create_time`, `update_time`, `delete_time`) VALUES ('37070764171595776', '角色(谢丽)', '角色描述', '2024-04-12 15:05:58', '2024-04-12 15:05:58', NULL);
INSERT INTO `roles` (`role_id`, `role_name`, `description`, `create_time`, `update_time`, `delete_time`) VALUES ('37070985622458368', '角色', '角色描述111', '2024-04-12 15:06:51', '2024-04-12 15:07:57', NULL);
INSERT INTO `roles` (`role_id`, `role_name`, `description`, `create_time`, `update_time`, `delete_time`) VALUES ('37071736151216128', '角色(廖静)', '角色描述', '2024-04-12 15:09:50', '2024-04-12 15:09:50', NULL);
INSERT INTO `roles` (`role_id`, `role_name`, `description`, `create_time`, `update_time`, `delete_time`) VALUES ('37071759630929920', '角色(宋娟)', '角色描述', '2024-04-12 15:09:56', '2024-04-12 15:09:56', NULL);
INSERT INTO `roles` (`role_id`, `role_name`, `description`, `create_time`, `update_time`, `delete_time`) VALUES ('37071775535730688', '角色(张杰)', '角色描述', '2024-04-12 15:10:00', '2024-04-12 15:10:00', NULL);
INSERT INTO `roles` (`role_id`, `role_name`, `description`, `create_time`, `update_time`, `delete_time`) VALUES ('37075363863465984', '角色(熊艳)', '角色描述', '2024-04-12 15:24:15', '2024-04-12 15:24:15', NULL);
INSERT INTO `roles` (`role_id`, `role_name`, `description`, `create_time`, `update_time`, `delete_time`) VALUES ('37075401813528576', '角色(杨刚)', '角色描述', '2024-04-12 15:24:24', '2024-04-12 15:24:24', NULL);
INSERT INTO `roles` (`role_id`, `role_name`, `description`, `create_time`, `update_time`, `delete_time`) VALUES ('37075417013686272', '角色(谢刚)', '角色描述', '2024-04-12 15:24:28', '2024-04-12 15:24:28', NULL);
INSERT INTO `roles` (`role_id`, `role_name`, `description`, `create_time`, `update_time`, `delete_time`) VALUES ('37904208560656384', '超级管理员', '超级管理员，拥有所有权限', '2024-04-14 22:17:47', '2024-04-14 22:17:47', NULL);
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
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '37899586299236352');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '37899744348999680');
INSERT INTO `roles_interfaces` (`role_id`, `interface_id`) VALUES ('37904208560656384', '37899744348423240');
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
INSERT INTO `roles_pages` (`role_id`, `page_id`) VALUES ('37904208560656384', '40351462613585920');
INSERT INTO `roles_pages` (`role_id`, `page_id`) VALUES ('37904208560656384', '40351711566499840');
INSERT INTO `roles_pages` (`role_id`, `page_id`) VALUES ('37904208560656384', '40352343601975296');
INSERT INTO `roles_pages` (`role_id`, `page_id`) VALUES ('37904208560656384', '40696900508192768');
INSERT INTO `roles_pages` (`role_id`, `page_id`) VALUES ('37904208560656384', '40696722552262656');
INSERT INTO `roles_pages` (`role_id`, `page_id`) VALUES ('37904208560656384', '40696900508192768');
INSERT INTO `roles_pages` (`role_id`, `page_id`) VALUES ('37904208560656384', '40352749044371456');
INSERT INTO `roles_pages` (`role_id`, `page_id`) VALUES ('37904208560656384', '40704268730109952');
INSERT INTO `roles_pages` (`role_id`, `page_id`) VALUES ('37904208560656384', '40708567220621312');
INSERT INTO `roles_pages` (`role_id`, `page_id`) VALUES ('37904208560656384', '40710393135370240');
INSERT INTO `roles_pages` (`role_id`, `page_id`) VALUES ('37904208560656384', '40350365404631040');
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
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('34975442167402496', '崔静', '$2a$10$49pnDbnjUMsX4h0DCWaRMOfj0F1JO0EgHkcXJkdt8klH7F/ntl4jO', '霞', NULL, '2024-04-25 18:02:36', '2024-04-26 17:48:11', NULL, 1, '37552846962429952', 0);
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('34975446479147008', '胡勇', '$2a$10$MXo0Yog/uH7RbUqJ3HvipOTSwduAMPmwdW5haoMio3UaccCHAPAk6', '霞', NULL, '2024-04-06 20:19:56', '2024-04-26 11:47:03', NULL, 1, '34915090499112960', 0);
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('34980631960096768', '何杰', '$2a$10$kPjaoLdiPnZitTo.elv8KuzJGIQGChpQuMJfVyLs7yHtVMkjE1qcG', '杰', NULL, '2024-04-23 15:40:14', '2024-04-26 10:13:36', NULL, 1, '34915090499112960', 1);
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('34983961516052480', '马军', '$2a$10$mtXq7n6JzmV6xjdhw.8z1u8h/PH4qFv3B4S7zVsuNCtn/le32y/A2', '涛', NULL, '2024-04-06 20:53:46', '2024-04-26 17:56:24', NULL, 1, NULL, 0);
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('34984339355734016', '江军', '$2a$10$iO49d9JkqWKjo0Q8vB3dHedgNeDdKhakaIrEyXwIgGRzdelbztWeW', '刚', NULL, '2024-04-06 20:55:16', '2024-04-26 15:38:03', NULL, 1, '34915090499112960', 0);
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('35246648728752128', '江杰', '$2a$10$wxCOVB2FZTYCv6W4403Q7O6wEBQ4FgDVOmfhhNKZyNMsKsxkAghGS', '军', NULL, '2024-04-07 14:17:35', '2024-04-26 10:14:53', NULL, 1, '34915090499112960', 0);
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('35246706782113792', '吕平', '$2a$10$TnoWAKSdXtMZQnCUL0kh0.duwVkK56S7cH.jYLP/TW7OG5XFcmLry', '芳', NULL, '2024-04-07 14:17:49', '2024-04-07 14:17:49', NULL, 1, '34915090499112960', 0);
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('35246811849428992', '秦秀英', '$2a$10$I3ao7iWlD3Dw4hYbhBBL7und0Nl8udoirtkztrj8jlLl9dqUDg6XC', '涛', NULL, '2024-04-07 14:18:14', '2024-04-07 14:18:14', NULL, 1, '34915090499112960', 0);
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('37091545895145472', '123', '$2a$10$GET9D1rxaYcxOt6pv53hdugFt9PlaT.eyvoDvtZxkRR2iFkrghwPC', '123', NULL, '2024-04-12 16:29:07', '2024-04-12 16:29:07', '2024-04-12 08:29:08', 1, '34915090499112960', 0);
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('41065561194500096', '444', '$2a$10$kPjaoLdiPnZitTo.elv8KuzJGIQGChpQuMJfVyLs7yHtVMkjE1qcG', '444', NULL, '2024-04-23 15:39:52', '2024-04-26 10:19:36', '2024-04-26 02:19:37', 1, '34915090499112960', 0);
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('42152163400486912', '111321', '$2a$10$UIxxNRnUBrmVkXKcIWzepegytx.ishlGli/w87DU/BX4prFsPOqie', '111321', NULL, '2024-04-26 15:37:38', '2024-04-26 15:38:45', '2024-04-26 07:38:45', 1, '34915090499112960', 0);
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('42176535586672640', '尹强', '$2a$10$LOCwA2bGIs1J06NN/67tS.H1jlm8AnPsa9i3Ya5OOFNyfsQtvJ51e', '静', NULL, '2024-04-26 17:14:29', '2024-04-26 17:14:29', NULL, 1, '34915090499112960', 0);
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('42176543551655936', '白杰', '$2a$10$S5p1IIty3ZKwm6dK0TSAr.nbMxpLfo5e.u9OH7ubv3ImWQbWaIn62', '军', NULL, '2024-04-26 17:14:31', '2024-04-26 17:14:31', NULL, 1, '34915090499112960', 0);
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('42176547674656768', '毛杰', '$2a$10$elrjsXa.wPWVbZ7Q1W4VgOIi3.CuYgBkVudG1boRS43KVlGFxF/u6', '霞', NULL, '2024-04-26 17:14:32', '2024-04-26 17:14:32', NULL, 1, '34915090499112960', 0);
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('42176551164317696', '武娟', '$2a$10$wwksXAsEq2QIwEG1.5cLNetE3F/apNAs7hA4TlxM7nEgu3zg346FS', '平', NULL, '2024-04-26 17:14:33', '2024-04-26 17:14:33', NULL, 1, '34915090499112960', 0);
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('42176554746253312', '龚洋', '$2a$10$cRBE/eWmXwmW5aIi9/ekiOr825esm2NnZIjfbGiG3lhlHgOIcq5pe', '秀英', NULL, '2024-04-26 17:14:34', '2024-04-26 17:14:34', NULL, 1, '34915090499112960', 0);
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('42176558038781952', '薛平', '$2a$10$Gthyz9drEZtBEJzK/IzJZebbetfo1yv5VvaIOp5iWLaSSC2OOSFNi', '芳', NULL, '2024-04-26 17:14:35', '2024-04-26 17:14:35', NULL, 1, '34915090499112960', 0);
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('42176560429535232', '吕霞', '$2a$10$ODpfMLa9CaHL4GGmStCpNOySCRr6Eg5e5uzrOjAJTqoRUwDnWV0EK', '洋', NULL, '2024-04-26 17:14:35', '2024-04-26 17:14:35', NULL, 1, '34915090499112960', 0);
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('42176562862231552', '薛秀兰', '$2a$10$8QoDAPoDNqZsJUu6JWsXluyv7ZIEx8t0OUgdOlefqZB3bbUWtoiFq', '霞', NULL, '2024-04-26 17:14:36', '2024-04-26 17:14:36', NULL, 1, '34915090499112960', 0);
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('42176566075068416', '尹军', '$2a$10$XjfSUqydcOjY5dPQn7TuR.rQs29buB.tKNquX3Ot8gu6A95HX62nG', '娜', NULL, '2024-04-26 17:14:36', '2024-04-26 17:14:36', NULL, 1, '34915090499112960', 0);
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('42176569917050880', '武静', '$2a$10$bXNSNUcwxEssX15XPRc1SeNw5X89ZE/9NSSx6nRt4pxNNi942NpvG', '超', NULL, '2024-04-26 17:14:37', '2024-04-26 17:14:37', NULL, 1, '34915090499112960', 0);
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('42176573113110528', '黎超', '$2a$10$598qfLMqda/kMHSDnuN8AuofnxN7Wkcz1yq6uIU02er/T2EheRVNi', '敏', NULL, '2024-04-26 17:14:38', '2024-04-26 17:14:38', NULL, 1, '34915090499112960', 0);
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('42176576447582208', '黎涛', '$2a$10$W9jF/P5U3I0oFnwNkEEPO.mxv1h1Mp7tsHpgIgLwXYRlqOjO0VmW2', '刚', NULL, '2024-04-26 17:14:39', '2024-04-26 17:14:39', NULL, 1, '34915090499112960', 0);
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('42176746128150528', '徐静', '$2a$10$7ASKj5WLAtnrepmufl/2Mu3yVrzHP9ZdZGNtLaSMp6LY7I6hT7xVG', '秀兰', NULL, '2024-04-26 17:15:19', '2024-04-26 17:15:19', NULL, 1, '34915090499112960', 0);
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('42176749340987392', '杜霞', '$2a$10$EY5YKBg/6hBvKcVd8FB57e4sfoPS.Rrz1nWiARmBt.xkh8gJqo8/C', '娟', NULL, '2024-04-26 17:15:20', '2024-04-26 17:15:20', NULL, 1, '34915090499112960', 0);
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('42176752511881216', '丁洋', '$2a$10$XcVGFlvmaZtvd0jJ/QXd6O5UBifUgBzScjEY.2x5XDlHCdUHDACAm', '娟', NULL, '2024-04-26 17:15:21', '2024-04-26 17:15:21', NULL, 1, '34915090499112960', 0);
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('42176755577917440', '贺强', '$2a$10$n5xWNfTj/yrHASvBXBA78OnXKSf4oarHvXmB8R4GgLYddZwTQgcjm', '超', NULL, '2024-04-26 17:15:22', '2024-04-26 17:15:22', NULL, 1, '34915090499112960', 0);
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('42176759533146112', '刘明', '$2a$10$OviHY/pPt8h60GTIbdbbquUhxZvZlZZIHKt7ZXrEjfIgQl3sJ9MGm', '丽', NULL, '2024-04-26 17:15:23', '2024-04-26 17:15:23', NULL, 1, '34915090499112960', 0);
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('42176762481741824', '乔霞', '$2a$10$rNgBP5Svd7.1RVqtL8u6fu0eabyq/vxYfsJGdqumS4g97QoDBSh52', '明', NULL, '2024-04-26 17:15:23', '2024-04-26 17:15:23', NULL, 1, '34915090499112960', 0);
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('42176765866545152', '董洋', '$2a$10$QMkTDwa/vjlj10xdia/9mOYzwwGZzPdkzu7owaaALOyxcLyhfzJ/S', '杰', NULL, '2024-04-26 17:15:24', '2024-04-26 17:15:24', NULL, 1, '34915090499112960', 0);
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('42176768555094016', '汪超', '$2a$10$hxx/zxXIHWMZNt7d8JuAGet94xPNtAUtFEpA./M/zWtdQ6.wlhnNu', '娜', NULL, '2024-04-26 17:15:25', '2024-04-26 17:57:48', NULL, 1, '34915090499112960', 0);
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('42176771625324544', '夏秀兰', '$2a$10$0H/gc2yJZKat7BSIORba7OHTHOspve6mYgjOXdJtzGlyzxPpkUcyy', '军', NULL, '2024-04-26 17:15:26', '2024-04-26 17:15:26', NULL, 1, '34915090499112960', 0);
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('42176774129324032', '孟娟', '$2a$10$AAM6ZAIapVevbhDW36h8rO58jUbSFbLiJFJEtsTGawLm1ulxVvK02', '娜', NULL, '2024-04-26 17:15:26', '2024-04-26 17:15:26', NULL, 1, '34915090499112960', 0);
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('42176801304219648', '锺丽', '$2a$10$6/NC6SOvI/YwdZ4MTXQJPeEMco4kp76m/NBO36.JlNam/s/YhWGqW', '军', NULL, '2024-04-26 17:15:33', '2024-04-26 17:15:33', NULL, 1, '34915090499112960', 0);
INSERT INTO `users` (`id`, `account`, `password`, `nickname`, `avatar`, `create_time`, `update_time`, `delete_time`, `status`, `department_id`, `is_admin`) VALUES ('42176804705800192', '康明', '$2a$10$NQH0K6W.QHiMW5b/kDkOD.sFj9Tq6nF26rvFbkMSmlhqSgfUQisbu', '艳', NULL, '2024-04-26 17:15:33', '2024-04-26 17:15:33', NULL, 1, '34915090499112960', 0);
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
  CONSTRAINT `users_roles_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  CONSTRAINT `users_roles_ibfk_2` FOREIGN KEY (`role_id`) REFERENCES `roles` (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of users_roles
-- ----------------------------
BEGIN;
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('34975442167402496', '32448651817127936');
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('34975442167402496', '32448655961100288');
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('34975442167402496', '32506178231603200');
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('34975442167402496', '32506184078462976');
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('34975442167402496', '32506187438100480');
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('34975442167402496', '34696529495199744');
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('34975442167402496', '34974688459362304');
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('34975442167402496', '34974697074462720');
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('35246648728752128', '34974697074462720');
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('37091545895145472', '34974697074462720');
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('41065561194500096', '34974697074462720');
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('42152163400486912', '34974697074462720');
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('34975442167402496', '34974700945805312');
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('35246648728752128', '34974700945805312');
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('35246706782113792', '34974700945805312');
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('37091545895145472', '34974700945805312');
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('41065561194500096', '34974700945805312');
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('34975442167402496', '35247440273608704');
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('34975442167402496', '37904208560656384');
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES ('34980631960096768', '37904208560656384');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
