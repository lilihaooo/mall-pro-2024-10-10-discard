-- MySQL dump 10.13  Distrib 9.0.1, for Linux (aarch64)
--
-- Host: 47.109.31.106    Database: xtaoke
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
-- Table structure for table `collect`
--

DROP TABLE IF EXISTS `collect`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `collect` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '''ID''',
  `user_id` bigint NOT NULL COMMENT '''用户ID''',
  `goods_id` bigint NOT NULL COMMENT '''商品ID''',
  `created_at` datetime DEFAULT NULL COMMENT '''创建时间''',
  `title` varchar(255) NOT NULL COMMENT '''商品标题''',
  `image_id` bigint DEFAULT NULL COMMENT '''图片ID''',
  `post_coupon_price` decimal(10,2) DEFAULT NULL COMMENT '''优惠券后的价格''',
  `commission_value` decimal(10,2) DEFAULT NULL COMMENT '''佣金金额''',
  `commission_rate` int DEFAULT NULL COMMENT '''佣金比率（整数）''',
  `coupon_value` decimal(10,2) DEFAULT NULL COMMENT '''优惠券金额''',
  `coupon_end_time` datetime DEFAULT NULL COMMENT '''优惠券结束时间''',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=700955 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `collect`
--

LOCK TABLES `collect` WRITE;
/*!40000 ALTER TABLE `collect` DISABLE KEYS */;
INSERT INTO `collect` VALUES (70239,1,70239,'2024-10-03 11:11:05','【南国食品】网红零食去核椰奶枣100g*3袋  大枣特产零食小吃',885,51.40,10.30,20,8.50,'2024-11-01 23:59:59'),(700954,1,700954,'2024-10-03 23:18:25','办公室巧克力棒巧克力饼干巧克力酱饼干棒宝宝磨牙休闲浪漫七夕礼',1185,1.50,0.30,20,8.40,'2024-11-01 23:59:59');
/*!40000 ALTER TABLE `collect` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-10-09 13:41:25
