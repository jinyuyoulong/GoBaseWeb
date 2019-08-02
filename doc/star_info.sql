/*
 Navicat Premium Data Transfer

 Source Server         : local_root
 Source Server Type    : MySQL
 Source Server Version : 80015
 Source Host           : localhost
 Source Database       : superstar

 Target Server Type    : MySQL
 Target Server Version : 80015
 File Encoding         : utf-8

 Date: 08/02/2019 13:12:44 PM
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `star_info`
-- ----------------------------
DROP TABLE IF EXISTS `star_info`;
CREATE TABLE `star_info` (
  `id` int(10) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name_zh` varchar(50) NOT NULL COMMENT '中文名',
  `name_en` varchar(50) NOT NULL COMMENT '英文名',
  `avatar` varchar(255) NOT NULL COMMENT '头像',
  `birthday` varchar(50) NOT NULL COMMENT '出生日期',
  `height` int(10) NOT NULL COMMENT '身高，单位cm',
  `weight` int(10) NOT NULL COMMENT '体重，单位g',
  `club` varchar(50) NOT NULL COMMENT '俱乐部',
  `jersy` varchar(50) NOT NULL COMMENT '球衣号码以及主打位置',
  `country` varchar(50) NOT NULL COMMENT '国籍',
  `birthaddress` varchar(255) NOT NULL COMMENT '出生地',
  `feature` varchar(255) NOT NULL COMMENT '个人特点',
  `moreinfo` text NOT NULL COMMENT '更多介绍',
  `sys_status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态，默认值 0 正常，1 删除',
  `sys_created` int(10) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `sys_updated` int(10) NOT NULL DEFAULT '0' COMMENT '最后修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 COMMENT='球星信息';

SET FOREIGN_KEY_CHECKS = 1;
