CREATE TABLE `%s`											/* Class Name based */
(
  `SID` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,		/* System ID */
  `PropertyId` BIGINT(20) UNSIGNED NOT NULL,				/* Property ID */
  `ObjectId` BIGINT(20) UNSIGNED NOT NULL,					/* Class Object ID */
  `CID` BIGINT(20) UNSIGNED NOT NULL,						/* Class ID */
  `Value` %s,												/* Property Value */
  `ToC` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,		/* Time of Creation */
  `ToU` DATETIME,											/* Time of Update */
  PRIMARY KEY (`SID`),
  UNIQUE KEY `SID_UNIQUE` (`SID`),
  UNIQUE KEY `ObjectId_UNIQUE` (`ObjectId`),
  KEY `ObjectId` (`ObjectId`),
  
  CONSTRAINT `FK_%s_PID_TO_%s_PID`			/* This -> Properties */
    FOREIGN KEY (`PropertyId`)
    REFERENCES `%s` (`PropertyId`)							/* Properties Table */
    ON DELETE RESTRICT
    ON UPDATE RESTRICT,
	
  CONSTRAINT `FK_%s_OID_TO_%s_OID`				/* This -> Class */
    FOREIGN KEY (`ObjectId`)
    REFERENCES `%s` (`ObjectId`)							/* Class Table */
    ON DELETE RESTRICT
    ON UPDATE RESTRICT,
  
  CONSTRAINT `FK_%s_CID_TO_%s_CID`							/* This -> Classes */
    FOREIGN KEY (`CID`)
    REFERENCES `%s` (`CID`)									/* Classes Table */
    ON DELETE RESTRICT
    ON UPDATE RESTRICT
	
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
