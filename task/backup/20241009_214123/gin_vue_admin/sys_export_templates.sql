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
-- Table structure for table `sys_export_templates`
--

DROP TABLE IF EXISTS `sys_export_templates`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_export_templates` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `db_name` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '数据库名称',
  `name` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '模板名称',
  `table_name` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '表名称',
  `template_id` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '模板标识',
  `template_info` text COLLATE utf8mb4_general_ci,
  `limit` bigint DEFAULT NULL COMMENT '导出限制',
  `order` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '排序',
  PRIMARY KEY (`id`),
  KEY `idx_sys_export_templates_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_export_templates`
--

LOCK TABLES `sys_export_templates` WRITE;
/*!40000 ALTER TABLE `sys_export_templates` DISABLE KEYS */;
INSERT INTO `sys_export_templates` VALUES (1,'2024-10-01 15:44:58.919','2024-10-01 15:44:58.919',NULL,'','api','sys_apis','api','{\n\"path\":\"路径\",\n\"method\":\"方法（大写）\",\n\"description\":\"方法介绍\",\n\"api_group\":\"方法分组\"\n}',0,'');
/*!40000 ALTER TABLE `sys_export_templates` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-10-09 13:42:53
