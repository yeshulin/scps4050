/*
Navicat MySQL Data Transfer

Source Server         : 172.29.4.87
Source Server Version : 50629
Source Host           : 172.29.4.87:3306
Source Database       : 4050

Target Server Type    : MYSQL
Target Server Version : 50629
File Encoding         : 65001

Date: 2017-05-27 18:47:02
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `members`
-- ----------------------------
DROP TABLE IF EXISTS `members`;
CREATE TABLE `members` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `openid` varchar(200) DEFAULT NULL COMMENT '微信ID',
  `username` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `realname` varchar(255) DEFAULT NULL,
  `avatarurl` varchar(255) DEFAULT NULL COMMENT '用户头像',
  `sex` varchar(50) DEFAULT NULL COMMENT '性别',
  `bothtime` varchar(50) DEFAULT NULL COMMENT '出生时间',
  `zone` tinyint(4) DEFAULT '0' COMMENT '所属区域',
  `address` varchar(255) DEFAULT NULL COMMENT '联系地址',
  `email` varchar(255) DEFAULT NULL,
  `phone` varchar(255) DEFAULT NULL,
  `workaddress` varchar(255) DEFAULT NULL COMMENT '灵活就业地址',
  `worktype` varchar(255) DEFAULT NULL COMMENT '灵活就业形式',
  `isverify` int(4) DEFAULT NULL COMMENT '是否审核通过用户',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注 ',
  `addtime` int(11) DEFAULT NULL,
  `updatetime` int(11) DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=MyISAM AUTO_INCREMENT=93 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of members
-- ----------------------------
INSERT INTO `members` VALUES ('40', null, 'yeshulin1', 'RE7+OOlq', '叶树林', null, null, null, null, null, '1914875404@qq.com', '13512341234', null, null, null, null, '1494651921', '1494651921');
INSERT INTO `members` VALUES ('60', null, 'liqing', 'RE7+OOlq', '李清', null, '', '', '1', '', '1914875404@qq.com', '13512344321', '', '', '0', '', '1495510178', '1495510178');
INSERT INTO `members` VALUES ('61', null, 'apply', 'RE7+OOlq', '申请人员', null, '', '', '2', '', '1914875404@qq.com', '13512341234', '', '', '-1', '343443', '1495529419', '1495529419');
INSERT INTO `members` VALUES ('80', '', '', 'RE7+OOlq', '', '', '', '', '0', '', '', '', '', '', '0', '', '1495878975', '1495878975');
INSERT INTO `members` VALUES ('81', '', '', 'RE7+OOlq', '', '', '', '', '0', '', '', '', '', '', '0', '', '1495879041', '1495879041');
INSERT INTO `members` VALUES ('82', '', '', 'RE7+OOlq', '', '', '', '', '0', '', '', '', '', '', '0', '', '1495879214', '1495879214');
INSERT INTO `members` VALUES ('83', '', '', 'RE7+OOlq', '', '', '', '', '0', '', '', '', '', '', '0', '', '1495879283', '1495879283');
INSERT INTO `members` VALUES ('84', '', '', 'RE7+OOlq', '', '', '', '', '0', '', '', '', '', '', '0', '', '1495879700', '1495879700');
INSERT INTO `members` VALUES ('85', '', '', 'RE7+OOlq', '', '', '', '', '0', '', '', '', '', '', '0', '', '1495879730', '1495879730');
INSERT INTO `members` VALUES ('86', '', '', 'RE7+OOlq', '', '', '', '', '0', '', '', '', '', '', '0', '', '1495879737', '1495879737');
INSERT INTO `members` VALUES ('87', '', '', 'RE7+OOlq', '', '', '', '', '0', '', '', '', '', '', '0', '', '1495881177', '1495881177');
INSERT INTO `members` VALUES ('88', '', '', 'RE7+OOlq', '', '', '', '', '0', '', '', '', '', '', '0', '', '1495881206', '1495881206');
INSERT INTO `members` VALUES ('89', '', '', 'RE7+OOlq', '', '', '', '', '0', '', '', '', '', '', '0', '', '1495881245', '1495881245');
INSERT INTO `members` VALUES ('90', '', '', 'RE7+OOlq', '', '', '', '', '0', '', '', '', '', '', '0', '', '1495881250', '1495881250');
INSERT INTO `members` VALUES ('91', 'oSEQM0Y_zAMtRsj0rmBgfh001rEI', '', 'RE7+OOlq', '', '', '', '', '0', '', '', '', '', '', '0', '', '1495881478', '1495881478');
INSERT INTO `members` VALUES ('92', 'oSEQM0Y_zAMtRsj0rmBgfh001rEI', '', 'RE7+OOlq', '', '', '', '', '0', '', '', '', '', '', '0', '', '1495881487', '1495881487');

-- ----------------------------
-- Table structure for `news`
-- ----------------------------
DROP TABLE IF EXISTS `news`;
CREATE TABLE `news` (
  `id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT,
  `catid` smallint(5) unsigned DEFAULT '0' COMMENT '分类ID',
  `title` varchar(80) NOT NULL DEFAULT '' COMMENT '标题',
  `thumb` varchar(100) NOT NULL DEFAULT '' COMMENT '缩略图',
  `keywords` char(40) NOT NULL DEFAULT '' COMMENT '关键词',
  `description` mediumtext NOT NULL COMMENT '描述',
  `content` text COMMENT '内容',
  `sort` tinyint(3) unsigned DEFAULT '0' COMMENT '排序',
  `status` tinyint(2) unsigned NOT NULL DEFAULT '1' COMMENT '状态',
  `username` char(20) NOT NULL COMMENT '添加用户',
  `addtime` int(11) unsigned NOT NULL DEFAULT '0',
  `updatetime` int(11) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `status` (`status`,`sort`,`id`),
  KEY `listorder` (`catid`,`status`,`sort`,`id`),
  KEY `catid` (`catid`,`status`,`id`)
) ENGINE=MyISAM AUTO_INCREMENT=1050 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of news
-- ----------------------------
INSERT INTO `news` VALUES ('1047', '1', '测试内容1', '', '测试内容2', '测试内容3343443', '<p>测试内容4</p>', '1', '1', '', '1495098531', '1495103815');
INSERT INTO `news` VALUES ('1049', '1', '转转为什么没有变得更好？', '23.jpg', '交网络', '的社交网络也确实有利于转转用户之间进行交易，但这种交易更多是基于熟人社交关系之间，但二手交易最根本核心在于撬动闲置物品需求，而这绝不是一种熟人关系的交易，就连姚劲波自己也承认了这个观点。\n\n姚劲波曾撰文谈为何', '<p class=\"text\">4月18日，58同城宣布与腾讯达成协议，后者将向58集团旗下的二手交易平台「转转」投资2亿美元。腾讯的入局加剧了二手交易市场的竞争局势，在一个月后的今天，或许我们是时候来复盘了。</p><p class=\"text\"><span>营收停止增长后，转转成为一根救命稻草</span></p><p class=\"text\"><br></p>', '2', '1', '', '1495104001', '1495105292');

-- ----------------------------
-- Table structure for `newstype`
-- ----------------------------
DROP TABLE IF EXISTS `newstype`;
CREATE TABLE `newstype` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `catname` varchar(50) DEFAULT NULL COMMENT '分类名称',
  `pid` int(11) DEFAULT '0' COMMENT '父ID',
  `sort` tinyint(4) DEFAULT NULL,
  `addtime` int(11) DEFAULT NULL,
  `updatetime` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of newstype
-- ----------------------------
INSERT INTO `newstype` VALUES ('1', '新闻动态', '0', '1', '1495028014', '1495028014');

-- ----------------------------
-- Table structure for `node`
-- ----------------------------
DROP TABLE IF EXISTS `node`;
CREATE TABLE `node` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(20) DEFAULT NULL COMMENT '节点名称',
  `title` varchar(20) DEFAULT NULL COMMENT '节点标题',
  `status` tinyint(4) DEFAULT NULL COMMENT '节点状态',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `sort` tinyint(4) DEFAULT NULL COMMENT '排序',
  `pid` tinyint(4) DEFAULT '0' COMMENT '父ID',
  `level` tinyint(4) DEFAULT NULL COMMENT '级别',
  `type` varchar(20) DEFAULT NULL COMMENT '类型',
  `group_id` tinyint(4) DEFAULT NULL COMMENT '分组',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of node
-- ----------------------------
INSERT INTO `node` VALUES ('1', 'user', '用户管理', '1', '用户管理', '1', '0', '0', 'model', '0');
INSERT INTO `node` VALUES ('2', 'role', '角色管理', '1', '角色管理', '2', '0', '0', 'model', '0');

-- ----------------------------
-- Table structure for `node_access`
-- ----------------------------
DROP TABLE IF EXISTS `node_access`;
CREATE TABLE `node_access` (
  `role_id` int(11) DEFAULT NULL,
  `node_id` int(11) DEFAULT NULL,
  `level` tinyint(4) DEFAULT NULL,
  `pid` tinyint(4) DEFAULT NULL,
  `module` varchar(10) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of node_access
-- ----------------------------

-- ----------------------------
-- Table structure for `role`
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(20) DEFAULT NULL COMMENT '角色名称',
  `status` tinyint(4) DEFAULT NULL COMMENT '状态',
  `remark` varchar(100) DEFAULT NULL COMMENT 'P备注',
  `addtime` int(11) DEFAULT NULL,
  `updatetime` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of role
-- ----------------------------
INSERT INTO `role` VALUES ('1', '系统管理员', '1', '系统管理员', '1494643853', '1494643853');
INSERT INTO `role` VALUES ('2', '乡镇管理员', '1', '乡镇管理员', '1494650587', '1494650587');

-- ----------------------------
-- Table structure for `role_member`
-- ----------------------------
DROP TABLE IF EXISTS `role_member`;
CREATE TABLE `role_member` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_id` int(11) DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of role_member
-- ----------------------------
INSERT INTO `role_member` VALUES ('1', '1', '40');
INSERT INTO `role_member` VALUES ('18', '1', '60');

-- ----------------------------
-- Table structure for `signs`
-- ----------------------------
DROP TABLE IF EXISTS `signs`;
CREATE TABLE `signs` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `years` varchar(10) DEFAULT NULL COMMENT '年度',
  `months` varchar(10) DEFAULT NULL COMMENT '月度',
  `userid` int(11) DEFAULT NULL COMMENT '签到用户',
  `postion` varchar(255) DEFAULT NULL COMMENT '签到位置',
  `photos` varchar(255) DEFAULT NULL COMMENT '照片',
  `isverify` tinyint(4) DEFAULT NULL COMMENT '是否审核',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `addtime` int(11) DEFAULT NULL COMMENT '签到时间',
  `updatetime` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of signs
-- ----------------------------
INSERT INTO `signs` VALUES ('1', '2323', '', '0', '', '', '-1', '3434342', '0', '1495455495');
INSERT INTO `signs` VALUES ('9', '', '', '0', '', '', '0', '', '1495787124', '1495787124');
INSERT INTO `signs` VALUES ('10', '', '', '0', '', '', '0', '', '1495787168', '1495787168');
INSERT INTO `signs` VALUES ('11', '', '', '0', '', '', '0', '', '1495787198', '1495787198');
INSERT INTO `signs` VALUES ('12', '', '', '0', '', '', '0', '', '1495787204', '1495787204');
INSERT INTO `signs` VALUES ('13', '', '', '0', '', '', '0', '', '1495787226', '1495787226');
INSERT INTO `signs` VALUES ('14', '', '', '0', '', '', '0', '', '1495787245', '1495787245');
INSERT INTO `signs` VALUES ('15', '', '', '0', '', '', '0', '', '1495787288', '1495787288');
INSERT INTO `signs` VALUES ('16', '', '', '0', '', '', '0', '', '1495787304', '1495787304');
INSERT INTO `signs` VALUES ('17', '', '', '0', '', '', '0', '', '1495787399', '1495787399');
INSERT INTO `signs` VALUES ('18', '', '', '0', '', '', '0', '', '1495787411', '1495787411');
INSERT INTO `signs` VALUES ('19', '', '', '0', '', '', '0', '', '1495787488', '1495787488');
INSERT INTO `signs` VALUES ('20', '', '', '0', '', '', '0', '', '1495787581', '1495787581');
INSERT INTO `signs` VALUES ('21', '', '', '0', '', '', '0', '', '1495787687', '1495787687');
INSERT INTO `signs` VALUES ('22', '', '', '0', '', '', '0', '', '1495787702', '1495787702');
INSERT INTO `signs` VALUES ('23', '', '', '0', '', '', '0', '', '1495787731', '1495787731');
INSERT INTO `signs` VALUES ('24', '', '', '0', '', '', '0', '', '1495787807', '1495787807');
INSERT INTO `signs` VALUES ('25', '', '', '0', '', '', '0', '', '1495787837', '1495787837');
INSERT INTO `signs` VALUES ('26', '', '', '0', '', '', '0', '', '1495787855', '1495787855');

-- ----------------------------
-- Table structure for `zones`
-- ----------------------------
DROP TABLE IF EXISTS `zones`;
CREATE TABLE `zones` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `zonename` varchar(255) DEFAULT NULL COMMENT '乡镇名称',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of zones
-- ----------------------------
INSERT INTO `zones` VALUES ('1', '锦江乡');
INSERT INTO `zones` VALUES ('2', '观音镇');
INSERT INTO `zones` VALUES ('3', '江口镇');
INSERT INTO `zones` VALUES ('4', '谢家镇');
INSERT INTO `zones` VALUES ('5', '保胜乡');
INSERT INTO `zones` VALUES ('6', '武阳乡');
INSERT INTO `zones` VALUES ('7', '义和乡');
INSERT INTO `zones` VALUES ('8', '公义镇');
INSERT INTO `zones` VALUES ('9', '凤鸣镇');
INSERT INTO `zones` VALUES ('10', '青龙镇');
INSERT INTO `zones` VALUES ('11', '牧马镇');
INSERT INTO `zones` VALUES ('12', '彭溪镇');
INSERT INTO `zones` VALUES ('13', '黄丰镇');
