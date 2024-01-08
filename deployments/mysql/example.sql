
-- MySQL dump 10.13  Distrib 5.7.36, for Linux (x86_64)
--
-- ------------------------------------------------------
-- Server version	5.7.36-0ubuntu0.18.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/ `example` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

USE `example`;


CREATE TABLE `table_template` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,

  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '名称',

  `created_time` datetime(3) NOT NULL DEFAULT '0001-01-01 00:00:00.000' COMMENT '创建时间',
  `created_by` bigint(20) NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `updated_time` datetime(3) NOT NULL DEFAULT '0001-01-01 00:00:00.000' COMMENT '更新时间',
  `updated_by` bigint(20) NOT NULL DEFAULT '0' COMMENT '更新者ID',
  `deleted_time` datetime(3) NOT NULL DEFAULT '0001-01-01 00:00:00.000' COMMENT '删除时间',
  `deleted_by` bigint(20) NOT NULL DEFAULT '0' COMMENT '删除者ID',
  UNIQUE KEY `table_template|name-deleted_time-UIDX` (`name`, `deleted_time`) USING BTREE,
  KEY `table_template|created_time-IDX` (`created_time`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='建表模板';
LOCK TABLES `table_template` WRITE;
INSERT INTO `table_template` (id, name) VALUES (1,'建表模板，仅作参考');
UNLOCK TABLES;


CREATE TABLE `user_account` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,

  `account` varchar(50) NOT NULL DEFAULT '' COMMENT '帐户名',
  `password` varchar(512) NOT NULL DEFAULT '' COMMENT '哈希后的密码',
  `salt` varchar(64) NOT NULL DEFAULT '0' COMMENT '密码哈希过程中加的盐',
  `status` tinyint(3) NOT NULL DEFAULT '0' COMMENT '状态，0：禁用, 1：启用',
  `tenant_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '租户ID',
  `is_admin` tinyint(3) NOT NULL DEFAULT '0' COMMENT '是否是管理员，0：否, 1：是',
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '名称',
  `avatar` varchar(500) NOT NULL DEFAULT '' COMMENT '头像URL',
  `phone` varchar(30) NOT NULL DEFAULT '' COMMENT '手机号',
  `organization` varchar(500) NOT NULL DEFAULT '' COMMENT '组织',
  `gender` tinyint(3) NOT NULL DEFAULT 0 COMMENT '性别，0：未知，1：男，2：女',

  `created_time` datetime(3) NOT NULL DEFAULT '0001-01-01 00:00:00.000' COMMENT '创建时间',
  `created_by` bigint(20) NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `updated_time` datetime(3) NOT NULL DEFAULT '0001-01-01 00:00:00.000' COMMENT '更新时间',
  `updated_by` bigint(20) NOT NULL DEFAULT '0' COMMENT '更新者ID',
  `deleted_time` datetime(3) NOT NULL DEFAULT '0001-01-01 00:00:00.000' COMMENT '删除时间',
  `deleted_by` bigint(20) NOT NULL DEFAULT '0' COMMENT '删除者ID',
  UNIQUE KEY `user_account|account-UIDX` (`account`) USING BTREE,
  KEY `user_account|tenant_id-IDX` (`tenant_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户帐户';
LOCK TABLES `user_account` WRITE;
INSERT INTO `user_account` (id, account, password, `status`, tenant_id, is_admin) VALUES (1,'admin','12345678', 1,1,1);
INSERT INTO `user_account` (id, account, password, `status`, tenant_id, is_admin) VALUES (2,'uploader','12345678', 1,1,0);
UNLOCK TABLES;
