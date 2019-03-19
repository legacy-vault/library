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

// column.go.

package storage

// Storage Column Settings and Methods.

import (
	"fmt"
	"strings"

	"xxx/Database"
)

// Database Identifiers' Parameters.
const DbColumnNameSid = "SID"
const DbColumnTypeSid = "BIGINT(20) UNSIGNED NOT NULL"
const DbColumnNameCid = "CID"
const DbColumnTypeCid = "BIGINT(20) UNSIGNED NOT NULL"
const DbColumnNameClassname = "Name"
const DbColumnTypeClassname = "VARCHAR(255) NOT NULL"
const DbColumnNameToc = "ToC"
const DbColumnTypeToc = "DATETIME NOT NULL"
const DbColumnNameTou = "ToU"
const DbColumnTypeTou = "DATETIME"
const DbColumnNameObjectid = "ObjectId"
const DbColumnTypeObjectid = "BIGINT(20) UNSIGNED NOT NULL"
const DbColumnNamePropertyid = "PropertyId"
const DbColumnTypePropertyid = "BIGINT(20) UNSIGNED NOT NULL"
const DbColumnNamePropertyname = "Name"
const DbColumnTypePropertyname = "VARCHAR(255) NOT NULL"
const DbColumnNamePropertydescription = "Description"
const DbColumnTypePropertydescription = "TEXT NOT NULL"
const DbColumnNamePropertytype = "Type"
const DbColumnTypePropertytype = "VARCHAR(255) NOT NULL"
const DbColumnNamePropertyValue = "Value"

// Returns the Settings for a Class Table.
func (this Storage) columnSettingsForTableClass() []database.DatabaseTableColumnSettings {

	const ColumnsCountExpected = 5

	var columnsSettingsExpected []database.DatabaseTableColumnSettings

	columnsSettingsExpected =
		make([]database.DatabaseTableColumnSettings, ColumnsCountExpected)
	columnsSettingsExpected[0] = database.DatabaseTableColumnSettings{
		Name: DbColumnNameSid,
		Type: DbColumnTypeSid,
	}
	columnsSettingsExpected[1] = database.DatabaseTableColumnSettings{
		Name: DbColumnNameObjectid,
		Type: DbColumnTypeObjectid,
	}
	columnsSettingsExpected[2] = database.DatabaseTableColumnSettings{
		Name: DbColumnNameCid,
		Type: DbColumnTypeCid,
	}
	columnsSettingsExpected[3] = database.DatabaseTableColumnSettings{
		Name: DbColumnNameToc,
		Type: DbColumnTypeToc,
	}
	columnsSettingsExpected[4] = database.DatabaseTableColumnSettings{
		Name: DbColumnNameTou,
		Type: DbColumnTypeTou,
	}

	return columnsSettingsExpected
}

// Returns the Settings for a Classes Table.
func (this Storage) columnSettingsForTableClasses() []database.DatabaseTableColumnSettings {

	const ColumnsCountExpected = 5

	var columnsSettingsExpected []database.DatabaseTableColumnSettings

	columnsSettingsExpected =
		make([]database.DatabaseTableColumnSettings, ColumnsCountExpected)
	columnsSettingsExpected[0] = database.DatabaseTableColumnSettings{
		Name: DbColumnNameSid,
		Type: DbColumnTypeSid,
	}
	columnsSettingsExpected[1] = database.DatabaseTableColumnSettings{
		Name: DbColumnNameCid,
		Type: DbColumnTypeCid,
	}
	columnsSettingsExpected[2] = database.DatabaseTableColumnSettings{
		Name: DbColumnNameClassname,
		Type: DbColumnTypeClassname,
	}
	columnsSettingsExpected[3] = database.DatabaseTableColumnSettings{
		Name: DbColumnNameToc,
		Type: DbColumnTypeToc,
	}
	columnsSettingsExpected[4] = database.DatabaseTableColumnSettings{
		Name: DbColumnNameTou,
		Type: DbColumnTypeTou,
	}

	return columnsSettingsExpected
}

