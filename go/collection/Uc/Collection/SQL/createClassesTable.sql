CREATE TABLE `%s`
(
  `SID` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,		/* System ID */
  `CID` BIGINT(20) UNSIGNED NOT NULL,						/* Class ID */
  `Name` VARCHAR(255) COLLATE utf8_unicode_ci NOT NULL,		/* Class Name */
  `ToC` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,		/* Time of Creation */
  `ToU` DATETIME,											/* Time of Update */
  PRIMARY KEY (`SID`),
  UNIQUE KEY `SID_UNIQUE` (`SID`),
  UNIQUE KEY `CID_UNIQUE` (`CID`),
  UNIQUE KEY `Name_UNIQUE` (`Name`),
  KEY `CID` (`CID`),
  KEY `Name` (`Name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
