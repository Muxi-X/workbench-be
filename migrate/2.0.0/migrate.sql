-- ----------------------------
-- 2.0.0 DB 迁移 SQL
-- ----------------------------


-- ----------------------------
-- status 服务
-- ----------------------------

-- status 点赞表
CREATE TABLE `user2status` (
	`id` int(11) NOT NULL AUTO_INCREMENT,
	`user_id` int(11) DEFAULT NULL,
	`status_id` int(11) DEFAULT NULL,
	PRIMARY KEY (`id`),
	KEY `user_id` (`user_id`),
	KEY `status_id` (`status_id`),
	UNIQUE KEY `user_status` (`user_id`,`status_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;


-- ----------------------------
-- user 服务
-- ----------------------------

-- token 黑名单
DROP TABLE IF EXISTS `blacklist`;
CREATE TABLE `blacklist` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `token` varchar(255) DEFAULT "" NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `expires_at` int(11) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_token` (`token`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;


-- ----------------------------
-- project 服务
-- ----------------------------

-- project 文件存储重构