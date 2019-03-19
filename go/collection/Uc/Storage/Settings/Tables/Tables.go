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

// Tables.go.

package tables

// Collection Settings Tables.

import (
	"errors"
	"fmt"

	"xxx/Common"
	"xxx/Database/Identifier"
	"xxx/Errorz"
	"xxx/Utf8"
)

// There are several Types of Tables in the Collection:
//	1.	System Table with Classes List, usually it is '_classes'.
//	2.	Tables for each Class, usually named as 'planet' or 'image'.
//	3.	Table with Property Types for a Class, like 'image_propertytypes'.
//	4.	Table with Property Values for a Property Type,
//		e.g.: 'image_property_size'.

const ErrorReporter = "Tables"

// Separator used in complex Table Names like 'image_property_size'.
const NameSeparator = "_"
const SystemNamePrefix = '_'

// Table Name.
const NameClassesDefault = string(SystemNamePrefix) + "classes"
const NamePostfixPropertyTypesDefault = "properties"
const NamePrefixPropertyDefault = "property"

// Errors.
const ErrFormatSystemTableNameBad = "Bad System Table Name: '%v'"
const ErrFormatTableNameBad = "Bad Table Name: '%v'"
const ErrFormatTableNamePartBad = "Bad Table Name Part: '%v'"

// Database Table Settings: Names.
type Tables struct {

	// Name of the main Table with Classes List.
	// The first Symbol must be '_', other Symbols are small ASCII Letters.
	// Usually, '_classes'.
	classes string

	// Postfix used in Tables with Property Types for a Class.
	// Only small ASCII Letters are allowed.
	// Usually, 'properties' or 'attributes'.
	propertyTypes string

	// Prefix used in Tables with Property Values for a Property Type.
	// Only small ASCII Letters are allowed.
	// Usually, 'property' or 'attribute'.
	propertyPrefix string
}

// Checks the Table Settings.
func (this *Tables) Check() error {

	var err error

	// 1. Classes Table Name.
	err = SystemTableNameCheck(this.classes)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// 2. Property Type Table Name Postfix.
	err = TableNamePartCheck(this.propertyTypes)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// 3. Property Table Name Prefix.
	err = TableNamePartCheck(this.propertyPrefix)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	return nil
}

// Returns the 'classes' Field.
func (this Tables) Classes() string {

	return this.classes
}

// Returns the 'propertyPrefix' Field.
func (this Tables) PropertyPrefix() string {

	return this.propertyPrefix
}

// Returns the 'propertyTypes' Field.
func (this Tables) PropertyTypes() string {

	return this.propertyTypes
}

// Sets all the Fields.
func (this *Tables) SetAll(
	classesTableName string,
	propertyTypesTableNamePostfix string,
	propertyTablesNamePrefix string,
) error {

	var err error

	// Fool's Check.
	if this == nil {
		err = errors.New(common.ErrMethodOwnerDoesNotExist)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// 1. Classes Table Name.
	if len(classesTableName) == 0 {
		this.classes = NameClassesDefault
	} else {
		err = SystemTableNameCheck(classesTableName)
		if err != nil {
			err = errorz.Report(ErrorReporter, err)
			return err
		}
		this.classes = classesTableName
	}

	// 2. Property Type Table Name Postfix.
	if len(propertyTypesTableNamePostfix) == 0 {
		this.propertyTypes = NamePostfixPropertyTypesDefault
	} else {
		err = TableNamePartCheck(propertyTypesTableNamePostfix)
		if err != nil {
			err = errorz.Report(ErrorReporter, err)
			return err
		}
		this.propertyTypes = propertyTypesTableNamePostfix
	}

	// 3. Property Table Name Prefix.
	if len(propertyTablesNamePrefix) == 0 {
		this.propertyPrefix = NamePrefixPropertyDefault
	} else {
		err = TableNamePartCheck(propertyTablesNamePrefix)
		if err != nil {
			err = errorz.Report(ErrorReporter, err)
			return err
		}
		this.propertyPrefix = propertyTablesNamePrefix
	}

	return nil
}

// Checks System Table Name Syntax.
func SystemTableNameCheck(
	systemTableName string,
) error {

	var err error
	var id identifier.Identifier
	var ok bool

	id.Name = systemTableName
	ok = id.IsGood(true)
	if !ok {
		err = fmt.Errorf(
			ErrFormatSystemTableNameBad,
			systemTableName,
		)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	return nil
}

// Checks ordinary Table Name Syntax.
func TableNameCheck(
	tableName string,
) error {

	var err error
	var id identifier.Identifier
	var ok bool

	id.Name = tableName
	ok = id.IsGood(false)
	if !ok {
		err = fmt.Errorf(
			ErrFormatTableNameBad,
			tableName,
		)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	return nil
}

// Checks a Table Name Part Syntax.
func TableNamePartCheck(
	tableNamePart string,
) error {

	var err error
	var ok bool

	ok = utf8.StringHasOnlyASCIILetterOrNumber(tableNamePart)
	if !ok {
		err = fmt.Errorf(
			ErrFormatTableNamePartBad,
			tableNamePart,
		)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	return nil
}
