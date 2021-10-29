-- MySQL dump 10.13  Distrib 5.7.37, for Linux (x86_64)
--
-- Host: localhost    Database: myapp
-- ------------------------------------------------------
-- Server version       5.7.37

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Current Database: `myapp`
--

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `myapp` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `myapp`;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `created_time` datetime(3) NOT NULL DEFAULT '0001-01-01 00:00:00.000' COMMENT '创建时间',
  `created_by` int(11) NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `updated_time` datetime(3) NOT NULL DEFAULT '0001-01-01 00:00:00.000' COMMENT '更新时间',
  `updated_by` int(11) NOT NULL DEFAULT '0' COMMENT '更新者ID',
  `deleted_time` datetime(3) NOT NULL DEFAULT '0001-01-01 00:00:00.000' COMMENT '删除时间',
  `deleted_by` int(11) NOT NULL DEFAULT '0' COMMENT '删除者ID',
  `deleted` int(11) NOT NULL DEFAULT '0' COMMENT '是否已删除，0：否，1：是',
  `version` int(11) NOT NULL DEFAULT '0' COMMENT '版本号',
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '姓名',
  `nick_name` varchar(100) NOT NULL DEFAULT '' COMMENT '昵称',
  `avatar` varchar(500) NOT NULL DEFAULT '' COMMENT '头像URL',
  `phone` varchar(50) NOT NULL DEFAULT '' COMMENT '手机号',
  `status` tinyint(3) NOT NULL DEFAULT '0' COMMENT '状态，0：正常',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='用户档案信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'0001-01-01 00:00:00.000',0,'0001-01-01 00:00:00.000',0,'0001-01-01 00:00:00.000',0,0,0,'admin','admin','','',0);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_account`
--

