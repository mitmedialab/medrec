CREATE DATABASE  IF NOT EXISTS `medrec-v1` /*!40100 DEFAULT CHARACTER SET latin1 */;
USE `medrec-v1`;
-- MySQL dump 10.13  Distrib 5.7.17, for macos10.12 (x86_64)
--
-- Host: localhost    Database: medrec-v1
-- ------------------------------------------------------
-- Server version	5.7.20

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `document_info`
--

DROP TABLE IF EXISTS `document_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `document_info` (
  `DocumentID` int(11) NOT NULL,
  `PatientID` int(11) NOT NULL,
  `PracticeID` int(11) NOT NULL,
  `RecvdDateTime` datetime NOT NULL,
  `DocDateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`DocumentID`),
  UNIQUE KEY `DocumentID` (`DocumentID`),
  KEY `PatientID` (`PatientID`),
  CONSTRAINT `document_info_ibfk_1` FOREIGN KEY (`PatientID`) REFERENCES `patient_info` (`PatientID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `document_info`
--

LOCK TABLES `document_info` WRITE;
/*!40000 ALTER TABLE `document_info` DISABLE KEYS */;
INSERT INTO `document_info` VALUES (145,1,1234567,'2017-01-01 00:00:00','2017-11-15 11:05:35'),(1224,3,1234567,'2017-01-01 00:00:00','2017-11-14 18:50:11'),(1246,2,1234567,'2017-01-01 00:00:00','2017-11-14 18:48:17'),(12245,1,1234567,'2017-01-01 00:00:00','2017-11-14 18:52:40'),(12246,2,1234567,'2017-01-01 00:00:00','2017-11-14 18:49:25'),(12346,2,1234567,'2017-01-01 00:00:00','2017-11-14 18:46:35'),(25554,2,1234567,'2017-01-01 00:00:00','2017-11-19 17:03:22'),(34334,3,1234567,'2017-01-01 00:00:00','2017-11-20 11:16:42'),(122545,1,1234567,'2017-01-01 00:00:00','2017-11-15 01:40:50'),(255234,1,1234567,'2017-01-01 00:00:00','2017-11-19 16:32:07'),(276234,1,1234567,'2017-01-01 00:00:00','2017-11-19 16:31:59'),(276847,1,1234567,'2017-01-01 00:00:00','2017-11-16 16:11:02'),(444334,3,1234567,'2017-01-01 00:00:00','2017-11-20 11:16:51'),(445534,2,1234567,'2017-01-01 00:00:00','2017-11-20 11:17:01'),(1234567,1,1234567,'2017-01-01 00:00:00','2017-11-14 17:55:04'),(2544334,3,1234567,'2017-01-01 00:00:00','2017-11-20 11:16:33'),(2555334,3,1234567,'2017-01-01 00:00:00','2017-11-19 17:03:32'),(12345678,2,1234567,'2017-01-01 00:00:00','2017-11-14 17:55:39');
/*!40000 ALTER TABLE `document_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `encounter_info`
--

DROP TABLE IF EXISTS `encounter_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `encounter_info` (
  `DocumentID` int(11) NOT NULL,
  `VisitCode` varchar(30) NOT NULL,
  `VisitCodeDisplayName` varchar(30) NOT NULL,
  `NormalisedCodeSystemName` varchar(30) NOT NULL,
  `NormalisedDate` datetime NOT NULL,
  `ProviderID` int(11) NOT NULL,
  KEY `DocumentID` (`DocumentID`),
  CONSTRAINT `encounter_info_ibfk_1` FOREIGN KEY (`DocumentID`) REFERENCES `document_info` (`DocumentID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `encounter_info`
--

LOCK TABLES `encounter_info` WRITE;
/*!40000 ALTER TABLE `encounter_info` DISABLE KEYS */;
/*!40000 ALTER TABLE `encounter_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `immunization_info`
--

DROP TABLE IF EXISTS `immunization_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `immunization_info` (
  `DocumentID` int(11) NOT NULL,
  `ProviderID` int(11) NOT NULL,
  `ImmunizationCode` varchar(30) NOT NULL,
  `NormalizedCodeSystemName` varchar(30) NOT NULL,
  `ImmunizationName` varchar(30) NOT NULL,
  `NormalisedImmunizationDate` datetime NOT NULL,
  `AdministrationStatus` varchar(30) NOT NULL,
  `Dosage` float NOT NULL,
  `DosageUnit` varchar(30) NOT NULL,
  `RouteCode` varchar(30) NOT NULL,
  `RouteName` varchar(30) NOT NULL,
  KEY `DocumentID` (`DocumentID`),
  CONSTRAINT `immunization_info_ibfk_1` FOREIGN KEY (`DocumentID`) REFERENCES `document_info` (`DocumentID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `immunization_info`
--

LOCK TABLES `immunization_info` WRITE;
/*!40000 ALTER TABLE `immunization_info` DISABLE KEYS */;
/*!40000 ALTER TABLE `immunization_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `medication_info`
--

DROP TABLE IF EXISTS `medication_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `medication_info` (
  `DocumentID` int(11) NOT NULL,
  `MedicationCode` varchar(30) NOT NULL,
  `NormalizedCodeSystemName` varchar(30) NOT NULL,
  `MedicationName` varchar(30) NOT NULL,
  `MedicationStatus` varchar(30) NOT NULL,
  `NormalisedStartDate` datetime NOT NULL,
  `NormalisedEndDate` datetime NOT NULL,
  `DosageInterval` float NOT NULL,
  `DosagePeriod` varchar(30) NOT NULL,
  `DosageUnit` varchar(30) NOT NULL,
  `DosageQuantity` float NOT NULL,
  KEY `DocumentID` (`DocumentID`),
  CONSTRAINT `medication_info_ibfk_1` FOREIGN KEY (`DocumentID`) REFERENCES `document_info` (`DocumentID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `medication_info`
--

LOCK TABLES `medication_info` WRITE;
/*!40000 ALTER TABLE `medication_info` DISABLE KEYS */;
/*!40000 ALTER TABLE `medication_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `patient_info`
--

DROP TABLE IF EXISTS `patient_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `patient_info` (
  `PatientID` int(11) NOT NULL,
  `LastName` varchar(30) DEFAULT NULL,
  `FirstName` varchar(30) DEFAULT NULL,
  `MiddleName` varchar(30) DEFAULT NULL,
  `Gender` varchar(30) DEFAULT NULL,
  `DOB` date DEFAULT NULL,
  `Address1` varchar(30) DEFAULT NULL,
  `Address2` varchar(30) DEFAULT NULL,
  `City` varchar(30) DEFAULT NULL,
  `State` varchar(30) DEFAULT NULL,
  `Zip` int(11) DEFAULT NULL,
  `Telecom1` int(11) DEFAULT NULL,
  `Race` varchar(30) DEFAULT NULL,
  `Ethnicity` varchar(30) DEFAULT NULL,
  `EducationLevel` varchar(30) DEFAULT NULL,
  `Lang` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`PatientID`),
  UNIQUE KEY `PatientID` (`PatientID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `patient_info`
--

LOCK TABLES `patient_info` WRITE;
/*!40000 ALTER TABLE `patient_info` DISABLE KEYS */;
INSERT INTO `patient_info` VALUES (1,'Smith','John',NULL,'M','1984-03-15',NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL),(2,'Smith','Jane',NULL,'F','1984-03-15',NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL),(3,'Smith','Jane',NULL,'F','1986-03-15',NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL),(4,'Smith','Jane',NULL,'F','1986-03-15',NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL),(5,'Jones','Jane',NULL,'F','1986-03-15',NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL),(6,'Jones','Jane',NULL,'F','1986-03-15',NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL),(7,'Jones','Jake',NULL,'F','1986-03-15',NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL),(8,'Jones','Jake',NULL,'F','1986-03-15',NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL);
/*!40000 ALTER TABLE `patient_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `result_info`
--

DROP TABLE IF EXISTS `result_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `result_info` (
  `ResultsInfoID` int(11) NOT NULL,
  `DocID` int(11) NOT NULL,
  `ResultCode` int(11) NOT NULL,
  `NormalisedCodeSystemName` varchar(30) NOT NULL,
  `LabName` varchar(30) NOT NULL,
  `LabStatus` varchar(30) NOT NULL,
  `NormalisedObservationDate` datetime NOT NULL,
  PRIMARY KEY (`ResultsInfoID`),
  UNIQUE KEY `ResultsInfoID` (`ResultsInfoID`),
  KEY `DocID` (`DocID`),
  CONSTRAINT `result_info_ibfk_1` FOREIGN KEY (`DocID`) REFERENCES `document_info` (`DocumentID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `result_info`
--

LOCK TABLES `result_info` WRITE;
/*!40000 ALTER TABLE `result_info` DISABLE KEYS */;
/*!40000 ALTER TABLE `result_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `resultdetail_info`
--

DROP TABLE IF EXISTS `resultdetail_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `resultdetail_info` (
  `ResultsDetailID` int(11) NOT NULL,
  `ResultsInfoID` int(11) NOT NULL,
  `ComponentName` varchar(30) NOT NULL,
  `LabCode` varchar(30) NOT NULL,
  `LabStatus` varchar(30) NOT NULL,
  `NormalisedObservationDate` datetime NOT NULL,
  `ResultValue` varchar(30) NOT NULL,
  `ResultRange` varchar(30) NOT NULL,
  `Unit` varchar(30) NOT NULL,
  PRIMARY KEY (`ResultsDetailID`),
  UNIQUE KEY `ResultsDetailID` (`ResultsDetailID`),
  KEY `ResultsInfoID` (`ResultsInfoID`),
  CONSTRAINT `resultdetail_info_ibfk_1` FOREIGN KEY (`ResultsInfoID`) REFERENCES `result_info` (`ResultsInfoID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `resultdetail_info`
--

LOCK TABLES `resultdetail_info` WRITE;
/*!40000 ALTER TABLE `resultdetail_info` DISABLE KEYS */;
/*!40000 ALTER TABLE `resultdetail_info` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2018-01-25 15:10:20
