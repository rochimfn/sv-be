CREATE DATABASE `article`;

CREATE TABLE `posts` (
  `Id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `Title` varchar(100) DEFAULT NULL,
  `Content` text,
  `Category` varchar(100) DEFAULT NULL,
  `Created_date` timestamp NULL DEFAULT NULL,
  `Updated_date` timestamp NULL DEFAULT NULL,
  `Status` varchar(100) DEFAULT NULL COMMENT 'Publish,Draft,Thrash',
  PRIMARY KEY (`Id`)
);