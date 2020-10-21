# ************************************************************
# Sequel Pro SQL dump
# Version 5446
#
# https://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 8.0.19)
# Database: gocore
# Generation Time: 2020-10-21 13:04:16 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
SET NAMES utf8mb4;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table sys_api
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_api`;

CREATE TABLE `sys_api` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '物理主键',
  `pid` int DEFAULT NULL COMMENT '父级ID',
  `name` varchar(50) DEFAULT NULL COMMENT '接口名称',
  `path` varchar(128) DEFAULT NULL COMMENT '接口地址',
  `method` varchar(20) DEFAULT NULL COMMENT '接口请求方式',
  `sort` int DEFAULT NULL COMMENT '排序',
  `type` int DEFAULT NULL COMMENT '类型0目录1接口',
  `status` int DEFAULT '1' COMMENT '接口状态',
  `create_by` varchar(20) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(20) DEFAULT NULL COMMENT '修改人',
  `update_time` datetime DEFAULT NULL COMMENT '修改时间',
  `is_deleted` int DEFAULT '0' COMMENT '是否删除逻辑标识',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `sys_api` WRITE;
/*!40000 ALTER TABLE `sys_api` DISABLE KEYS */;

INSERT INTO `sys_api` (`id`, `pid`, `name`, `path`, `method`, `sort`, `type`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `is_deleted`)
VALUES
	(2,0,'基础接口','/base','',0,0,1,'admin','2020-10-13 15:13:39','admin','2020-10-13 15:13:39',0),
	(3,2,'登录接口','/base/login','POST',0,1,1,'admin','2020-10-13 15:14:13','admin','2020-10-13 15:14:13',0),
	(4,2,'退出接口','/base/logout','POST',1,1,1,'admin','2020-10-13 15:14:57','admin','2020-10-13 15:14:57',0),
	(5,2,'用户信息接口','/base/userinfo','GET',2,1,1,'admin','2020-10-13 15:15:32','admin','2020-10-13 15:18:34',1);

/*!40000 ALTER TABLE `sys_api` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_dept
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_dept`;

CREATE TABLE `sys_dept` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '物理主键',
  `dept_name` varchar(20) NOT NULL COMMENT '部门名称',
  `pid` int DEFAULT NULL COMMENT '父级ID',
  `leader_id` int DEFAULT NULL COMMENT '主管ID',
  `sort` int DEFAULT NULL COMMENT '排序',
  `status` int DEFAULT '1' COMMENT '部门状态',
  `create_by` varchar(20) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(20) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL COMMENT '修改时间',
  `is_deleted` int DEFAULT '0' COMMENT '是否删除逻辑标识',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `sys_dept` WRITE;
/*!40000 ALTER TABLE `sys_dept` DISABLE KEYS */;

INSERT INTO `sys_dept` (`id`, `dept_name`, `pid`, `leader_id`, `sort`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `is_deleted`)
VALUES
	(1,'华讯科技',0,1,0,1,'-','2020-10-09 22:01:09','-','2020-10-09 22:01:09',0),
	(2,'研发部',1,1,1,1,'-','2020-10-09 22:01:09','-','2020-10-09 22:01:09',0),
	(3,'运维部',1,1,3,0,'-','2020-10-09 22:01:09','huazi','2020-10-16 14:31:58',0),
	(4,'人力资源部',1,1,2,1,'-','2020-10-09 22:14:23','-','2020-10-09 22:14:23',0),
	(5,'前端开发组',2,1,1,1,'-','2020-10-09 22:15:34','-','2020-10-09 22:15:34',0),
	(6,'后端开发组',2,1,0,1,'-','2020-10-09 22:15:47','-','2020-10-09 22:15:47',0),
	(7,'场外运维部',3,NULL,1,0,'-','2020-10-11 16:02:10','huazi','2020-10-16 14:22:59',0),
	(8,'aaa',3,2,0,0,'huazi','2020-10-16 14:23:35','huazi','2020-10-16 14:23:35',1),
	(9,'场内运维部',3,7,2,0,'huazi','2020-10-16 14:33:40','huazi','2020-10-16 14:33:40',0),
	(10,'测试运维部',3,7,2,0,'-','2020-10-21 00:11:05','-','2020-10-21 00:11:05',0);

