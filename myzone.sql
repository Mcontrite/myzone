/*
Navicat MySQL Data Transfer

Source Server         : myaliyun-xg
Source Server Version : 80015
Source Host           : 47.91.212.4:3306
Source Database       : myzone

Target Server Type    : MYSQL
Target Server Version : 80015
File Encoding         : 65001
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for z_attach
-- ----------------------------
DROP TABLE IF EXISTS `z_attach`;
CREATE TABLE `z_attach` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '附件id',
  `article_id` int(11) NOT NULL DEFAULT '0' COMMENT '文章id',
  `reply_id` int(11) NOT NULL DEFAULT '0' COMMENT '回复id',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户id',
  `filesize` int(8) unsigned NOT NULL DEFAULT '0' COMMENT '文件尺寸，单位字节',
  `width` mediumint(8) unsigned NOT NULL DEFAULT '0' COMMENT 'width > 0 则为图片',
  `height` mediumint(8) unsigned NOT NULL DEFAULT '0',
  `filename` char(120) NOT NULL DEFAULT '' COMMENT '文件名称，会过滤，并且截断，保存后的文件名，不包含URL前缀 upload_url',
  `orgfilename` char(120) NOT NULL DEFAULT '' COMMENT '上传的原文件名',
  `filetype` char(7) NOT NULL DEFAULT '' COMMENT 'image/txt/zip，小图标显示',
  `comment` char(100) NOT NULL DEFAULT '' COMMENT '文件注释 方便于搜索',
  `downloads_cnt` int(11) NOT NULL DEFAULT '0' COMMENT '下载次数',
  `isimage` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否为图片',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `downloads_num` int(11) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `pid` (`reply_id`),
  KEY `uid` (`user_id`),
  KEY `idx_z_attach_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='论坛附件表  只能按照从上往下的方式查找和删除！ 此表如果大，可以考虑通过 aid 分区。';


-- ----------------------------
-- Table structure for z_cache
-- ----------------------------
DROP TABLE IF EXISTS `z_cache`;
CREATE TABLE `z_cache` (
  `k` char(32) NOT NULL DEFAULT '',
  `v` mediumtext NOT NULL,
  `expiry` int(11) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`k`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='缓存表，用来保存临时数据';

-- ----------------------------
-- Records of z_cache
-- ----------------------------


-- ----------------------------
-- Table structure for z_group
-- ----------------------------
DROP TABLE IF EXISTS `z_group`;
CREATE TABLE `z_group` (
  `id` smallint(6) unsigned NOT NULL,
  `name` char(20) NOT NULL DEFAULT '' COMMENT '用户组名称',
  `allowread` int(11) NOT NULL DEFAULT '0' COMMENT '允许访问',
  `allowarticle` int(11) NOT NULL DEFAULT '0' COMMENT '允许发文章',
  `allowsaying` int(11) NOT NULL DEFAULT '0' COMMENT '允许发说说',
  `allowreply` int(11) NOT NULL DEFAULT '0' COMMENT '允许回复',
  `allowcomment` int(11) NOT NULL DEFAULT '0' COMMENT '允许评论',
  `allowattach` int(11) NOT NULL DEFAULT '0' COMMENT '允许上传文件',
  `allowdown` int(11) NOT NULL DEFAULT '0' COMMENT '允许下载文件',
  `allowupdate` int(11) NOT NULL DEFAULT '0' COMMENT '允许编辑',
  `allowdelete` int(11) NOT NULL DEFAULT '0',
  `allowdeleteuser` int(11) NOT NULL DEFAULT '0',
  `allowviewip` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '允许查看用户敏感信息',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_z_group_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户组';

-- ----------------------------
-- Records of z_group
-- ----------------------------
INSERT INTO `z_group` VALUES ('0', '游客', '1', '0', '0', '1', '1', '0', '1', '0', '0', '0', '0', null, null, null);
INSERT INTO `z_group` VALUES ('1', '管理员', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', null, null, null);


-- ----------------------------
-- Table structure for z_kv
-- ----------------------------
DROP TABLE IF EXISTS `z_kv`;
CREATE TABLE `z_kv` (
  `k` char(32) NOT NULL DEFAULT '',
  `v` mediumtext NOT NULL,
  `expiry` int(11) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`k`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='持久的 key value 数据存储, ttserver, mysql';


-- ----------------------------
-- Table structure for z_my_favourite
-- ----------------------------
DROP TABLE IF EXISTS `z_my_favourite`;
CREATE TABLE `z_my_favourite` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `user_id` int(11) DEFAULT '0',
  `article_id` int(11) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_z_my_favourite_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;


-- ----------------------------
-- Table structure for z_myreply
-- ----------------------------
DROP TABLE IF EXISTS `z_my_reply`;
CREATE TABLE `z_my_reply` (
  `user_id` int(11) unsigned NOT NULL DEFAULT '0',
  `article_id` int(11) unsigned NOT NULL DEFAULT '0',
  `reply_id` int(11) unsigned NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`user_id`,`reply_id`),
  KEY `tid` (`article_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='我的回复';


-- ----------------------------
-- Table structure for z_mycomment
-- ----------------------------
DROP TABLE IF EXISTS `z_my_comment`;
CREATE TABLE `z_my_comment` (
  `user_id` int(11) unsigned NOT NULL DEFAULT '0',
  `saying_id` int(11) unsigned NOT NULL DEFAULT '0',
  `comment_id` int(11) unsigned NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`user_id`,`comment_id`),
  KEY `tid` (`saying_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='我的回复';


-- ----------------------------
-- Table structure for z_myarticle
-- ----------------------------
DROP TABLE IF EXISTS `z_my_article`;
CREATE TABLE `z_my_article` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) unsigned NOT NULL DEFAULT '0',
  `article_id` int(11) unsigned NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_z_my_article_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='我的文章，每个文章不管回复多少次，只记录一次。大表，需要分区';


-- ----------------------------
-- Table structure for z_mysaying
-- ----------------------------
DROP TABLE IF EXISTS `z_my_saying`;
CREATE TABLE `z_my_saying` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) unsigned NOT NULL DEFAULT '0',
  `saying_id` int(11) unsigned NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_z_my_saying_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='我的文章，每个文章不管回复多少次，只记录一次。大表，需要分区';

-- ----------------------------
-- Table structure for z_reply
-- ----------------------------
DROP TABLE IF EXISTS `z_reply`;
CREATE TABLE `z_reply` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '回复id',
  `article_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '文章id',
  `user_id` int(11) unsigned NOT NULL DEFAULT '0',
  `isfirst` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '是否为首帖，与 article.firstpid 呼应',
  `userip` varchar(50) NOT NULL DEFAULT '0' COMMENT '发帖时用户ip ip2long()',
  `images_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '附件中包含的图片数',
  `files_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '附件中包含的文件数',
  `doctype` tinyint(3) NOT NULL DEFAULT '0' COMMENT '类型，0: html, 1: txt; 2: markdown; 3: ubb',
  `quote_reply_id` int(11) NOT NULL DEFAULT '0' COMMENT '引用哪个 pid，可能不存在',
  `message` longtext NOT NULL COMMENT '内容，用户提示的原始数据',
  `message_fmt` longtext NOT NULL COMMENT '内容，存放的过滤后的html内容，可以定期清理，减肥',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `tid` (`article_id`,`id`),
  KEY `uid` (`user_id`),
  KEY `idx_z_reply_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='论坛回复数据';

-- ----------------------------

-- ----------------------------
-- Table structure for z_comment
-- ----------------------------
DROP TABLE IF EXISTS `z_comment`;
CREATE TABLE `z_comment` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '评论id',
  `saying_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '说说id',
  `user_id` int(11) unsigned NOT NULL DEFAULT '0',
  `isfirst` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '是否为首帖，与 saying.firstpid 呼应',
  `userip` varchar(50) NOT NULL DEFAULT '0' COMMENT '发帖时用户ip ip2long()',
  `quote_comment_id` int(11) NOT NULL DEFAULT '0' COMMENT '引用哪个 pid，可能不存在',
  `message` longtext NOT NULL COMMENT '内容，用户提示的原始数据',
  `message_fmt` longtext NOT NULL COMMENT '内容，存放的过滤后的html内容，可以定期清理，减肥',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `tid` (`saying_id`,`id`),
  KEY `uid` (`user_id`),
  KEY `idx_z_comment_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='论坛评论数据';

-- ----------------------------
-- ----------------------------
-- Table structure for z_queue
-- ----------------------------
DROP TABLE IF EXISTS `z_queue`;
CREATE TABLE `z_queue` (
  `id` int(11) unsigned NOT NULL DEFAULT '0',
  `value` int(11) NOT NULL DEFAULT '0',
  `expiry` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '过期时间',
  UNIQUE KEY `queueid` (`id`,`value`),
  KEY `expiry` (`expiry`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='临时队列，用来保存临时数据';

-- ----------------------------
-- Records of z_queue
-- ----------------------------

-- ----------------------------
-- Table structure for z_session
-- ----------------------------
DROP TABLE IF EXISTS `z_session`;
CREATE TABLE `z_session` (
  `id` char(32) NOT NULL DEFAULT '0' COMMENT '随机生成 id 不能重复 uniqueid() 13 位',
  `user_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '用户id 未登录为 0，可以重复',
  `url` char(32) NOT NULL DEFAULT '' COMMENT '当前访问 url',
  `ip` int(11) unsigned NOT NULL DEFAULT '0',
  `useragent` char(128) NOT NULL DEFAULT '',
  `data` char(255) NOT NULL DEFAULT '',
  `bigdata` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否有大数据',
  `last_date` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '上次活动时间',
  PRIMARY KEY (`id`),
  KEY `ip` (`ip`),
  KEY `uid_last_date` (`user_id`,`last_date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='session 表 缓存到 runtime 表。提高遍历效率\r\n';

-- ----------------------------
-- Records of z_session
-- ----------------------------

-- ----------------------------
-- Table structure for z_session_data
-- ----------------------------
DROP TABLE IF EXISTS `z_session_data`;
CREATE TABLE `z_session_data` (
  `session_id` char(32) NOT NULL DEFAULT '0',
  `last_date` int(11) unsigned NOT NULL DEFAULT '0',
  `data` text NOT NULL,
  PRIMARY KEY (`session_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of z_session_data
-- ----------------------------

-- ----------------------------
-- Table structure for z_table_day
-- ----------------------------
DROP TABLE IF EXISTS `z_table_day`;
CREATE TABLE `z_table_day` (
  `year` smallint(11) unsigned NOT NULL DEFAULT '0' COMMENT '年',
  `month` tinyint(11) unsigned NOT NULL DEFAULT '0' COMMENT '月',
  `day` tinyint(11) unsigned NOT NULL DEFAULT '0' COMMENT '日',
  `create_date` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '时间戳',
  `table` char(16) NOT NULL DEFAULT '' COMMENT '表名',
  `maxid` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '最大ID',
  `count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '总数',
  PRIMARY KEY (`year`,`month`,`day`,`table`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='系统表';

-- ----------------------------
-- Records of z_table_day
-- ----------------------------

-- ----------------------------
-- Table structure for z_article
-- ----------------------------
DROP TABLE IF EXISTS `z_article`;
CREATE TABLE `z_article` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '文章id',
  `user_id` int(11) unsigned NOT NULL DEFAULT '0',
  `userip` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '0' COMMENT '发帖时用户ip ip2long()，主要用来清理',
  `title` char(128) NOT NULL DEFAULT '' COMMENT ' 文章',
  `last_date` timestamp NULL DEFAULT NULL COMMENT '最后回复时间',
  `views_cnt` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '查看次数, 剥离出去，单独的服务，避免 cache 失效',
  `replys_cnt` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '回复数',
  `images_num` tinyint(6) NOT NULL DEFAULT '0' COMMENT '附件中包含的图片数',
  `files_num` tinyint(6) NOT NULL DEFAULT '0' COMMENT '附件中包含的文件数',
  `first_reply_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '首贴 pid',
  `last_user_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '最近参与的 uid',
  `last_reply_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '最后回复的 pid',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `favourite_cnt` int(11) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `lastpid` (`last_reply_id`),
  KEY `fid` (`id`),
  KEY `fid_2` (`last_reply_id`),
  KEY `idx_z_article_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='论坛文章';
-- ----------------------------
-- Table structure for z_saying
-- ----------------------------
DROP TABLE IF EXISTS `z_saying`;
CREATE TABLE `z_saying` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '文章id',
  `user_id` int(11) unsigned NOT NULL DEFAULT '0',
  `userip` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '0' COMMENT '发帖时用户ip ip2long()，主要用来清理',
  `last_date` timestamp NULL DEFAULT NULL COMMENT '最后回复时间',
  `views_cnt` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '查看次数, 剥离出去，单独的服务，避免 cache 失效',
  `comments_cnt` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '回复数',
  `first_comment_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '首贴 pid',
  `last_user_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '最近参与的 uid',
  `last_comment_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '最后回复的 pid',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `lastpid` (`last_comment_id`),
  KEY `fid_2` (`last_comment_id`),
  KEY `idx_z_saying_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='论坛说说';

-- ----------------------------
-- Table structure for z_user
-- ----------------------------
DROP TABLE IF EXISTS `z_user`;
CREATE TABLE `z_user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户编号',
  `group_id` smallint(6) unsigned NOT NULL DEFAULT '0' COMMENT '用户组编号',
  `username` char(32) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` char(100) NOT NULL DEFAULT '' COMMENT '密码',
  `articles_cnt` int(11) NOT NULL DEFAULT '0' COMMENT '发帖数',
  `sayings_cnt` int(11) NOT NULL DEFAULT '0' COMMENT '说说数',
  `replys_cnt` int(11) NOT NULL DEFAULT '0' COMMENT '回复数',
  `comments_cnt` int(11) NOT NULL DEFAULT '0' COMMENT '评论数',
  `create_ip` varchar(20) NOT NULL DEFAULT '0' COMMENT '创建时IP',
  `create_date` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `login_ip` varchar(20) NOT NULL DEFAULT '0' COMMENT '登录时IP',
  `login_date` timestamp NULL DEFAULT NULL COMMENT '登录时间',
  `logins_cnt` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '登录次数',
  `avatar` varchar(200) NOT NULL DEFAULT '/static/img/avatar.png' COMMENT '用户最后更新图像时间',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `favourite_cnt` int(11) DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  KEY `gid` (`group_id`),
  KEY `idx_z_user_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of z_user
-- ----------------------------
INSERT INTO `z_user` (`id`, `group_id`, `username`, `password`, `articles_cnt`, `sayings_cnt`, `replys_cnt`, `comments_cnt`, `create_ip`, `create_date`, `login_ip`, `login_date`, `logins_cnt`, `avatar`, `created_at`, `updated_at`, `deleted_at`, `favourite_cnt`)
 VALUES ('1', '1', 'admin', '$2a$10$zzjAmJrsR0hk8UBbL9P3OOTLBBNEjtME1G5s3Vl2./.TwHrroDwkm', '0', '0', '0', '0', '0', '0', '0', '2019-08-21 13:49:12', '0', '/static/img/avatar.png', '2019-08-21 13:49:12', '2019-08-21 13:49:12', NULL, '0');
