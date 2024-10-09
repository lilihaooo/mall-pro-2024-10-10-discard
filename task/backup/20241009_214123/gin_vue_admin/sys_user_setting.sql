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
-- Table structure for table `sys_user_setting`
--

DROP TABLE IF EXISTS `sys_user_setting`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_user_setting` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '''系统设置ID''',
  `user_id` bigint NOT NULL COMMENT '''用户ID''',
  `weakness` tinyint(1) DEFAULT '0' COMMENT '''是否启用色弱模式''',
  `grey` tinyint(1) DEFAULT '0' COMMENT '''是否启用灰色模式''',
  `primary_color` varchar(100) COLLATE utf8mb4_general_ci NOT NULL COMMENT '''主色调''',
  `show_tabs` tinyint(1) DEFAULT '1' COMMENT '''是否显示标签页''',
  `dark_mode` varchar(100) COLLATE utf8mb4_general_ci NOT NULL COMMENT '''暗黑模式''',
  `layout_side_width` bigint NOT NULL COMMENT '''侧边栏宽度''',
  `layout_side_collapsed_width` bigint NOT NULL COMMENT '''侧边栏折叠宽度''',
  `layout_side_item_height` bigint NOT NULL COMMENT '''侧边栏项高度''',
  `show_watermark` tinyint(1) DEFAULT '0' COMMENT '''是否显示水印''',
  `side_mode` varchar(100) COLLATE utf8mb4_general_ci NOT NULL COMMENT '''侧边栏模式''',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_user_setting`
--

LOCK TABLES `sys_user_setting` WRITE;
/*!40000 ALTER TABLE `sys_user_setting` DISABLE KEYS */;
INSERT INTO `sys_user_setting` VALUES (1,1,0,0,'#EB2F96',1,'light',256,80,48,0,'head');
/*!40000 ALTER TABLE `sys_user_setting` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-10-09 13:42:55
