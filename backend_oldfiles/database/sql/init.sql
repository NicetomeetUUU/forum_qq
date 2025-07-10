CREATE DATABASE IF NOT EXISTS qq_forum DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE qq_forum;

-- 用户表
CREATE TABLE IF NOT EXISTS `user` (
    `id` BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT 'user id',
    `email` VARCHAR(255) NOT NULL UNIQUE DEFAULT '' COMMENT 'email',
    `password` VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'password',
    `username` VARCHAR(255) NOT NULL UNIQUE DEFAULT '' COMMENT 'username',
    `avatar` VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'avatar',
    `signature` VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'signature',
    `birthday` DATE DEFAULT NULL COMMENT 'birthday',
    `phone` VARCHAR(20) DEFAULT NULL COMMENT 'phone',
    `role` ENUM('user', 'admin', 'moderator') NOT NULL DEFAULT 'user' COMMENT 'role',
    `status` ENUM('active', 'inactive') NOT NULL DEFAULT 'active' COMMENT 'status',
    `is_deleted` TINYINT(1) NOT NULL DEFAULT 0,
    `last_login_at` DATETIME DEFAULT NULL COMMENT 'last login time',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_email (email),
    INDEX idx_username (username),
    INDEX idx_phone (phone)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='user table';

-- 分类表
CREATE TABLE IF NOT EXISTS `category` (
    `id` BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT 'category id',
    `name` VARCHAR(100) NOT NULL COMMENT 'category name',
    `description` TEXT COMMENT 'category description',
    `parent_id` BIGINT DEFAULT NULL COMMENT 'parent category id',
    `sort_order` INT NOT NULL DEFAULT 0 COMMENT 'sort order',
    `is_active` TINYINT(1) NOT NULL DEFAULT 1 COMMENT 'is active',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_parent_id (parent_id),
    INDEX idx_sort_order (sort_order)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='category table';

-- 帖子表
CREATE TABLE IF NOT EXISTS `post` (
    `id` bigint AUTO_INCREMENT PRIMARY KEY COMMENT 'post id',
    `title` VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'title',
    `content` TEXT NOT NULL COMMENT 'content',
    `user_id` BIGINT NOT NULL COMMENT 'user id',
    `category_id` BIGINT DEFAULT NULL COMMENT 'category id',
    `view_count` INT NOT NULL DEFAULT 0 COMMENT 'view count',
    `like_count` INT NOT NULL DEFAULT 0 COMMENT 'like count',
    `comment_count` INT NOT NULL DEFAULT 0 COMMENT 'comment count',
    `status` ENUM('published', 'draft', 'deleted') NOT NULL DEFAULT 'published' COMMENT 'status',
    `is_top` TINYINT(1) NOT NULL DEFAULT 0 COMMENT 'is top',
    `is_hot` TINYINT(1) NOT NULL DEFAULT 0 COMMENT 'is hot',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created time',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'updated time',
    INDEX idx_user_id (user_id),
    INDEX idx_category_id (category_id),
    INDEX idx_status (status),
    INDEX idx_is_top (is_top),
    INDEX idx_created_at (created_at)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='post table';

-- 评论表
CREATE TABLE IF NOT EXISTS `comment` (
    `id` bigint AUTO_INCREMENT PRIMARY KEY COMMENT 'comment id',
    `content` TEXT NOT NULL COMMENT 'content',
    `user_id` BIGINT NOT NULL COMMENT 'user id',
    `post_id` BIGINT NOT NULL COMMENT 'post id',
    `parent_id` BIGINT DEFAULT NULL COMMENT 'parent comment id',
    `like_count` INT NOT NULL DEFAULT 0 COMMENT 'like count',
    `status` ENUM('published', 'deleted', 'hidden') NOT NULL DEFAULT 'published' COMMENT 'status',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created time',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'updated time',
    INDEX idx_post_id (post_id),
    INDEX idx_parent_id (parent_id),
    INDEX idx_user_id (user_id),
    INDEX idx_status (status)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='comment table';

-- 用户点赞表
CREATE TABLE IF NOT EXISTS `user_like` (
    `id` BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT 'like id',
    `user_id` BIGINT NOT NULL COMMENT 'user id',
    `target_type` ENUM('post', 'comment') NOT NULL COMMENT 'target type',
    `target_id` BIGINT NOT NULL COMMENT 'target id',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE KEY uk_user_target (user_id, target_type, target_id),
    INDEX idx_target (target_type, target_id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='user like table';

-- 用户关注表
CREATE TABLE IF NOT EXISTS `user_follow` (
    `id` BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT 'follow id',
    `follower_id` BIGINT NOT NULL COMMENT 'follower user id',
    `following_id` BIGINT NOT NULL COMMENT 'following user id',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE KEY uk_follow (follower_id, following_id),
    INDEX idx_follower (follower_id),
    INDEX idx_following (following_id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='user follow table';

-- 添加外键约束
ALTER TABLE `post` ADD CONSTRAINT `fk_post_user` 
FOREIGN KEY (`user_id`) REFERENCES `user`(`id`) ON DELETE CASCADE;

ALTER TABLE `post` ADD CONSTRAINT `fk_post_category` 
FOREIGN KEY (`category_id`) REFERENCES `category`(`id`) ON DELETE SET NULL;

ALTER TABLE `comment` ADD CONSTRAINT `fk_comment_user` 
FOREIGN KEY (`user_id`) REFERENCES `user`(`id`) ON DELETE CASCADE;

ALTER TABLE `comment` ADD CONSTRAINT `fk_comment_post` 
FOREIGN KEY (`post_id`) REFERENCES `post`(`id`) ON DELETE CASCADE;

ALTER TABLE `comment` ADD CONSTRAINT `fk_comment_parent` 
FOREIGN KEY (`parent_id`) REFERENCES `comment`(`id`) ON DELETE CASCADE;

ALTER TABLE `user_like` ADD CONSTRAINT `fk_like_user` 
FOREIGN KEY (`user_id`) REFERENCES `user`(`id`) ON DELETE CASCADE;

ALTER TABLE `user_follow` ADD CONSTRAINT `fk_follow_follower` 
FOREIGN KEY (`follower_id`) REFERENCES `user`(`id`) ON DELETE CASCADE;

ALTER TABLE `user_follow` ADD CONSTRAINT `fk_follow_following` 
FOREIGN KEY (`following_id`) REFERENCES `user`(`id`) ON DELETE CASCADE;

ALTER TABLE `category` ADD CONSTRAINT `fk_category_parent` 
FOREIGN KEY (`parent_id`) REFERENCES `category`(`id`) ON DELETE SET NULL;
