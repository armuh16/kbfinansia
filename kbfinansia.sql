-- MySQL dump 10.13  Distrib 8.0.26, for macos11 (x86_64)
--
-- Host: localhost    Database: kbfinansia
-- ------------------------------------------------------
-- Server version	8.2.0

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `assets`
--

DROP TABLE IF EXISTS `assets`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `assets` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `tenor_id` int NOT NULL,
  `contract_number` int NOT NULL,
  `on_the_road` int NOT NULL,
  `admin_fee` int NOT NULL,
  `installment` int NOT NULL,
  `interest` float NOT NULL,
  `asset_name` varchar(20) NOT NULL,
  `grand_total` int NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `tenor_id` (`tenor_id`),
  CONSTRAINT `assets_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  CONSTRAINT `assets_ibfk_2` FOREIGN KEY (`tenor_id`) REFERENCES `tenors` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `assets`
--

LOCK TABLES `assets` WRITE;
/*!40000 ALTER TABLE `assets` DISABLE KEYS */;
INSERT INTO `assets` VALUES (1,3,8,123456789,1000000,100000,6,10,'Handphone',1118333,'2023-12-26 12:02:25','2023-12-26 12:02:25',NULL),(2,3,7,123456789,500000,50000,3,10,'Makeup',568333,'2023-12-26 12:03:39','2023-12-26 12:03:39',NULL),(3,3,6,123456789,500000,50000,2,10,'Sepeda',577500,'2023-12-26 12:04:47','2023-12-26 12:04:47',NULL),(4,3,5,123456789,500000,50000,1,10,'Setrika',605000,'2023-12-26 12:05:45','2023-12-26 12:05:45',NULL),(5,2,1,123456789,60000,5000,1,10,'Pulsa',71500,'2023-12-26 12:07:25','2023-12-26 12:07:25',NULL),(6,2,2,123456789,100000,10000,2,10,'Popok',115500,'2023-12-26 12:08:12','2023-12-26 12:08:12',NULL),(7,2,3,123456789,300000,30000,3,10,'Dispenser',341000,'2023-12-26 12:08:45','2023-12-26 12:08:45',NULL),(8,2,4,123456789,500000,50000,6,10,'Handphone',559167,'2023-12-26 12:09:11','2023-12-26 12:09:11',NULL);
/*!40000 ALTER TABLE `assets` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `goose_db_version`
--

DROP TABLE IF EXISTS `goose_db_version`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `goose_db_version` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `version_id` bigint NOT NULL,
  `is_applied` tinyint(1) NOT NULL,
  `tstamp` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=64 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goose_db_version`
--

LOCK TABLES `goose_db_version` WRITE;
/*!40000 ALTER TABLE `goose_db_version` DISABLE KEYS */;
INSERT INTO `goose_db_version` VALUES (1,0,1,'2023-12-21 07:23:42'),(60,1,1,'2023-12-26 11:41:01'),(61,2,1,'2023-12-26 11:41:01'),(62,3,1,'2023-12-26 11:41:01'),(63,4,1,'2023-12-26 11:41:01');
/*!40000 ALTER TABLE `goose_db_version` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tenors`
--

DROP TABLE IF EXISTS `tenors`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tenors` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `tenor` int NOT NULL,
  `limit` int NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `tenors_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tenors`
--

LOCK TABLES `tenors` WRITE;
/*!40000 ALTER TABLE `tenors` DISABLE KEYS */;
INSERT INTO `tenors` VALUES (1,2,1,40000,'2023-12-26 11:55:04','2023-12-26 12:07:25',NULL),(2,2,2,100000,'2023-12-26 11:55:12','2023-12-26 12:08:12',NULL),(3,2,3,200000,'2023-12-26 11:55:48','2023-12-26 12:08:45',NULL),(4,2,6,200000,'2023-12-26 11:55:54','2023-12-26 12:09:11',NULL),(5,3,1,500000,'2023-12-26 11:56:10','2023-12-26 12:05:45',NULL),(6,3,2,700000,'2023-12-26 11:56:15','2023-12-26 12:04:47',NULL),(7,3,3,1000000,'2023-12-26 11:56:36','2023-12-26 12:03:39',NULL),(8,3,6,1000000,'2023-12-26 11:56:46','2023-12-26 12:02:25',NULL);
/*!40000 ALTER TABLE `tenors` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_details`
--

DROP TABLE IF EXISTS `user_details`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_details` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `nik` int NOT NULL,
  `full_name` varchar(255) NOT NULL,
  `legal_name` varchar(255) NOT NULL,
  `place_of_birth` varchar(20) NOT NULL,
  `date_of_birth` varchar(20) NOT NULL,
  `salary` int NOT NULL,
  `user_ktp` varchar(255) NOT NULL,
  `user_photo` varchar(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `user_details_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_details`
--

LOCK TABLES `user_details` WRITE;
/*!40000 ALTER TABLE `user_details` DISABLE KEYS */;
INSERT INTO `user_details` VALUES (1,2,123456789,'Budi','Budi saja','Jakarta','2023-11-10',10000000,'https://t.ly/VlChe','https://t.ly/VlChe','2023-12-26 11:58:33','2023-12-26 11:58:33',NULL),(2,3,123456789,'Annisa','Annisa saja','Jakarta','2023-11-10',10000000,'https://t.ly/VlChe','https://t.ly/VlChe','2023-12-26 11:58:52','2023-12-26 11:58:52',NULL);
/*!40000 ALTER TABLE `user_details` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `role` int NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'Admin','$2a$12$yT.dJTZnu4FRJq9zXw0mBOA/xmZHJPVi5ni13Zk9Pn6E0QmwKkZTu','admin@mail.com',1,'2023-12-26 11:41:01','2023-12-26 11:41:01',NULL),(2,'Budi','$2a$12$yT.dJTZnu4FRJq9zXw0mBOA/xmZHJPVi5ni13Zk9Pn6E0QmwKkZTu','user1@mail.com',2,'2023-12-26 11:41:01','2023-12-26 11:41:01',NULL),(3,'Annisa','$2a$12$yT.dJTZnu4FRJq9zXw0mBOA/xmZHJPVi5ni13Zk9Pn6E0QmwKkZTu','user2@mail.com',2,'2023-12-26 11:41:01','2023-12-26 11:41:01',NULL);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-12-26 20:20:50
