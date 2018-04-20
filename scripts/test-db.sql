# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: localhost (MySQL 5.7.20)
# Database: medrecV1
# Generation Time: 2017-10-27 02:27:26 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table document_info
# ------------------------------------------------------------

USE medrec-v1;

DROP TABLE IF EXISTS `document_info`;

CREATE TABLE `document_info` (
  `DocumentID` int(11) NOT NULL,
  `PatientID` int(11) NOT NULL,
  `PracticeID` int(11) NOT NULL,
  `RecvdDateTime` datetime NOT NULL,
  `DocDateTime` datetime NOT NULL,
  PRIMARY KEY (`DocumentID`),
  UNIQUE KEY `DocumentID` (`DocumentID`),
  KEY `PatientID` (`PatientID`),
  CONSTRAINT `document_info_ibfk_1` FOREIGN KEY (`PatientID`) REFERENCES `patient_info` (`PatientID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;



# Dump of table encounter_info
# ------------------------------------------------------------

DROP TABLE IF EXISTS `encounter_info`;

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



# Dump of table immunization_info
# ------------------------------------------------------------

DROP TABLE IF EXISTS `immunization_info`;

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



# Dump of table medication_info
# ------------------------------------------------------------

DROP TABLE IF EXISTS `medication_info`;

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



# Dump of table patient_info
# ------------------------------------------------------------

DROP TABLE IF EXISTS `patient_info`;

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



# Dump of table result_info
# ------------------------------------------------------------

DROP TABLE IF EXISTS `result_info`;

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



# Dump of table resultdetail_info
# ------------------------------------------------------------

DROP TABLE IF EXISTS `resultdetail_info`;

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




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