/*!40000 ALTER TABLE `sys_dept` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_dict_data
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_dict_data`;

CREATE TABLE `sys_dict_data` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '物理主键',
  `dict_type_id` int NOT NULL COMMENT '字典类型ID',
  `dict_label` varchar(50) NOT NULL COMMENT '字典数据名称',
  `dict_value` varchar(50) NOT NULL COMMENT '字典数据值',
  `description` varchar(255) DEFAULT NULL COMMENT '字典数据描述',
  `sort` int DEFAULT NULL COMMENT '排序',
  `status` int DEFAULT '1' COMMENT '字典数据状态',
  `create_by` varchar(20) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(20) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL COMMENT '修改时间',
  `is_deleted` int DEFAULT '0' COMMENT '是否删除逻辑标识',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `sys_dict_data` WRITE;
/*!40000 ALTER TABLE `sys_dict_data` DISABLE KEYS */;

INSERT INTO `sys_dict_data` (`id`, `dict_type_id`, `dict_label`, `dict_value`, `description`, `sort`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `is_deleted`)
VALUES
	(1,1,'0','男','男性',0,1,'-',NULL,'-',NULL,0),
	(2,1,'1','女','女性',1,1,'-',NULL,'-',NULL,0),
	(3,1,'2','未知','未知',2,1,'-',NULL,'-',NULL,0),
	(4,2,'0','禁用','禁用',0,1,'-',NULL,'-',NULL,0),
	(5,2,'1','启用','启用',1,1,'-',NULL,'-',NULL,0),
	(6,3,'0','否','否',0,1,'-',NULL,'-',NULL,0),
	(7,3,'1','是','是',1,1,'-',NULL,'-',NULL,0),
	(9,5,'高','height','',4,1,'-','2020-10-12 23:54:27','-','2020-10-12 23:54:27',0);

/*!40000 ALTER TABLE `sys_dict_data` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_dict_type
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_dict_type`;

CREATE TABLE `sys_dict_type` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '物理主键',
  `dict_name` varchar(20) NOT NULL COMMENT '字典名称',
  `dict_type` varchar(20) NOT NULL COMMENT '字典类型',
  `status` int DEFAULT '1' COMMENT '字典状态',
  `description` varchar(128) DEFAULT NULL COMMENT '描述',
  `sort` int DEFAULT NULL COMMENT '排序',
  `create_by` varchar(20) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '修改人',
  `update_time` datetime DEFAULT NULL COMMENT '修改时间',
  `is_deleted` int DEFAULT '0' COMMENT '是否删除逻辑标识',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `sys_dict_type` WRITE;
/*!40000 ALTER TABLE `sys_dict_type` DISABLE KEYS */;

INSERT INTO `sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `description`, `sort`, `create_by`, `create_time`, `update_by`, `update_time`, `is_deleted`)
VALUES
	(1,'性别','sex',0,NULL,0,'-',NULL,'huazi','2020-10-21 14:11:39',0),
	(2,'系统状态','status',1,NULL,1,'-',NULL,'-',NULL,0),
	(3,'结果','result',1,NULL,2,'-',NULL,'-',NULL,0),
	(5,'身材','body',1,NULL,3,'admin','2020-10-12 22:50:49','admin','2020-10-12 22:52:46',1),
	(6,'测试类型','test',0,'13141131412414124',1,'huazi','2020-10-21 14:19:28','huazi','2020-10-21 14:19:43',1);

/*!40000 ALTER TABLE `sys_dict_type` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_login_log
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_login_log`;

