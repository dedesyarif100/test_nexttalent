-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               5.7.38 - MySQL Community Server (GPL)
-- Server OS:                    Linux
-- HeidiSQL Version:             12.1.0.6537
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- Dumping database structure for nawadata
CREATE DATABASE IF NOT EXISTS `nawadata` /*!40100 DEFAULT CHARACTER SET latin1 */;
USE `nawadata`;

-- Dumping structure for table nawadata.employment_types
CREATE TABLE IF NOT EXISTS `employment_types` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;

-- Dumping data for table nawadata.employment_types: ~3 rows (approximately)
INSERT IGNORE INTO `employment_types` (`id`, `name`) VALUES
	(1, 'Contract'),
	(2, 'Internship'),
	(3, 'Fulltime');

-- Dumping structure for table nawadata.positions
CREATE TABLE IF NOT EXISTS `positions` (
  `id` bigint(50) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `job_description` longtext,
  `minimum_qualification` varchar(200) DEFAULT NULL,
  `minimum_experience` varchar(100) DEFAULT NULL,
  `skills` longtext,
  `benefit` longtext,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=latin1;

-- Dumping data for table nawadata.positions: ~9 rows (approximately)
INSERT IGNORE INTO `positions` (`id`, `name`, `job_description`, `minimum_qualification`, `minimum_experience`, `skills`, `benefit`) VALUES
	(1, 'Database Administrator', '- Menguasai ms sql server. -Experience 2 years', 'Pendidikan Terakhir D3/S1 jurusan terkait', '2 years', '-SQL SERVER, -MYSQL', 'TES'),
	(2, 'Product Owner', '- Menguasai ms sql server. -Experience 2 years', 'Pendidikan Terakhir D3/S1 jurusan terkait', '2 years', '-SQL SERVER, -MYSQL', 'TES'),
	(3, 'Fullstack Developer', '- Menguasai ms sql server. -Experience 2 years', 'Pendidikan Terakhir D3/S1 jurusan terkait', '2 years', '-SQL SERVER, -MYSQL', 'TES'),
	(4, 'PHP Developer', '- Menguasai ms sql server. -Experience 2 years', 'Pendidikan Terakhir D3/S1 jurusan terkait', '2 years', '-SQL SERVER, -MYSQL', 'TES'),
	(5, 'React Native Developer', '- Menguasai ms sql server. -Experience 2 years', 'Pendidikan Terakhir D3/S1 jurusan terkait', '2 years', '-SQL SERVER, -MYSQL', 'TES'),
	(6, 'Golang Developer', '- Menguasai ms sql server. -Experience 2 years', 'Pendidikan Terakhir D3/S1 jurusan terkait', '2 years', '-SQL SERVER, -MYSQL', 'TES'),
	(7, 'iOS Developer', '- Menguasai ms sql server. -Experience 2 years', 'Pendidikan Terakhir D3/S1 jurusan terkait', '2 years', '-SQL SERVER, -MYSQL', 'TES'),
	(8, 'Project Support', '- Menguasai ms sql server. -Experience 2 years', 'Pendidikan Terakhir D3/S1 jurusan terkait', '2 years', '-SQL SERVER, -MYSQL', 'TES'),
	(9, 'UI UX Designer', '- Menguasai ms sql server. -Experience 2 years', 'Pendidikan Terakhir D3/S1 jurusan terkait', '2 years', '-SQL SERVER, -MYSQL', 'TES');

-- Dumping structure for table nawadata.position_levels
CREATE TABLE IF NOT EXISTS `position_levels` (
  `id` bigint(50) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=latin1;

-- Dumping data for table nawadata.position_levels: ~5 rows (approximately)
INSERT IGNORE INTO `position_levels` (`id`, `name`) VALUES
	(1, 'CEO'),
	(2, 'Manajer'),
	(3, 'Supervisor'),
	(4, 'Pegawai Non Manajemen'),
	(5, 'Lulusan Baru');

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