// Returns the Settings for a Class Properties Table.
func (this Storage) columnSettingsForTableClassProperties() []database.DatabaseTableColumnSettings {

	const ColumnsCountExpected = 8

	var columnsSettingsExpected []database.DatabaseTableColumnSettings

	columnsSettingsExpected =
		make([]database.DatabaseTableColumnSettings, ColumnsCountExpected)
	columnsSettingsExpected[0] = database.DatabaseTableColumnSettings{
		Name: DbColumnNameSid,
		Type: DbColumnTypeSid,
	}
	columnsSettingsExpected[1] = database.DatabaseTableColumnSettings{
		Name: DbColumnNamePropertyid,
		Type: DbColumnTypePropertyid,
	}
	columnsSettingsExpected[2] = database.DatabaseTableColumnSettings{
		Name: DbColumnNamePropertyname,
		Type: DbColumnTypePropertyname,
	}
	columnsSettingsExpected[3] = database.DatabaseTableColumnSettings{
		Name: DbColumnNamePropertydescription,
		Type: DbColumnTypePropertydescription,
	}
	columnsSettingsExpected[4] = database.DatabaseTableColumnSettings{
		Name: DbColumnNamePropertytype,
		Type: DbColumnTypePropertytype,
	}
	columnsSettingsExpected[5] = database.DatabaseTableColumnSettings{
		Name: DbColumnNameCid,
		Type: DbColumnTypeCid,
	}
	columnsSettingsExpected[6] = database.DatabaseTableColumnSettings{
		Name: DbColumnNameToc,
		Type: DbColumnTypeToc,
	}
	columnsSettingsExpected[7] = database.DatabaseTableColumnSettings{
		Name: DbColumnNameTou,
		Type: DbColumnTypeTou,
	}

	return columnsSettingsExpected
}

// Returns the Settings for a Class Property Table.
func (this Storage) columnSettingsForTableClassProperty(
	dbColumnTypePropertyValue string,
) []database.DatabaseTableColumnSettings {

	const ColumnsCountExpected = 7

	var columnsSettingsExpected []database.DatabaseTableColumnSettings

	columnsSettingsExpected =
		make([]database.DatabaseTableColumnSettings, ColumnsCountExpected)
	columnsSettingsExpected[0] = database.DatabaseTableColumnSettings{
		Name: DbColumnNameSid,
		Type: DbColumnTypeSid,
	}
	columnsSettingsExpected[1] = database.DatabaseTableColumnSettings{
		Name: DbColumnNamePropertyid,
		Type: DbColumnTypePropertyid,
	}
	columnsSettingsExpected[2] = database.DatabaseTableColumnSettings{
		Name: DbColumnNameObjectid,
		Type: DbColumnTypeObjectid,
	}
	columnsSettingsExpected[3] = database.DatabaseTableColumnSettings{
		Name: DbColumnNameCid,
		Type: DbColumnTypeCid,
	}
	columnsSettingsExpected[4] = database.DatabaseTableColumnSettings{
		Name: DbColumnNamePropertyValue,
		Type: dbColumnTypePropertyValue,
	}
	columnsSettingsExpected[5] = database.DatabaseTableColumnSettings{
		Name: DbColumnNameToc,
		Type: DbColumnTypeToc,
	}
	columnsSettingsExpected[6] = database.DatabaseTableColumnSettings{
		Name: DbColumnNameTou,
		Type: DbColumnTypeTou,
	}

	return columnsSettingsExpected
}