CREATE TABLE `sys_login_log` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '物理主键',
  `user_name` varchar(20) DEFAULT NULL COMMENT '用户名称',
  `ip_address` varchar(20) DEFAULT NULL COMMENT 'ip地址',
  `ip_location` varchar(50) DEFAULT NULL COMMENT 'ip所属区域',
  `browser` varchar(255) DEFAULT NULL COMMENT '浏览器',
  `os` varchar(50) DEFAULT NULL COMMENT '操作系统',
  `result` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '登录结果信息',
  `description` varchar(255) DEFAULT NULL COMMENT '具体描述（浏览器user-agent）',
  `login_time` datetime DEFAULT NULL COMMENT '登录时间',
  `is_deleted` int DEFAULT '0' COMMENT '是否删除逻辑标识',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table sys_menu
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_menu`;

CREATE TABLE `sys_menu` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '物理主键',
  `pid` int DEFAULT NULL COMMENT '父级ID',
  `name` varchar(20) DEFAULT NULL COMMENT '菜单名称',
  `title` varchar(40) DEFAULT NULL COMMENT '标题',
  `icon` varchar(20) DEFAULT NULL COMMENT '图标',
  `path` varchar(128) DEFAULT NULL COMMENT '路径',
  `component` varchar(128) DEFAULT NULL COMMENT '组件地址（组件名称）',
  `permission` varchar(100) DEFAULT NULL COMMENT '权限',
  `visible` int DEFAULT '1' COMMENT '显示',
  `cache` int DEFAULT '0' COMMENT '缓存',
  `sort` int DEFAULT NULL COMMENT '排序',
  `type` int DEFAULT NULL COMMENT '菜单类型 0目录1菜单2按钮',
  `create_by` varchar(20) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(20) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL COMMENT '修改时间',
  `is_deleted` int DEFAULT '0' COMMENT '是否删除逻辑标识',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `sys_menu` WRITE;
/*!40000 ALTER TABLE `sys_menu` DISABLE KEYS */;

INSERT INTO `sys_menu` (`id`, `pid`, `name`, `title`, `icon`, `path`, `component`, `permission`, `visible`, `cache`, `sort`, `type`, `create_by`, `create_time`, `update_by`, `update_time`, `is_deleted`)
VALUES
	(1,0,'system','系统管理','dashboard','/system','PageView','system',1,0,0,0,'-',NULL,'huazi','2020-10-21 09:49:52',0),
	(2,1,'user','用户管理','user','/system/user','/views/user/index','system:user',1,0,0,1,'-',NULL,'-',NULL,0),
	(4,2,'useradd','新增用户','','','','system:user:add',1,0,0,2,'-',NULL,'-',NULL,0),
	(5,2,'userupdate','编辑用户','','','','system:user:edit',1,0,1,2,'-',NULL,'-',NULL,0),
	(6,2,'userdelete','删除用户','','','','system:user:delete',1,0,2,2,'-',NULL,'-',NULL,0),
	(7,2,'userlist','查询用户','','','','system:user:list',0,1,3,2,'admin','2020-10-11 21:51:36','admin','2020-10-11 21:53:43',1),
	(8,0,'dashboard','主页','dashboard','/dashboard','PageView','dashboard',1,0,0,0,NULL,NULL,NULL,NULL,0),
	(9,8,'dashboard','主页','dashboard','/dashboard/index','/views/dashboard/index','dashboard:index',1,0,0,1,'-',NULL,'-',NULL,0),
	(10,0,'exception','异常','exception','/exception','PageView','exception',1,0,0,0,'-',NULL,'-',NULL,0),
	(11,10,'exception403','403','exception','/exception/403','/views/exception/403','exception:403',1,0,0,1,'-',NULL,'-',NULL,0),
	(12,10,'exception500','500','exception','/exception/500','/views/exception/500','exception:500',1,0,1,1,'-',NULL,'-',NULL,0),
	(13,1,'role','角色管理','role','/system/role','/view/role/index','system:role',1,0,1,1,'admin',NULL,'admin',NULL,0),
	(14,13,'roleadd','新增角色','','','','system:role:add',1,0,2,2,'-',NULL,'-',NULL,0),
	(15,13,'roleedit','编辑角色','','','','system:role:eidt',1,0,3,2,NULL,NULL,NULL,NULL,0),
	(16,13,'roledelete','删除角色','','','','system:role:delete',1,0,4,2,NULL,NULL,NULL,NULL,0),
	(17,14,'rolelist','查询角色','','','','system:role:list',1,0,5,2,NULL,NULL,NULL,NULL,0),
	(18,1,'post','岗位管理','post','/system/post','/views/post/index','system:post',1,0,2,1,'admin',NULL,'admin',NULL,0),
	(19,18,'postadd','新增岗位','','','','system:post:add',1,0,0,2,'-',NULL,'-',NULL,0),
	(20,18,'postedit','编辑岗位','','','','system:post:edit',1,0,1,2,'-',NULL,'-',NULL,0),
	(21,18,'postdelete','删除岗位','','','','system:post:delete',1,0,2,2,'-',NULL,'-',NULL,0),
	(22,18,'postlist','查询岗位','','','','system:post:list',1,0,3,2,'-',NULL,'-',NULL,0),
	(23,1,'menu','菜单管理','cascader','/system/menu','/views/menu/index','system:menu',1,0,3,1,'-','2020-10-20 14:50:56','-','2020-10-20 14:51:32',0),
	(24,23,'menuadd','添加菜单','code','','','system:menu:add',0,0,0,2,'huazi','2020-10-20 17:18:56','huazi','2020-10-20 17:18:56',0),
	(25,23,'menuedit','编辑菜单','404','','','system:menu:edit',0,0,1,2,'huazi','2020-10-20 17:22:02','huazi','2020-10-20 17:22:02',0),
	(26,23,'menudelete','删除菜单','cascader','','','system:menu:delete',0,0,2,2,'huazi','2020-10-20 17:23:15','huazi','2020-10-20 17:23:15',0),
	(27,23,'menulist','查询菜单','clipboard','','','system:menu:list',1,0,3,2,'huazi','2020-10-20 17:26:12','huazi','2020-10-20 17:26:52',0),
	(28,0,'log','日志管理','dashboard','/log','PageView','log',1,0,0,0,'huazi','2020-10-20 18:56:32','huazi','2020-10-21 09:49:59',0),
	(29,28,'loginlog','登录日志','peoples','/log/loginlog','/views/loginlog/index','log:loginlog',1,0,0,1,'huazi','2020-10-20 18:57:22','huazi','2020-10-20 18:57:22',0),
	(30,28,'operlog','操作日志','people','/log/operlog','/views/operlog/index','log:operlog',1,0,0,1,'huazi','2020-10-20 18:58:13','huazi','2020-10-20 18:58:13',0),
	(31,1,'dicttype','字典类型管理','dashboard','/system/dicttype','/views/dict/index','system:dicttype',1,0,5,1,'huazi','2020-10-21 10:34:11','huazi','2020-10-21 14:29:52',0),
	(32,31,'adddicttype','新增字典类型','dashboard','','','system:dict:addtype',1,0,0,2,'huazi','2020-10-21 10:35:08','huazi','2020-10-21 10:35:08',0),
	(33,31,'editdicttype','编辑字典类型','dashboard','','','system:dict:edittype',1,0,1,2,'huazi','2020-10-21 10:35:44','huazi','2020-10-21 10:35:44',0),
	(34,31,'deletedicttype','删除字典类型','dashboard','','','system:dict:deletetype',1,0,2,2,'huazi','2020-10-21 13:43:21','huazi','2020-10-21 13:43:21',0),
	(35,31,'listdicttype','查询字典类型','dashboard','','','system:dict:listtype',1,0,3,2,'huazi','2020-10-21 13:44:23','huazi','2020-10-21 13:44:23',0),
	(36,40,'adddictval','新增字典值','dashboard','','','system:dictval:addval',1,0,4,2,'huazi','2020-10-21 13:45:09','huazi','2020-10-21 14:33:18',0),
	(37,40,'editdictval','编辑字典值','','','','system:dictval:editval',1,0,5,2,'huazi','2020-10-21 13:46:19','huazi','2020-10-21 14:33:35',0),
	(38,40,'deletedictval','删除字典值','','','','system:dictval:deleteval',1,0,6,2,'huazi','2020-10-21 13:46:59','huazi','2020-10-21 14:33:46',0),
	(39,40,'listdictval','查询字典值','','','','system:dictval:listval',1,0,7,2,'huazi','2020-10-21 13:47:38','huazi','2020-10-21 14:33:56',0),
	(40,1,'dictval','字典数据管理','dashboard','/system/dictval','/views/dict/data','system:dictval',1,0,4,1,'huazi','2020-10-21 14:32:45','huazi','2020-10-21 14:32:45',0);

/*!40000 ALTER TABLE `sys_menu` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_oper_log
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_oper_log`;

