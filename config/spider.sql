/*
SQLyog Ultimate v12.09 (64 bit)
MySQL - 8.0.31 : Database - lj_spider
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`lj_spider` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `lj_spider`;

/*Table structure for table `house_info` */

DROP TABLE IF EXISTS `house_info`;

CREATE TABLE `house_info` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'id',
  `house_id` varchar(512) COLLATE utf8mb4_bin NOT NULL COMMENT '链家房屋id',
  `title` varchar(512) COLLATE utf8mb4_bin NOT NULL COMMENT '房源标题',
  `house_type` varchar(256) COLLATE utf8mb4_bin NOT NULL COMMENT '几居室',
  `acreage` float NOT NULL COMMENT '面积',
  `face` varchar(32) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '朝向',
  `floor` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '楼层分布',
  `total_price` int NOT NULL COMMENT '总价',
  `unit_price` int NOT NULL COMMENT '单价',
  `area` varchar(64) COLLATE utf8mb4_bin NOT NULL COMMENT '区域',
  `plate` varchar(64) COLLATE utf8mb4_bin NOT NULL COMMENT '板块',
  `residential_area_name` varchar(512) COLLATE utf8mb4_bin NOT NULL COMMENT '小区名称',
  `house_birth` varchar(64) COLLATE utf8mb4_bin NOT NULL COMMENT '建成年限',
  `total_height` int NOT NULL COMMENT '楼层总高',
  `publish_time` varchar(256) COLLATE utf8mb4_bin NOT NULL COMMENT '上架时间',
  `created_time_month` varchar(64) COLLATE utf8mb4_bin NOT NULL COMMENT '创建时间-月',
  `created_time_day` int NOT NULL COMMENT '创建时间-天',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
