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
-- Table structure for table `planet`
--

DROP TABLE IF EXISTS `planet`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `planet` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `rotation_period` int DEFAULT NULL,
  `orbital_period` int DEFAULT NULL,
  `diameter` int DEFAULT NULL,
  `climate` varchar(100) DEFAULT NULL,
  `gravity` varchar(100) DEFAULT NULL,
  `terrain` varchar(100) DEFAULT NULL,
  `surface_water` int DEFAULT NULL,
  `population` bigint DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `planet`
--

LOCK TABLES `planet` WRITE;
/*!40000 ALTER TABLE `planet` DISABLE KEYS */;
INSERT INTO `planet` VALUES (1,'Tatooine',23,304,10465,'arid','1 standard','desert',1,200000,'2023-10-14 12:12:17'),(2,'Alderaan',24,364,12500,'temperate','1 Standard','grasslands, mountains',40,2000000000,'2023-10-14 12:12:17'),(3,'Yavin IV',24,4818,10200,'temperate, tropical','1 standard','jungle, rainforests',8,1000,'2023-10-14 12:12:17'),(4,'Hoth',23,549,7200,'Frozen','1.1 Standard','tundra, ice caves, mountain ranges',100,NULL,'2023-10-14 12:12:17'),(5,'Dagobah',23,341,8900,'Murky','n/a','Swamp, Jungles',8,NULL,'2023-10-14 12:12:17'),(6,'Bespin',12,5110,118000,'temperate','1.5 (surface), 1 standard (Cloud City)','gas giant',0,6000000,'2023-10-14 12:12:17'),(7,'Endor',18,402,4900,'temperate','0.85 standard','forests, mountains, lakes',8,30000000,'2023-10-14 12:12:17'),(8,'Naboo',26,312,12120,'temperate','1 Standards','grassy hills, swamps, forests, mountains',12,4500000000,'2023-10-14 12:12:17'),(9,'Coruscant',24,368,12240,'temparate','1 standard','Cityscape, Mountains',NULL,1000000000000,'2023-10-14 12:12:17'),(10,'Kamino',27,463,19720,'temperate','1 standard','ocean',100,1000000000,'2023-10-14 12:12:17');
/*!40000 ALTER TABLE `planet` ENABLE KEYS */;
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
