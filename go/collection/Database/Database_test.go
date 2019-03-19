// ============================================================================
//
// Copyright © 2019 by McArcher.
//
// All rights reserved. No part of this publication may be reproduced,
// distributed, or transmitted in any form or by any means, including
// photocopying, recording, or other electronic or mechanical methods,
// without the prior written permission of the publisher, except in the case
// of brief quotations embodied in critical reviews and certain other
// noncommercial uses permitted by copyright law. For permission requests,
// write to the publisher, addressed “Copyright Protected Material” at the
// address below.
//
// ============================================================================
//
// Web Site:		'https://github.com/legacy-vault'.
// Author:			McArcher.
// Creation Date:	2019-03-19.
// Web Site Address is an Address in the global Computer Internet Network.
//
// ============================================================================

// Database.go.

package database

// Database.

import (
	"database/sql"
	"fmt"
	"strconv"
	"testing"

	"xxx/Database/Class/Value"
	"xxx/Database/Configuration/Dsn"
	"xxx/Database/Settings"
	"xxx/Test"
)

func Test_Database_Prepare(
	t *testing.T,
) {

	const QueryDeleteTableA = "DROP TABLE IF EXISTS `existent_table`;"
	const QueryCreateTableA = "CREATE TABLE `existent_table` (" +
		"`size` int(10) unsigned NOT NULL," +
		"`name` varchar(45) COLLATE utf8_unicode_ci DEFAULT NULL," +
		"PRIMARY KEY (`size`)," +
		"UNIQUE KEY `ID_UNIQUE` (`size`)" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;"
	const QueryDeleteTableB = "DROP TABLE IF EXISTS `_existent_table`;"
	const QueryCreateTableB = "CREATE TABLE `_existent_table` (" +
		"`size` int(10) unsigned NOT NULL," +
		"`name` varchar(45) COLLATE utf8_unicode_ci DEFAULT NULL," +
		"PRIMARY KEY (`size`)," +
		"UNIQUE KEY `ID_UNIQUE` (`size`)" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;"
	const QueryDeleteTableC = "DROP TABLE IF EXISTS `non_existent_table`;"
	const QueryDeleteTableD = "DROP TABLE IF EXISTS `_non_existent_system_table`;"

	var aTest *test.Test
	var connection *sql.DB
	var dbSettings settings.Settings
	var err error
	var theDsn string

	aTest = test.New(t)

	dbSettings = settings.Settings{
		Type:       "...",
		HostName:   "localhost",
		PortNumber: 3306,
		UserName:   "test",
		Password:   "test",
		Database:   "test",
		Journal:    true,
	}

	// Connect to the Test Database.
	theDsn = fmt.Sprintf(
		dsn.FormatMysql,
		dbSettings.UserName,
		dbSettings.Password,
		dsn.DriverMysqlConnectionType,
		dbSettings.HostName,
		strconv.FormatUint(uint64(dbSettings.PortNumber), 10),
		dbSettings.Database,
		"",
	)
	connection, err = sql.Open("mysql", theDsn)
	aTest.CheckError(err)
	err = connection.Ping()
	aTest.CheckError(err)

	// Create Tables which must exist before Test Start.
	_, err = connection.Exec(QueryDeleteTableA)
	aTest.CheckError(err)
	_, err = connection.Exec(QueryCreateTableA)
	aTest.CheckError(err)
	_, err = connection.Exec(QueryDeleteTableB)
	aTest.CheckError(err)
	_, err = connection.Exec(QueryCreateTableB)
	aTest.CheckError(err)

	// Delete Tables which must not exist before Test Start.
	_, err = connection.Exec(QueryDeleteTableC)
	aTest.CheckError(err)
	_, err = connection.Exec(QueryDeleteTableD)
	aTest.CheckError(err)

	// Disconnect from the Test Database.
	err = connection.Close()
	aTest.CheckError(err)
}

func Test_Database_All(
	t *testing.T,
) {

	// Notes.
	//
	//	During the Test, in the Test Database...
	//	1. System Table '_existent_table' must exist,
	//	2. Table 'existent_table' must exist,
	//	3. System Table '_non_existent_system_table' must not exist,
	//	4. Table 'non_existent_table' must not exist.

	const ErrExpectedErrorMissed = "Expected Error has not been received"
	const ErrConnect = "Connection Error"
	const ErrConfigure = "Configuration Error"
	const ErrTableExists = "Error in 'TableExists'"
	const ErrDisconnect = "Disconnection Error"

	const TableNameExistentA = "existent_table"
	const TableNameExistentB = "_existent_table"

	const TableNameNonExistentA = "non_existent_table"
	const TableNameNonExistentB = "_non_existent_system_table"

	const TableNameNonSystemBadA = "_bad_name_test"

	const DatabaseTypeBad = "Non-Existent-Type"

	var aTest *test.Test
	var columnsSettings []DatabaseTableColumnSettings
	var db *Database
	var dbSettings settings.Settings
	var err error
	var tableExists bool

	aTest = test.New(t)

	dbSettings = settings.Settings{
		Type:       "...",
		HostName:   "localhost",
		PortNumber: 3306,
		UserName:   "test",
		Password:   "test",
		Database:   "test",
		Parameters: "",
		Journal:    true,
	}
	dbSettings.Type = DatabaseTypeBad
	db, err = New(dbSettings.Type)
	if err == nil {
		aTest.Stop(ErrExpectedErrorMissed)
	}

	dbSettings.Type = value.MysqlAlias
	db, err = New(dbSettings.Type)
	aTest.CheckError(err)

	err = db.Configure(dbSettings)
	aTest.CheckError(err)
	if db.IsConfigured() != true {
		aTest.Stop(ErrConfigure)
	}

	err = db.Connect()
	aTest.CheckError(err)
	if db.IsConnected() != true {
		aTest.Stop(ErrConnect)
	}

	tableExists, err = db.TableExists(TableNameNonExistentA)
	aTest.CheckError(err)
	if tableExists != false {
		aTest.Stop(ErrTableExists)
	}

	tableExists, err = db.TableExists(TableNameExistentA)
	aTest.CheckError(err)
	if tableExists != true {
		aTest.Stop(ErrTableExists)
	}

	tableExists, err = db.TableExists(TableNameNonSystemBadA)
	if err == nil {
		aTest.Stop(ErrExpectedErrorMissed)
	}

	tableExists, err = db.SystemTableExists(TableNameExistentB)
	aTest.CheckError(err)
	if tableExists != true {
		aTest.Stop(ErrTableExists)
	}

	tableExists, err = db.SystemTableExists(TableNameNonExistentB)
	aTest.CheckError(err)
	if tableExists != false {
		aTest.Stop(ErrTableExists)
	}

	columnsSettings, err = db.ShowColumns(TableNameExistentA)
	aTest.CheckError(err)
	columnsSettings = columnsSettings // Manual Check.

	err = db.Disconnect()
	aTest.CheckError(err)
	if db.IsConnected() != false {
		aTest.Stop(ErrDisconnect)
	}
}
