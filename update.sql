-- --------------------------------
-- Table structure for user2status
-- --------------------------------
-- 用户-进度 点赞表
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
-- 项目回收站表，一个项目一个回收站，回收站的文件包括文件、文档、文件夹、文档夹
DROP TABLE IF EXISTS `trashbin`;
CREATE TABLE `trashbin` (
	`id` int(11) NOT NULL AUTO_INCREMENT,
	`create_time` varchar(50) DEFAULT NULL,
    `delete_time` varchar(50) DEFAULT NULL,
	`file_id` int(11) DEFAULT NULL COMMENT "文件的id，包括文件、文档、文件夹、文档夹",
	`file_type` tinyint(4) DEFAULT NUll COMMENT "文件的类型，1-doc 2-file 3-docFolder 4-fileFolder",
	`name` varchar(255) DEFAULT NULL COMMENT "文件名",
	`re` tinyint(1) DEFAULT NULL COMMENT "标记回收站内文件是否被删除 0-未删除 1-删除 删除文件只需将 re 置 1",
	`expires_at` int(11) unsigned DEFAULT NULL COMMENT "过期时间，由定时任务使用，过期将自动删除",
	PRIMARY KEY (`id`),
	UNIQUE KEY `idx_id_type` (`file_id`,`file_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;

-- add children for project, foldersformds, foldersforfiles
-- 以下四个字段用来表示当前节点下一层的文件节点，存储形式都是 <id>-<is_folder> 形式
-- id 表示文件的 id，is-folder 表示是不是文件
-- 这里不需要标志是文件和文档，在文件夹下的只能是文件，文档夹下的只能是文档
ALTER TABLE `projects` ADD `file_children` TEXT DEFAULT NULL;

ALTER TABLE `projects` ADD `doc_children` TEXT DEFAULT NULL;

ALTER TABLE `foldersforfiles` ADD `children` TEXT DEFAULT NULL;

ALTER TABLE `foldersformds` ADD `children` TEXT DEFAULT NULL;

-- add father_id for doc, file, folder
-- 以下四个字段用来标志当前节点的父节点
-- father_id 为 0 表示父节点是项目
ALTER TABLE `docs` ADD `father_id` int(11) DEFAULT 0;

ALTER TABLE `files` ADD `father_id` int(11) DEFAULT 0;

ALTER TABLE `foldersformds` ADD `father_id` int(11) DEFAULT 0;

ALTER TABLE `foldersforfiles` ADD `father_id` int(11) DEFAULT 0;

-- add index for user2projects index
ALTER TABLE `user2projects` ADD UNIQUE INDEX(`user_id`,`project_id`);

-- add last_edit_time for doc
ALTER TABLE `docs` ADD `last_edit_time` varchar(30) DEFAULT NULL;

-- add project soft delete
-- 使用 gorm 提供好的软删除
ALTER TABLE `projects` ADD `deleted_at` datetime DEFAULT NULL;

-- ----------------------------
-- Table structure for user2attentions
-- ----------------------------

DROP TABLE IF EXISTS `user2attentions`;
CREATE TABLE `user2attentions` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `user_id` int(11) DEFAULT NULL,
    `doc_id` int(11) DEFAULT NULL,
    `time_day` varchar(20) DEFAULT NULL,
    `time_hm` varchar(20) DEFAULT NULL,
    `file_kind` int(11) DEFAULT NULL COMMENT "file 的类型，包括 doc 和 file，0-doc 1-file",
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;

-- add creator_id for projects
ALTER TABLE `projects` ADD COLUMN `creator_id` int(11);
alter table `user2attentions` add column `file_kind` int(11) DEFAULT NULL;