DROP TABLE IF EXISTS `user_account`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_account` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `created_time` datetime(3) NOT NULL DEFAULT '0001-01-01 00:00:00.000' COMMENT '创建时间',
  `created_by` int(11) NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `updated_time` datetime(3) NOT NULL DEFAULT '0001-01-01 00:00:00.000' COMMENT '更新时间',
  `updated_by` int(11) NOT NULL DEFAULT '0' COMMENT '更新者ID',
  `deleted_time` datetime(3) NOT NULL DEFAULT '0001-01-01 00:00:00.000' COMMENT '删除时间',
  `deleted_by` int(11) NOT NULL DEFAULT '0' COMMENT '删除者ID',
  `deleted` int(11) NOT NULL DEFAULT '0' COMMENT '是否已删除，0：否，1：是',
  `version` int(11) NOT NULL DEFAULT '0' COMMENT '版本号',
  `account` varchar(50) NOT NULL DEFAULT '' COMMENT '帐户名',
  `password` varchar(100) NOT NULL DEFAULT '' COMMENT '密码',
  `salt` varchar(100) NOT NULL DEFAULT '' COMMENT '盐',
  `status` tinyint(3) NOT NULL DEFAULT '0' COMMENT '状态，0：正常',
  `uid` bigint(20) NOT NULL DEFAULT '0' COMMENT '关联的用户ID',
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_account_account_IDX` (`account`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='帐户';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_account`
--

LOCK TABLES `user_account` WRITE;
/*!40000 ALTER TABLE `user_account` DISABLE KEYS */;
INSERT INTO `user_account` VALUES (1,'0001-01-01 00:00:00.000',0,'0001-01-01 00:00:00.000',0,'0001-01-01 00:00:00.000',0,0,0,'admin','92795f1ff549d008eadf1a4de4eaccfd','',0,1);
/*!40000 ALTER TABLE `user_account` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_login_token`
--

DROP TABLE IF EXISTS `user_login_token`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_login_token` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `created_time` datetime(3) NOT NULL DEFAULT '0001-01-01 00:00:00.000' COMMENT '创建时间',
  `created_by` int(11) NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `updated_time` datetime(3) NOT NULL DEFAULT '0001-01-01 00:00:00.000' COMMENT '更新时间',
  `updated_by` int(11) NOT NULL DEFAULT '0' COMMENT '更新者ID',
  `deleted_time` datetime(3) NOT NULL DEFAULT '0001-01-01 00:00:00.000' COMMENT '删除时间',
  `deleted_by` int(11) NOT NULL DEFAULT '0' COMMENT '删除者ID',
  `deleted` int(11) NOT NULL DEFAULT '0' COMMENT '是否已删除，0：否，1：是',
  `version` int(11) NOT NULL DEFAULT '0' COMMENT '版本号',
  `uid` bigint(20) NOT NULL DEFAULT '0' COMMENT '关联的用户ID',
  `token` varchar(3000) NOT NULL DEFAULT '' COMMENT '密码',
  `expire_at` datetime(3) NOT NULL DEFAULT '0001-01-01 00:00:00.000' COMMENT '过期时间',
  `status` tinyint(3) NOT NULL DEFAULT '0' COMMENT '状态，0：正常',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='登录令牌';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_login_token`
--

LOCK TABLES `user_login_token` WRITE;
/*!40000 ALTER TABLE `user_login_token` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_login_token` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `dicts`
--

DROP TABLE IF EXISTS `dicts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `dicts` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键 主键ID',
  `created_by` int(11) NOT NULL DEFAULT '0' COMMENT '创建人 创建用户ID',
  `created_time` datetime(3) NOT NULL DEFAULT '0001-01-01 00:00:00.000' COMMENT '创建时间',
  `updated_by` int(11) NOT NULL DEFAULT '0' COMMENT '更新人 最后修改用户ID',
  `updated_time` datetime(3) NOT NULL DEFAULT '0001-01-01 00:00:00.000' COMMENT '更新时间',
  `deleted_by` int(11) NOT NULL DEFAULT '0' COMMENT '删除人 删除用户ID',
  `deleted_time` datetime(3) NOT NULL DEFAULT '0001-01-01 00:00:00.000' COMMENT '删除时间',
  `deleted` tinyint(4) NOT NULL DEFAULT '0' COMMENT '删除标识',
  `version` int(11) NOT NULL DEFAULT '0' COMMENT '版本号',
  `tenant_id` int(11) NOT NULL DEFAULT '0' COMMENT '租户ID',
  `name` varchar(1000) NOT NULL COMMENT '字典项的名称',
  `code` varchar(68) NOT NULL DEFAULT '' COMMENT '字典项的编码',
  `ordinal` int(11) NOT NULL DEFAULT '0' COMMENT '字典项的序号',
  `remark` varchar(128) NOT NULL DEFAULT '' COMMENT '字典项的说明',
  `category` varchar(32) NOT NULL DEFAULT '' COMMENT '字典项的类别',
  `category_desc` varchar(128) NOT NULL DEFAULT '' COMMENT '字典项的类别的说明',
  `parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '父级字典项的ID，0代表无父级字典。',
  PRIMARY KEY (`id`),
  UNIQUE KEY `dicts_UN` (`code`,`deleted_time`,`tenant_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='通用字典表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `dicts`
--

LOCK TABLES `dicts` WRITE;
/*!40000 ALTER TABLE `dicts` DISABLE KEYS */;
/*!40000 ALTER TABLE `dicts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `upload`
--

DROP TABLE IF EXISTS `upload`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `upload` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `created_time` datetime(3) NOT NULL DEFAULT '0001-01-01 00:00:00.000' COMMENT '创建时间',
  `created_by` int(11) NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `updated_time` datetime(3) NOT NULL DEFAULT '0001-01-01 00:00:00.000' COMMENT '更新时间',
  `updated_by` int(11) NOT NULL DEFAULT '0' COMMENT '更新者ID',
  `deleted_time` datetime(3) NOT NULL DEFAULT '0001-01-01 00:00:00.000' COMMENT '删除时间',
  `deleted_by` int(11) NOT NULL DEFAULT '0' COMMENT '删除者ID',
  `deleted` int(11) NOT NULL DEFAULT '0' COMMENT '是否已删除，0：否，1：是',
  `version` int(11) NOT NULL DEFAULT '0' COMMENT '版本号',
  `file_name` varchar(200) NOT NULL DEFAULT '' COMMENT '文件名',
  `file_size` int(11) NOT NULL DEFAULT '0' COMMENT '文件大小',
  `upload_offset` int(11) NOT NULL DEFAULT '0' COMMENT '下一分块偏移量',
  `is_complete` tinyint(3) NOT NULL DEFAULT '0' COMMENT '整个文件（不是当前分块）是否上传完成，0：否，1：是',
  `file_path` varchar(1000) NOT NULL DEFAULT '' COMMENT '文件保存路径，绝对路径',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文件上传记录表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `upload`
--

LOCK TABLES `upload` WRITE;
/*!40000 ALTER TABLE `upload` DISABLE KEYS */;
/*!40000 ALTER TABLE `upload` ENABLE KEYS */;
UNLOCK TABLES;

/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-02-10  2:59:02