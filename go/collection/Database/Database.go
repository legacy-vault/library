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
	"errors"
	"fmt"
	"strings"

	"xxx/Database/Class"
	"xxx/Database/Class/Value"
	databaseCollection "xxx/Database/Collection"
	"xxx/Database/Configuration"
	"xxx/Database/Connection"
	"xxx/Database/Identifier"
	"xxx/Database/Journal"
	"xxx/Database/Mysql"
	"xxx/Database/Settings"
	"xxx/Errorz"
	myStrings "xxx/Strings"
)

const ErrorReporter = "Database"

const MsgConnectionEstablished = "Database Connection has been established."
const MsgConnectionClosed = "Database Connection has been closed."

const ErrFormatBadIdentifier = "Bad Identifier: %s"
const ErrFormatBadSystemIdentifier = "Bad System Identifier: %s"
const ErrMethodOwnerDoesNotExist = "Method Owner does not exist"
const ErrFormatNullableFlag = "'Nullable' Flag has unsupported Value '%v'"
const ErrFormatMultipleObjectPropertyValues = "Multiple Object Property Values."
const ErrFormatReturnColumnCount = "Number of Columns in the Database Response " +
	"is bad. Column Count expected: '%v', received: '%v'."
const ErrFormatReturnColumnType = "Column Type in the Database Response " +
	"is bad. Column Index: '%v'."

const QueryFormatShowTablesLike = "SHOW TABLES LIKE '%s';"
const QueryFormatShowColumnsFrom = "SHOW COLUMNS FROM `%s`;"

const TypePostfixNotNull = " " + "NOT NULL"

type Database struct {
	class         class.Class
	configuration configuration.Configuration
	connection    connection.Connection
	journal       journal.Journal
}

type DatabaseTableColumnSettings struct {
	Name string
	Type string
}

