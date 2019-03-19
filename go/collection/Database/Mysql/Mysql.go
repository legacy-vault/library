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

// Mysql.go.

package mysql

import (
	"database/sql"
	"fmt"

	"xxx/Errorz"
)

// MySQL Verification.

// The Settings of MySQL Commands Output for the MySQL Community Server v8.0.

const ErrorReporter = "Mysql"

const OutputColumnsCountShowColumns = 6

type ColumnSettingsMysql struct {
	Name    string
	Type    string
	Null    string
	Key     string
	Default *string //TODO:?
	Extra   string
}

const DbColumnTypeChar = "CHAR"
const DbColumnTypeText = "TEXT"
const DbColumnTypeVarchar = "VARCHAR"

const OutputColumnNameShowColumnsA = "Field"
const OutputColumnNameShowColumnsB = "Type"
const OutputColumnNameShowColumnsC = "Null"
const OutputColumnNameShowColumnsD = "Key"
const OutputColumnNameShowColumnsE = "Default"
const OutputColumnNameShowColumnsF = "Extra"

const ErrFormatShowColumnsCount = "'SHOW COLUMNS' Command Error. " +
	"Output Columns Count Mismatch: %v / %v."
const ErrFormatShowColumnsNames = "'SHOW COLUMNS' Command Error. " +
	"Output Column Names Mismatch. Expected '%v', received '%v'."
const ErrFormatShowColumnsTypes = "'SHOW COLUMNS' Command Error. " +
	"Output Column Types Mismatch. Expected '%v', received '%v'."

// Scans and parses the Output of the 'SHOW COLUMNS' SQL Command.
func ParseOutputOfShowColumnsCmd(
	rows *sql.Rows,
) ([]ColumnSettingsMysql, error) {

	var err error
	var result []ColumnSettingsMysql
	var row ColumnSettingsMysql

	result = make([]ColumnSettingsMysql, 0)

	for rows.Next() {
		err = rows.Scan(
			&row.Name,
			&row.Type,
			&row.Null,
			&row.Key,
			&row.Default,
			&row.Extra,
		)
		if err != nil {
			err = errorz.Report(ErrorReporter, err)
			return result, err
		}
		result = append(result, row)
	}

	return result, nil
}

// Verifies the Output of the 'SHOW COLUMNS' SQL Command.
// Checks the Number and Names of Output Columns.
func VerifyOutputOfShowColumnsCmd(
	rows *sql.Rows,
) error {

	var err error
	var outputColumnNames []string
	var outputColumnNamesExpected []string
	var outputColumnsCount int
	var outputColumnTypes []*sql.ColumnType
	var outputColumnTypesExpected []string

	// Preparations.
	outputColumnNamesExpected = make([]string, OutputColumnsCountShowColumns)
	outputColumnNamesExpected[0] = OutputColumnNameShowColumnsA
	outputColumnNamesExpected[1] = OutputColumnNameShowColumnsB
	outputColumnNamesExpected[2] = OutputColumnNameShowColumnsC
	outputColumnNamesExpected[3] = OutputColumnNameShowColumnsD
	outputColumnNamesExpected[4] = OutputColumnNameShowColumnsE
	outputColumnNamesExpected[5] = OutputColumnNameShowColumnsF
	//
	outputColumnTypesExpected =
		make([]string, OutputColumnsCountShowColumns)
	outputColumnTypesExpected[0] = DbColumnTypeVarchar
	outputColumnTypesExpected[1] = DbColumnTypeText
	outputColumnTypesExpected[2] = DbColumnTypeVarchar
	outputColumnTypesExpected[3] = DbColumnTypeChar
	outputColumnTypesExpected[4] = DbColumnTypeText
	outputColumnTypesExpected[5] = DbColumnTypeVarchar
	//
	outputColumnNames, err = rows.Columns()
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}
	outputColumnTypes, err = rows.ColumnTypes()
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// 1. Count the Output Columns' Number.
	outputColumnsCount = len(outputColumnNames)
	if outputColumnsCount != OutputColumnsCountShowColumns {
		err = fmt.Errorf(
			ErrFormatShowColumnsCount,
			outputColumnsCount,
			OutputColumnsCountShowColumns,
		)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// 2. Check the Output Columns' Names.
	for i := 0; i < outputColumnsCount; i++ {
		if outputColumnNames[i] != outputColumnNamesExpected[i] {
			err = fmt.Errorf(
				ErrFormatShowColumnsNames,
				outputColumnNamesExpected[i],
				outputColumnNames[i],
			)
			err = errorz.Report(ErrorReporter, err)
			return err
		}
	}

	// 3. Check the Output Columns' Types.
	for i := 0; i < outputColumnsCount; i++ {
		if outputColumnTypes[i].DatabaseTypeName() !=
			outputColumnTypesExpected[i] {
			err = fmt.Errorf(
				ErrFormatShowColumnsTypes,
				outputColumnTypesExpected[i],
				outputColumnTypes[i].DatabaseTypeName(),
			)
			err = errorz.Report(ErrorReporter, err)
			return err
		}
	}

	return nil
}
