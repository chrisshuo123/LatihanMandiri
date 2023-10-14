-- MySQL dump 10.13  Distrib 8.0.31, for Win64 (x86_64)
--
-- Host: localhost    Database: swapi_db
-- ------------------------------------------------------
-- Server version	8.0.31

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
-- Table structure for table `starships`
--

DROP TABLE IF EXISTS `starships`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `starships` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` longtext NOT NULL,
  `model` longtext,
  `manufacturer` longtext,
  `cost_in_credits` bigint DEFAULT NULL,
  `length` bigint DEFAULT NULL,
  `max_atmosphering_speed` bigint DEFAULT NULL,
  `crew` longtext,
  `passengers` bigint DEFAULT NULL,
  `cargo_capacity` bigint DEFAULT NULL,
  `consumables` longtext,
  `hyperdrive_rating` bigint DEFAULT NULL,
  `MGLT` int DEFAULT NULL,
  `starship_class` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `starships`
--

LOCK TABLES `starships` WRITE;
/*!40000 ALTER TABLE `starships` DISABLE KEYS */;
INSERT INTO `starships` VALUES (1,'CR90 Corvette','CR90 Corvette','Corellian Engineering Corporation',3500000,150,950,'30-165',600,3000000,'1 year',2,60,'corvette'),(2,'Star Destroyer','Imperial I-class Star Destroyer','Kuat Drive Yards',150000000,2,975,'47.060',NULL,36000000,'2 years',2,60,'Star Destroyer'),(3,'Sentinel-class landing craft','Sentinel-class landing craft','Sienar Fleet Systems, Cyngus Spaceworks',240000,38,1000,'5',75,180000,'1 month',1,70,'landing craft'),(4,'Death Star','DS-1 Orbital Battle Station','Imperial Department of Military Research, Sienar Fleet Systems',1000000000000,120000,NULL,'342.953',843342,1000000000000,'3 years',4,10,'Deep Space Mobile Battlestation'),(5,'Millennium Falcon','YT-1300 light freighter','Corellian Engineering Corporation',100000,34,1050,'4',6,100000,'2 months',0,75,'Light freighter');
/*!40000 ALTER TABLE `starships` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-10-14 20:18:35
