CREATE TABLE `%s`											/* Class Name */
(
  `SID` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,		/* System ID */
  `ObjectId` BIGINT(20) UNSIGNED NOT NULL,					/* Object ID */
  `CID` BIGINT(20) UNSIGNED NOT NULL,						/* Class ID [FK] */
  `ToC` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,		/* Time of Creation */
  `ToU` DATETIME,											/* Time of Update */
  PRIMARY KEY (`SID`),
  UNIQUE KEY `SID_UNIQUE` (`SID`),
  UNIQUE KEY `ObjectId_UNIQUE` (`ObjectId`),
  KEY `ObjectId` (`ObjectId`),
  
  CONSTRAINT `FK_%s_CID_TO_%s_CID`							/* This -> Classes */
    FOREIGN KEY (`CID`)
    REFERENCES `%s` (`CID`)									/* Classes Table */
    ON DELETE RESTRICT
    ON UPDATE RESTRICT
	
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
