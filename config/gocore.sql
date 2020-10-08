/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50729
Source Host           : localhost:3306
Source Database       : gocore

Target Server Type    : MYSQL
Target Server Version : 50729
File Encoding         : 65001

Date: 2020-10-08 18:33:25
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for sys_api
-- ----------------------------
DROP TABLE IF EXISTS `sys_api`;
CREATE TABLE `sys_api` (
  `id` int(20) NOT NULL AUTO_INCREMENT COMMENT '物理主键',
  `pid` int(20) DEFAULT NULL COMMENT '父级ID',
  `name` varchar(50) DEFAULT NULL COMMENT '接口名称',
  `path` varchar(128) DEFAULT NULL COMMENT '接口地址',
  `method` varchar(20) DEFAULT NULL COMMENT '接口请求方式',
  `sort` int(4) DEFAULT NULL COMMENT '排序',
  `status` int(1) DEFAULT '1' COMMENT '接口状态',
  `create_by` varchar(20) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(20) DEFAULT NULL COMMENT '修改人',
  `update_time` datetime DEFAULT NULL COMMENT '修改时间',
  `is_deleted` int(1) DEFAULT '0' COMMENT '是否删除逻辑标识',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for sys_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_dept`;
CREATE TABLE `sys_dept` (
  `id` int(20) NOT NULL AUTO_INCREMENT COMMENT '物理主键',
  `dept_name` varchar(20) NOT NULL COMMENT '部门名称',
  `pid` int(20) DEFAULT NULL COMMENT '父级ID',
  `leader_id` int(20) DEFAULT NULL COMMENT '主管ID',
  `sort` int(4) DEFAULT NULL COMMENT '排序',
  `status` int(1) DEFAULT '1' COMMENT '部门状态',
  `create_by` varchar(20) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(20) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL COMMENT '修改时间',
  `is_deleted` int(1) DEFAULT '0' COMMENT '是否删除逻辑标识',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for sys_dict_data
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_data`;
CREATE TABLE `sys_dict_data` (
  `id` int(20) NOT NULL AUTO_INCREMENT COMMENT '物理主键',
  `dict_type_id` int(20) NOT NULL COMMENT '字典类型ID',
  `dict_label` varchar(50) NOT NULL COMMENT '字典数据名称',
  `dict_value` varchar(50) NOT NULL COMMENT '字典数据值',
  `description` varchar(255) DEFAULT NULL COMMENT '字典数据描述',
  `status` int(1) DEFAULT '1' COMMENT '字典数据状态',
  `create_by` varchar(20) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(20) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL COMMENT '修改时间',
  `is_deleted` int(1) DEFAULT '0' COMMENT '是否删除逻辑标识',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for sys_dict_type
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_type`;
CREATE TABLE `sys_dict_type` (
  `id` int(20) NOT NULL AUTO_INCREMENT COMMENT '物理主键',
  `dict_name` varchar(20) NOT NULL COMMENT '字典名称',
  `dict_type` varchar(20) NOT NULL COMMENT '字典类型',
  `status` int(1) DEFAULT '1' COMMENT '字典状态',
  `create_by` varchar(20) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(20) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL COMMENT '修改时间',
  `is_deleted` int(1) DEFAULT '0' COMMENT '是否删除逻辑标识',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for sys_login_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_login_log`;
CREATE TABLE `sys_login_log` (
  `id` int(20) NOT NULL AUTO_INCREMENT COMMENT '物理主键',
  `user_name` varchar(20) DEFAULT NULL COMMENT '用户名称',
  `ip_address` varchar(20) DEFAULT NULL COMMENT 'ip地址',
  `ip_location` varchar(50) DEFAULT NULL COMMENT 'ip所属区域',
  `browser` varchar(255) DEFAULT NULL COMMENT '浏览器',
  `os` varchar(50) DEFAULT NULL COMMENT '操作系统',
  `result` varchar(50) DEFAULT NULL COMMENT '登录结果信息',
  `description` varchar(255) DEFAULT NULL COMMENT '具体描述（浏览器user-agent）',
  `login_time` datetime DEFAULT NULL COMMENT '登录时间',
  `is_deleted` int(1) DEFAULT '0' COMMENT '是否删除逻辑标识',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=47 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu` (
  `id` int(20) NOT NULL AUTO_INCREMENT COMMENT '物理主键',
  `pid` int(20) DEFAULT NULL COMMENT '父级ID',
  `name` varchar(20) DEFAULT NULL COMMENT '菜单名称',
  `title` varchar(40) DEFAULT NULL COMMENT '标题',
  `icon` varchar(20) DEFAULT NULL COMMENT '图标',
  `path` varchar(128) DEFAULT NULL COMMENT '路径',
  `component` varchar(128) DEFAULT NULL COMMENT '组件地址（组件名称）',
  `permission` varchar(100) DEFAULT NULL COMMENT '权限',
  `visible` int(1) DEFAULT '1' COMMENT '显示',
  `cache` int(1) DEFAULT '0' COMMENT '缓存',
  `sort` int(4) DEFAULT NULL COMMENT '排序',
  `create_by` varchar(20) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(20) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL COMMENT '修改时间',
  `is_deleted` int(1) DEFAULT '0' COMMENT '是否删除逻辑标识',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for sys_oper_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_oper_log`;
CREATE TABLE `sys_oper_log` (
  `id` int(20) NOT NULL AUTO_INCREMENT COMMENT '物理主键',
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
  `result` varchar(50) DEFAULT NULL COMMENT '操作结果信息',
  `oper_time` datetime DEFAULT NULL COMMENT '操作时间',
  `is_deleted` int(1) DEFAULT '0' COMMENT '是否删除逻辑标识',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for sys_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_post`;
CREATE TABLE `sys_post` (
  `id` int(20) NOT NULL AUTO_INCREMENT COMMENT '物理主键',
  `post_code` varchar(40) DEFAULT NULL COMMENT '岗位编码',
  `post_name` varchar(40) DEFAULT NULL COMMENT '岗位名称',
  `description` varchar(255) DEFAULT NULL COMMENT '岗位描述',
  `sort` int(4) DEFAULT NULL COMMENT '排序',
  `status` int(1) DEFAULT '1' COMMENT '岗位状态',
  `create_by` varchar(20) DEFAULT NULL COMMENT '创建人',
  `create_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_by` varchar(20) DEFAULT NULL COMMENT '修改人',
  `update_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `is_deleted` int(1) DEFAULT '0' COMMENT '是否删除逻辑标识',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role` (
  `id` int(20) NOT NULL AUTO_INCREMENT COMMENT '物理主键',
  `role_code` varchar(40) DEFAULT NULL COMMENT '角色编码',
  `role_name` varchar(40) DEFAULT NULL COMMENT '角色名称',
  `description` varchar(255) DEFAULT NULL COMMENT '角色描述',
  `sort` int(4) DEFAULT NULL COMMENT '排序',
  `status` int(1) DEFAULT '1' COMMENT '角色状态',
  `create_by` varchar(20) DEFAULT NULL COMMENT '创建人',
  `create_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_by` varchar(20) DEFAULT NULL COMMENT '修改人',
  `update_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `is_deleted` int(1) DEFAULT '0' COMMENT '是否删除逻辑标识',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for sys_role_api
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_api`;
CREATE TABLE `sys_role_api` (
  `id` int(20) NOT NULL AUTO_INCREMENT COMMENT '物理主键',
  `role_id` int(20) DEFAULT NULL COMMENT '角色ID',
  `api_id` int(20) DEFAULT NULL COMMENT '接口ID',
  `create_by` varchar(20) DEFAULT NULL COMMENT '创建人',
  `create_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_by` varchar(20) DEFAULT NULL COMMENT '修改人',
  `update_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `is_deleted` int(1) DEFAULT '0' COMMENT '是否删除逻辑标识',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu` (
  `id` int(20) NOT NULL AUTO_INCREMENT COMMENT '物理主键',
  `role_id` int(20) DEFAULT NULL COMMENT '角色ID',
  `menu_id` int(20) DEFAULT NULL COMMENT '菜单ID',
  `create_by` varchar(20) DEFAULT NULL COMMENT '创建人',
  `create_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_by` varchar(20) DEFAULT NULL COMMENT '修改人',
  `update_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `is_deleted` int(1) DEFAULT '0' COMMENT '是否删除逻辑标识',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
  `id` int(20) NOT NULL AUTO_INCREMENT COMMENT '物理主键',
  `uuid` varchar(100) DEFAULT NULL COMMENT '用户唯一标识',
  `user_code` varchar(20) NOT NULL COMMENT '用户号',
  `user_name` varchar(100) DEFAULT NULL COMMENT '用户昵称',
  `avatar` varchar(255) DEFAULT NULL COMMENT '用户头像',
  `mobile` varchar(11) DEFAULT NULL COMMENT '手机号',
  `gender` varchar(1) DEFAULT '0' COMMENT '性别',
  `password` varchar(128) DEFAULT NULL COMMENT '密码，需要加密',
  `dept_id` int(20) DEFAULT NULL COMMENT '部门ID',
  `role_id` int(20) DEFAULT NULL COMMENT '角色ID',
  `post_id` int(20) DEFAULT NULL COMMENT '岗位ID',
  `status` int(1) DEFAULT '1' COMMENT '用户状态',
  `create_by` varchar(50) DEFAULT NULL COMMENT '创建人',
  `create_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_by` varchar(50) DEFAULT NULL COMMENT '更新人',
  `update_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `is_deleted` int(1) DEFAULT '0' COMMENT '是否删除逻辑标识',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