// Verifies the received Column Settings.
// Checks Columns Count, Column Names & Types.
func (this Storage) verifyColumnSettings(
	settingsExpected []database.DatabaseTableColumnSettings,
	settingsReceived []database.DatabaseTableColumnSettings,
) error {

	var err error
	var expectedColumnsCount int
	var receivedColumnsCount int

	// Check Columns Count.
	receivedColumnsCount = len(settingsReceived)
	expectedColumnsCount = len(settingsExpected)
	if receivedColumnsCount != expectedColumnsCount {
		err = fmt.Errorf(
			ErrFormatColumnsCount,
			expectedColumnsCount,
			receivedColumnsCount,
		)
		return err
	}

	// Check Column Names.
	for i := 0; i < receivedColumnsCount; i++ {
		colNameExpected := strings.ToLower(settingsExpected[i].Name)
		colNameReceived := strings.ToLower(settingsReceived[i].Name)
		if colNameReceived != colNameExpected {
			err = fmt.Errorf(
				ErrFormatColumnName,
				colNameExpected,
				colNameReceived,
			)
			return err
		}
	}

	// Check Column Types.
	for i := 0; i < receivedColumnsCount; i++ {
		colTypeExpected := strings.ToLower(settingsExpected[i].Type)
		colTypeReceived := strings.ToLower(settingsReceived[i].Type)
		if colTypeReceived != colTypeExpected {
			err = fmt.Errorf(
				ErrFormatColumnType,
				colTypeExpected,
				colTypeReceived,
			)
			return err
		}
	}

	return nil
}

// Verifies Columns For the Class Table.
// The List of Columns is hard-coded into an
// SQL Command File 'createClassTable.sql'.
func (this Storage) verifyColumnsForTableClass(
	columnsSettings []database.DatabaseTableColumnSettings,
) error {

	var columnsSettingsExpected []database.DatabaseTableColumnSettings
	var err error

	// Compare received Column Settings with proper Settings.
	columnsSettingsExpected = this.columnSettingsForTableClass()
	err = this.verifyColumnSettings(
		columnsSettingsExpected,
		columnsSettings,
	)
	if err != nil {
		return err
	}

	return nil
}

// Verifies Columns For the Classes Table.
// The List of Columns is hard-coded into an
// SQL Command File 'createClassesTable.sql'.
func (this Storage) verifyColumnsForTableClasses(
	columnsSettings []database.DatabaseTableColumnSettings,
) error {

	var columnsSettingsExpected []database.DatabaseTableColumnSettings
	var err error

	// Compare received Column Settings with proper Settings.
	columnsSettingsExpected = this.columnSettingsForTableClasses()
	err = this.verifyColumnSettings(
		columnsSettingsExpected,
		columnsSettings,
	)
	if err != nil {
		return err
	}

	return nil
}

// Verifies Columns For the Class Properties Table.
// The List of Columns is hard-coded into an
// SQL Command File 'createClassPropertiesTable.sql'.
func (this Storage) verifyColumnsForTableClassProperties(
	columnsSettings []database.DatabaseTableColumnSettings,
) error {

	var columnsSettingsExpected []database.DatabaseTableColumnSettings
	var err error

	// Compare received Column Settings with proper Settings.
	columnsSettingsExpected = this.columnSettingsForTableClassProperties()
	err = this.verifyColumnSettings(
		columnsSettingsExpected,
		columnsSettings,
	)
	if err != nil {
		return err
	}

	return nil
}

// Verifies Columns For the Class Property Table.
// The List of Columns is hard-coded into an
// SQL Command File 'createClassPropertyTable.sql'.
func (this Storage) verifyColumnsForTableClassProperty(
	columnsSettings []database.DatabaseTableColumnSettings,
	dbColumnTypePropertyValue string,
) error {

	var columnsSettingsExpected []database.DatabaseTableColumnSettings
	var err error

	// Compare received Column Settings with proper Settings.
	columnsSettingsExpected = this.columnSettingsForTableClassProperty(
		dbColumnTypePropertyValue,
	)
	err = this.verifyColumnSettings(
		columnsSettingsExpected,
		columnsSettings,
	)
	if err != nil {
		return err
	}

	return nil
}