CREATE TABLE `sys_oper_log` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '物理主键',
  `oper_name` varchar(20) DEFAULT NULL COMMENT '操作人',
  `oper_title` varchar(50) DEFAULT NULL COMMENT '操作描述',
  `method` varchar(10) DEFAULT NULL COMMENT '请求类型',
  `path` varchar(100) DEFAULT NULL COMMENT '请求地址',
  `params` varchar(255) DEFAULT NULL COMMENT '请求参数(json字符串)',
  `latency_time` varchar(20) DEFAULT NULL COMMENT '响应时间',
  `ip_address` varchar(20) DEFAULT NULL COMMENT 'ip地址',
  `ip_location` varchar(50) DEFAULT NULL COMMENT 'ip所属区域',
  `browser` varchar(255) DEFAULT NULL COMMENT '浏览器',
  `os` varchar(50) DEFAULT NULL COMMENT '操作系统',
  `result` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '操作结果信息',
  `oper_time` datetime DEFAULT NULL COMMENT '操作时间',
  `is_deleted` int DEFAULT '0' COMMENT '是否删除逻辑标识',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table sys_post
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_post`;

CREATE TABLE `sys_post` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '物理主键',
  `post_code` varchar(40) DEFAULT NULL COMMENT '岗位编码',
  `post_name` varchar(40) DEFAULT NULL COMMENT '岗位名称',
  `description` varchar(255) DEFAULT NULL COMMENT '岗位描述',
  `sort` int DEFAULT NULL COMMENT '排序',
  `status` int DEFAULT '1' COMMENT '岗位状态',
  `create_by` varchar(20) DEFAULT NULL COMMENT '创建人',
  `create_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_by` varchar(20) DEFAULT NULL COMMENT '修改人',
  `update_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `is_deleted` int DEFAULT '0' COMMENT '是否删除逻辑标识',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `sys_post` WRITE;
