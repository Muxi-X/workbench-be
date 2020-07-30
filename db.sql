-- ----------------------------
-- Table structure for status
-- ----------------------------
DROP TABLE IF EXISTS `status`;
CREATE TABLE `status` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `content` text,
  `title` varchar(20) DEFAULT NULL,
  `time` varchar(50) DEFAULT NULL,
  `like` int(11) DEFAULT NULL,
  `comment` int(11) DEFAULT NULL,
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
  `role` int(11) DEFAULT NULL,
  `email_service` tinyint(1) DEFAULT NULL,
  `message` tinyint(1) DEFAULT NULL,
  `team_id` int(11) DEFAULT NULL,
  `group_id` int(11) DEFAULT NULL,
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
  `action` varchar(20) DEFAULT NULL,
  `source_kindid` int(11) DEFAULT NULL,
  `source_objectname` varchar(100) DEFAULT NULL,
  `source_objectid` int(11) DEFAULT NULL,
  `source_projectname` varchar(100) DEFAULT NULL,
  `source_projectid` int(11) DEFAULT NULL,
  `timeday` varchar(20) DEFAULT NULL,
  `timehm` varchar(20) DEFAULT NULL,
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
  `filetree` text,
  `doctree` text,
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
  `re` tinyint(1) DEFAULT NULL,
  `top` tinyint(1) DEFAULT NULL,
  `create_time` varchar(30) DEFAULT NULL,
  `delete_time` varchar(30) DEFAULT NULL,
  `editor_id` int(11) DEFAULT NULL,
  `creator_id` int(11) DEFAULT NULL,
  `project_id` int(11) DEFAULT NULL,
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
  `re` tinyint(1) DEFAULT NULL,
  `top` tinyint(1) DEFAULT NULL,
  `create_time` varchar(30) DEFAULT NULL,
  `delete_time` varchar(30) DEFAULT NULL,
  `creator_id` int(11) DEFAULT NULL,
  `project_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `creator_id` (`creator_id`),
  KEY `project_id` (`project_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;


DROP TABLE IF EXISTS `user2files`;
CREATE TABLE `user2files` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT NULL,
  `file_id` int(11) DEFAULT NULL,
  `file_kind` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;


-- 文件-文件夹
DROP TABLE IF EXISTS `foldersforfiles`;
CREATE TABLE `foldersforfiles` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL,
  `create_time` varchar(30) DEFAULT NULL,
  `create_id` int(11) DEFAULT NULL,
  `project_id` int(11) DEFAULT NULL,
  `re` tinyint(1) DEFAULT NULL,
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
  `project_id` int(11) DEFAULT NULL,
  `re` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `create_id` (`create_id`),
  KEY `project_id` (`project_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;
