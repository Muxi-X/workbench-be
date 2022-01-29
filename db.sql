CREATE DATABASE workbench;

use workbench;

-- ----------------------------
-- Table structure for status
-- ----------------------------
DROP TABLE IF EXISTS `status`;
CREATE TABLE `status` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `content` text,
  `title` varchar(20) DEFAULT NULL,
  `time` varchar(50) DEFAULT NULL,
  `like` int(11) DEFAULT NULL COMMENT '点赞数',
  `comment` int(11) DEFAULT NULL COMMENT '评论数',
  `user_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(20) DEFAULT NULL,
  `real_name` varchar(20) DEFAULT NULL,
  `email` varchar(35) DEFAULT NULL,
  `avatar` text,
  `tel` varchar(15) DEFAULT NULL,
  `role` int(11) DEFAULT NULL COMMENT '权限 0-无权限用户 1-普通用户 3-管理员 4-超级管理员',
  `email_service` tinyint(1) DEFAULT NULL,
  `message` tinyint(1) DEFAULT NULL,
  `team_id` int(11) DEFAULT NULL COMMENT '团队 id，木犀团队是1',
  `group_id` int(11) DEFAULT NULL COMMENT '组别 id 1-产品 2-前端 3-后端 4-安卓 5-设计',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  UNIQUE KEY `email` (`email`),
  KEY `team_id` (`team_id`),
  KEY `group_id` (`group_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for feeds
-- ----------------------------

DROP TABLE IF EXISTS `feeds`;
CREATE TABLE `feeds` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `userid` int(11) DEFAULT NULL,
  `username` varchar(100) DEFAULT NULL,
  `useravatar` varchar(200) DEFAULT NULL,
  `action` varchar(20) DEFAULT NULL COMMENT '动作，存储如 <创建>、<编辑>、<删除>、<评论>、<加入> 等常量字符串',
  `source_kindid` int(11) DEFAULT NULL COMMENT '动态的类型 1 -> 团队，2 -> 项目，3 -> 文档，4 -> 文件，6 -> 进度（5 不使用）',
  `source_objectname` varchar(100) DEFAULT NULL COMMENT 'object 包括 status、file、doc、等，这里是它们的名字',
  `source_objectid` int(11) DEFAULT NULL COMMENT '对象的 id',
  `source_projectname` varchar(100) DEFAULT NULL COMMENT '如果是 file/doc，这里存储其项目名，否则为 NULL',
  `source_projectid` int(11) DEFAULT NULL COMMENT '如果是 file/doc，这里存项目id，否则为 NULL',
  `timeday` varchar(20) DEFAULT NULL COMMENT '日期',
  `timehm` varchar(20) DEFAULT NULL COMMENT '时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for projects
-- ----------------------------

