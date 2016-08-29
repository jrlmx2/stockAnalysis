# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.5.42)
# Database: sa
# Generation Time: 2016-08-29 18:35:26 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table quotes
# ------------------------------------------------------------

DROP TABLE IF EXISTS `quotes`;

CREATE TABLE `quotes` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `ask` float NOT NULL,
  `askvolume` int(11) NOT NULL,
  `bid` float NOT NULL,
  `bidvolume` int(11) NOT NULL,
  `symbol_id` int(11) NOT NULL,
  `timestamp` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`id`),
  KEY `symbol` (`symbol_id`),
  KEY `timestamp` (`timestamp`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table symbols
# ------------------------------------------------------------

DROP TABLE IF EXISTS `symbols`;

CREATE TABLE `symbols` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `symbol` varchar(10) NOT NULL DEFAULT '',
  `created_at` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table trades
# ------------------------------------------------------------

DROP TABLE IF EXISTS `trades`;

CREATE TABLE `trades` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `last` float NOT NULL,
  `symbol_id` int(11) NOT NULL,
  `timestamp` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `tradedvolume` int(11) NOT NULL,
  `vwa` float NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table watchlist
# ------------------------------------------------------------

DROP TABLE IF EXISTS `watchlist`;

CREATE TABLE `watchlist` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(200) NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table watchlist_options
# ------------------------------------------------------------

DROP TABLE IF EXISTS `watchlist_options`;

CREATE TABLE `watchlist_options` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `watchlist_id` int(11) NOT NULL,
  `option_type` int(11) NOT NULL,
  `option_value` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `watchlist_id` (`watchlist_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table watchlist_symbols
# ------------------------------------------------------------

DROP TABLE IF EXISTS `watchlist_symbols`;

CREATE TABLE `watchlist_symbols` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `symbol_id` int(11) NOT NULL,
  `watchlist_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `symbol_id_2` (`symbol_id`,`watchlist_id`),
  KEY `symbol_id` (`symbol_id`),
  KEY `watchlist_id` (`watchlist_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
