-- MySQL dump 10.13  Distrib 9.0.1, for Linux (aarch64)
--
-- Host: 47.109.31.106    Database: gin_vue_admin
-- ------------------------------------------------------
-- Server version	8.2.0

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `sys_base_menus`
--

DROP TABLE IF EXISTS `sys_base_menus`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_base_menus` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `menu_level` bigint unsigned DEFAULT NULL,
  `parent_id` bigint unsigned DEFAULT NULL COMMENT '父菜单ID',
  `path` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '路由path',
  `name` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '路由name',
  `hidden` tinyint(1) DEFAULT NULL COMMENT '是否在列表隐藏',
  `component` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '对应前端文件路径',
  `sort` bigint DEFAULT NULL COMMENT '排序标记',
  `active_name` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '附加属性',
  `keep_alive` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  `default_menu` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  `title` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '附加属性',
  `icon` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '附加属性',
  `close_tab` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  PRIMARY KEY (`id`),
  KEY `idx_sys_base_menus_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=39 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_base_menus`
--

LOCK TABLES `sys_base_menus` WRITE;
/*!40000 ALTER TABLE `sys_base_menus` DISABLE KEYS */;
INSERT INTO `sys_base_menus` VALUES (1,'2024-10-01 15:44:58.213','2024-10-01 15:44:58.213',NULL,0,0,'dashboard','dashboard',0,'view/dashboard/index.vue',1,'',0,0,'仪表盘','odometer',0),(2,'2024-10-01 15:44:58.213','2024-10-01 15:44:58.213','2024-10-02 15:47:01.694',0,0,'about','about',0,'view/about/index.vue',9,'',0,0,'关于我们','info-filled',0),(3,'2024-10-01 15:44:58.213','2024-10-01 15:44:58.213',NULL,0,0,'admin','superAdmin',0,'view/superAdmin/index.vue',3,'',0,0,'超级管理员','user',0),(4,'2024-10-01 15:44:58.213','2024-10-01 15:44:58.213',NULL,0,3,'authority','authority',0,'view/superAdmin/authority/authority.vue',1,'',0,0,'角色管理','avatar',0),(5,'2024-10-01 15:44:58.213','2024-10-01 15:44:58.213',NULL,0,3,'menu','menu',0,'view/superAdmin/menu/menu.vue',2,'',1,0,'菜单管理','tickets',0),(6,'2024-10-01 15:44:58.213','2024-10-01 15:44:58.213',NULL,0,3,'api','api',0,'view/superAdmin/api/api.vue',3,'',1,0,'api管理','platform',0),(7,'2024-10-01 15:44:58.213','2024-10-01 15:44:58.213',NULL,0,3,'user','user',0,'view/superAdmin/user/user.vue',4,'',0,0,'用户管理','coordinate',0),(8,'2024-10-01 15:44:58.213','2024-10-01 15:44:58.213',NULL,0,3,'dictionary','dictionary',0,'view/superAdmin/dictionary/sysDictionary.vue',5,'',0,0,'字典管理','notebook',0),(9,'2024-10-01 15:44:58.213','2024-10-01 15:44:58.213',NULL,0,3,'operation','operation',0,'view/superAdmin/operation/sysOperationRecord.vue',6,'',0,0,'操作历史','pie-chart',0),(10,'2024-10-01 15:44:58.213','2024-10-01 15:44:58.213',NULL,0,0,'person','person',1,'view/person/person.vue',4,'',0,0,'个人信息','message',0),(11,'2024-10-01 15:44:58.213','2024-10-01 15:44:58.213',NULL,0,0,'example','example',0,'view/example/index.vue',7,'',0,0,'示例文件','management',0),(12,'2024-10-01 15:44:58.213','2024-10-01 15:44:58.213',NULL,0,11,'upload','upload',0,'view/example/upload/upload.vue',5,'',0,0,'媒体库（上传下载）','upload',0),(13,'2024-10-01 15:44:58.213','2024-10-01 15:44:58.213',NULL,0,11,'breakpoint','breakpoint',0,'view/example/breakpoint/breakpoint.vue',6,'',0,0,'断点续传','upload-filled',0),(14,'2024-10-01 15:44:58.213','2024-10-01 15:44:58.213',NULL,0,11,'customer','customer',0,'view/example/customer/customer.vue',7,'',0,0,'客户列表（资源示例）','avatar',0),(15,'2024-10-01 15:44:58.213','2024-10-01 15:44:58.213',NULL,0,0,'systemTools','systemTools',0,'view/systemTools/index.vue',5,'',0,0,'系统工具','tools',0),(16,'2024-10-01 15:44:58.213','2024-10-01 15:44:58.213',NULL,0,15,'autoCode','autoCode',0,'view/systemTools/autoCode/index.vue',1,'',1,0,'代码生成器','cpu',0),(17,'2024-10-01 15:44:58.213','2024-10-01 15:44:58.213',NULL,0,15,'formCreate','formCreate',0,'view/systemTools/formCreate/index.vue',2,'',1,0,'表单生成器','magic-stick',0),(18,'2024-10-01 15:44:58.213','2024-10-01 15:44:58.213',NULL,0,15,'system','system',0,'view/systemTools/system/system.vue',3,'',0,0,'系统配置','operation',0),(19,'2024-10-01 15:44:58.213','2024-10-01 15:44:58.213',NULL,0,15,'autoCodeAdmin','autoCodeAdmin',0,'view/systemTools/autoCodeAdmin/index.vue',1,'',0,0,'自动化代码管理','magic-stick',0),(20,'2024-10-01 15:44:58.213','2024-10-01 15:44:58.213',NULL,0,15,'autoCodeEdit/:id','autoCodeEdit',1,'view/systemTools/autoCode/index.vue',0,'',0,0,'自动化代码-${id}','magic-stick',0),(21,'2024-10-01 15:44:58.213','2024-10-01 15:44:58.213',NULL,0,15,'autoPkg','autoPkg',0,'view/systemTools/autoPkg/autoPkg.vue',0,'',0,0,'自动化package','folder',0),(22,'2024-10-01 15:44:58.213','2024-10-01 15:44:58.213',NULL,0,0,'https://www.gin-vue-admin.com','https://www.gin-vue-admin.com',0,'/',0,'',0,0,'官方网站','customer-gva',0),(23,'2024-10-01 15:44:58.213','2024-10-01 15:44:58.213',NULL,0,0,'state','state',0,'view/system/state.vue',8,'',0,0,'服务器状态','cloudy',0),(24,'2024-10-01 15:44:58.213','2024-10-01 15:44:58.213',NULL,0,0,'plugin','plugin',0,'view/routerHolder.vue',6,'',0,0,'插件系统','cherry',0),(25,'2024-10-01 15:44:58.213','2024-10-01 15:44:58.213',NULL,0,24,'https://plugin.gin-vue-admin.com/','https://plugin.gin-vue-admin.com/',0,'https://plugin.gin-vue-admin.com/',0,'',0,0,'插件市场','shop',0),(26,'2024-10-01 15:44:58.213','2024-10-01 15:44:58.213',NULL,0,24,'installPlugin','installPlugin',0,'view/systemTools/installPlugin/index.vue',1,'',0,0,'插件安装','box',0),(27,'2024-10-01 15:44:58.213','2024-10-01 15:44:58.213',NULL,0,24,'autoPlug','autoPlug',0,'view/systemTools/autoPlug/autoPlug.vue',2,'',0,0,'插件模板','folder',0),(28,'2024-10-01 15:44:58.213','2024-10-01 15:44:58.213',NULL,0,24,'pubPlug','pubPlug',0,'view/systemTools/pubPlug/pubPlug.vue',3,'',0,0,'打包插件','files',0),(29,'2024-10-01 15:44:58.213','2024-10-01 15:44:58.213',NULL,0,24,'plugin-email','plugin-email',0,'plugin/email/view/index.vue',4,'',0,0,'邮件插件','message',0),(30,'2024-10-01 15:44:58.213','2024-10-01 15:44:58.213',NULL,0,15,'exportTemplate','exportTemplate',0,'view/systemTools/exportTemplate/exportTemplate.vue',10,'',0,0,'表格模板','reading',0),(31,'2024-10-02 15:51:02.954','2024-10-02 16:16:31.024',NULL,0,0,'goods','goods',0,'view/routerHolder.vue',2,'',0,0,'商品管理','box',0),(32,'2024-10-02 15:53:01.972','2024-10-02 16:18:41.617',NULL,0,31,'imitate','imitate',0,'view/product/data/index.vue',10,'',0,0,'数据模拟','magic-stick',0),(33,'2024-10-02 15:57:44.939','2024-10-02 16:17:46.594',NULL,0,31,'crawl','crawl',0,'view/product/getGoods/index.vue',2,'',0,0,'数据爬取','bowl',0),(34,'2024-10-02 15:59:03.367','2024-10-02 16:18:03.313',NULL,0,31,'category','category',0,'view/product/category/category.vue',9,'',0,0,'分类管理','grid',0),(35,'2024-10-02 16:05:43.608','2024-10-02 16:16:45.591',NULL,0,31,'all','all',0,'view/product/goods/list/index.vue',1,'',0,0,'全部商品','goods',0),(36,'2024-10-02 16:08:15.737','2024-10-02 16:15:38.698',NULL,0,0,'search','search',0,'view/product/search/goods_search/index.vue',1,'',0,0,'商品搜索','search',0),(37,'2024-10-02 21:03:39.835','2024-10-02 21:08:58.866',NULL,0,0,'goodsDetail','goodsDetail',1,'view/product/search/goods_detail/index.vue',0,'',0,0,'商品详情','',0),(38,'2024-10-02 21:57:04.799','2024-10-02 21:57:04.799',NULL,0,0,'goodsManage','goodsManage',1,'view/product/goods/manage/index.vue',0,'',0,0,'商品管理','',0);
/*!40000 ALTER TABLE `sys_base_menus` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-10-09 13:42:49