DROP TABLE IF EXISTS `projects`;
CREATE TABLE `projects` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) DEFAULT NULL,
  `intro` varchar(100) DEFAULT NULL,
  `time` varchar(50) DEFAULT NULL,
  `count` int(11) DEFAULT NULL,
  `team_id` int(11) DEFAULT NULL,
  `filetree` text COMMENT '旧版的子文件树，现已弃用，详情见 update.sql',
  `doctree` text COMMENT '旧版的子文档树，现已弃用，详情见 update.sql',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `team_id` (`team_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;

DROP TABLE IF EXISTS `user2projects`;
CREATE TABLE `user2projects` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT NULL,
  `project_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for applys
-- ----------------------------

DROP TABLE IF EXISTS `applys`;
CREATE TABLE `applys` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for groups
-- ----------------------------

DROP TABLE IF EXISTS `groups`;
CREATE TABLE `groups` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(10) DEFAULT NULL,
  `order` int(11) DEFAULT NULL,
  `count` int(11) DEFAULT NULL,
  `leader` int(11) DEFAULT NULL,
  `time` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  UNIQUE KEY `order` (`order`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for teams
-- ----------------------------

DROP TABLE IF EXISTS `teams`;
CREATE TABLE `teams` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(10) DEFAULT NULL,
  `count` int(11) DEFAULT NULL,
  `time` varchar(50) DEFAULT NULL,
  `creator` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;


-- ----------------------------
-- Table structure for comments
-- ----------------------------

DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `kind` int(11) DEFAULT NULL,
  `content` text,
  `time` varchar(50) DEFAULT NULL,
  `creator` int(11) DEFAULT NULL,
  `doc_id` int(11) DEFAULT NULL,
  `file_id` int(11) DEFAULT NULL,
  `statu_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `doc_id` (`doc_id`),
  KEY `file_id` (`file_id`),
  KEY `statu_id` (`statu_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;


-- --------------------------------------------
-- Table structure for docs, files and folders
-- --------------------------------------------

DROP TABLE IF EXISTS `docs`;
CREATE TABLE `docs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `filename` varchar(150) DEFAULT NULL,
  `content` text,
  `re` tinyint(1) DEFAULT NULL COMMENT '标志是否删除，0-未删除 1-删除 删除时只要将 re 置为 1',
  `top` tinyint(1) DEFAULT NULL,
  `create_time` varchar(30) DEFAULT NULL,
  `delete_time` varchar(30) DEFAULT NULL,
  `editor_id` int(11) DEFAULT NULL,
  `creator_id` int(11) DEFAULT NULL,
  `project_id` int(11) DEFAULT NULL COMMENT '此文档所属的项目 id',
  PRIMARY KEY (`id`),
  KEY `editor_id` (`editor_id`),
  KEY `creator_id` (`creator_id`),
  KEY `project_id` (`project_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;


DROP TABLE IF EXISTS `files`;
CREATE TABLE `files` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `url` varchar(150) DEFAULT NULL,
  `filename` varchar(150) DEFAULT NULL,
  `realname` varchar(150) DEFAULT NULL,
  `re` tinyint(1) DEFAULT NULL COMMENT '标志是否删除，0-未删除 1-删除 删除时只要将 re 置为 1',
  `top` tinyint(1) DEFAULT NULL,
  `create_time` varchar(30) DEFAULT NULL,
  `delete_time` varchar(30) DEFAULT NULL,
  `creator_id` int(11) DEFAULT NULL,
  `project_id` int(11) DEFAULT NULL COMMENT '此文件所属的项目 id',
  PRIMARY KEY (`id`),
  KEY `creator_id` (`creator_id`),
  KEY `project_id` (`project_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- 用户-文件关注表
-- !! user2files to attentions;
DROP TABLE IF EXISTS `user2files`;
CREATE TABLE `user2files` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT NULL,
  `file_id` int(11) DEFAULT NULL COMMENT '文件的 id，这里文件包括 doc 和 file',
  `file_kind` int(11) DEFAULT NULL COMMENT 'file 的类型，包括 doc 和 file，1-doc 2-file',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- 文件-文件夹
DROP TABLE IF EXISTS `foldersforfiles`;
CREATE TABLE `foldersforfiles` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL,
  `create_time` varchar(30) DEFAULT NULL,
  `create_id` int(11) DEFAULT NULL,
  `project_id` int(11) DEFAULT NULL COMMENT '该文件夹所属项目 id',
  `re` tinyint(1) DEFAULT NULL COMMENT '标志是否删除，0-未删除 1-删除 删除时只要将 re 置为 1',
  PRIMARY KEY (`id`),
  KEY `create_id` (`create_id`),
  KEY `project_id` (`project_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;


-- 文档-文件夹
DROP TABLE IF EXISTS `foldersformds`;
CREATE TABLE `foldersformds` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL,
  `create_time` varchar(30) DEFAULT NULL,
  `create_id` int(11) DEFAULT NULL,
  `project_id` int(11) DEFAULT NULL COMMENT '该文档夹所属项目 id',
  `re` tinyint(1) DEFAULT NULL COMMENT '标志是否删除，0-未删除 1-删除 删除时只要将 re 置为 1',
  PRIMARY KEY (`id`),
  KEY `create_id` (`create_id`),
  KEY `project_id` (`project_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;


-- token 黑名单
DROP TABLE IF EXISTS `blacklist`;
CREATE TABLE `blacklist` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `token` varchar(255) DEFAULT '' NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `expires_at` int(11) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_token` (`token`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;
