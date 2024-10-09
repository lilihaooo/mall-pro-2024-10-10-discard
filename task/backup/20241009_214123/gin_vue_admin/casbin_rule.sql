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
-- Table structure for table `casbin_rule`
--

DROP TABLE IF EXISTS `casbin_rule`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `casbin_rule` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v0` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v1` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v2` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v3` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v4` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v5` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB AUTO_INCREMENT=7018 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `casbin_rule`
--

LOCK TABLES `casbin_rule` WRITE;
/*!40000 ALTER TABLE `casbin_rule` DISABLE KEYS */;
INSERT INTO `casbin_rule` VALUES (6899,'p','1','/api/createApi','POST','','',''),(6900,'p','1','/api/deleteApi','POST','','',''),(6903,'p','1','/api/deleteApisByIds','DELETE','','',''),(6906,'p','1','/api/freshCasbin','GET','','',''),(6904,'p','1','/api/getAllApis','POST','','',''),(6901,'p','1','/api/getApiById','POST','','',''),(6905,'p','1','/api/getApiList','POST','','',''),(6902,'p','1','/api/updateApi','POST','','',''),(6949,'p','1','/authority/copyAuthority','POST','','',''),(6946,'p','1','/authority/createAuthority','POST','','',''),(6947,'p','1','/authority/deleteAuthority','POST','','',''),(6951,'p','1','/authority/getAuthorityList','POST','','',''),(6950,'p','1','/authority/setDataAuthority','POST','','',''),(6948,'p','1','/authority/updateAuthority','PUT','','',''),(6973,'p','1','/authorityBtn/canRemoveAuthorityBtn','POST','','',''),(6971,'p','1','/authorityBtn/getAuthorityBtn','POST','','',''),(6972,'p','1','/authorityBtn/setAuthorityBtn','POST','','',''),(6940,'p','1','/autoCode/createPackage','POST','','',''),(6943,'p','1','/autoCode/createPlug','POST','','',''),(6939,'p','1','/autoCode/createTemp','POST','','',''),(6942,'p','1','/autoCode/delPackage','POST','','',''),(6959,'p','1','/autoCode/delSysHistory','POST','','',''),(6937,'p','1','/autoCode/getColumn','GET','','',''),(6935,'p','1','/autoCode/getDB','GET','','',''),(6957,'p','1','/autoCode/getMeta','POST','','',''),(6941,'p','1','/autoCode/getPackage','POST','','',''),(6960,'p','1','/autoCode/getSysHistory','POST','','',''),(6936,'p','1','/autoCode/getTables','GET','','',''),(6944,'p','1','/autoCode/installPlugin','POST','','',''),(6938,'p','1','/autoCode/preview','POST','','',''),(6945,'p','1','/autoCode/pubPlug','POST','','',''),(6958,'p','1','/autoCode/rollback','POST','','',''),(6896,'p','1','/base/captcha','POST','','',''),(6895,'p','1','/base/login','POST','','',''),(7015,'p','1','/brand/list','GET','','',''),(6934,'p','1','/casbin/getPolicyPathByAuthorityId','POST','','',''),(6933,'p','1','/casbin/updateCasbin','POST','','',''),(6996,'p','1','/category/','GET','','',''),(6985,'p','1','/customer/customer','DELETE','','',''),(6986,'p','1','/customer/customer','GET','','',''),(6983,'p','1','/customer/customer','POST','','',''),(6984,'p','1','/customer/customer','PUT','','',''),(6987,'p','1','/customer/customerList','GET','','',''),(6992,'p','1','/fileUploadAndDownload/breakpointContinue','POST','','',''),(6994,'p','1','/fileUploadAndDownload/breakpointContinueFinish','POST','','',''),(6990,'p','1','/fileUploadAndDownload/deleteFile','POST','','',''),(6991,'p','1','/fileUploadAndDownload/editFileName','POST','','',''),(6993,'p','1','/fileUploadAndDownload/findFile','GET','','',''),(6989,'p','1','/fileUploadAndDownload/getFileList','POST','','',''),(6995,'p','1','/fileUploadAndDownload/removeChunk','POST','','',''),(6988,'p','1','/fileUploadAndDownload/upload','POST','','',''),(7008,'p','1','/goods/batchCancelCollect','POST','','',''),(7007,'p','1','/goods/cancelCollect','POST','','',''),(7006,'p','1','/goods/collect','POST','','',''),(7002,'p','1','/goods/image/delete','GET','','',''),(7003,'p','1','/goods/image/setCover','POST','','',''),(7001,'p','1','/goods/image/upload','POST','','',''),(7000,'p','1','/goods/info','GET','','',''),(7009,'p','1','/goods/myCollect','GET','','',''),(7010,'p','1','/goods/myPromotion','GET','','',''),(7011,'p','1','/goods/promotion','POST','','',''),(6999,'p','1','/goods/search/list','GET','','',''),(7004,'p','1','/goods/updateBaseInfo','POST','','',''),(7005,'p','1','/goods/updateCouponInfo','POST','','',''),(7012,'p','1','/goods/v2/search/list','GET','','',''),(6998,'p','1','/grab/log','GET','','',''),(6997,'p','1','/grab/record','GET','','',''),(6898,'p','1','/init/checkdb','POST','','',''),(6897,'p','1','/init/initdb','POST','','',''),(6907,'p','1','/jwt/jsonInBlacklist','POST','','',''),(6920,'p','1','/menu/addBaseMenu','POST','','',''),(6921,'p','1','/menu/addMenuAuthority','POST','','',''),(6922,'p','1','/menu/deleteBaseMenu','POST','','',''),(6928,'p','1','/menu/getBaseMenuById','POST','','',''),(6926,'p','1','/menu/getBaseMenuTree','POST','','',''),(6924,'p','1','/menu/getMenu','POST','','',''),(6927,'p','1','/menu/getMenuAuthority','POST','','',''),(6925,'p','1','/menu/getMenuList','POST','','',''),(6923,'p','1','/menu/updateBaseMenu','POST','','',''),(7016,'p','1','/search/hotSearch/list','GET','','',''),(7017,'p','1','/search/suggestion/list','GET','','',''),(7013,'p','1','/shop/list','GET','','',''),(6952,'p','1','/sysDictionary/createSysDictionary','POST','','',''),(6953,'p','1','/sysDictionary/deleteSysDictionary','DELETE','','',''),(6955,'p','1','/sysDictionary/findSysDictionary','GET','','',''),(6956,'p','1','/sysDictionary/getSysDictionaryList','GET','','',''),(6954,'p','1','/sysDictionary/updateSysDictionary','PUT','','',''),(6966,'p','1','/sysDictionaryDetail/createSysDictionaryDetail','POST','','',''),(6967,'p','1','/sysDictionaryDetail/deleteSysDictionaryDetail','DELETE','','',''),(6969,'p','1','/sysDictionaryDetail/findSysDictionaryDetail','GET','','',''),(6970,'p','1','/sysDictionaryDetail/getSysDictionaryDetailList','GET','','',''),(6968,'p','1','/sysDictionaryDetail/updateSysDictionaryDetail','PUT','','',''),(6974,'p','1','/sysExportTemplate/createSysExportTemplate','POST','','',''),(6975,'p','1','/sysExportTemplate/deleteSysExportTemplate','DELETE','','',''),(6976,'p','1','/sysExportTemplate/deleteSysExportTemplateByIds','DELETE','','',''),(6981,'p','1','/sysExportTemplate/exportExcel','GET','','',''),(6982,'p','1','/sysExportTemplate/exportTemplate','GET','','',''),(6979,'p','1','/sysExportTemplate/findSysExportTemplate','GET','','',''),(6980,'p','1','/sysExportTemplate/getSysExportTemplateList','GET','','',''),(6978,'p','1','/sysExportTemplate/importExcel','POST','','',''),(6977,'p','1','/sysExportTemplate/updateSysExportTemplate','PUT','','',''),(6961,'p','1','/sysOperationRecord/createSysOperationRecord','POST','','',''),(6962,'p','1','/sysOperationRecord/deleteSysOperationRecord','DELETE','','',''),(6963,'p','1','/sysOperationRecord/deleteSysOperationRecordByIds','DELETE','','',''),(6964,'p','1','/sysOperationRecord/findSysOperationRecord','GET','','',''),(6965,'p','1','/sysOperationRecord/getSysOperationRecordList','GET','','',''),(6931,'p','1','/system/getServerInfo','POST','','',''),(6929,'p','1','/system/getSystemConfig','POST','','',''),(6932,'p','1','/system/reloadSystem','POST','','',''),(6930,'p','1','/system/setSystemConfig','POST','','',''),(7014,'p','1','/tag/list','GET','','',''),(6908,'p','1','/user/admin_register','POST','','',''),(6909,'p','1','/user/changePassword','POST','','',''),(6911,'p','1','/user/deleteUser','DELETE','','',''),(6916,'p','1','/user/getSetting','GET','','',''),(6919,'p','1','/user/getUserInfo','GET','','',''),(6918,'p','1','/user/getUserList','POST','','',''),(6915,'p','1','/user/resetPassword','POST','','',''),(6913,'p','1','/user/setSelfInfo','PUT','','',''),(6917,'p','1','/user/setSetting','POST','','',''),(6914,'p','1','/user/setUserAuthorities','POST','','',''),(6910,'p','1','/user/setUserAuthority','POST','','',''),(6912,'p','1','/user/setUserInfo','PUT','','','');
/*!40000 ALTER TABLE `casbin_rule` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-10-09 13:42:40
