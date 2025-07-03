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
    `role` ENUM('user', 'admin', 'moderator') NOT NULL DEFAULT 'user' COMMENT 'role',
    `status` ENUM('active', 'inactive') NOT NULL DEFAULT 'active' COMMENT 'status',
    `is_deleted` TINYINT(1) NOT NULL DEFAULT 0,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='user table';

-- 帖子表
CREATE TABLE IF NOT EXISTS `post` (
    `id` bigint AUTO_INCREMENT PRIMARY KEY COMMENT 'post id',
    `title` VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'title',
    `content` TEXT NOT NULL COMMENT 'content',
    `user_id` BIGINT NOT NULL COMMENT 'user id',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created time',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'updated time',
    INDEX idx_user_id (user_id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='post table';

-- 评论表
CREATE TABLE IF NOT EXISTS `comment` (
    `id` bigint AUTO_INCREMENT PRIMARY KEY COMMENT 'comment id',
    `content` TEXT NOT NULL COMMENT 'content',
    `user_id` BIGINT NOT NULL COMMENT 'user id',
    `post_id` BIGINT NOT NULL COMMENT 'post id',
    `parent_id` BIGINT DEFAULT NULL COMMENT 'parent comment id',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created time',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'updated time',
    INDEX idx_post_id (post_id),
    INDEX idx_parent_id (parent_id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='comment table';
