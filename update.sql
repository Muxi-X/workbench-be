-- --------------------------------
-- Table structure for user2status
-- --------------------------------

DROP TABLE IF EXISTS `user2status`;
CREATE TABLE `user2status` (
	`id` int(11) NOT NULL AUTO_INCREMENT,
	`user_id` int(11) DEFAULT NULL,
	`status_id` int(11) DEFAULT NULL,
	PRIMARY KEY (`id`),
	KEY `user_id` (`user_id`),
	KEY `status_id` (`status_id`),
	UNIQUE KEY `user_status` (`user_id`,`status_id`) 
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;

-- -----------------------------
-- Table structure for trashbin
-- -----------------------------

DROP TABLE IF EXISTS `trashbin`;
CREATE TABLE `trashbin` (
	`id` int(11) NOT NULL AUTO_INCREMENT,
	`file_id` int(11) DEFAULT NULL,
	`file_type` tinyint(4) DEFAULT NUll,
	`name` varchar(255) DEFAULT NULL,
	`re` tinyint(1) DEFAULT NULL,
	`expires_at` int(11) unsigned DEFAULT NULL,
	PRIMARY KEY (`id`),
	UNIQUE KEY `idx_id_type` (`file_id`,`file_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;

-- add children for project, foldersformds, foldersforfiles
ALTER TABLE `projects` ADD `file_children` TEXT DEFAULT NULL;

ALTER TABLE `projects` ADD `doc_children` TEXT DEFAULT NULL;

ALTER TABLE `foldersforfiles` ADD `children` TEXT DEFAULT NULL;

ALTER TABLE `foldersformds` ADD `children` TEXT DEFAULT NULL;

-- add index for user2projects index
ALTER TABLE `user2projects` ADD UNIQUE INDEX(`user_id`,`project_id`);

-- add last_edit_time for doc
ALTER TABLE `docs` ADD `last_edit_time` varchar(30) DEFAULT NULL;

-- add project soft delete
ALTER TABLE `projects` ADD `deleted_at` timestamp DEFAULT NULL;