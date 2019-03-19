CREATE TABLE `%s`											/* Class Name based */
(
  `SID` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,		/* System ID */
  `PropertyId` BIGINT(20) UNSIGNED NOT NULL,				/* Property ID */
  `Name` VARCHAR(255) COLLATE utf8_unicode_ci NOT NULL,		/* Property Name */
  `Description` TEXT COLLATE utf8_unicode_ci NOT NULL,		/* Property Description */
  `Type` VARCHAR(255) COLLATE utf8_unicode_ci NOT NULL,		/* Property Type */
  `CID` BIGINT(20) UNSIGNED NOT NULL,						/* Class ID */
  `ToC` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,		/* Time of Creation */
  `ToU` DATETIME,											/* Time of Update */
  PRIMARY KEY (`SID`),
  UNIQUE KEY `SID_UNIQUE` (`SID`),
  UNIQUE KEY `PropertyId_UNIQUE` (`PropertyId`),
  UNIQUE KEY `Name_UNIQUE` (`Name`),
  KEY `PropertyId` (`PropertyId`),
  KEY `Name` (`Name`),
  
  CONSTRAINT `FK_%s_CID_TO_%s_CID`							/* This -> Classes */
    FOREIGN KEY (`CID`)
    REFERENCES `%s` (`CID`)									/* Classes Table */
    ON DELETE RESTRICT
    ON UPDATE RESTRICT
	
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
