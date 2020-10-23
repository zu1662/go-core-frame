# ************************************************************
# Sequel Pro SQL dump
# Version 5446
#
# https://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 8.0.19)
# Database: gocore
# Generation Time: 2020-10-23 10:18:21 +0000
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
	(5,2,'用户信息接口','/base/userinfo','GET',2,1,1,'admin','2020-10-13 15:15:32','admin','2020-10-13 15:18:34',1),
	(6,2,'获取验证码','/base/captcha','GET',2,1,1,'admin','2020-10-23 13:08:19','admin','2020-10-23 13:08:19',0),
	(7,2,'获取用户信息','/base/userinfo','GET',3,1,1,'admin','2020-10-23 13:08:56','admin','2020-10-23 13:08:56',0),
	(8,0,'日志接口','/log','',1,0,1,'admin','2020-10-23 13:10:12','admin','2020-10-23 13:10:12',0),
	(9,8,'获取登录日志详情','/log/loginloginfo','GET',0,1,1,'admin','2020-10-23 13:10:42','admin','2020-10-23 13:10:42',0),
	(10,8,'登录日志列表','/log/loginloglist','GET',1,1,1,'admin','2020-10-23 13:11:06','admin','2020-10-23 13:11:06',0),
	(11,8,'删除登录日志','/log/deleteloginlog','DELETE',2,1,1,'admin','2020-10-23 13:11:37','admin','2020-10-23 13:11:37',0),
	(12,8,'清空登录日志','/log/cleanloginlog','DELETE',3,1,1,'admin','2020-10-23 13:12:29','admin','2020-10-23 13:12:29',0),
	(13,8,'操作日志详情','/log/operloginfo','GET',4,1,1,'admin','2020-10-23 13:12:52','admin','2020-10-23 13:12:52',0),
	(14,8,'操作日志列表','/log/operloglist','GET',5,1,1,'admin','2020-10-23 13:14:03','admin','2020-10-23 13:14:03',0),
	(15,8,'删除操作日志','/log/deleteoperlog','DELETE',6,1,1,'admin','2020-10-23 13:14:25','admin','2020-10-23 13:14:59',0),
	(16,8,'清空操作日志','/log/cleanoperlog','DELETE',7,1,1,'admin','2020-10-23 13:14:44','common','2020-10-23 18:04:08',0),
	(17,0,'用户接口','/user','',2,0,1,'admin','2020-10-23 13:17:45','admin','2020-10-23 13:17:45',0),
	(18,17,'用户详情','/user/info','GET',0,1,1,'admin','2020-10-23 13:18:08','admin','2020-10-23 13:18:08',0),
	(19,17,'用户分页列表','/user/list','GET',1,1,1,'admin','2020-10-23 13:18:42','admin','2020-10-23 13:18:42',0),
	(20,17,'用户全部列表','/user/listall','GET',2,1,1,'admin','2020-10-23 13:19:36','admin','2020-10-23 13:19:36',0),
	(21,17,'更新用户信息','/user/update','PUT',0,1,1,'admin','2020-10-23 13:20:03','admin','2020-10-23 13:20:03',0),
	(22,17,'用户重置密码','/user/resetpsw','PUT',3,1,1,'admin','2020-10-23 13:20:30','admin','2020-10-23 13:20:30',0),
	(23,17,'删除用户','/user/delete','DELETE',4,1,1,'admin','2020-10-23 13:20:53','admin','2020-10-23 13:21:19',0),
	(24,17,'新增用户','/user/add','POST',5,1,1,'admin','2020-10-23 13:21:10','admin','2020-10-23 13:21:22',0),
	(25,0,'部门接口','/dept','',3,0,1,'admin','2020-10-23 13:22:09','admin','2020-10-23 13:22:09',0),
	(26,25,'部门详情','/dept/info','GET',0,1,1,'admin','2020-10-23 13:22:34','admin','2020-10-23 13:22:34',0),
	(27,25,'部门树结构','/dept/tree','GET',1,1,1,'admin','2020-10-23 13:22:53','admin','2020-10-23 13:22:53',0),
	(28,25,'部门更新','/dept/update','PUT',2,1,1,'admin','2020-10-23 13:23:08','admin','2020-10-23 13:23:34',0),
	(29,25,'部门删除','/dept/delete','DELETE',3,1,1,'admin','2020-10-23 13:23:51','admin','2020-10-23 13:23:51',0),
	(30,25,'部门新增','/dept/add','POST',4,1,1,'admin','2020-10-23 13:24:14','admin','2020-10-23 13:24:14',0),
	(31,0,'岗位接口','/post','',4,0,1,'admin','2020-10-23 13:24:51','admin','2020-10-23 13:24:51',0),
	(32,31,'岗位详情','/post/info','GET',0,1,1,'admin','2020-10-23 13:25:13','admin','2020-10-23 13:25:13',0),
	(33,31,'岗位分页列表','/post/list','GET',1,1,1,'admin','2020-10-23 13:25:40','admin','2020-10-23 13:25:40',0),
	(34,31,'岗位全部列表','/post/listall','GET',2,1,1,'admin','2020-10-23 13:26:01','admin','2020-10-23 13:26:01',0),
	(35,31,'岗位更新','/post/update','PUT',3,1,1,'admin','2020-10-23 13:26:20','admin','2020-10-23 13:26:45',0),
	(36,31,'岗位删除','/post/delete','DELETE',4,1,1,'admin','2020-10-23 13:26:35','admin','2020-10-23 13:26:35',0),
	(37,31,'岗位新增','/post/add','POST',5,1,1,'admin','2020-10-23 13:27:03','admin','2020-10-23 13:27:03',0),
	(38,0,'菜单接口','/menu','',5,0,1,'admin','2020-10-23 13:28:10','admin','2020-10-23 13:28:10',0),
	(39,38,'菜单详情','/menu/info','GET',0,1,1,'admin','2020-10-23 13:28:28','admin','2020-10-23 13:28:28',0),
	(40,38,'菜单树结构','/menu/tree','GET',1,1,1,'admin','2020-10-23 13:28:49','admin','2020-10-23 13:28:49',0),
	(41,38,'菜单更新','/menu/update','PUT',2,1,1,'admin','2020-10-23 13:29:09','admin','2020-10-23 13:29:23',0),
	(42,38,'菜单删除','/menu/delete','DELETE',3,1,1,'admin','2020-10-23 13:29:41','admin','2020-10-23 13:29:41',0),
	(43,38,'菜单新增','/menu/add','POST',4,1,1,'admin','2020-10-23 13:29:57','admin','2020-10-23 13:29:57',0),
	(44,0,'角色接口','/role','',6,0,1,'admin','2020-10-23 13:30:24','admin','2020-10-23 13:30:24',0),
	(45,44,'角色详情','/role/info','GET',0,1,1,'admin','2020-10-23 13:30:43','admin','2020-10-23 13:30:43',0),
	(46,44,'角色分页列表','/role/list','GET',1,1,1,'admin','2020-10-23 13:31:00','admin','2020-10-23 13:31:00',0),
	(47,44,'角色全部列表','/role/listall','GET',2,1,1,'admin','2020-10-23 13:31:18','admin','2020-10-23 13:31:18',0),
	(48,44,'角色更新','/role/update','PUT',3,1,1,'admin','2020-10-23 13:31:41','admin','2020-10-23 13:31:41',0),
	(49,44,'角色删除','/role/delete','DELETE',4,1,1,'admin','2020-10-23 13:31:58','admin','2020-10-23 13:31:58',0),
	(50,44,'角色新增','/role/add','POST',5,1,1,'admin','2020-10-23 13:32:15','admin','2020-10-23 13:32:15',0),
	(51,0,'角色菜单接口','/rolemenu','',7,0,1,'admin','2020-10-23 13:33:08','admin','2020-10-23 13:33:08',0),
	(52,51,'角色菜单列表','/rolemenu/list','GET',0,1,1,'admin','2020-10-23 13:33:40','admin','2020-10-23 13:33:40',0),
	(53,51,'角色菜单更新','/rolemenu/update','POST',1,1,1,'admin','2020-10-23 13:34:01','admin','2020-10-23 13:34:11',0),
	(54,0,'字典接口','/dict','',8,0,1,'-','2020-10-23 14:07:45','-','2020-10-23 14:07:45',0),
	(55,54,'字典类型列表','/dict/dicttypelist','GET',2,1,1,'-','2020-10-23 14:08:21','-','2020-10-23 14:10:00',0),
	(56,54,'字典类型详情','/dict/dicttype','GET',1,1,1,'-','2020-10-23 14:08:47','-','2020-10-23 14:08:47',0),
	(57,54,'字典全部值获取','/dict/dictmap','GET',0,1,1,'-','2020-10-23 14:09:44','-','2020-10-23 14:09:44',0),
	(58,54,'字典类型更新','/dict/dicttypeupdate','PUT',3,1,1,'-','2020-10-23 14:10:20','-','2020-10-23 14:10:20',0),
	(59,54,'字典类型删除','/dict/dicttypedelete','DELETE',4,1,1,'-','2020-10-23 14:10:44','-','2020-10-23 14:10:44',0),
	(60,54,'字典类型新增','/dict/dicttypeadd','POST',5,1,1,'-','2020-10-23 14:11:07','-','2020-10-23 14:11:07',0),
	(61,54,'字典值列表','/dict/dictdatalist','GET',6,1,1,'-','2020-10-23 14:11:36','-','2020-10-23 14:11:36',0),
	(62,54,'字典值详情','/dict/dictdata','GET',7,1,1,'-','2020-10-23 14:12:02','-','2020-10-23 14:12:02',0),
	(63,54,'字典值更新','/dict/dictdataupdate','PUT',8,1,1,'-','2020-10-23 14:12:22','-','2020-10-23 14:12:22',0),
	(64,54,'字典值删除','/dict/dictdatadelete','DELETE',9,1,1,'-','2020-10-23 14:12:42','-','2020-10-23 14:12:42',0),
	(65,54,'字典值新增','/dict/dictdataadd','POST',10,1,1,'-','2020-10-23 14:13:01','-','2020-10-23 14:13:01',0),
	(66,0,'API接口','/interface','',9,0,1,'-','2020-10-23 14:14:16','-','2020-10-23 14:14:16',0),
	(67,66,'API详情','/interface/info','GET',0,1,1,'-','2020-10-23 14:14:38','-','2020-10-23 14:14:38',0),
	(68,66,'API树形结构','/interface/tree','GET',1,1,1,'-','2020-10-23 14:15:02','-','2020-10-23 14:15:02',0),
	(69,66,'API更新','/interface/update','PUT',2,1,1,'-','2020-10-23 14:15:24','-','2020-10-23 14:15:24',0),
	(70,66,'API删除','/interface/delete','DELETE',3,1,1,'-','2020-10-23 14:15:40','admin','2020-10-23 18:06:18',0),
	(71,66,'API新增','/interface/add','POST',3,1,1,'-','2020-10-23 14:16:02','-','2020-10-23 14:16:08',0),
	(72,0,'角色API接口','/roleapi','',10,0,1,'-','2020-10-23 14:16:49','-','2020-10-23 14:16:49',0),
	(73,72,'角色API列表','/roleapi/list','GET',0,1,1,'-','2020-10-23 14:17:18','-','2020-10-23 14:17:18',0),
	(74,72,'角色API更新','/roleapi/update','POST',1,1,1,'-','2020-10-23 14:17:37','-','2020-10-23 14:17:37',0);

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
	(3,'运维部',1,1,3,0,'-','2020-10-09 22:01:09','admin','2020-10-16 14:31:58',0),
	(4,'人力资源部',1,1,2,1,'-','2020-10-09 22:14:23','-','2020-10-09 22:14:23',0),
	(5,'前端开发组',2,1,1,1,'-','2020-10-09 22:15:34','-','2020-10-09 22:15:34',0),
	(6,'后端开发组',2,1,0,1,'-','2020-10-09 22:15:47','-','2020-10-09 22:15:47',0),
	(7,'场外运维部',3,NULL,1,0,'-','2020-10-11 16:02:10','admin','2020-10-16 14:22:59',0),
	(9,'场内运维部',3,7,2,0,'admin','2020-10-16 14:33:40','admin','2020-10-16 14:33:40',0);

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
	(1,1,'男','0','男性',0,1,'-',NULL,'-',NULL,0),
	(2,1,'女','1','女性',1,1,'-',NULL,'-',NULL,0),
	(3,1,'未知','2','未知',2,1,'-',NULL,'-',NULL,0),
	(4,2,'禁用','0','禁用',0,1,'-',NULL,'-',NULL,0),
	(5,2,'启用','1','启用',1,1,'-',NULL,'-',NULL,0),
	(6,3,'否','0','否',0,1,'-',NULL,'-',NULL,0),
	(7,3,'是','1','是',1,1,'-',NULL,'-',NULL,0),
	(10,7,'显示','1','显示',0,1,'admin','2020-10-23 10:56:26','admin','2020-10-23 10:56:26',0),
	(11,7,'隐藏','0','隐藏',1,1,'admin','2020-10-23 10:56:38','admin','2020-10-23 10:56:38',0),
	(12,8,'缓存','1','缓存',0,1,'admin','2020-10-23 11:00:19','admin','2020-10-23 11:00:19',0),
	(13,8,'不缓存','0','不缓存',1,1,'admin','2020-10-23 11:00:30','admin','2020-10-23 11:00:30',0);

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
	(1,'性别','SEX',0,NULL,0,'-',NULL,'huazi','2020-10-21 14:11:39',0),
	(2,'系统状态','STATUS',1,NULL,1,'-',NULL,'-',NULL,0),
	(3,'结果','RESULT',1,NULL,2,'-',NULL,'-',NULL,0),
	(7,'显示','VISIBLE',1,'是否显示',3,'admin','2020-10-23 10:44:57','admin','2020-10-23 10:44:57',0),
	(8,'缓存','CACHE',1,'菜单路由缓存',4,'admin','2020-10-23 11:00:04','admin','2020-10-23 11:00:04',0);

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
	(40,1,'dictval','字典数据管理','dashboard','/system/dictval','/views/dict/data','system:dictval',1,0,4,1,'huazi','2020-10-21 14:32:45','huazi','2020-10-21 14:32:45',0),
	(41,1,'interface','接口管理','dashboard','/system/interface','/views/interface/index','system:interface',1,0,6,1,'admin','2020-10-23 18:12:24','admin','2020-10-23 18:12:24',0),
	(42,41,'interfacelist','查询接口','','','','system:interface:list',1,0,0,2,'admin','2020-10-23 18:14:41','admin','2020-10-23 18:15:07',0),
	(43,41,'interfaceadd','新增接口','','','','system:interface:add',1,0,1,2,'admin','2020-10-23 18:15:42','admin','2020-10-23 18:15:42',0),
	(44,41,'interfaceedit','更新接口','','','','system:interface:edit',1,0,2,2,'admin','2020-10-23 18:16:19','admin','2020-10-23 18:16:19',0),
	(45,41,'interfacedelete','删除接口','','','','system:interface:delete',1,0,3,2,'admin','2020-10-23 18:16:47','admin','2020-10-23 18:16:47',0);

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
	(3,'CTO','技术总监','技术总监',1,0,'huazi','2020-10-20 12:17:27','huazi','2020-10-20 12:17:27',1);

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
	(1,'admin','系统管理员','系统管理员，拥有所有权限',0,1,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(2,'common','普通操作员','普通操作员，拥有部分权限',1,1,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(3,'tester','测试操作员','测试操作员，拥有测试权限',2,1,'-','2020-10-23 17:48:34','admin','2020-10-23 18:17:21',0);

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

LOCK TABLES `sys_role_api` WRITE;
/*!40000 ALTER TABLE `sys_role_api` DISABLE KEYS */;

INSERT INTO `sys_role_api` (`id`, `role_id`, `api_id`, `create_by`, `create_time`, `update_by`, `update_time`, `is_deleted`)
VALUES
	(226,1,3,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(227,1,4,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(228,1,6,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(229,1,7,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(230,1,9,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(231,1,10,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(232,1,11,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(233,1,12,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(234,1,13,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(235,1,14,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(236,1,15,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(237,1,16,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(238,1,21,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(239,1,18,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(240,1,19,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(241,1,20,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(242,1,22,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(243,1,23,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(244,1,24,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(245,1,26,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(246,1,27,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(247,1,28,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(248,1,29,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(249,1,30,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(250,1,32,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(251,1,33,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(252,1,34,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(253,1,35,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(254,1,36,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(255,1,37,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(256,1,39,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(257,1,40,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(258,1,41,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(259,1,42,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(260,1,43,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(261,1,45,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(262,1,46,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(263,1,47,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(264,1,48,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(265,1,49,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(266,1,50,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(267,1,52,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(268,1,53,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(269,1,57,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(270,1,56,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(271,1,55,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(272,1,58,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(273,1,59,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(274,1,60,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(275,1,61,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(276,1,62,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(277,1,63,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(278,1,64,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(279,1,65,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(280,1,67,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(281,1,68,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(282,1,69,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(283,1,71,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(284,1,70,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(285,1,73,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(286,1,74,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(287,2,3,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(288,2,4,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(289,2,6,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(290,2,7,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(291,2,9,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(292,2,10,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(293,2,13,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(294,2,14,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(295,2,18,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(296,2,19,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(297,2,20,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(298,2,26,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(299,2,27,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(300,2,32,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(301,2,33,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(302,2,34,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(303,2,39,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(304,2,40,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(305,2,45,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(306,2,46,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(307,2,47,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(308,2,52,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(309,2,57,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(310,2,56,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(311,2,55,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(312,2,61,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(313,2,62,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(314,2,67,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(315,2,68,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(316,2,73,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(317,3,3,'-','2020-10-23 17:48:34','admin','2020-10-23 18:17:21',0),
	(318,3,4,'-','2020-10-23 17:48:34','admin','2020-10-23 18:17:21',0),
	(319,3,6,'-','2020-10-23 17:48:34','admin','2020-10-23 18:17:21',0),
	(320,3,7,'-','2020-10-23 17:48:34','admin','2020-10-23 18:17:21',0),
	(321,3,9,'-','2020-10-23 17:48:34','admin','2020-10-23 18:17:21',0),
	(322,3,10,'-','2020-10-23 17:48:34','admin','2020-10-23 18:17:21',0);

/*!40000 ALTER TABLE `sys_role_api` ENABLE KEYS */;
UNLOCK TABLES;


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

LOCK TABLES `sys_role_menu` WRITE;
/*!40000 ALTER TABLE `sys_role_menu` DISABLE KEYS */;

INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `create_by`, `create_time`, `update_by`, `update_time`, `is_deleted`)
VALUES
	(95,5,1,'huazi','2020-10-20 11:16:02','huazi','2020-10-20 11:38:31',0),
	(96,5,2,'huazi','2020-10-20 11:16:02','huazi','2020-10-20 11:38:31',0),
	(97,5,4,'huazi','2020-10-20 11:16:02','huazi','2020-10-20 11:38:31',0),
	(98,5,5,'huazi','2020-10-20 11:16:02','huazi','2020-10-20 11:38:31',0),
	(99,5,6,'huazi','2020-10-20 11:16:02','huazi','2020-10-20 11:38:31',0),
	(100,5,13,'huazi','2020-10-20 11:16:02','huazi','2020-10-20 11:38:31',0),
	(101,5,14,'huazi','2020-10-20 11:16:02','huazi','2020-10-20 11:38:31',0),
	(102,5,17,'huazi','2020-10-20 11:16:02','huazi','2020-10-20 11:38:31',0),
	(103,5,15,'huazi','2020-10-20 11:16:02','huazi','2020-10-20 11:38:31',0),
	(104,5,16,'huazi','2020-10-20 11:16:02','huazi','2020-10-20 11:38:31',0),
	(105,5,8,'huazi','2020-10-20 11:16:02','huazi','2020-10-20 11:38:31',0),
	(106,5,9,'huazi','2020-10-20 11:16:02','huazi','2020-10-20 11:38:31',0),
	(107,5,10,'huazi','2020-10-20 11:16:02','huazi','2020-10-20 11:38:31',0),
	(108,5,11,'huazi','2020-10-20 11:16:02','huazi','2020-10-20 11:38:31',0),
	(109,5,12,'huazi','2020-10-20 11:16:02','huazi','2020-10-20 11:38:31',0),
	(793,1,1,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(794,1,2,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(795,1,4,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(796,1,5,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(797,1,6,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(798,1,13,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(799,1,14,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(800,1,17,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(801,1,15,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(802,1,16,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(803,1,18,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(804,1,19,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(805,1,20,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(806,1,21,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(807,1,22,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(808,1,23,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(809,1,24,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(810,1,25,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(811,1,26,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(812,1,27,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(813,1,40,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(814,1,36,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(815,1,37,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(816,1,38,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(817,1,39,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(818,1,31,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(819,1,32,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(820,1,33,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(821,1,34,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(822,1,35,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(823,1,41,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(824,1,42,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(825,1,43,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(826,1,44,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(827,1,45,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(828,1,8,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(829,1,9,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(830,1,10,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(831,1,11,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(832,1,12,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(833,1,28,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(834,1,30,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(835,1,29,'-','2020-10-09 22:01:09','admin','2020-10-23 18:16:58',0),
	(836,2,1,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(837,2,2,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(838,2,4,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(839,2,5,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(840,2,6,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(841,2,13,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(842,2,14,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(843,2,17,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(844,2,15,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(845,2,16,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(846,2,18,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(847,2,19,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(848,2,20,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(849,2,21,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(850,2,22,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(851,2,23,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(852,2,24,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(853,2,25,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(854,2,26,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(855,2,27,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(856,2,40,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(857,2,36,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(858,2,37,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(859,2,38,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(860,2,39,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(861,2,31,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(862,2,32,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(863,2,33,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(864,2,34,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(865,2,35,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(866,2,41,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(867,2,42,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(868,2,43,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(869,2,44,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(870,2,45,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(871,2,8,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(872,2,9,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(873,2,10,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(874,2,11,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(875,2,12,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(876,2,28,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(877,2,30,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(878,2,29,'-','2020-10-09 22:01:09','admin','2020-10-23 18:17:10',0),
	(879,3,8,'-','2020-10-23 17:48:34','admin','2020-10-23 18:17:21',0),
	(880,3,9,'-','2020-10-23 17:48:34','admin','2020-10-23 18:17:21',0),
	(881,3,28,'-','2020-10-23 17:48:34','admin','2020-10-23 18:17:21',0),
	(882,3,30,'-','2020-10-23 17:48:34','admin','2020-10-23 18:17:21',0),
	(883,3,29,'-','2020-10-23 17:48:34','admin','2020-10-23 18:17:21',0);

/*!40000 ALTER TABLE `sys_role_menu` ENABLE KEYS */;
UNLOCK TABLES;


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
	(1,'e10adc3949ba59abbe56e057f20f883e','admin','管理员账号','https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png','19951299999','0','zu1662@qq.com','cdf4a007e2b02a0c49fc9b7ccfbb8a10c644f635e1765dcf2a7ab794ddc7edac',1,1,1,1,NULL,'0','2020-10-23 11:06:57','0','2020-10-23 11:06:57',0),
	(2,'qeqwrqwe12314131weq1123qweq','test','测试账号','https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png','18236888848','1','abc@dd.com','cdf4a007e2b02a0c49fc9b7ccfbb8a10c644f635e1765dcf2a7ab794ddc7edac',1,3,1,1,NULL,'0','2020-10-23 11:07:10','huazi','2020-10-23 11:07:10',0),
	(7,'7d782588-6c7c-4e84-8846-f4e9d61f1637','common','普通人员账号','https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png','19955599999','1','huazi@qq.com','cdf4a007e2b02a0c49fc9b7ccfbb8a10c644f635e1765dcf2a7ab794ddc7edac',4,2,1,1,NULL,'-','2020-10-23 11:06:52','huazi','2020-10-23 11:06:52',0);

/*!40000 ALTER TABLE `sys_user` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
