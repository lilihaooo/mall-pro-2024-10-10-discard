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
-- Table structure for table `sys_dictionary_details`
--

DROP TABLE IF EXISTS `sys_dictionary_details`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_dictionary_details` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `label` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '展示值',
  `value` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '字典值',
  `extend` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '扩展值',
  `status` tinyint(1) DEFAULT NULL COMMENT '启用状态',
  `sort` bigint DEFAULT NULL COMMENT '排序标记',
  `sys_dictionary_id` bigint unsigned DEFAULT NULL COMMENT '关联标记',
  PRIMARY KEY (`id`),
  KEY `idx_sys_dictionary_details_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_dictionary_details`
--

LOCK TABLES `sys_dictionary_details` WRITE;
/*!40000 ALTER TABLE `sys_dictionary_details` DISABLE KEYS */;
INSERT INTO `sys_dictionary_details` VALUES (1,'2024-10-01 15:44:57.439','2024-10-01 15:44:57.439',NULL,'男','1','',1,1,1),(2,'2024-10-01 15:44:57.439','2024-10-01 15:44:57.439',NULL,'女','2','',1,2,1),(3,'2024-10-01 15:44:57.564','2024-10-01 15:44:57.564',NULL,'smallint','1','mysql',1,1,2),(4,'2024-10-01 15:44:57.564','2024-10-01 15:44:57.564',NULL,'mediumint','2','mysql',1,2,2),(5,'2024-10-01 15:44:57.564','2024-10-01 15:44:57.564',NULL,'int','3','mysql',1,3,2),(6,'2024-10-01 15:44:57.564','2024-10-01 15:44:57.564',NULL,'bigint','4','mysql',1,4,2),(7,'2024-10-01 15:44:57.564','2024-10-01 15:44:57.564',NULL,'int2','5','pgsql',1,5,2),(8,'2024-10-01 15:44:57.564','2024-10-01 15:44:57.564',NULL,'int4','6','pgsql',1,6,2),(9,'2024-10-01 15:44:57.564','2024-10-01 15:44:57.564',NULL,'int6','7','pgsql',1,7,2),(10,'2024-10-01 15:44:57.564','2024-10-01 15:44:57.564',NULL,'int8','8','pgsql',1,8,2),(11,'2024-10-01 15:44:57.692','2024-10-01 15:44:57.692',NULL,'date','','',1,0,3),(12,'2024-10-01 15:44:57.692','2024-10-01 15:44:57.692',NULL,'time','1','mysql',1,1,3),(13,'2024-10-01 15:44:57.692','2024-10-01 15:44:57.692',NULL,'year','2','mysql',1,2,3),(14,'2024-10-01 15:44:57.692','2024-10-01 15:44:57.692',NULL,'datetime','3','mysql',1,3,3),(15,'2024-10-01 15:44:57.692','2024-10-01 15:44:57.692',NULL,'timestamp','5','mysql',1,5,3),(16,'2024-10-01 15:44:57.692','2024-10-01 15:44:57.692',NULL,'timestamptz','6','pgsql',1,5,3),(17,'2024-10-01 15:44:57.843','2024-10-01 15:44:57.843',NULL,'float','','',1,0,4),(18,'2024-10-01 15:44:57.843','2024-10-01 15:44:57.843',NULL,'double','1','mysql',1,1,4),(19,'2024-10-01 15:44:57.843','2024-10-01 15:44:57.843',NULL,'decimal','2','mysql',1,2,4),(20,'2024-10-01 15:44:57.843','2024-10-01 15:44:57.843',NULL,'numeric','3','pgsql',1,3,4),(21,'2024-10-01 15:44:57.843','2024-10-01 15:44:57.843',NULL,'smallserial','4','pgsql',1,4,4),(22,'2024-10-01 15:44:57.972','2024-10-01 15:44:57.972',NULL,'char','','',1,0,5),(23,'2024-10-01 15:44:57.972','2024-10-01 15:44:57.972',NULL,'varchar','1','mysql',1,1,5),(24,'2024-10-01 15:44:57.972','2024-10-01 15:44:57.972',NULL,'tinyblob','2','mysql',1,2,5),(25,'2024-10-01 15:44:57.972','2024-10-01 15:44:57.972',NULL,'tinytext','3','mysql',1,3,5),(26,'2024-10-01 15:44:57.972','2024-10-01 15:44:57.972',NULL,'text','4','mysql',1,4,5),(27,'2024-10-01 15:44:57.972','2024-10-01 15:44:57.972',NULL,'blob','5','mysql',1,5,5),(28,'2024-10-01 15:44:57.972','2024-10-01 15:44:57.972',NULL,'mediumblob','6','mysql',1,6,5),(29,'2024-10-01 15:44:57.972','2024-10-01 15:44:57.972',NULL,'mediumtext','7','mysql',1,7,5),(30,'2024-10-01 15:44:57.972','2024-10-01 15:44:57.972',NULL,'longblob','8','mysql',1,8,5),(31,'2024-10-01 15:44:57.972','2024-10-01 15:44:57.972',NULL,'longtext','9','mysql',1,9,5),(32,'2024-10-01 15:44:58.081','2024-10-01 15:44:58.081',NULL,'tinyint','1','mysql',1,0,6),(33,'2024-10-01 15:44:58.081','2024-10-01 15:44:58.081',NULL,'bool','2','pgsql',1,0,6);
/*!40000 ALTER TABLE `sys_dictionary_details` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-10-09 13:42:51
