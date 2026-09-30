/*
 Navicat Premium Dump SQL

 Source Server         : lo_mysql8
 Source Server Type    : MySQL
 Source Server Version : 80039 (8.0.39)
 Source Host           : localhost:3306
 Source Schema         : vben-go

 Target Server Type    : MySQL
 Target Server Version : 80039 (8.0.39)
 File Encoding         : 65001

 Date: 20/01/2026 15:04:16
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for demo_link
-- ----------------------------
DROP TABLE IF EXISTS `demo_link`;
CREATE TABLE `demo_link`  (
  `id` bigint NOT NULL COMMENT '主键ID',
  `avtag` bit(1) NULL DEFAULT NULL COMMENT '可用标记',
  `crtim` datetime(6) NULL DEFAULT NULL COMMENT '创建时间',
  `cruid` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '创建人ID',
  `name` varchar(126) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
  `notes` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `uptim` datetime(6) NULL DEFAULT NULL COMMENT '更新时间',
  `upuid` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新人ID',
  `catid` bigint NULL DEFAULT NULL COMMENT '所属分类ID',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `IDX1ia94fj0wmefg9508r7bfymml`(`cruid` ASC) USING BTREE,
  INDEX `IDX6r6t02bhx3d8gvsv3j0lru14t`(`upuid` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '关联主表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of demo_link
-- ----------------------------

-- ----------------------------
-- Table structure for demo_link_actor
-- ----------------------------
DROP TABLE IF EXISTS `demo_link_actor`;
CREATE TABLE `demo_link_actor`  (
  `mid` bigint NOT NULL,
  `aid` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  INDEX `FKdq9s6svrpxr3emvowyd1fhrtu`(`aid` ASC) USING BTREE,
  INDEX `FK2qebnnev5askpa4cwb2fv4dtn`(`mid` ASC) USING BTREE,
  CONSTRAINT `FK2qebnnev5askpa4cwb2fv4dtn` FOREIGN KEY (`mid`) REFERENCES `demo_link` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `FKdq9s6svrpxr3emvowyd1fhrtu` FOREIGN KEY (`aid`) REFERENCES `sys_actor` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of demo_link_actor
-- ----------------------------

-- ----------------------------
-- Table structure for demo_link_cate
-- ----------------------------
DROP TABLE IF EXISTS `demo_link_cate`;
CREATE TABLE `demo_link_cate`  (
  `id` bigint NOT NULL COMMENT '主键ID',
  `avtag` bit(1) NULL DEFAULT NULL COMMENT '可用标记 1启用，0禁用',
  `crtim` datetime(6) NULL DEFAULT NULL COMMENT '创建时间',
  `cruid` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '创建人ID',
  `label` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标签',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
  `notes` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `ornum` int NULL DEFAULT NULL COMMENT '排序号',
  `pid` bigint NULL DEFAULT NULL COMMENT '父分类ID',
  `tier` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '层级信息',
  `uptim` datetime(6) NULL DEFAULT NULL COMMENT '更新时间',
  `upuid` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新人ID',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `IDXayrvyxx7u0bcy0vk3llsvvic9`(`cruid` ASC) USING BTREE,
  INDEX `IDX4uodprn78bsjiblaj6c7r7i2x`(`upuid` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '关联分类表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of demo_link_cate
-- ----------------------------

-- ----------------------------
-- Table structure for demo_link_item
-- ----------------------------
DROP TABLE IF EXISTS `demo_link_item`;
CREATE TABLE `demo_link_item`  (
  `id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '主键ID',
  `maiid` bigint NULL DEFAULT NULL COMMENT '主表ID',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '子项目名称',
  `notes` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `ornum` int NULL DEFAULT NULL COMMENT '排序号',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `FK4uvdeb9iplndpt959bhxjr14j`(`maiid` ASC) USING BTREE,
  CONSTRAINT `FK4uvdeb9iplndpt959bhxjr14j` FOREIGN KEY (`maiid`) REFERENCES `demo_link` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '关联子表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of demo_link_item
-- ----------------------------

-- ----------------------------
-- Table structure for demo_single
-- ----------------------------
DROP TABLE IF EXISTS `demo_single`;
CREATE TABLE `demo_single`  (
  `id` bigint NOT NULL COMMENT '主键ID',
  `avtag` bit(1) NULL DEFAULT NULL COMMENT '可用标记',
  `crtim` datetime(6) NULL DEFAULT NULL COMMENT '创建时间',
  `cruid` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '创建人ID',
  `name` varchar(126) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
  `notes` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `uptim` datetime(6) NULL DEFAULT NULL COMMENT '更新时间',
  `upuid` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新人ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '单一主表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of demo_single
-- ----------------------------

-- ----------------------------
-- Table structure for demo_single_cate
-- ----------------------------
DROP TABLE IF EXISTS `demo_single_cate`;
CREATE TABLE `demo_single_cate`  (
  `id` bigint NOT NULL COMMENT '主键ID',
  `avtag` bit(1) NULL DEFAULT NULL COMMENT '可用标记 1启用，0禁用',
  `crtim` datetime(6) NULL DEFAULT NULL COMMENT '创建时间',
  `cruid` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '创建人ID',
  `label` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标签',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
  `notes` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `ornum` int NULL DEFAULT NULL COMMENT '排序号',
  `pid` bigint NULL DEFAULT NULL COMMENT '父分类ID',
  `tier` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '层级信息',
  `uptim` datetime(6) NULL DEFAULT NULL COMMENT '更新时间',
  `upuid` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新人ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '单一树表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of demo_single_cate
-- ----------------------------

-- ----------------------------
-- Table structure for mon_job
-- ----------------------------
DROP TABLE IF EXISTS `mon_job`;
CREATE TABLE `mon_job`  (
  `id` bigint NOT NULL,
  `avtag` bit(1) NULL DEFAULT NULL,
  `code` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `cron` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `crtim` datetime(6) NULL DEFAULT NULL,
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `notes` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `ornum` int NULL DEFAULT NULL,
  `retyp` int NULL DEFAULT NULL,
  `reurl` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of mon_job
-- ----------------------------

-- ----------------------------
-- Table structure for mon_job_log
-- ----------------------------
DROP TABLE IF EXISTS `mon_job_log`;
CREATE TABLE `mon_job_log`  (
  `id` bigint NOT NULL,
  `entim` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `msg` varchar(5000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `ret` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `sttim` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of mon_job_log
-- ----------------------------

-- ----------------------------
-- Table structure for mon_login_log
-- ----------------------------
DROP TABLE IF EXISTS `mon_login_log`;
CREATE TABLE `mon_login_log`  (
  `id` bigint NOT NULL COMMENT '主键ID',
  `browser` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '浏览器',
  `clkey` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '客户端',
  `detyp` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '设备类型',
  `himsg` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '提示消息',
  `loip` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '登录IP地址',
  `loloc` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '登录地点',
  `lotim` datetime(6) NULL DEFAULT NULL COMMENT '登录时间',
  `os` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '操作系统',
  `sutag` bit(1) NULL DEFAULT NULL COMMENT '登录状态',
  `tenid` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户编号',
  `username` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户账号',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '登录日志' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of mon_login_log
-- ----------------------------

-- ----------------------------
-- Table structure for mon_oper_log
-- ----------------------------
DROP TABLE IF EXISTS `mon_oper_log`;
CREATE TABLE `mon_oper_log`  (
  `id` bigint NOT NULL COMMENT '主键ID',
  `bapar` varchar(4000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '返回参数',
  `butyp` int NULL DEFAULT NULL COMMENT '业务类型',
  `cotim` bigint NULL DEFAULT NULL COMMENT '消耗时间',
  `ermsg` varchar(4000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '错误消息',
  `opdna` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '组织名称',
  `opip` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '操作IP地址',
  `oploc` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '操作地点',
  `opmod` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '操作模块',
  `optim` datetime(6) NULL DEFAULT NULL COMMENT '操作时间',
  `optyp` int NULL DEFAULT NULL COMMENT '操作类别',
  `opuna` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '操作人员',
  `remet` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '请求方法',
  `repar` varchar(4000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '请求参数',
  `reurl` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '请求url',
  `reway` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '请求方式',
  `sutag` bit(1) NULL DEFAULT NULL COMMENT '操作状态',
  `tenid` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户编号',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '操作日志' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of mon_oper_log
-- ----------------------------

-- ----------------------------
-- Table structure for sys_actor
-- ----------------------------
DROP TABLE IF EXISTS `sys_actor`;
CREATE TABLE `sys_actor`  (
  `id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '主键ID',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
  `type` int NULL DEFAULT NULL COMMENT '类型',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统参与者' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_actor
-- ----------------------------
INSERT INTO `sys_actor` VALUES ('g3001', '北京分公司管理组', 8);
INSERT INTO `sys_actor` VALUES ('g3002', '北京分公司销售员', 8);
INSERT INTO `sys_actor` VALUES ('o1000', 'XX科技', 1);
INSERT INTO `sys_actor` VALUES ('o1100', '北京分公司', 1);
INSERT INTO `sys_actor` VALUES ('o1110', '北京分公司销售部', 1);
INSERT INTO `sys_actor` VALUES ('o1111', '北京分公司销售部一组', 1);
INSERT INTO `sys_actor` VALUES ('o1112', '北京分公司销售部二组', 1);
INSERT INTO `sys_actor` VALUES ('o1120', '北京分公司人事部', 1);
INSERT INTO `sys_actor` VALUES ('o1130', '北京分公司财务部', 1);
INSERT INTO `sys_actor` VALUES ('o1140', '北京分公司综合部', 1);
INSERT INTO `sys_actor` VALUES ('o1200', '上海分公司', 1);
INSERT INTO `sys_actor` VALUES ('o1210', '上海分公司销售部', 1);
INSERT INTO `sys_actor` VALUES ('o1220', '上海分公司人事部', 1);
INSERT INTO `sys_actor` VALUES ('o1230', '上海分公司财务部', 1);
INSERT INTO `sys_actor` VALUES ('o1300', '广州分公司', 1);
INSERT INTO `sys_actor` VALUES ('o1310', '广州分公司综合部', 1);
INSERT INTO `sys_actor` VALUES ('o1320', '广州分公司销售部', 1);
INSERT INTO `sys_actor` VALUES ('o1330', '广州分公司人事部', 1);
INSERT INTO `sys_actor` VALUES ('p2001', '董事长', 4);
INSERT INTO `sys_actor` VALUES ('p2002', '北京分公司总经理', 4);
INSERT INTO `sys_actor` VALUES ('p2003', '北京分公司销售部长', 4);
INSERT INTO `sys_actor` VALUES ('p2004', '北京分公司销售经理', 4);
INSERT INTO `sys_actor` VALUES ('u1', '管理员', 2);
INSERT INTO `sys_actor` VALUES ('u2', '小狐狸', 2);
INSERT INTO `sys_actor` VALUES ('u3', '张三', 2);
INSERT INTO `sys_actor` VALUES ('u4', '李四', 2);
INSERT INTO `sys_actor` VALUES ('u5', '王五', 2);
INSERT INTO `sys_actor` VALUES ('u6', '赵六', 2);
INSERT INTO `sys_actor` VALUES ('u7', '孙七', 2);
INSERT INTO `sys_actor` VALUES ('u8', '周八', 2);
INSERT INTO `sys_actor` VALUES ('u9', '吴九', 2);

-- ----------------------------
-- Table structure for sys_api
-- ----------------------------
DROP TABLE IF EXISTS `sys_api`;
CREATE TABLE `sys_api`  (
  `id` bigint NOT NULL COMMENT '主键ID',
  `avtag` bit(1) NULL DEFAULT NULL COMMENT '可用标记',
  `code` bigint NULL DEFAULT NULL COMMENT '权限代码',
  `crtim` datetime(6) NULL DEFAULT NULL COMMENT '创建时间',
  `menid` bigint NULL DEFAULT NULL COMMENT '菜单ID',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
  `notes` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `ornum` int NULL DEFAULT NULL COMMENT '排序号',
  `perm` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '权限字符',
  `pos` int NULL DEFAULT NULL COMMENT '权限位',
  `type` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '权限类型',
  `uptim` datetime(6) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '权限接口' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_api
-- ----------------------------
INSERT INTO `sys_api` VALUES (101001, b'1', NULL, '2026-01-20 15:03:01.161000', 1010, '组织查询', NULL, 101001, 'sys:dept:query', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (101002, b'1', NULL, '2026-01-20 15:03:01.161000', 1010, '组织编辑', NULL, 101002, 'sys:dept:edit', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (101003, b'1', NULL, '2026-01-20 15:03:01.161000', 1010, '组织删除', NULL, 101003, 'sys:dept:delete', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (102001, b'1', NULL, '2026-01-20 15:03:01.161000', 1020, '用户查询', NULL, 102001, 'sys:user:query', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (102002, b'1', NULL, '2026-01-20 15:03:01.161000', 1020, '用户编辑', NULL, 102002, 'sys:user:edit', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (102003, b'1', NULL, '2026-01-20 15:03:01.161000', 1020, '用户删除', NULL, 102003, 'sys:user:delete', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (102004, b'1', NULL, '2026-01-20 15:03:01.161000', 1020, '用户启用禁用', NULL, 102004, 'sys:user:avtag', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (102005, b'1', NULL, '2026-01-20 15:03:01.161000', 1020, '用户密码修改', NULL, 102005, 'sys:user:password', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (103001, b'1', NULL, '2026-01-20 15:03:01.161000', 1030, '岗位查询', NULL, 103001, 'sys:post:query', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (103002, b'1', NULL, '2026-01-20 15:03:01.161000', 1030, '岗位编辑', NULL, 103002, 'sys:post:edit', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (103003, b'1', NULL, '2026-01-20 15:03:01.161000', 1030, '岗位删除', NULL, 103003, 'sys:post:delete', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (104001, b'1', NULL, '2026-01-20 15:03:01.161000', 1040, '群组查询', NULL, 104001, 'sys:group:query', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (104002, b'1', NULL, '2026-01-20 15:03:01.161000', 1040, '群组编辑', NULL, 104002, 'sys:group:edit', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (104003, b'1', NULL, '2026-01-20 15:03:01.161000', 1040, '群组删除', NULL, 104003, 'sys:group:delete', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (104004, b'1', NULL, '2026-01-20 15:03:01.161000', 1040, '群组分类查询', NULL, 104004, 'sys:groupc:query', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (104005, b'1', NULL, '2026-01-20 15:03:01.161000', 1040, '群组分类编辑', NULL, 104005, 'sys:groupc:edit', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (104006, b'1', NULL, '2026-01-20 15:03:01.161000', 1040, '群组分类删除', NULL, 104006, 'sys:groupc:delete', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (105001, b'1', NULL, '2026-01-20 15:03:01.161000', 1050, '菜单查询', NULL, 105001, 'sys:menu:query', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (105002, b'1', NULL, '2026-01-20 15:03:01.161000', 1050, '菜单编辑', NULL, 105002, 'sys:menu:edit', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (105003, b'1', NULL, '2026-01-20 15:03:01.161000', 1050, '菜单删除', NULL, 105003, 'sys:menu:delete', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (106001, b'1', NULL, '2026-01-20 15:03:01.161000', 1060, '接口查询', NULL, 106001, 'sys:api:query', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (106002, b'1', NULL, '2026-01-20 15:03:01.161000', 1060, '接口编辑', NULL, 106002, 'sys:api:edit', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (106003, b'1', NULL, '2026-01-20 15:03:01.161000', 1060, '接口删除', NULL, 106003, 'sys:api:delete', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (107001, b'1', NULL, '2026-01-20 15:03:01.161000', 1070, '角色查询', NULL, 107001, 'sys:role:query', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (107002, b'1', NULL, '2026-01-20 15:03:01.161000', 1070, '角色编辑', NULL, 107002, 'sys:role:edit', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (107003, b'1', NULL, '2026-01-20 15:03:01.161000', 1070, '角色删除', NULL, 107003, 'sys:role:delete', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (108001, b'1', NULL, '2026-01-20 15:03:01.161000', 1080, '参数查询', NULL, 108001, 'sys:config:query', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (108002, b'1', NULL, '2026-01-20 15:03:01.161000', 1080, '参数编辑', NULL, 108002, 'sys:config:edit', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (108003, b'1', NULL, '2026-01-20 15:03:01.161000', 1080, '参数删除', NULL, 108003, 'sys:config:delete', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (109001, b'1', NULL, '2026-01-20 15:03:01.161000', 1090, '通知查询', NULL, 109001, 'sys:notice:query', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (109002, b'1', NULL, '2026-01-20 15:03:01.161000', 1090, '通知编辑', NULL, 109002, 'sys:notice:edit', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (109003, b'1', NULL, '2026-01-20 15:03:01.161000', 1090, '通知删除', NULL, 109003, 'sys:notice:delete', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (201001, b'1', NULL, '2026-01-20 15:03:01.161000', 2010, '在线用户查询', NULL, 201001, 'mon:online:query', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (201002, b'1', NULL, '2026-01-20 15:03:01.161000', 2010, '在线用户强退', NULL, 201002, 'mon:online:delete', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (202001, b'1', NULL, '2026-01-20 15:03:01.161000', 2020, '登录日志查询', NULL, 202001, 'mon:login:query', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (202002, b'1', NULL, '2026-01-20 15:03:01.161000', 2020, '登录日志删除', NULL, 202002, 'mon:login:delete', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (203001, b'1', NULL, '2026-01-20 15:03:01.161000', 2030, '操作日志查询', NULL, 203001, 'mon:oper:query', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (203002, b'1', NULL, '2026-01-20 15:03:01.161000', 2030, '操作日志删除', NULL, 203002, 'mon:oper:delete', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (204001, b'1', NULL, '2026-01-20 15:03:01.161000', 2040, '服务器信息查询', NULL, 204001, 'mon:server:query', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (205001, b'1', NULL, '2026-01-20 15:03:01.161000', 2050, '缓存信息查询', NULL, 205001, 'mon:cache:query', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (206001, b'1', NULL, '2026-01-20 15:03:01.161000', 2060, '定时任务查询', NULL, 206001, 'mon:job:query', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (206002, b'1', NULL, '2026-01-20 15:03:01.161000', 2060, '定时任务修改', NULL, 206002, 'mon:job:edit', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (206003, b'1', NULL, '2026-01-20 15:03:01.161000', 2060, '定时任务执行', NULL, 206003, 'mon:job:run', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (206101, b'1', NULL, '2026-01-20 15:03:01.161000', 2061, '定时任务日志查询', NULL, 206101, 'mon:jobl:query', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (206102, b'1', NULL, '2026-01-20 15:03:01.161000', 2061, '定时任务日志删除', NULL, 206102, 'mon:jobl:delete', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (301001, b'1', NULL, '2026-01-20 15:03:01.161000', 3010, '字典查询', NULL, 301001, 'tool:dcit:query', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (301002, b'1', NULL, '2026-01-20 15:03:01.161000', 3010, '字典修改', NULL, 301002, 'tool:dict:edit', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (301003, b'1', NULL, '2026-01-20 15:03:01.161000', 3010, '字典删除', NULL, 301003, 'tool:dict:delete', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (301004, b'1', NULL, '2026-01-20 15:03:01.161000', 3010, '字典数据查询', NULL, 301004, 'tool:dictd:query', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (301005, b'1', NULL, '2026-01-20 15:03:01.161000', 3010, '字典数据编辑', NULL, 301005, 'tool:dictd:edit', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (301006, b'1', NULL, '2026-01-20 15:03:01.161000', 3010, '字典数据删除', NULL, 301006, 'tool:dictd:delete', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (302001, b'1', NULL, '2026-01-20 15:03:01.161000', 3020, '编号查询', NULL, 302001, 'tool:num:query', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (302002, b'1', NULL, '2026-01-20 15:03:01.161000', 3020, '编号编辑', NULL, 302002, 'tool:num:edit', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (302003, b'1', NULL, '2026-01-20 15:03:01.161000', 3020, '编号删除', NULL, 302003, 'tool:num:delete', NULL, NULL, '2026-01-20 15:03:01.161000');
INSERT INTO `sys_api` VALUES (603001, b'1', NULL, '2026-01-20 15:03:01.223000', 6030, '流程查询', NULL, 603001, 'bpmbus:main:query', NULL, NULL, '2026-01-20 15:03:01.223000');
INSERT INTO `sys_api` VALUES (603002, b'1', NULL, '2026-01-20 15:03:01.223000', 6030, '流程新增', NULL, 603002, 'bpmbus:main:add', NULL, NULL, '2026-01-20 15:03:01.223000');
INSERT INTO `sys_api` VALUES (603003, b'1', NULL, '2026-01-20 15:03:01.223000', 6030, '流程编辑', NULL, 603003, 'bpmbus:main:edit', NULL, NULL, '2026-01-20 15:03:01.223000');
INSERT INTO `sys_api` VALUES (801001, b'1', NULL, '2026-01-20 15:03:01.435000', 8010, '单一主表-查询', NULL, 801001, 'demo:single:query', NULL, NULL, '2026-01-20 15:03:01.435000');
INSERT INTO `sys_api` VALUES (801002, b'1', NULL, '2026-01-20 15:03:01.435000', 8010, '单一主表-新增', NULL, 801002, 'demo:single:add', NULL, NULL, '2026-01-20 15:03:01.435000');
INSERT INTO `sys_api` VALUES (801003, b'1', NULL, '2026-01-20 15:03:01.435000', 8010, '单一主表-修改', NULL, 801003, 'demo:single:edit', NULL, NULL, '2026-01-20 15:03:01.435000');
INSERT INTO `sys_api` VALUES (801004, b'1', NULL, '2026-01-20 15:03:01.435000', 8010, '单一主表-删除', NULL, 801004, 'demo:single:remove', NULL, NULL, '2026-01-20 15:03:01.435000');
INSERT INTO `sys_api` VALUES (802001, b'1', NULL, '2026-01-20 15:03:01.435000', 8020, '单一树表-查询', NULL, 802001, 'demo:singlec:query', NULL, NULL, '2026-01-20 15:03:01.435000');
INSERT INTO `sys_api` VALUES (802002, b'1', NULL, '2026-01-20 15:03:01.435000', 8020, '单一树表-新增', NULL, 802002, 'demo:singlec:add', NULL, NULL, '2026-01-20 15:03:01.435000');
INSERT INTO `sys_api` VALUES (802003, b'1', NULL, '2026-01-20 15:03:01.435000', 8020, '单一树表-修改', NULL, 802003, 'demo:singlec:edit', NULL, NULL, '2026-01-20 15:03:01.435000');
INSERT INTO `sys_api` VALUES (802004, b'1', NULL, '2026-01-20 15:03:01.435000', 8020, '单一树表-删除', NULL, 802004, 'demo:singlec:remove', NULL, NULL, '2026-01-20 15:03:01.435000');
INSERT INTO `sys_api` VALUES (803001, b'1', NULL, '2026-01-20 15:03:01.435000', 8030, '关联主表-查询', NULL, 803001, 'demo:link:query', NULL, NULL, '2026-01-20 15:03:01.435000');
INSERT INTO `sys_api` VALUES (803002, b'1', NULL, '2026-01-20 15:03:01.435000', 8030, '关联主表-新增', NULL, 803002, 'demo:link:add', NULL, NULL, '2026-01-20 15:03:01.435000');
INSERT INTO `sys_api` VALUES (803003, b'1', NULL, '2026-01-20 15:03:01.435000', 8030, '关联主表-修改', NULL, 803003, 'demo:link:edit', NULL, NULL, '2026-01-20 15:03:01.435000');
INSERT INTO `sys_api` VALUES (803004, b'1', NULL, '2026-01-20 15:03:01.435000', 8030, '关联主表-删除', NULL, 803004, 'demo:link:remove', NULL, NULL, '2026-01-20 15:03:01.435000');
INSERT INTO `sys_api` VALUES (803011, b'1', NULL, '2026-01-20 15:03:01.435000', 8030, '关联树表-查询', NULL, 803011, 'demo:linkc:query', NULL, NULL, '2026-01-20 15:03:01.435000');
INSERT INTO `sys_api` VALUES (803012, b'1', NULL, '2026-01-20 15:03:01.435000', 8030, '关联树表-新增', NULL, 803012, 'demo:linkc:add', NULL, NULL, '2026-01-20 15:03:01.435000');
INSERT INTO `sys_api` VALUES (803013, b'1', NULL, '2026-01-20 15:03:01.435000', 8030, '关联树表-修改', NULL, 803013, 'demo:linkc:edit', NULL, NULL, '2026-01-20 15:03:01.435000');
INSERT INTO `sys_api` VALUES (803014, b'1', NULL, '2026-01-20 15:03:01.435000', 8030, '关联树表-删除', NULL, 803014, 'demo:linkc:remove', NULL, NULL, '2026-01-20 15:03:01.435000');

-- ----------------------------
-- Table structure for sys_config
-- ----------------------------
DROP TABLE IF EXISTS `sys_config`;
CREATE TABLE `sys_config`  (
  `id` bigint NOT NULL COMMENT '主键ID',
  `avtag` bit(1) NULL DEFAULT NULL,
  `crtim` datetime(6) NULL DEFAULT NULL,
  `intag` bit(1) NULL DEFAULT NULL,
  `kenam` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '参数键名',
  `keval` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '参数键值',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '参数名称',
  `notes` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `ornum` int NULL DEFAULT NULL,
  `uptim` datetime(6) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统参数' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_config
-- ----------------------------
INSERT INTO `sys_config` VALUES (1, b'1', '2026-01-20 15:03:01.374000', b'1', 'sys.user.initPassword', '123456', '用户管理-账号初始密码', NULL, NULL, '2026-01-20 15:03:01.374000');
INSERT INTO `sys_config` VALUES (2, b'1', '2026-01-20 15:03:01.374000', b'1', 'sys.account.registerUser', 'false', '账号自助-是否开启用户注册功能', NULL, NULL, '2026-01-20 15:03:01.374000');

-- ----------------------------
-- Table structure for sys_corp
-- ----------------------------
DROP TABLE IF EXISTS `sys_corp`;
CREATE TABLE `sys_corp`  (
  `id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '主键ID',
  `avtag` bit(1) NULL DEFAULT NULL COMMENT '可用标记',
  `catid` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '分类ID',
  `crtim` datetime(6) NULL DEFAULT NULL COMMENT '创建时间',
  `cruid` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '创建人ID',
  `label` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标签',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
  `notes` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `ornum` int NULL DEFAULT NULL COMMENT '排序号',
  `type` int NULL DEFAULT NULL COMMENT '公司类型',
  `uptim` datetime(6) NULL DEFAULT NULL COMMENT '更新时间',
  `upuid` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新人ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '公司' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_corp
-- ----------------------------

-- ----------------------------
-- Table structure for sys_corp_cate
-- ----------------------------
DROP TABLE IF EXISTS `sys_corp_cate`;
CREATE TABLE `sys_corp_cate`  (
  `id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '主键ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
  `ornum` int NULL DEFAULT NULL COMMENT '排序号',
  `pid` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '父ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '公司分类' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_corp_cate
-- ----------------------------

-- ----------------------------
-- Table structure for sys_group
-- ----------------------------
DROP TABLE IF EXISTS `sys_group`;
CREATE TABLE `sys_group`  (
  `id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '主键ID',
  `avtag` bit(1) NULL DEFAULT NULL COMMENT '可用标记',
  `catid` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '分类ID',
  `crtim` datetime(6) NULL DEFAULT NULL COMMENT '创建时间',
  `cruid` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '创建人ID',
  `label` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标签',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
  `notes` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `ornum` int NULL DEFAULT NULL COMMENT '排序号',
  `uptim` datetime(6) NULL DEFAULT NULL COMMENT '更新时间',
  `upuid` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新人ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统群组' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_group
-- ----------------------------
INSERT INTO `sys_group` VALUES ('g3001', b'1', NULL, '2026-01-20 15:03:01.088000', NULL, NULL, '北京分公司管理组', NULL, 3001, NULL, NULL);
INSERT INTO `sys_group` VALUES ('g3002', b'1', NULL, '2026-01-20 15:03:01.093000', NULL, NULL, '北京分公司销售员', NULL, 3002, NULL, NULL);

-- ----------------------------
-- Table structure for sys_group_actor
-- ----------------------------
DROP TABLE IF EXISTS `sys_group_actor`;
CREATE TABLE `sys_group_actor`  (
  `gid` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `aid` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  INDEX `FKip1xv3f82g83rcx576crjxgu4`(`aid` ASC) USING BTREE,
  INDEX `FKe960yje7qotse51ycr7b4qcf3`(`gid` ASC) USING BTREE,
  CONSTRAINT `FKe960yje7qotse51ycr7b4qcf3` FOREIGN KEY (`gid`) REFERENCES `sys_group` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `FKip1xv3f82g83rcx576crjxgu4` FOREIGN KEY (`aid`) REFERENCES `sys_actor` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_group_actor
-- ----------------------------
INSERT INTO `sys_group_actor` VALUES ('g3001', 'u4');
INSERT INTO `sys_group_actor` VALUES ('g3001', 'u5');
INSERT INTO `sys_group_actor` VALUES ('g3001', 'p2004');
INSERT INTO `sys_group_actor` VALUES ('g3001', 'o1140');
INSERT INTO `sys_group_actor` VALUES ('g3002', 'u8');
INSERT INTO `sys_group_actor` VALUES ('g3002', 'u9');

-- ----------------------------
-- Table structure for sys_group_cate
-- ----------------------------
DROP TABLE IF EXISTS `sys_group_cate`;
CREATE TABLE `sys_group_cate`  (
  `id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '主键ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
  `ornum` int NULL DEFAULT NULL COMMENT '排序号',
  `pid` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '父ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统群组分类' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_group_cate
-- ----------------------------

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu`  (
  `id` bigint NOT NULL COMMENT '主键ID',
  `avtag` bit(1) NULL DEFAULT NULL COMMENT '可用标记 1启用，0禁用',
  `catag` bit(1) NULL DEFAULT NULL COMMENT '缓存标记',
  `comp` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '组件路径',
  `crtim` datetime(6) NULL DEFAULT NULL COMMENT '创建时间',
  `cruid` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `icon` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '图标',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
  `notes` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `ornum` int NULL DEFAULT NULL COMMENT '排序号',
  `outag` bit(1) NULL DEFAULT NULL COMMENT '外链标记',
  `param` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '路由参数',
  `path` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '路由路径',
  `pid` bigint NULL DEFAULT NULL COMMENT '父ID',
  `shtag` bit(1) NULL DEFAULT NULL COMMENT '显示标记',
  `type` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '类型',
  `uptim` datetime(6) NULL DEFAULT NULL COMMENT '更新时间',
  `upuid` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统菜单' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
INSERT INTO `sys_menu` VALUES (1000, b'1', b'0', 'Layout', '2026-01-20 15:03:01.104000', NULL, 'tdesign:system-setting', '系统管理', NULL, 1000, b'0', NULL, 'sys', 0, b'1', '1', '2026-01-20 15:03:01.104000', NULL);
INSERT INTO `sys_menu` VALUES (1010, b'1', b'0', 'sys/org/index', '2026-01-20 15:03:01.104000', NULL, 'mingcute:department-line', '组织管理', NULL, 1010, b'0', NULL, 'org', 1000, b'1', '2', '2026-01-20 15:03:01.104000', NULL);
INSERT INTO `sys_menu` VALUES (1020, b'1', b'0', 'sys/user/index', '2026-01-20 15:03:01.104000', NULL, 'ant-design:user-outlined', '用户管理', NULL, 1020, b'0', NULL, 'user', 1000, b'1', '2', '2026-01-20 15:03:01.104000', NULL);
INSERT INTO `sys_menu` VALUES (1021, b'1', b'1', 'sys/user/tedit', '2026-01-20 15:03:01.104000', NULL, 'mingcute:user-edit-line', '用户编辑', NULL, 1021, b'0', NULL, 'user/edit', 1000, b'0', '2', '2026-01-20 15:03:01.104000', NULL);
INSERT INTO `sys_menu` VALUES (1030, b'1', b'0', 'sys/post/index', '2026-01-20 15:03:01.104000', NULL, 'icon-park-outline:appointment', '岗位管理', NULL, 1030, b'0', NULL, 'post', 1000, b'1', '2', '2026-01-20 15:03:01.104000', NULL);
INSERT INTO `sys_menu` VALUES (1040, b'1', b'0', 'sys/group/index', '2026-01-20 15:03:01.104000', NULL, 'material-symbols:group-outline-rounded', '群组管理', NULL, 1040, b'0', NULL, 'group', 1000, b'1', '2', '2026-01-20 15:03:01.104000', NULL);
INSERT INTO `sys_menu` VALUES (1050, b'1', b'0', 'sys/menu/index', '2026-01-20 15:03:01.104000', NULL, 'ri:menu-fold-2-fill', '菜单管理', NULL, 1050, b'0', NULL, 'menu', 1000, b'1', '2', '2026-01-20 15:03:01.104000', NULL);
INSERT INTO `sys_menu` VALUES (1060, b'1', b'0', 'sys/api/index', '2026-01-20 15:03:01.104000', NULL, 'ant-design:api-outlined', '接口管理', NULL, 1060, b'0', NULL, 'api', 1000, b'1', '2', '2026-01-20 15:03:01.104000', NULL);
INSERT INTO `sys_menu` VALUES (1070, b'1', b'0', 'sys/role/index', '2026-01-20 15:03:01.104000', NULL, 'eos-icons:role-binding-outlined', '角色管理', NULL, 1070, b'0', NULL, 'role', 1000, b'1', '2', '2026-01-20 15:03:01.104000', NULL);
INSERT INTO `sys_menu` VALUES (1071, b'1', b'1', 'sys/role/edit', '2026-01-20 15:03:01.104000', NULL, 'oui:app-users-roles', '角色编辑', NULL, 1071, b'0', NULL, 'role/edit', 1000, b'0', '2', '2026-01-20 15:03:01.104000', NULL);
INSERT INTO `sys_menu` VALUES (1080, b'1', b'0', 'sys/config/index', '2026-01-20 15:03:01.104000', NULL, 'ant-design:setting-outlined', '参数设置', NULL, 1080, b'0', NULL, 'config', 1000, b'1', '2', '2026-01-20 15:03:01.104000', NULL);
INSERT INTO `sys_menu` VALUES (1090, b'1', b'0', 'sys/notice/index', '2026-01-20 15:03:01.104000', NULL, 'fe:notice-push', '通知公告', NULL, 1090, b'0', NULL, 'notice', 1000, b'1', '2', '2026-01-20 15:03:01.104000', NULL);
INSERT INTO `sys_menu` VALUES (2000, b'1', b'0', 'Layout', '2026-01-20 15:03:01.104000', NULL, 'eos-icons:monitoring', '监控中心', NULL, 2000, b'0', NULL, 'mon', 0, b'1', '1', '2026-01-20 15:03:01.104000', NULL);
INSERT INTO `sys_menu` VALUES (2010, b'1', b'0', 'mon/online/user/index', '2026-01-20 15:03:01.104000', NULL, 'oui:online', '在线用户', NULL, 2010, b'0', NULL, 'online', 2000, b'1', '2', '2026-01-20 15:03:01.104000', NULL);
INSERT INTO `sys_menu` VALUES (2020, b'1', b'0', 'mon/login/log/index', '2026-01-20 15:03:01.104000', NULL, 'uiw:login', '登录日志', NULL, 2020, b'0', NULL, 'login', 2000, b'1', '2', '2026-01-20 15:03:01.104000', NULL);
INSERT INTO `sys_menu` VALUES (2030, b'1', b'0', 'mon/oper/log/index', '2026-01-20 15:03:01.104000', NULL, 'icon-park-outline:reverse-operation-in', '操作日志', NULL, 2030, b'0', NULL, 'oper', 2000, b'1', '2', '2026-01-20 15:03:01.104000', NULL);
INSERT INTO `sys_menu` VALUES (2040, b'1', b'0', 'mon/server/index', '2026-01-20 15:03:01.104000', NULL, 'mdi:server-outline', '服务监控', NULL, 2040, b'0', NULL, 'server', 2000, b'1', '2', '2026-01-20 15:03:01.104000', NULL);
INSERT INTO `sys_menu` VALUES (2050, b'1', b'0', 'mon/cache/index', '2026-01-20 15:03:01.104000', NULL, 'octicon:cache-24', '缓存监控', NULL, 2050, b'0', NULL, 'cache', 2000, b'1', '2', '2026-01-20 15:03:01.104000', NULL);
INSERT INTO `sys_menu` VALUES (2060, b'1', b'0', 'mon/job/main/index', '2026-01-20 15:03:01.104000', NULL, 'streamline:task-list', '定时任务', NULL, 2060, b'0', NULL, 'job', 2000, b'1', '2', '2026-01-20 15:03:01.104000', NULL);
INSERT INTO `sys_menu` VALUES (2061, b'1', b'0', 'mon/job/log/index', '2026-01-20 15:03:01.104000', NULL, 'ix:log', '任务日志', NULL, 2061, b'0', NULL, 'jobl', 2000, b'0', '2', '2026-01-20 15:03:01.104000', NULL);
INSERT INTO `sys_menu` VALUES (3000, b'1', b'0', 'Layout', '2026-01-20 15:03:01.104000', NULL, 'ant-design:tool-outlined', '辅助工具', NULL, 3000, b'0', NULL, 'tool', 0, b'1', '1', '2026-01-20 15:03:01.104000', NULL);
INSERT INTO `sys_menu` VALUES (3010, b'1', b'0', 'tool/dict/index', '2026-01-20 15:03:01.104000', NULL, 'fluent-mdl2:dictionary', '字典工具', NULL, 3010, b'0', NULL, 'dict', 3000, b'1', '2', '2026-01-20 15:03:01.104000', NULL);
INSERT INTO `sys_menu` VALUES (3020, b'1', b'0', 'tool/num/index', '2026-01-20 15:03:01.104000', NULL, 'streamline-sharp:steps-number', '编号工具', NULL, 3020, b'0', NULL, 'num', 3000, b'1', '2', '2026-01-20 15:03:01.104000', NULL);
INSERT INTO `sys_menu` VALUES (3030, b'1', b'0', 'tool/oss/main/index', '2026-01-20 15:03:01.104000', NULL, 'mdi:file-outline', '文件工具', NULL, 3030, b'0', NULL, 'oss', 3000, b'1', '2', '2026-01-20 15:03:01.104000', NULL);
INSERT INTO `sys_menu` VALUES (3040, b'1', b'0', 'tool/form/index', '2026-01-20 15:03:01.104000', NULL, 'fluent:form-20-regular', '在线表单', NULL, 3040, b'0', NULL, 'form', 3000, b'1', '2', '2026-01-20 15:03:01.104000', NULL);
INSERT INTO `sys_menu` VALUES (3041, b'1', b'1', 'tool/form/edit', '2026-01-20 15:03:01.104000', NULL, 'fluent:form-20-regular', '在线表单', NULL, 3041, b'0', NULL, 'form/edit', 3000, b'0', '2', '2026-01-20 15:03:01.104000', NULL);
INSERT INTO `sys_menu` VALUES (3050, b'1', b'0', 'tool/code/index', '2026-01-20 15:03:01.104000', NULL, 'humbleicons:code', '代码生成', NULL, 3050, b'0', NULL, 'code', 3000, b'1', '2', '2026-01-20 15:03:01.104000', NULL);
INSERT INTO `sys_menu` VALUES (3051, b'1', b'0', 'tool/code/edit', '2026-01-20 15:03:01.104000', NULL, 'humbleicons:code', '代码生成', NULL, 3051, b'0', NULL, 'code/edit', 3000, b'0', '2', '2026-01-20 15:03:01.104000', NULL);
INSERT INTO `sys_menu` VALUES (6000, b'1', b'0', 'Layout', '2026-01-20 15:03:01.150000', NULL, 'streamline-sharp:text-flow-rows', '流程管理', NULL, 6000, b'0', NULL, 'bpm', 0, b'1', '1', '2026-01-20 15:03:01.150000', NULL);
INSERT INTO `sys_menu` VALUES (6010, b'1', b'0', 'bpm/bus/cate/index', '2026-01-20 15:03:01.150000', NULL, 'tabler:category-plus', '流程分类', NULL, 6010, b'0', NULL, 'busc', 6000, b'1', '2', '2026-01-20 15:03:01.150000', NULL);
INSERT INTO `sys_menu` VALUES (6020, b'1', b'0', 'bpm/bus/tmpl/index', '2026-01-20 15:03:01.150000', NULL, 'carbon:prompt-template', '流程模板', NULL, 6020, b'0', NULL, 'bust', 6000, b'1', '2', '2026-01-20 15:03:01.150000', NULL);
INSERT INTO `sys_menu` VALUES (6021, b'1', b'1', 'bpm/bus/tmpl/edit', '2026-01-20 15:03:01.150000', NULL, 'carbon:prompt-template', '流程模板编辑', NULL, 6021, b'0', NULL, 'bust/edit', 6000, b'0', '2', '2026-01-20 15:03:01.150000', NULL);
INSERT INTO `sys_menu` VALUES (6030, b'1', b'0', 'bpm/bus/main/index', '2026-01-20 15:03:01.150000', NULL, 'ri:instance-line', '流程清单', NULL, 6030, b'0', NULL, 'bus', 6000, b'1', '2', '2026-01-20 15:03:01.150000', NULL);
INSERT INTO `sys_menu` VALUES (6031, b'1', b'1', 'bpm/bus/main/edit', '2026-01-20 15:03:01.150000', NULL, 'ri:instance-line', '流程编辑', NULL, 6031, b'0', NULL, 'bus/edit', 6000, b'0', '2', '2026-01-20 15:03:01.150000', NULL);
INSERT INTO `sys_menu` VALUES (6032, b'1', b'1', 'bpm/bus/main/view', '2026-01-20 15:03:01.150000', NULL, 'ri:instance-line', '流程查看', NULL, 6032, b'0', NULL, 'bus/view', 6000, b'0', '2', '2026-01-20 15:03:01.150000', NULL);
INSERT INTO `sys_menu` VALUES (6040, b'1', b'0', 'bpm/todo/index', '2026-01-20 15:03:01.150000', NULL, 'ri:todo-line', '流程待办', NULL, 6040, b'0', NULL, 'todo', 6000, b'1', '2', '2026-01-20 15:03:01.150000', NULL);
INSERT INTO `sys_menu` VALUES (6050, b'1', b'0', 'bpm/role/tree/index', '2026-01-20 15:03:01.150000', NULL, 'mdi:workflow-outline', '流程角色', NULL, 6050, b'0', NULL, 'role', 6000, b'1', '2', '2026-01-20 15:03:01.150000', NULL);
INSERT INTO `sys_menu` VALUES (6051, b'1', b'0', 'bpm/role/node/index', '2026-01-20 15:03:01.150000', NULL, 'mdi:workflow-outline', '流程角色节点', NULL, 6051, b'0', NULL, 'role/node', 6000, b'0', '2', '2026-01-20 15:03:01.150000', NULL);
INSERT INTO `sys_menu` VALUES (8000, b'1', b'0', 'Layout', '2026-01-20 15:03:01.429000', NULL, 'hugeicons:star', '使用案例', NULL, 8000, b'0', NULL, 'demo', 0, b'1', '1', '2026-01-20 15:03:01.429000', NULL);
INSERT INTO `sys_menu` VALUES (8010, b'1', b'0', 'demo/single/main/index', '2026-01-20 15:03:01.429000', NULL, 'pajamas:work-item-requirement', '单一主表案例', NULL, 8010, b'0', NULL, 'single', 8000, b'1', '2', '2026-01-20 15:03:01.429000', NULL);
INSERT INTO `sys_menu` VALUES (8020, b'1', b'0', 'demo/single/cate/index', '2026-01-20 15:03:01.429000', NULL, 'pajamas:work-item-requirement', '单一树表案例', NULL, 8020, b'0', NULL, 'singlec', 8000, b'1', '2', '2026-01-20 15:03:01.429000', NULL);
INSERT INTO `sys_menu` VALUES (8030, b'1', b'0', 'demo/link/index', '2026-01-20 15:03:01.429000', NULL, 'pajamas:work-item-requirement', '关联主分子案例', NULL, 8030, b'0', NULL, 'link', 8000, b'1', '2', '2026-01-20 15:03:01.429000', NULL);

-- ----------------------------
-- Table structure for sys_notice
-- ----------------------------
DROP TABLE IF EXISTS `sys_notice`;
CREATE TABLE `sys_notice`  (
  `id` bigint NOT NULL COMMENT '主键ID',
  `avtag` bit(1) NULL DEFAULT NULL COMMENT '可用标记',
  `cont` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '公告内容',
  `crtim` datetime(6) NULL DEFAULT NULL COMMENT '创建时间',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '公告标题',
  `notes` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `ornum` int NULL DEFAULT NULL COMMENT '排序号',
  `type` int NULL DEFAULT NULL COMMENT '公告类型（1通知 2公告）',
  `uptim` datetime(6) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '通知公告' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_notice
-- ----------------------------
INSERT INTO `sys_notice` VALUES (1, b'1', '系统将于今天晚上20点到22点进行停机维护，请提前做好工作安排', '2026-01-20 15:03:01.379000', '系统停机公告', NULL, 1, 1, '2026-01-20 15:03:01.379000');

-- ----------------------------
-- Table structure for sys_org
-- ----------------------------
DROP TABLE IF EXISTS `sys_org`;
CREATE TABLE `sys_org`  (
  `id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '主键ID',
  `avtag` bit(1) NULL DEFAULT NULL COMMENT '可用标记',
  `crtim` datetime(6) NULL DEFAULT NULL COMMENT '创建时间',
  `cruid` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '创建人ID',
  `label` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标签',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
  `notes` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `ornum` int NULL DEFAULT NULL COMMENT '排序号',
  `pid` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '父ID',
  `tier` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '层级',
  `type` int NULL DEFAULT NULL COMMENT '组织类型',
  `uptim` datetime(6) NULL DEFAULT NULL COMMENT '更新时间',
  `upuid` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新人ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统组织（公司，机构，部门）' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_org
-- ----------------------------
INSERT INTO `sys_org` VALUES ('o1000', b'1', '2026-01-20 15:03:00.881000', NULL, NULL, 'XX科技', NULL, 1000, NULL, '_o1000_', 1, NULL, NULL);
INSERT INTO `sys_org` VALUES ('o1100', b'1', '2026-01-20 15:03:00.881000', NULL, NULL, '北京分公司', NULL, 1100, 'o1000', '_o1000_o1100_', 2, NULL, NULL);
INSERT INTO `sys_org` VALUES ('o1110', b'1', '2026-01-20 15:03:00.881000', NULL, NULL, '北京分公司销售部', NULL, 1110, 'o1100', '_o1000_o1100_o1110_', 8, NULL, NULL);
INSERT INTO `sys_org` VALUES ('o1111', b'1', '2026-01-20 15:03:00.881000', NULL, NULL, '北京分公司销售部一组', NULL, 1111, 'o1110', '_o1000_o1100_o1110_o1111_', 8, NULL, NULL);
INSERT INTO `sys_org` VALUES ('o1112', b'1', '2026-01-20 15:03:00.881000', NULL, NULL, '北京分公司销售部二组', NULL, 1112, 'o1110', '_o1000_o1100_o1110_o1112_', 8, NULL, NULL);
INSERT INTO `sys_org` VALUES ('o1120', b'1', '2026-01-20 15:03:00.881000', NULL, NULL, '北京分公司人事部', NULL, 1120, 'o1100', '_o1000_o1100_o1120_', 8, NULL, NULL);
INSERT INTO `sys_org` VALUES ('o1130', b'1', '2026-01-20 15:03:00.881000', NULL, NULL, '北京分公司财务部', NULL, 1130, 'o1100', '_o1000_o1100_o1130_', 8, NULL, NULL);
INSERT INTO `sys_org` VALUES ('o1140', b'1', '2026-01-20 15:03:00.881000', NULL, NULL, '北京分公司综合部', NULL, 1140, 'o1100', '_o1000_o1100_o1410_', 8, NULL, NULL);
INSERT INTO `sys_org` VALUES ('o1200', b'1', '2026-01-20 15:03:00.881000', NULL, NULL, '上海分公司', NULL, 1200, 'o1000', '_o1000_o1200_', 2, NULL, NULL);
INSERT INTO `sys_org` VALUES ('o1210', b'1', '2026-01-20 15:03:00.881000', NULL, NULL, '上海分公司销售部', NULL, 1210, 'o1200', '_o1000_o1200_o1210_', 8, NULL, NULL);
INSERT INTO `sys_org` VALUES ('o1220', b'1', '2026-01-20 15:03:00.881000', NULL, NULL, '上海分公司人事部', NULL, 1220, 'o1200', '_o1000_o1200_o1220_', 8, NULL, NULL);
INSERT INTO `sys_org` VALUES ('o1230', b'1', '2026-01-20 15:03:00.881000', NULL, NULL, '上海分公司财务部', NULL, 1230, 'o1200', '_o1000_o1200_o1230_', 8, NULL, NULL);
INSERT INTO `sys_org` VALUES ('o1300', b'1', '2026-01-20 15:03:00.881000', NULL, NULL, '广州分公司', NULL, 1300, 'o1000', '_o1000_o1300_', 2, NULL, NULL);
INSERT INTO `sys_org` VALUES ('o1310', b'1', '2026-01-20 15:03:00.881000', NULL, NULL, '广州分公司综合部', NULL, 1310, 'o1300', '_o1000_o1300_o1310_', 8, NULL, NULL);
INSERT INTO `sys_org` VALUES ('o1320', b'1', '2026-01-20 15:03:00.881000', NULL, NULL, '广州分公司销售部', NULL, 1320, 'o1300', '_o1000_o1300_o1320_', 8, NULL, NULL);
INSERT INTO `sys_org` VALUES ('o1330', b'1', '2026-01-20 15:03:00.881000', NULL, NULL, '广州分公司人事部', NULL, 1330, 'o1300', '_o1000_o1300_o1330_', 8, NULL, NULL);

-- ----------------------------
-- Table structure for sys_oss
-- ----------------------------
DROP TABLE IF EXISTS `sys_oss`;
CREATE TABLE `sys_oss`  (
  `oss_id` bigint NOT NULL,
  `create_by` bigint NULL DEFAULT NULL,
  `create_dept` bigint NULL DEFAULT NULL,
  `create_time` datetime(6) NULL DEFAULT NULL,
  `ext1` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `file_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `file_suffix` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `original_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `search_value` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `service` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `tenant_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `update_by` bigint NULL DEFAULT NULL,
  `update_time` datetime(6) NULL DEFAULT NULL,
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`oss_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_oss
-- ----------------------------

-- ----------------------------
-- Table structure for sys_oss_config
-- ----------------------------
DROP TABLE IF EXISTS `sys_oss_config`;
CREATE TABLE `sys_oss_config`  (
  `oss_config_id` bigint NOT NULL,
  `access_key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `access_policy` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `bucket_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `config_key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `create_by` bigint NULL DEFAULT NULL,
  `create_dept` bigint NULL DEFAULT NULL,
  `create_time` datetime(6) NULL DEFAULT NULL,
  `domain` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `endpoint` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `ext1` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `is_https` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `prefix` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `region` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `search_value` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `secret_key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `status` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `update_by` bigint NULL DEFAULT NULL,
  `update_time` datetime(6) NULL DEFAULT NULL,
  PRIMARY KEY (`oss_config_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_oss_config
-- ----------------------------

-- ----------------------------
-- Table structure for sys_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_post`;
CREATE TABLE `sys_post`  (
  `id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '主键ID',
  `avtag` bit(1) NULL DEFAULT NULL COMMENT '可用标记',
  `crtim` datetime(6) NULL DEFAULT NULL COMMENT '创建时间',
  `cruid` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '创建人ID',
  `label` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标签',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
  `notes` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `orgid` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '组织ID',
  `ornum` int NULL DEFAULT NULL COMMENT '排序号',
  `tier` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '层级',
  `uptim` datetime(6) NULL DEFAULT NULL COMMENT '更新时间',
  `upuid` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新人ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统岗位' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_post
-- ----------------------------
INSERT INTO `sys_post` VALUES ('p2001', b'1', '2026-01-20 15:03:01.046000', NULL, NULL, '董事长', NULL, 'o1000', 2001, '_o1000_p2001_', NULL, NULL);
INSERT INTO `sys_post` VALUES ('p2002', b'1', '2026-01-20 15:03:01.056000', NULL, NULL, '北京分公司总经理', NULL, 'o1100', 2002, '_o1000_o1100_p2002_', NULL, NULL);
INSERT INTO `sys_post` VALUES ('p2003', b'1', '2026-01-20 15:03:01.056000', NULL, NULL, '北京分公司销售部长', NULL, 'o1110', 2003, '_o1000_o1100_o1110_p2003_', NULL, NULL);
INSERT INTO `sys_post` VALUES ('p2004', b'1', '2026-01-20 15:03:01.056000', NULL, NULL, '北京分公司销售经理', NULL, 'o1111', 2004, '_o1000_o1100_o1110_o1111_p2004_', NULL, NULL);

-- ----------------------------
-- Table structure for sys_post_actor
-- ----------------------------
DROP TABLE IF EXISTS `sys_post_actor`;
CREATE TABLE `sys_post_actor`  (
  `pid` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `aid` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  INDEX `FK7kcaee2ekwfo9jc1s94r21n3e`(`aid` ASC) USING BTREE,
  INDEX `FKbrjugfpca985cebsdwqhefj35`(`pid` ASC) USING BTREE,
  CONSTRAINT `FK7kcaee2ekwfo9jc1s94r21n3e` FOREIGN KEY (`aid`) REFERENCES `sys_actor` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `FKbrjugfpca985cebsdwqhefj35` FOREIGN KEY (`pid`) REFERENCES `sys_post` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_post_actor
-- ----------------------------
INSERT INTO `sys_post_actor` VALUES ('p2001', 'u3');
INSERT INTO `sys_post_actor` VALUES ('p2002', 'u4');
INSERT INTO `sys_post_actor` VALUES ('p2003', 'u5');
INSERT INTO `sys_post_actor` VALUES ('p2004', 'u6');
INSERT INTO `sys_post_actor` VALUES ('p2004', 'u7');

-- ----------------------------
-- Table structure for sys_rece
-- ----------------------------
DROP TABLE IF EXISTS `sys_rece`;
CREATE TABLE `sys_rece`  (
  `id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '主键ID',
  `aid` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '最近使用的系统参与者ID',
  `uptim` datetime(6) NULL DEFAULT NULL COMMENT '最近使用时间',
  `useid` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统参与者-最近访问记录' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_rece
-- ----------------------------

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role`  (
  `id` bigint NOT NULL COMMENT '主键ID',
  `avtag` bit(1) NULL DEFAULT NULL COMMENT '可用标记',
  `crtim` datetime(6) NULL DEFAULT NULL COMMENT '创建时间',
  `cruid` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '创建人ID',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '角色名称',
  `notes` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `ornum` int NULL DEFAULT NULL COMMENT '排序号',
  `scope` int NULL DEFAULT NULL COMMENT '数据权限',
  `type` int NULL DEFAULT NULL COMMENT '角色类型',
  `uptim` datetime(6) NULL DEFAULT NULL COMMENT '更新时间',
  `upuid` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新人ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '权限角色' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES (1, b'1', '2026-01-20 15:03:01.229000', NULL, '管理员', '拥有所有权限', 1, NULL, NULL, '2026-01-20 15:03:01.229000', NULL);
INSERT INTO `sys_role` VALUES (2, b'1', '2026-01-20 15:03:01.303000', NULL, '普通用户', '只包含流程使用权限', 2, NULL, NULL, '2026-01-20 15:03:01.303000', NULL);

-- ----------------------------
-- Table structure for sys_role_actor
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_actor`;
CREATE TABLE `sys_role_actor`  (
  `rid` bigint NOT NULL,
  `aid` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  INDEX `FKfny0antg5w2x6no0kw63c6xm3`(`aid` ASC) USING BTREE,
  INDEX `FKi7d9sn34a15ly7ie76kqng40r`(`rid` ASC) USING BTREE,
  CONSTRAINT `FKfny0antg5w2x6no0kw63c6xm3` FOREIGN KEY (`aid`) REFERENCES `sys_actor` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `FKi7d9sn34a15ly7ie76kqng40r` FOREIGN KEY (`rid`) REFERENCES `sys_role` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role_actor
-- ----------------------------
INSERT INTO `sys_role_actor` VALUES (1, 'u2');
INSERT INTO `sys_role_actor` VALUES (1, 'u3');
INSERT INTO `sys_role_actor` VALUES (1, 'u4');
INSERT INTO `sys_role_actor` VALUES (1, 'u5');
INSERT INTO `sys_role_actor` VALUES (2, 'u6');
INSERT INTO `sys_role_actor` VALUES (2, 'u7');
INSERT INTO `sys_role_actor` VALUES (2, 'u8');
INSERT INTO `sys_role_actor` VALUES (2, 'u9');

-- ----------------------------
-- Table structure for sys_role_api
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_api`;
CREATE TABLE `sys_role_api`  (
  `rid` bigint NOT NULL,
  `aid` bigint NOT NULL,
  INDEX `FKjn3td9tgpjp62deyom8b6i5hv`(`aid` ASC) USING BTREE,
  INDEX `FKfbfuye05ikel5hpyimdf9h0mk`(`rid` ASC) USING BTREE,
  CONSTRAINT `FKfbfuye05ikel5hpyimdf9h0mk` FOREIGN KEY (`rid`) REFERENCES `sys_role` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `FKjn3td9tgpjp62deyom8b6i5hv` FOREIGN KEY (`aid`) REFERENCES `sys_api` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role_api
-- ----------------------------
INSERT INTO `sys_role_api` VALUES (1, 101001);
INSERT INTO `sys_role_api` VALUES (1, 101002);
INSERT INTO `sys_role_api` VALUES (1, 101003);
INSERT INTO `sys_role_api` VALUES (1, 102001);
INSERT INTO `sys_role_api` VALUES (1, 102002);
INSERT INTO `sys_role_api` VALUES (1, 102003);
INSERT INTO `sys_role_api` VALUES (1, 102004);
INSERT INTO `sys_role_api` VALUES (1, 102005);
INSERT INTO `sys_role_api` VALUES (1, 103001);
INSERT INTO `sys_role_api` VALUES (1, 103002);
INSERT INTO `sys_role_api` VALUES (1, 103003);
INSERT INTO `sys_role_api` VALUES (1, 104001);
INSERT INTO `sys_role_api` VALUES (1, 104002);
INSERT INTO `sys_role_api` VALUES (1, 104003);
INSERT INTO `sys_role_api` VALUES (1, 104004);
INSERT INTO `sys_role_api` VALUES (1, 104005);
INSERT INTO `sys_role_api` VALUES (1, 104006);
INSERT INTO `sys_role_api` VALUES (1, 105001);
INSERT INTO `sys_role_api` VALUES (1, 105002);
INSERT INTO `sys_role_api` VALUES (1, 105003);
INSERT INTO `sys_role_api` VALUES (1, 106001);
INSERT INTO `sys_role_api` VALUES (1, 106002);
INSERT INTO `sys_role_api` VALUES (1, 106003);
INSERT INTO `sys_role_api` VALUES (1, 107001);
INSERT INTO `sys_role_api` VALUES (1, 107002);
INSERT INTO `sys_role_api` VALUES (1, 107003);
INSERT INTO `sys_role_api` VALUES (1, 108001);
INSERT INTO `sys_role_api` VALUES (1, 108002);
INSERT INTO `sys_role_api` VALUES (1, 108003);
INSERT INTO `sys_role_api` VALUES (1, 109001);
INSERT INTO `sys_role_api` VALUES (1, 109002);
INSERT INTO `sys_role_api` VALUES (1, 109003);
INSERT INTO `sys_role_api` VALUES (1, 201001);
INSERT INTO `sys_role_api` VALUES (1, 201002);
INSERT INTO `sys_role_api` VALUES (1, 202001);
INSERT INTO `sys_role_api` VALUES (1, 202002);
INSERT INTO `sys_role_api` VALUES (1, 203001);
INSERT INTO `sys_role_api` VALUES (1, 203002);
INSERT INTO `sys_role_api` VALUES (1, 204001);
INSERT INTO `sys_role_api` VALUES (1, 205001);
INSERT INTO `sys_role_api` VALUES (1, 206001);
INSERT INTO `sys_role_api` VALUES (1, 206002);
INSERT INTO `sys_role_api` VALUES (1, 206003);
INSERT INTO `sys_role_api` VALUES (1, 206101);
INSERT INTO `sys_role_api` VALUES (1, 206102);
INSERT INTO `sys_role_api` VALUES (1, 301001);
INSERT INTO `sys_role_api` VALUES (1, 301002);
INSERT INTO `sys_role_api` VALUES (1, 301003);
INSERT INTO `sys_role_api` VALUES (1, 301004);
INSERT INTO `sys_role_api` VALUES (1, 301005);
INSERT INTO `sys_role_api` VALUES (1, 301006);
INSERT INTO `sys_role_api` VALUES (1, 302001);
INSERT INTO `sys_role_api` VALUES (1, 302002);
INSERT INTO `sys_role_api` VALUES (1, 302003);
INSERT INTO `sys_role_api` VALUES (1, 603001);
INSERT INTO `sys_role_api` VALUES (1, 603002);
INSERT INTO `sys_role_api` VALUES (1, 603003);
INSERT INTO `sys_role_api` VALUES (2, 603001);
INSERT INTO `sys_role_api` VALUES (2, 603002);
INSERT INTO `sys_role_api` VALUES (2, 603003);

-- ----------------------------
-- Table structure for sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu`  (
  `rid` bigint NOT NULL,
  `mid` bigint NOT NULL,
  INDEX `FK5grhomnyrb2nkm20ee7a2iv92`(`mid` ASC) USING BTREE,
  INDEX `FK5pc0747orubmrx2oe86newexy`(`rid` ASC) USING BTREE,
  CONSTRAINT `FK5grhomnyrb2nkm20ee7a2iv92` FOREIGN KEY (`mid`) REFERENCES `sys_menu` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `FK5pc0747orubmrx2oe86newexy` FOREIGN KEY (`rid`) REFERENCES `sys_role` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role_menu
-- ----------------------------
INSERT INTO `sys_role_menu` VALUES (1, 1000);
INSERT INTO `sys_role_menu` VALUES (1, 1010);
INSERT INTO `sys_role_menu` VALUES (1, 1020);
INSERT INTO `sys_role_menu` VALUES (1, 1021);
INSERT INTO `sys_role_menu` VALUES (1, 1030);
INSERT INTO `sys_role_menu` VALUES (1, 1040);
INSERT INTO `sys_role_menu` VALUES (1, 1050);
INSERT INTO `sys_role_menu` VALUES (1, 1060);
INSERT INTO `sys_role_menu` VALUES (1, 1070);
INSERT INTO `sys_role_menu` VALUES (1, 1071);
INSERT INTO `sys_role_menu` VALUES (1, 1080);
INSERT INTO `sys_role_menu` VALUES (1, 1090);
INSERT INTO `sys_role_menu` VALUES (1, 2000);
INSERT INTO `sys_role_menu` VALUES (1, 2010);
INSERT INTO `sys_role_menu` VALUES (1, 2020);
INSERT INTO `sys_role_menu` VALUES (1, 2030);
INSERT INTO `sys_role_menu` VALUES (1, 2040);
INSERT INTO `sys_role_menu` VALUES (1, 2050);
INSERT INTO `sys_role_menu` VALUES (1, 2060);
INSERT INTO `sys_role_menu` VALUES (1, 2061);
INSERT INTO `sys_role_menu` VALUES (1, 3000);
INSERT INTO `sys_role_menu` VALUES (1, 3010);
INSERT INTO `sys_role_menu` VALUES (1, 3020);
INSERT INTO `sys_role_menu` VALUES (1, 3030);
INSERT INTO `sys_role_menu` VALUES (1, 3040);
INSERT INTO `sys_role_menu` VALUES (1, 3041);
INSERT INTO `sys_role_menu` VALUES (1, 3050);
INSERT INTO `sys_role_menu` VALUES (1, 3051);
INSERT INTO `sys_role_menu` VALUES (1, 6000);
INSERT INTO `sys_role_menu` VALUES (1, 6010);
INSERT INTO `sys_role_menu` VALUES (1, 6020);
INSERT INTO `sys_role_menu` VALUES (1, 6021);
INSERT INTO `sys_role_menu` VALUES (1, 6030);
INSERT INTO `sys_role_menu` VALUES (1, 6031);
INSERT INTO `sys_role_menu` VALUES (1, 6032);
INSERT INTO `sys_role_menu` VALUES (1, 6040);
INSERT INTO `sys_role_menu` VALUES (1, 6050);
INSERT INTO `sys_role_menu` VALUES (1, 6051);
INSERT INTO `sys_role_menu` VALUES (2, 6000);
INSERT INTO `sys_role_menu` VALUES (2, 6010);
INSERT INTO `sys_role_menu` VALUES (2, 6020);
INSERT INTO `sys_role_menu` VALUES (2, 6021);
INSERT INTO `sys_role_menu` VALUES (2, 6030);
INSERT INTO `sys_role_menu` VALUES (2, 6031);
INSERT INTO `sys_role_menu` VALUES (2, 6032);
INSERT INTO `sys_role_menu` VALUES (2, 6040);
INSERT INTO `sys_role_menu` VALUES (2, 6050);
INSERT INTO `sys_role_menu` VALUES (2, 6051);

-- ----------------------------
-- Table structure for sys_social
-- ----------------------------
DROP TABLE IF EXISTS `sys_social`;
CREATE TABLE `sys_social`  (
  `id` bigint NOT NULL COMMENT '主键ID',
  `access_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `access_token` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `auth_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `create_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `create_org` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `create_time` datetime(6) NULL DEFAULT NULL,
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `expire_in` int NOT NULL,
  `id_token` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `mac_algorithm` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `mac_key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `nick_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `oauth_token` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `oauth_token_secret` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `open_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `refresh_token` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `scope` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `source` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `token_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `union_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `update_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `update_time` datetime(6) NULL DEFAULT NULL,
  `user_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '三方登录' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_social
-- ----------------------------

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
  `id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '主键ID',
  `arcod` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '地区编号',
  `arnam` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '地区名称',
  `avatar` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '头像url',
  `avtag` bit(1) NULL DEFAULT NULL COMMENT '可用标记',
  `catag` bit(1) NULL DEFAULT NULL COMMENT '缓存标记',
  `crtim` datetime(6) NULL DEFAULT NULL COMMENT '创建时间',
  `cruid` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '创建人ID',
  `email` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '邮箱',
  `gender` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '性别',
  `job` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '职务',
  `jotyp` bit(1) NULL DEFAULT NULL COMMENT '直接可加还是需要同意后再加',
  `label` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标签',
  `loip` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '登录IP',
  `lotim` datetime(6) NULL DEFAULT NULL COMMENT '登录时间',
  `monum` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '手机号',
  `name` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '姓名（昵称）',
  `notes` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'notes',
  `oftim` datetime(6) NULL DEFAULT NULL COMMENT '最后离开时间',
  `orgid` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '组织ID',
  `ornum` int NULL DEFAULT NULL COMMENT '排序号',
  `password` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '密码',
  `tier` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '层级',
  `type` int NULL DEFAULT NULL COMMENT '用户类别',
  `uptim` datetime(6) NULL DEFAULT NULL COMMENT '更新时间',
  `upuid` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新人ID',
  `username` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户名（账号）',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统用户' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES ('u1', NULL, NULL, NULL, b'1', b'0', '2026-01-20 15:03:01.014000', NULL, 'vben@qq.com', '1', NULL, NULL, NULL, NULL, NULL, '13812345678', '管理员', '管理员不给修改', NULL, 'o1000', 1, '$2a$10$09f8rxsX4tbj1CZla2MSOuiwHwp5QAPUzbp5whnoZEFK4/xplNwZq', '_o1000_u1_', NULL, NULL, NULL, 'admin');
INSERT INTO `sys_user` VALUES ('u2', NULL, NULL, NULL, b'1', b'0', '2026-01-20 15:03:01.014000', NULL, NULL, '2', NULL, NULL, NULL, NULL, NULL, '13912345678', '小狐狸', NULL, NULL, 'o1000', 2, '$2a$10$09f8rxsX4tbj1CZla2MSOuiwHwp5QAPUzbp5whnoZEFK4/xplNwZq', '_o1000_u2_', NULL, NULL, NULL, 'vben');
INSERT INTO `sys_user` VALUES ('u3', NULL, NULL, NULL, b'1', b'0', '2026-01-20 15:03:01.014000', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '张三', NULL, NULL, 'o1000', 3, '$2a$10$09f8rxsX4tbj1CZla2MSOuiwHwp5QAPUzbp5whnoZEFK4/xplNwZq', '_o1000_u3_', NULL, NULL, NULL, 'zs');
INSERT INTO `sys_user` VALUES ('u4', NULL, NULL, NULL, b'1', b'0', '2026-01-20 15:03:01.014000', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '李四', NULL, NULL, 'o1100', 4, '$2a$10$09f8rxsX4tbj1CZla2MSOuiwHwp5QAPUzbp5whnoZEFK4/xplNwZq', '_o1000_o1100_u4_', NULL, NULL, NULL, 'ls');
INSERT INTO `sys_user` VALUES ('u5', NULL, NULL, NULL, b'1', b'0', '2026-01-20 15:03:01.014000', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '王五', NULL, NULL, 'o1110', 5, '$2a$10$09f8rxsX4tbj1CZla2MSOuiwHwp5QAPUzbp5whnoZEFK4/xplNwZq', '_o1000_o1100_o1110_u5_', NULL, NULL, NULL, 'ww');
INSERT INTO `sys_user` VALUES ('u6', NULL, NULL, NULL, b'1', b'0', '2026-01-20 15:03:01.014000', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '赵六', NULL, NULL, 'o1111', 6, '$2a$10$09f8rxsX4tbj1CZla2MSOuiwHwp5QAPUzbp5whnoZEFK4/xplNwZq', '_o1000_o1100_o1110_o1111_u6_', NULL, NULL, NULL, 'zl');
INSERT INTO `sys_user` VALUES ('u7', NULL, NULL, NULL, b'1', b'0', '2026-01-20 15:03:01.014000', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '孙七', NULL, NULL, 'o1111', 7, '$2a$10$09f8rxsX4tbj1CZla2MSOuiwHwp5QAPUzbp5whnoZEFK4/xplNwZq', '_o1000_o1100_o1110_o1111_u7_', NULL, NULL, NULL, 'sq');
INSERT INTO `sys_user` VALUES ('u8', NULL, NULL, NULL, b'1', b'0', '2026-01-20 15:03:01.014000', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '周八', NULL, NULL, 'o1111', 8, '$2a$10$09f8rxsX4tbj1CZla2MSOuiwHwp5QAPUzbp5whnoZEFK4/xplNwZq', '_o1000_o1100_o1110_o1111_u8_', NULL, NULL, NULL, 'zb');
INSERT INTO `sys_user` VALUES ('u9', NULL, NULL, NULL, b'1', b'0', '2026-01-20 15:03:01.014000', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '吴九', NULL, NULL, 'o1111', 9, '$2a$10$09f8rxsX4tbj1CZla2MSOuiwHwp5QAPUzbp5whnoZEFK4/xplNwZq', '_o1000_o1100_o1110_o1111_u9_', NULL, NULL, NULL, 'wj');

-- ----------------------------
-- Table structure for tool_code_field
-- ----------------------------
DROP TABLE IF EXISTS `tool_code_field`;
CREATE TABLE `tool_code_field`  (
  `id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '主键ID',
  `length` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '字段长度',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '字段名称',
  `notes` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `ornum` int NULL DEFAULT NULL COMMENT '排序号',
  `remark` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '字段描述',
  `type` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '字段类型',
  `tabid` bigint NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `FK9smrfy4b11ekpadloyrj26kly`(`tabid` ASC) USING BTREE,
  CONSTRAINT `FK9smrfy4b11ekpadloyrj26kly` FOREIGN KEY (`tabid`) REFERENCES `tool_code_table` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '代码生成-字段' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tool_code_field
-- ----------------------------

-- ----------------------------
-- Table structure for tool_code_table
-- ----------------------------
DROP TABLE IF EXISTS `tool_code_table`;
CREATE TABLE `tool_code_table`  (
  `id` bigint NOT NULL COMMENT '主键ID',
  `avtag` bit(1) NULL DEFAULT NULL COMMENT '可用标记',
  `crtim` datetime(6) NULL DEFAULT NULL COMMENT '创建时间',
  `cruid` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '创建人ID',
  `name` varchar(126) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
  `notes` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `uptim` datetime(6) NULL DEFAULT NULL COMMENT '更新时间',
  `upuid` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新人ID',
  `addbt` bit(1) NULL DEFAULT NULL COMMENT '新增按钮',
  `baent` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '继承基类',
  `bunam` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '实体类',
  `delbt` bit(1) NULL DEFAULT NULL COMMENT '删除按钮',
  `edtyp` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '编辑页类型',
  `expbt` bit(1) NULL DEFAULT NULL COMMENT '导出按钮',
  `impbt` bit(1) NULL DEFAULT NULL COMMENT '导入按钮',
  `orfie` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '排序字段',
  `ornum` int NULL DEFAULT NULL COMMENT '排序号',
  `ortyp` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'orm类型',
  `panam` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '基础包名',
  `pecol` int NULL DEFAULT NULL COMMENT '每行列数',
  `pmeid` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '上级菜单ID',
  `porid` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '所属门户ID',
  `remark` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '表描述',
  `rotyp` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '路由类型',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '代码生成-表格' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tool_code_table
-- ----------------------------

-- ----------------------------
-- Table structure for tool_dict
-- ----------------------------
DROP TABLE IF EXISTS `tool_dict`;
CREATE TABLE `tool_dict`  (
  `id` bigint NOT NULL COMMENT '主键ID',
  `avtag` bit(1) NULL DEFAULT NULL COMMENT '可用标记',
  `catid` bigint NULL DEFAULT NULL COMMENT '分类ID',
  `code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '字典代码',
  `crtim` datetime(6) NULL DEFAULT NULL COMMENT '创建时间',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
  `notes` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
  `ornum` int NULL DEFAULT NULL COMMENT '排序号',
  `uptim` datetime(6) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '字典信息' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tool_dict
-- ----------------------------
INSERT INTO `tool_dict` VALUES (1, NULL, NULL, 'sys_oper_type', '2026-01-20 15:03:01.387000', '操作类型', NULL, 1, NULL);
INSERT INTO `tool_dict` VALUES (2, NULL, NULL, 'sys_grant_type', '2026-01-20 15:03:01.387000', '授权类型', NULL, 2, NULL);
INSERT INTO `tool_dict` VALUES (3, NULL, NULL, 'sys_device_type', '2026-01-20 15:03:01.387000', '设备类型', NULL, 3, NULL);

-- ----------------------------
-- Table structure for tool_dict_cate
-- ----------------------------
DROP TABLE IF EXISTS `tool_dict_cate`;
CREATE TABLE `tool_dict_cate`  (
  `id` bigint NOT NULL COMMENT '主键ID',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
  `ornum` int NULL DEFAULT NULL COMMENT '排序号',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '字典分类' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tool_dict_cate
-- ----------------------------

-- ----------------------------
-- Table structure for tool_dict_data
-- ----------------------------
DROP TABLE IF EXISTS `tool_dict_data`;
CREATE TABLE `tool_dict_data`  (
  `id` bigint NOT NULL COMMENT '主键ID',
  `avtag` bit(1) NULL DEFAULT NULL COMMENT '可用标记',
  `crtim` datetime(6) NULL DEFAULT NULL COMMENT '创建时间',
  `dalab` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '数据标签',
  `daval` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '数据键值',
  `detag` bit(1) NULL DEFAULT NULL COMMENT '默认标记',
  `dicid` bigint NULL DEFAULT NULL COMMENT '字典ID',
  `notes` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `ornum` int NULL DEFAULT NULL COMMENT '排序号',
  `shsty` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '显示样式',
  `uptim` datetime(6) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '字典数据' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tool_dict_data
-- ----------------------------
INSERT INTO `tool_dict_data` VALUES (101, b'1', '2026-01-20 15:03:01.387000', '新增', '1', NULL, 1, NULL, 1, 'primary', '2026-01-20 15:03:01.387000');
INSERT INTO `tool_dict_data` VALUES (102, b'1', '2026-01-20 15:03:01.387000', '修改', '2', NULL, 1, NULL, 2, 'primary', '2026-01-20 15:03:01.387000');
INSERT INTO `tool_dict_data` VALUES (103, b'1', '2026-01-20 15:03:01.387000', '删除', '3', NULL, 1, NULL, 3, 'danger', '2026-01-20 15:03:01.387000');
INSERT INTO `tool_dict_data` VALUES (104, b'1', '2026-01-20 15:03:01.387000', '授权', '4', NULL, 1, NULL, 4, 'primary', '2026-01-20 15:03:01.387000');
INSERT INTO `tool_dict_data` VALUES (105, b'1', '2026-01-20 15:03:01.387000', '导出', '5', NULL, 1, NULL, 5, 'warning', '2026-01-20 15:03:01.387000');
INSERT INTO `tool_dict_data` VALUES (106, b'1', '2026-01-20 15:03:01.387000', '导入', '6', NULL, 1, NULL, 6, 'warning', '2026-01-20 15:03:01.387000');
INSERT INTO `tool_dict_data` VALUES (107, b'1', '2026-01-20 15:03:01.387000', '强退', '7', NULL, 1, NULL, 7, 'danger', '2026-01-20 15:03:01.387000');
INSERT INTO `tool_dict_data` VALUES (108, b'1', '2026-01-20 15:03:01.387000', '生成代码', '8', NULL, 1, NULL, 8, 'warning', '2026-01-20 15:03:01.387000');
INSERT INTO `tool_dict_data` VALUES (109, b'1', '2026-01-20 15:03:01.387000', '清空数据', '9', NULL, 1, NULL, 9, 'danger', '2026-01-20 15:03:01.387000');
INSERT INTO `tool_dict_data` VALUES (199, b'1', '2026-01-20 15:03:01.387000', '其他操作', '0', NULL, 1, NULL, 99, 'warning', '2026-01-20 15:03:01.387000');
INSERT INTO `tool_dict_data` VALUES (201, b'1', '2026-01-20 15:03:01.387000', '密码认证', 'password', NULL, 2, NULL, 1, 'default', '2026-01-20 15:03:01.387000');
INSERT INTO `tool_dict_data` VALUES (202, b'1', '2026-01-20 15:03:01.387000', '短信认证', 'sms', NULL, 2, NULL, 2, 'default', '2026-01-20 15:03:01.387000');
INSERT INTO `tool_dict_data` VALUES (203, b'1', '2026-01-20 15:03:01.387000', '邮箱认证', 'email', NULL, 2, NULL, 3, 'default', '2026-01-20 15:03:01.387000');
INSERT INTO `tool_dict_data` VALUES (204, b'1', '2026-01-20 15:03:01.387000', '小程序认证', 'xcx', NULL, 2, NULL, 4, 'default', '2026-01-20 15:03:01.387000');
INSERT INTO `tool_dict_data` VALUES (205, b'1', '2026-01-20 15:03:01.387000', '三方登录认证', 'social', NULL, 2, NULL, 5, 'default', '2026-01-20 15:03:01.387000');
INSERT INTO `tool_dict_data` VALUES (301, b'1', '2026-01-20 15:03:01.387000', 'PC', 'PC', NULL, 3, NULL, 1, 'default', '2026-01-20 15:03:01.387000');
INSERT INTO `tool_dict_data` VALUES (302, b'1', '2026-01-20 15:03:01.387000', '安卓', 'android', NULL, 3, NULL, 2, 'default', '2026-01-20 15:03:01.387000');
INSERT INTO `tool_dict_data` VALUES (303, b'1', '2026-01-20 15:03:01.387000', '苹果', 'IOS', NULL, 3, NULL, 3, 'default', '2026-01-20 15:03:01.387000');
INSERT INTO `tool_dict_data` VALUES (304, b'1', '2026-01-20 15:03:01.387000', '小程序', 'XCX', NULL, 3, NULL, 4, 'default', '2026-01-20 15:03:01.387000');

-- ----------------------------
-- Table structure for tool_form
-- ----------------------------
DROP TABLE IF EXISTS `tool_form`;
CREATE TABLE `tool_form`  (
  `id` bigint NOT NULL COMMENT '主键ID',
  `avtag` bit(1) NULL DEFAULT NULL COMMENT '可用标记',
  `crtim` datetime(6) NULL DEFAULT NULL COMMENT '创建时间',
  `cruid` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '创建人ID',
  `name` varchar(126) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
  `notes` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `uptim` datetime(6) NULL DEFAULT NULL COMMENT '更新时间',
  `upuid` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新人ID',
  `frule` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '表单规则',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '在线表单' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tool_form
-- ----------------------------

-- ----------------------------
-- Table structure for tool_num
-- ----------------------------
DROP TABLE IF EXISTS `tool_num`;
CREATE TABLE `tool_num`  (
  `id` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '主键ID',
  `avtag` bit(1) NULL DEFAULT NULL COMMENT '可用标记',
  `crtim` datetime(6) NULL DEFAULT NULL COMMENT '创建时间',
  `cudat` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '当前日期',
  `label` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标签',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
  `nflag` bit(1) NULL DEFAULT NULL COMMENT '是否被修改过或新添加的',
  `notes` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `nulen` int NULL DEFAULT NULL COMMENT '编号长度',
  `numod` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '生成模式',
  `nunex` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '下一个编号',
  `nupre` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '编号前缀',
  `ornum` int NULL DEFAULT NULL COMMENT '排序号',
  `uptim` datetime(6) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '编号信息' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tool_num
-- ----------------------------
INSERT INTO `tool_num` VALUES ('BUS', b'1', '2026-01-20 15:03:01.413000', NULL, NULL, '流程编号', b'1', NULL, 4, 'YYYYMMDD', NULL, 'BUS', 1, '2026-01-20 15:03:01.413000');

-- ----------------------------
-- Table structure for tool_oss
-- ----------------------------
DROP TABLE IF EXISTS `tool_oss`;
CREATE TABLE `tool_oss`  (
  `id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '主键ID',
  `busid` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '业务ID',
  `crtim` datetime(6) NULL DEFAULT NULL COMMENT '创建时间',
  `filid` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '文件ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '文件名称',
  `type` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '类型（后缀）',
  `cruid` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '创建人',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `FK5x8t712epy32kps2h4ay2bxuc`(`cruid` ASC) USING BTREE,
  CONSTRAINT `FK5x8t712epy32kps2h4ay2bxuc` FOREIGN KEY (`cruid`) REFERENCES `sys_actor` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'OSS存储引用' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tool_oss
-- ----------------------------

-- ----------------------------
-- Table structure for tool_oss_file
-- ----------------------------
DROP TABLE IF EXISTS `tool_oss_file`;
CREATE TABLE `tool_oss_file`  (
  `id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '主键ID',
  `crtim` datetime(6) NULL DEFAULT NULL COMMENT '创建时间',
  `fsize` bigint NULL DEFAULT NULL COMMENT '文件大小',
  `md5` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '文件md5',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '存储地址',
  `service` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '服务商',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `IDX3s53j69dg34b9sfa5khxwv851`(`md5` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'OSS存储文件' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tool_oss_file
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
