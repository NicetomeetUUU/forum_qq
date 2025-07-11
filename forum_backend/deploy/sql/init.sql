SET NAMES utf8mb4;
SET CHARACTER SET utf8mb4;

CREATE DATABASE IF NOT EXISTS qq_forum DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE qq_forum;

-- 用户表
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
    `id` BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT 'auto_user_id',
    `email` VARCHAR(255) NOT NULL UNIQUE DEFAULT '' COMMENT 'email',
    `password` VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'password',
    `username` VARCHAR(255) NOT NULL UNIQUE DEFAULT '' COMMENT 'username',
    `avatar` VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'avatar url',
    `signature` VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'signature',
    `status` ENUM('active', 'inactive') NOT NULL DEFAULT 'active' COMMENT 'status',
    `last_login_time` DATETIME DEFAULT NULL COMMENT 'last login time',
    `created_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_email (email),
    INDEX idx_username (username)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='user table';

-- 管理员表
DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin` (
    `id` BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT 'auto_admin_id',
    `email` VARCHAR(255) NOT NULL UNIQUE DEFAULT '' COMMENT 'email',
    `password` VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'password',
    `username` VARCHAR(255) NOT NULL UNIQUE DEFAULT '' COMMENT 'username',
    `status` ENUM('active', 'inactive') NOT NULL DEFAULT 'active' COMMENT 'status',
    `last_login_time` DATETIME DEFAULT NULL COMMENT 'last login time',
    `created_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_email (email),
    INDEX idx_username (username)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='admin table';

-- 分类表
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
    `id` BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT 'auto_category_id',
    `name` VARCHAR(100) NOT NULL UNIQUE COMMENT 'category name',
    `description` TEXT COMMENT 'category description',
    `status` ENUM('active', 'inactive') NOT NULL DEFAULT 'active' COMMENT 'status',
    `created_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_name (name)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='category table';

-- 帖子表
DROP TABLE IF EXISTS `post`;
CREATE TABLE `post` (
    `id` BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT 'auto_post_id',
    `title` VARCHAR(255) NOT NULL UNIQUE COMMENT 'title',
    `content` TEXT NOT NULL COMMENT 'content',
    `user_id` BIGINT NOT NULL COMMENT 'user id',
    `view_count` INT NOT NULL DEFAULT 0 COMMENT 'view count',
    `like_count` INT NOT NULL DEFAULT 0 COMMENT 'like count',
    `comment_count` INT NOT NULL DEFAULT 0 COMMENT 'comment count',
    `status` ENUM('published', 'hidden') NOT NULL DEFAULT 'published' COMMENT 'status',
    `created_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created time',
    `updated_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'updated time',
    INDEX idx_user_id (user_id),
    INDEX idx_created_time (created_time)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='post table';

-- 帖子类别表
DROP TABLE IF EXISTS `post_category`;
CREATE TABLE `post_category` (
    `id` BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT 'auto_post_category_id',
    `post_id` BIGINT NOT NULL COMMENT 'post id',
    `category_id` BIGINT NOT NULL COMMENT 'category id',
    `created_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created time',
    UNIQUE KEY uk_post_category (post_id, category_id),
    INDEX idx_post_id (post_id),
    INDEX idx_category_id (category_id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='post category table';

-- 评论表
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
    `id` BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT 'auto_comment_id',
    `content` TEXT NOT NULL COMMENT 'content',
    `user_id` BIGINT NOT NULL COMMENT 'user id',
    `post_id` BIGINT NOT NULL COMMENT 'post id',
    `parent_id` BIGINT DEFAULT NULL COMMENT 'parent comment id',
    `like_count` INT NOT NULL DEFAULT 0 COMMENT 'like count',
    `status` ENUM('published', 'hidden') NOT NULL DEFAULT 'published' COMMENT 'status',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created time',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'updated time',
    INDEX idx_post_id (post_id),
    INDEX idx_parent_id (parent_id),
    INDEX idx_user_id (user_id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='comment table';

-- 用户点赞表
DROP TABLE IF EXISTS `user_like`;
CREATE TABLE `user_like` (
    `id` BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT 'auto_like_id',
    `user_id` BIGINT NOT NULL COMMENT 'user id',
    `target_type` ENUM('post', 'comment') NOT NULL COMMENT 'target type',
    `target_id` BIGINT NOT NULL COMMENT 'target id',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE KEY uk_user_target (user_id, target_type, target_id),
    INDEX idx_target (target_type, target_id),
    INDEX idx_user_id (user_id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='user like table';

-- 用户关注表
DROP TABLE IF EXISTS `user_follow`;
CREATE TABLE `user_follow` (
    `id` BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT 'auto_follow_id',
    `follower_id` BIGINT NOT NULL COMMENT 'follower user id',
    `following_id` BIGINT NOT NULL COMMENT 'following user id',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE KEY uk_follow (follower_id, following_id),
    INDEX idx_follower (follower_id),
    INDEX idx_following (following_id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='user follow table';

INSERT INTO `category` (`name`, `description`) VALUES
('默认分类', '新建帖子所在默认分类');