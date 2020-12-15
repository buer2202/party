/*
Navicat MySQL Data Transfer

Source Server         : 10.0.31.254
Source Server Version : 50730
Source Host           : 10.0.31.254:3306
Source Database       : party

Target Server Type    : MYSQL
Target Server Version : 50730
File Encoding         : 65001

Date: 2020-12-15 16:46:24
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for member
-- ----------------------------
DROP TABLE IF EXISTS `member`;
CREATE TABLE `member` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL COMMENT '用户ID',
  `nickname` varchar(10) NOT NULL COMMENT '成员昵称',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COMMENT='用户成员表';

-- ----------------------------
-- Table structure for party
-- ----------------------------
DROP TABLE IF EXISTS `party`;
CREATE TABLE `party` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL,
  `url_code` varchar(20) NOT NULL COMMENT '活动码，用于url链接',
  `name` varchar(50) NOT NULL COMMENT '活动名',
  `party_desc` varchar(255) NOT NULL COMMENT '活动发起描述',
  `confirm_desc` varchar(255) NOT NULL DEFAULT '' COMMENT '活动确认描述',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_url_code` (`url_code`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COMMENT='用户活动表';

-- ----------------------------
-- Table structure for party_member
-- ----------------------------
DROP TABLE IF EXISTS `party_member`;
CREATE TABLE `party_member` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `party_id` int(10) unsigned NOT NULL COMMENT '活动ID',
  `member_nickname` varchar(10) NOT NULL COMMENT '成员昵称',
  `join_people_num` tinyint(4) NOT NULL DEFAULT '0' COMMENT '参加人数',
  `can_join_date` date NOT NULL COMMENT '可参加时间',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=82 DEFAULT CHARSET=utf8mb4 COMMENT='活动成员参与表';

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `account` varchar(10) NOT NULL COMMENT '登录账号',
  `password` varchar(255) NOT NULL COMMENT '登录密码',
  `nickname` varchar(20) NOT NULL DEFAULT '' COMMENT '昵称',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';