/*!40000 ALTER TABLE `sys_post` DISABLE KEYS */;

INSERT INTO `sys_post` (`id`, `post_code`, `post_name`, `description`, `sort`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `is_deleted`)
VALUES
	(1,'CEO','董事长','集团首脑',0,1,'-','2020-10-09 22:04:19','-','2020-10-09 22:04:19',0),
	(2,'CFO','财务主管','管理公司的财务',1,0,'admin','2020-10-20 12:12:14','admin','2020-10-20 12:12:14',0),
	(3,'CTO','技术总监','技术总监',1,0,'huazi','2020-10-20 12:17:27','huazi','2020-10-20 12:17:27',1),
	(4,'CDO','运维主管','12314124',2,0,'huazi','2020-10-20 12:17:27','huazi','2020-10-20 12:17:27',1);

/*!40000 ALTER TABLE `sys_post` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_role
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_role`;

CREATE TABLE `sys_role` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '物理主键',
  `role_code` varchar(40) DEFAULT NULL COMMENT '角色编码',
  `role_name` varchar(40) DEFAULT NULL COMMENT '角色名称',
  `description` varchar(255) DEFAULT NULL COMMENT '角色描述',
  `sort` int DEFAULT NULL COMMENT '排序',
  `status` int DEFAULT '1' COMMENT '角色状态',
  `create_by` varchar(20) DEFAULT NULL COMMENT '创建人',
  `create_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_by` varchar(20) DEFAULT NULL COMMENT '修改人',
  `update_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `is_deleted` int DEFAULT '0' COMMENT '是否删除逻辑标识',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `sys_role` WRITE;
/*!40000 ALTER TABLE `sys_role` DISABLE KEYS */;

