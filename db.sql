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
) ENGINE=InnoDB AUTO_INCREMENT=7393 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;