// Adds a Record into the Journal.
func (this *Database) AddJournalRecord(
	record string,
) error {

	var err error

	// Fool's Check.
	if this == nil {
		err = errors.New(ErrMethodOwnerDoesNotExist)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	return this.addJournalRecord(record)
}

// Adds a Record into the Journal.
func (this *Database) addJournalRecord(
	record string,
) error {

	var err error

	if this.configuration.GetJournalIsOn() {
		err = this.journal.AddRecord(record)
		if err != nil {
			return err
		}
	}

	return nil
}

// Configures the Database.
func (this *Database) Configure(
	settings settings.Settings,
) error {

	var err error

	// Fool's Check.
	if this == nil {
		err = errors.New(ErrMethodOwnerDoesNotExist)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Check the Database Class Configuration State.
	if this.class.IsConfigured() != true {
		err = fmt.Errorf(
			class.ErrNoClass,
		)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Configure the DSN.
	err = this.configuration.UpdateDsn(
		settings.HostName,
		settings.PortNumber,
		settings.UserName,
		settings.Password,
		settings.Database,
		settings.Parameters,
	)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	err = this.configuration.SetJournalIsOn(settings.Journal)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	err = this.configuration.Finalize()
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	return nil
}

// Connects to the Database.
func (this *Database) Connect() error {

	var err error
	var errDouble error
	var errJournal error

	// Fool's Check.
	if this == nil {
		err = errors.New(ErrMethodOwnerDoesNotExist)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Connect.
	err = this.connection.Establish(
		this.class.GetValue(),
		this.configuration,
	)
	if err != nil {
		errJournal = this.addJournalRecord(err.Error())
		if errJournal != nil {
			errDouble = errorz.Combine(err, errJournal)
			errDouble = errorz.Report(ErrorReporter, errDouble)
			return errDouble
		}
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Report.
	errJournal = this.addJournalRecord(MsgConnectionEstablished)
	if errJournal != nil {
		errJournal = errorz.Report(ErrorReporter, errJournal)
		return errJournal
	}

	return nil
}

// Disconnects from the Database.
func (this *Database) Disconnect() error {

	var err error
	var errDouble error
	var errJournal error

	// Fool's Check.
	if this == nil {
		err = errors.New(ErrMethodOwnerDoesNotExist)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Disconnect.
	err = this.connection.Close(
		this.class.GetValue(),
	)
	if err != nil {
		errJournal = this.addJournalRecord(err.Error())
		if errJournal != nil {
			errDouble = errorz.Combine(err, errJournal)
			errDouble = errorz.Report(ErrorReporter, errDouble)
			return errDouble
		}
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	errJournal = this.addJournalRecord(MsgConnectionClosed)
	if errJournal != nil {
		errJournal = errorz.Report(ErrorReporter, errJournal)
		return errJournal
	}

	return nil
}

// Returns the 'configured' State.
func (this Database) IsConfigured() bool {

	if this.class.IsConfigured() != true {
		return false
	}

	if this.configuration.IsSet() != true {
		return false
	}

	return true
}

// Returns the 'connected' State.
func (this Database) IsConnected() bool {

	return this.connection.IsEstablished()
}

// Executes a Query which returns Classes.
func (this Database) QueryClasses(
	query string,
) ([]databaseCollection.DatabaseClass, error) {

	var databaseClass value.Value
	var err error
	var result []databaseCollection.DatabaseClass

	databaseClass = this.class.GetValue()
	switch databaseClass {

	case value.Mysql:

		// Read Classes
		result, err = this.queryClassesMysql(
			query,
		)
		if err != nil {
			err = errorz.Report(ErrorReporter, err)
			return result, err
		}

	default:

		// Database Class in unsupported.
		err = fmt.Errorf(
			class.ErrFormatClassValueAliasUnknown,
			databaseClass,
		)
		err = errorz.Report(ErrorReporter, err)
		return result, err
	}

	// Analyze the Results.
	return result, nil
}

// Executes a Query which returns Classes.
// MySQL Method Variant.
func (this Database) queryClassesMysql(
	query string,
) ([]databaseCollection.DatabaseClass, error) {

	var aClass databaseCollection.DatabaseClass
	var err error
	var errDouble error
	var errJournal error
	var result []databaseCollection.DatabaseClass
	var rows *sql.Rows

	// Execute a Query.
	rows, err = this.connection.Query(
		this.class.GetValue(),
		query,
	)
	if err != nil {
		errJournal = this.addJournalRecord(err.Error())
		if errJournal != nil {
			errDouble = errorz.Combine(err, errJournal)
			return result, errDouble
		}
		return result, err
	}

	// Get Response.
	result = make([]databaseCollection.DatabaseClass, 0)
	for rows.Next() {
		err = rows.Scan(
			&aClass.Id,
			&aClass.Name,
		)
		if err != nil {
			return result, err
		}
		result = append(result, aClass)
	}

	// Finalize the Execution.
	err = rows.Close()
	if err != nil {
		errJournal = this.addJournalRecord(err.Error())
		if errJournal != nil {
			errDouble = errorz.Combine(err, errJournal)
			return result, errDouble
		}
		return result, err
	}

	// Analyze the Results.
	return result, nil
}

// Executes a Query which returns a Class Object Property.
func (this Database) QueryClassObjectProperty(
	query string,
) (databaseCollection.DatabaseClassObjectProperty, error) {

	var databaseClass value.Value
	var err error
	var result databaseCollection.DatabaseClassObjectProperty

	databaseClass = this.class.GetValue()
	switch databaseClass {

	case value.Mysql:

		// Read Classes
		result, err = this.queryClassObjectPropertyMysql(query)
		if err != nil {
			err = errorz.Report(ErrorReporter, err)
			return result, err
		}

	default:

		// Database Class in unsupported.
		err = fmt.Errorf(
			class.ErrFormatClassValueAliasUnknown,
			databaseClass,
		)
		err = errorz.Report(ErrorReporter, err)
		return result, err
	}

	// Analyze the Results.
	return result, nil
}

// Executes a Query which returns a Class Object Property.
// MySQL Method Variant.
func (this Database) queryClassObjectPropertyMysql(
	query string,
) (databaseCollection.DatabaseClassObjectProperty, error) {

	// Query Parameters hard-coded in the SQL File.
	const ColumnCount = 4
	const ColumnIdxValue = 4 - 1

	var aClassObjectProperty databaseCollection.DatabaseClassObjectProperty
	var columnTypeValue sql.ColumnType
	var err error
	var errDouble error
	var errJournal error
	var result databaseCollection.DatabaseClassObjectProperty
	var results []databaseCollection.DatabaseClassObjectProperty
	var rows *sql.Rows
	var rowsColumnTypes []*sql.ColumnType

	// Execute a Query.
	rows, err = this.connection.Query(
		this.class.GetValue(),
		query,
	)
	if err != nil {
		errJournal = this.addJournalRecord(err.Error())
		if errJournal != nil {
			errDouble = errorz.Combine(err, errJournal)
			return result, errDouble
		}
		return result, err
	}

	// Get Response.
	results = make([]databaseCollection.DatabaseClassObjectProperty, 0)
	for rows.Next() {

		// Scan the Column Values.
		err = rows.Scan(
			&aClassObjectProperty.PropertyId,
			&aClassObjectProperty.ObjectId,
			&aClassObjectProperty.ClassId,
			&aClassObjectProperty.Value,
		)
		if err != nil {
			return result, err
		}

		// Add the raw Database Column Type (as Text) of the 'Value' Column to
		// the Collection Class Object Property. It will be later used to parse
		// the received Value into something more usable.
		rowsColumnTypes, err = rows.ColumnTypes()
		if err != nil {
			return result, err
		}
		if len(rowsColumnTypes) != ColumnCount {
			err = fmt.Errorf(
				ErrFormatReturnColumnCount,
				ColumnCount,
				len(rowsColumnTypes),
			)
			return result, err
		}
		if rowsColumnTypes[ColumnIdxValue] == nil {
			err = fmt.Errorf(
				ErrFormatReturnColumnType,
				ColumnIdxValue,
			)
			return result, err
		}
		columnTypeValue = *(rowsColumnTypes[ColumnIdxValue])
		aClassObjectProperty.ValueDbTypeFromDriver =
			columnTypeValue.DatabaseTypeName()

		// Save the Result.
		results = append(results, aClassObjectProperty)
	}
	if len(results) > 1 {
		err = fmt.Errorf(
			ErrFormatMultipleObjectPropertyValues,
		)
		return result, err
	}
	if len(results) == 1 {
		result = results[0]
		result.ValueIsSet = true
	} else {
		result.ValueIsSet = false
	}

	// Finalize the Execution.
	err = rows.Close()
	if err != nil {
		errJournal = this.addJournalRecord(err.Error())
		if errJournal != nil {
			errDouble = errorz.Combine(err, errJournal)
			return result, errDouble
		}
		return result, err
	}

	// Analyze the Results.
	return result, nil
}

// Executes a Query which returns Class Objects.
func (this Database) QueryClassObjects(
	query string,
) ([]databaseCollection.DatabaseClassObject, error) {

	var databaseClass value.Value
	var err error
	var result []databaseCollection.DatabaseClassObject

	databaseClass = this.class.GetValue()
	switch databaseClass {

	case value.Mysql:

		// Read Classes
		result, err = this.queryClassObjectsMysql(
			query,
		)
		if err != nil {
			err = errorz.Report(ErrorReporter, err)
			return result, err
		}

	default:

		// Database Class in unsupported.
		err = fmt.Errorf(
			class.ErrFormatClassValueAliasUnknown,
			databaseClass,
		)
		err = errorz.Report(ErrorReporter, err)
		return result, err
	}

	// Analyze the Results.
	return result, nil
}

// Executes a Query which returns Class Objects.
// MySQL Method Variant.
func (this Database) queryClassObjectsMysql(
	query string,
) ([]databaseCollection.DatabaseClassObject, error) {

	var aClassObject databaseCollection.DatabaseClassObject
	var err error
	var errDouble error
	var errJournal error
	var result []databaseCollection.DatabaseClassObject
	var rows *sql.Rows

	// Execute a Query.
	rows, err = this.connection.Query(
		this.class.GetValue(),
		query,
	)
	if err != nil {
		errJournal = this.addJournalRecord(err.Error())
		if errJournal != nil {
			errDouble = errorz.Combine(err, errJournal)
			return result, errDouble
		}
		return result, err
	}

	// Get Response.
	result = make([]databaseCollection.DatabaseClassObject, 0)
	for rows.Next() {
		err = rows.Scan(
			&aClassObject.Id,
			&aClassObject.ClassId,
		)
		if err != nil {
			return result, err
		}
		result = append(result, aClassObject)
	}

	// Finalize the Execution.
	err = rows.Close()
	if err != nil {
		errJournal = this.addJournalRecord(err.Error())
		if errJournal != nil {
			errDouble = errorz.Combine(err, errJournal)
			return result, errDouble
		}
		return result, err
	}

	// Analyze the Results.
	return result, nil
}

// Executes a Query which returns Class Properties.
func (this Database) QueryClassProperties(
	query string,
) ([]databaseCollection.DatabaseClassProperty, error) {

	var databaseClass value.Value
	var err error
	var result []databaseCollection.DatabaseClassProperty

	databaseClass = this.class.GetValue()
	switch databaseClass {

	case value.Mysql:

		// Read Classes
		result, err = this.queryClassPropertiesMysql(
			query,
		)
		if err != nil {
			err = errorz.Report(ErrorReporter, err)
			return result, err
		}

	default:

		// Database Class in unsupported.
		err = fmt.Errorf(
			class.ErrFormatClassValueAliasUnknown,
			databaseClass,
		)
		err = errorz.Report(ErrorReporter, err)
		return result, err
	}

	// Analyze the Results.
	return result, nil
}

// Executes a Query which returns Class Properties.
// MySQL Method Variant.
func (this Database) queryClassPropertiesMysql(
	query string,
) ([]databaseCollection.DatabaseClassProperty, error) {

	var aClassProperty databaseCollection.DatabaseClassProperty
	var err error
	var errDouble error
	var errJournal error
	var result []databaseCollection.DatabaseClassProperty
	var rows *sql.Rows

	// Execute a Query.
	rows, err = this.connection.Query(
		this.class.GetValue(),
		query,
	)
	if err != nil {
		errJournal = this.addJournalRecord(err.Error())
		if errJournal != nil {
			errDouble = errorz.Combine(err, errJournal)
			return result, errDouble
		}
		return result, err
	}

	// Get Response.
	result = make([]databaseCollection.DatabaseClassProperty, 0)
	for rows.Next() {
		err = rows.Scan(
			&aClassProperty.Id,
			&aClassProperty.Name,
			&aClassProperty.Description,
			&aClassProperty.DbType,
			&aClassProperty.ClassId,
		)
		if err != nil {
			return result, err
		}
		result = append(result, aClassProperty)
	}

	// Finalize the Execution.
	err = rows.Close()
	if err != nil {
		errJournal = this.addJournalRecord(err.Error())
		if errJournal != nil {
			errDouble = errorz.Combine(err, errJournal)
			return result, errDouble
		}
		return result, err
	}

	// Analyze the Results.
	return result, nil
}

// Shows Information about Columns in the Table.
func (this Database) ShowColumns(
	tableName string,
) ([]DatabaseTableColumnSettings, error) {

	var databaseClass value.Value
	var err error
	var result []DatabaseTableColumnSettings

	databaseClass = this.class.GetValue()
	switch databaseClass {

	case value.Mysql:

		// Get the List of Columns' Settings.
		result, err = this.showColumnsMysql(
			tableName,
		)
		if err != nil {
			err = errorz.Report(ErrorReporter, err)
			return result, err
		}

	default:

		// Database Class in unsupported.
		err = fmt.Errorf(
			class.ErrFormatClassValueAliasUnknown,
			databaseClass,
		)
		err = errorz.Report(ErrorReporter, err)
		return result, err
	}

	// Post-Processing.
	// Change Letter Case.
	for i := 0; i < len(result); i++ {
		result[i].Name = strings.ToLower(result[i].Name)
		result[i].Type = strings.ToUpper(result[i].Type)
	}

	return result, nil
}

// Shows Information about Columns in the Table.
// MySQL Method Variant.
func (this Database) showColumnsMysql(
	tableName string,
) ([]DatabaseTableColumnSettings, error) {

	const NullableYes = "YES"
	const NullableNo = "NO"

	var args []interface{}
	var columnsCount int
	var columnName string
	var columnHasNotnullAttribute bool
	var columnSettings []mysql.ColumnSettingsMysql
	var columnType string
	var databaseClass value.Value
	var err error
	var errDouble error
	var errJournal error
	var query string
	var result []DatabaseTableColumnSettings
	var rows *sql.Rows

	databaseClass = this.class.GetValue()

	// Create and execute the Query.
	query = fmt.Sprintf(QueryFormatShowColumnsFrom, tableName)
	rows, err = this.connection.Query(databaseClass, query, args...)
	if err != nil {
		return result, err
	}

	// Verify the Database Reply.
	err = mysql.VerifyOutputOfShowColumnsCmd(rows)
	if err != nil {
		return result, err
	}

	// Parse the Database Reply.
	columnSettings, err = mysql.ParseOutputOfShowColumnsCmd(rows)
	if err != nil {
		return result, err
	}

	// Finalize the Execution.
	err = rows.Close()
	if err != nil {
		errJournal = this.addJournalRecord(err.Error())
		if errJournal != nil {
			errDouble = errorz.Combine(err, errJournal)
			return result, errDouble
		}
		return result, err
	}

	// Convert Column Settings into a common Format.
	columnsCount = len(columnSettings)
	result = make([]DatabaseTableColumnSettings, columnsCount)
	for i := 0; i < columnsCount; i++ {

		columnName = columnSettings[i].Name
		columnType = columnSettings[i].Type

		switch columnSettings[i].Null {

		case NullableYes:
			columnHasNotnullAttribute = false

		case NullableNo:
			columnHasNotnullAttribute = true

		default:
			err = fmt.Errorf(
				ErrFormatNullableFlag,
				columnSettings[i].Null,
			)
			return result, err
		}
		if columnHasNotnullAttribute {
			columnType = columnType + TypePostfixNotNull
		}

		result[i] = DatabaseTableColumnSettings{
			Name: columnName,
			Type: columnType,
		}
	}

	return result, nil
}

// Executes a simple Query (which returns no Results).
func (this Database) SimpleQuery(
	query string,
) error {

	var databaseClass value.Value
	var err error

	databaseClass = this.class.GetValue()
	switch databaseClass {

	case value.Mysql:

		// Execute a simple Query.
		err = this.simpleQueryMysql(query)
		if err != nil {
			err = errorz.Report(ErrorReporter, err)
			return err
		}

	default:

		// Database Class in unsupported.
		err = fmt.Errorf(
			class.ErrFormatClassValueAliasUnknown,
			databaseClass,
		)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Analyze the Results.
	return nil
}

// Executes a simple Query (which returns no Results).
// MySQL Method Variant.
func (this Database) simpleQueryMysql(
	query string,
) error {

	var err error
	var errDouble error
	var errJournal error
	var rows *sql.Rows

	// Execute a Query.
	rows, err = this.connection.Query(
		this.class.GetValue(),
		query,
	)
	if err != nil {
		errJournal = this.addJournalRecord(err.Error())
		if errJournal != nil {
			errDouble = errorz.Combine(err, errJournal)
			return errDouble
		}
		return err
	}

	// Finalize the Execution.
	err = rows.Close()
	if err != nil {
		errJournal = this.addJournalRecord(err.Error())
		if errJournal != nil {
			errDouble = errorz.Combine(err, errJournal)
			return errDouble
		}
		return err
	}

	return nil
}

// Searches for a system Table.
// Table Name is checked.
func (this Database) SystemTableExists(
	tableName string,
) (bool, error) {

	var databaseClass value.Value
	var err error
	var id identifier.Identifier
	var tableExists bool

	// Check Table Name Syntax.
	id.Name = tableName
	if id.IsGood(true) != true {
		err = fmt.Errorf(
			ErrFormatBadSystemIdentifier,
			tableName,
		)
		if err != nil {
			err = errorz.Report(ErrorReporter, err)
			return false, err
		}
	}

	databaseClass = this.class.GetValue()
	switch databaseClass {

	case value.Mysql:

		// Check Table's Existence.
		tableExists, err = this.tableExistsMysql(tableName)
		if err != nil {
			err = errorz.Report(ErrorReporter, err)
			return false, err
		}

	default:

		// Database Class in unsupported.
		err = fmt.Errorf(
			class.ErrFormatClassValueAliasUnknown,
			databaseClass,
		)
		err = errorz.Report(ErrorReporter, err)
		return false, err
	}

	// Analyze the Results.
	if tableExists {
		return true, nil
	}

	return false, nil
}

// Searches for a normal Table.
// Table Name is checked.
func (this Database) TableExists(
	tableName string,
) (bool, error) {

	var databaseClass value.Value
	var err error
	var id identifier.Identifier
	var tableExists bool

	// Check Table Name Syntax.
	id.Name = tableName
	if id.IsGood(false) != true {
		err = fmt.Errorf(
			ErrFormatBadIdentifier,
			tableName,
		)
		if err != nil {
			err = errorz.Report(ErrorReporter, err)
			return false, err
		}
	}

	databaseClass = this.class.GetValue()
	switch databaseClass {

	case value.Mysql:

		// Check Table's Existence.
		tableExists, err = this.tableExistsMysql(tableName)
		if err != nil {
			err = errorz.Report(ErrorReporter, err)
			return false, err
		}

	default:

		// Database Class in unsupported.
		err = fmt.Errorf(
			class.ErrFormatClassValueAliasUnknown,
			databaseClass,
		)
		err = errorz.Report(ErrorReporter, err)
		return false, err
	}

	// Analyze the Results.
	if tableExists {
		return true, nil
	}

	return false, nil
}

// Searches for a Table.
// Table Name is not checked.
// MySQL Method Variant.
func (this Database) tableExistsMysql(
	tableName string,
) (bool, error) {

	var args []interface{}
	var databaseClass value.Value
	var err error
	var errDouble error
	var errJournal error
	var query string
	var rows *sql.Rows
	var tableExists bool
	var tableNameFound string
	var tableNamesFound []string

	// Preparation.
	databaseClass = this.class.GetValue()

	// Create and execute the Query.
	query = fmt.Sprintf(QueryFormatShowTablesLike, tableName)
	rows, err = this.connection.Query(databaseClass, query, args...)
	if err != nil {
		return false, err
	}

	// Get Response.
	tableNamesFound = make([]string, 0)
	for rows.Next() {
		err = rows.Scan(&tableNameFound)
		if err != nil {
			return false, err
		}
		tableNamesFound = append(tableNamesFound, tableNameFound)
	}
	tableExists = myStrings.ArrayContainsString(tableNamesFound, tableName)

	// Close the Query.
	err = rows.Close()
	if err != nil {
		errJournal = this.journal.AddRecord(err.Error())
		if errJournal != nil {
			errDouble = errorz.Combine(err, errJournal)
			return false, errDouble
		}
		return false, err
	}

	// Analyze the Results.
	if tableExists {
		return true, nil
	}

	return false, nil
}

// Creates a new Database with the specified Class.
func New(
	dbClass string,
) (*Database, error) {

	var err error
	var result *Database

	result = new(Database)

	// Set the Database Class.
	err = result.class.Set(dbClass)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return nil, err
	}

	return result, nil
}