INSERT INTO `sys_role` (`id`, `role_code`, `role_name`, `description`, `sort`, `status`, `create_by`, `create_time`, `update_by`, `update_time`, `is_deleted`)
VALUES
	(1,'admin','系统管理员','系统管理员，拥有所有权限',0,1,'-','2020-10-09 22:01:09','huazi','2020-10-21 14:34:10',0),
	(2,'common','普通操作员','普通操作员，拥有部分权限',1,1,'-','2020-10-09 22:01:09','huazi','2020-10-20 17:36:15',0),
	(3,'tester','测试操作员','测试操作员，拥有测试权限',2,0,'-','2020-10-20 11:44:46','huazi','2020-10-20 19:26:52',0),
	(4,'huazi111','华子啊','测试测试测试111',3,0,'admin','2020-10-20 11:42:48','huazi','2020-10-20 11:42:48',1),
	(5,'test001','test001','1234123',4,0,'huazi','2020-10-20 11:42:48','huazi','2020-10-20 11:42:48',1),
	(6,'test002','test002','1111111',5,0,'huazi','2020-10-20 11:39:15','huazi','2020-10-20 11:39:15',1);

/*!40000 ALTER TABLE `sys_role` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_role_api
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_role_api`;

CREATE TABLE `sys_role_api` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '物理主键',
  `role_id` int DEFAULT NULL COMMENT '角色ID',
  `api_id` int DEFAULT NULL COMMENT '接口ID',
  `create_by` varchar(20) DEFAULT NULL COMMENT '创建人',
  `create_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_by` varchar(20) DEFAULT NULL COMMENT '修改人',
  `update_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `is_deleted` int DEFAULT '0' COMMENT '是否删除逻辑标识',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table sys_role_menu
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_role_menu`;

CREATE TABLE `sys_role_menu` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '物理主键',
  `role_id` int DEFAULT NULL COMMENT '角色ID',
  `menu_id` int DEFAULT NULL COMMENT '菜单ID',
  `create_by` varchar(20) DEFAULT NULL COMMENT '创建人',
  `create_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_by` varchar(20) DEFAULT NULL COMMENT '修改人',
  `update_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `is_deleted` int DEFAULT '0' COMMENT '是否删除逻辑标识',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table sys_user
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_user`;

CREATE TABLE `sys_user` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '物理主键',
  `uuid` varchar(100) DEFAULT NULL COMMENT '用户唯一标识',
  `user_code` varchar(20) NOT NULL COMMENT '用户号',
  `user_name` varchar(100) DEFAULT NULL COMMENT '用户昵称',
  `avatar` varchar(255) DEFAULT NULL COMMENT '用户头像',
  `mobile` varchar(11) DEFAULT NULL COMMENT '手机号',
  `gender` varchar(1) DEFAULT '0' COMMENT '性别',
  `email` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '邮箱',
  `password` varchar(128) DEFAULT NULL COMMENT '密码，需要加密',
  `dept_id` int DEFAULT NULL COMMENT '部门ID',
  `role_id` int DEFAULT NULL COMMENT '角色ID',
  `post_id` int DEFAULT NULL COMMENT '岗位ID',
  `status` int DEFAULT '1' COMMENT '用户状态',
  `description` varchar(255) DEFAULT NULL COMMENT '描述',
  `create_by` varchar(50) DEFAULT NULL COMMENT '创建人',
  `create_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_by` varchar(50) DEFAULT NULL COMMENT '更新人',
  `update_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `is_deleted` int DEFAULT '0' COMMENT '是否删除逻辑标识',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `sys_user` WRITE;
/*!40000 ALTER TABLE `sys_user` DISABLE KEYS */;

INSERT INTO `sys_user` (`id`, `uuid`, `user_code`, `user_name`, `avatar`, `mobile`, `gender`, `email`, `password`, `dept_id`, `role_id`, `post_id`, `status`, `description`, `create_by`, `create_time`, `update_by`, `update_time`, `is_deleted`)
VALUES
	(1,'e10adc3949ba59abbe56e057f20f883e','admin','华子','https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png','19951299999','0','zu1662@qq.com','e10adc3949ba59abbe56e057f20f883e',1,1,1,1,NULL,'0','2020-10-15 13:38:07','0','2020-10-15 13:38:07',0),
	(2,'qeqwrqwe12314131weq1123qweq','test','测试账号','https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png','18236888848','1','abc@dd.com','e10adc3949ba59abbe56e057f20f883e',1,3,1,0,NULL,'0','2020-10-15 21:32:54','huazi','2020-10-15 21:32:55',0),
	(4,'123ewwqeqrqeqrqw','huazi111','华子1111','https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png','11111111111','1','huazi@qq.com','e10adc3949ba59abbe56e057f20f883e',1,3,1,1,NULL,'admin','2020-10-15 13:38:10','admin','2020-10-15 13:38:10',1),
	(7,'7d782588-6c7c-4e84-8846-f4e9d61f1637','huazi','嘿嘿嘿','https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png','19955599999','1','huazi@qq.com','cdf4a007e2b02a0c49fc9b7ccfbb8a10c644f635e1765dcf2a7ab794ddc7edac',4,1,1,1,NULL,'-','2020-10-15 20:13:08','-','2020-10-15 20:13:09',0),
	(8,'76e9fc7c-296d-467b-a72b-d7779383b94f','testa','testaa','','19955555555','0','a@qq.com','cdf4a007e2b02a0c49fc9b7ccfbb8a10c644f635e1765dcf2a7ab794ddc7edac',1,3,1,0,NULL,'-','2020-10-15 20:47:59','-','2020-10-15 20:47:59',1),
	(9,'66c9c89c-5ecf-4804-ada2-106b9f335713','testaa','testaaa','','19955555555','0','a@qq.com','cdf4a007e2b02a0c49fc9b7ccfbb8a10c644f635e1765dcf2a7ab794ddc7edac',1,3,1,0,NULL,'-','2020-10-15 20:48:33','-','2020-10-15 20:48:33',1),
	(10,'1418ce72-c48d-4f1a-a285-9d02d5669b30','testaab','testaaab','','19955555555','0','a@qq.com','cdf4a007e2b02a0c49fc9b7ccfbb8a10c644f635e1765dcf2a7ab794ddc7edac',1,3,1,0,NULL,'-','2020-10-15 20:48:33','-','2020-10-15 20:48:33',1),
	(11,'a2643ab2-4188-46e4-a898-cbad460f680f','testa','testa','','19966666666','0','q@qq.com','cdf4a007e2b02a0c49fc9b7ccfbb8a10c644f635e1765dcf2a7ab794ddc7edac',1,1,1,1,NULL,'-','2020-10-15 21:16:12','-','2020-10-15 21:16:12',1),
	(12,'5aca569a-9a82-4800-873d-aec9c4789ce5','testa','teata','','19951111111','0','q@qq.com','cdf4a007e2b02a0c49fc9b7ccfbb8a10c644f635e1765dcf2a7ab794ddc7edac',1,2,1,0,NULL,'-','2020-10-15 20:52:19','-','2020-10-15 20:52:19',1),
	(13,'c86b2a8e-b6ff-45f3-9300-cd9266453b74','teata','teata','','19981212122','0','1@1.com','b7ccfad2888f91c50b9a18b8e113c0aed274ab6a41fc5738c7105ae82ba2806e',1,1,1,0,NULL,'-','2020-10-15 21:16:12','huazi','2020-10-15 21:16:12',1);

/*!40000 ALTER TABLE `sys_user` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
