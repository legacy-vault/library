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

// table.go.

package storage

// Storage Methods for Collection Tables:
//
//	* Table Name Composing;
//	* Table Checks.

import (
	"fmt"
	"strings"

	"xxx/Database"
	"xxx/Uc/Collection/Class"
	"xxx/Uc/Collection/Class/Property"
	"xxx/Uc/Storage/Settings/Tables"
)

// Checks all the Tables of the Collection in the Database.
func (this Storage) CheckDbCollection(
	classes map[uint]class.Class,
) error {

	return this.checkDbCollection(classes)
}

// Checks all the Tables of the Collection in the Database.
func (this Storage) checkDbCollection(
	classes map[uint]class.Class,
) error {

	var err error

	// Check the Classes Table of the Database.
	err = this.checkDbTableClasses()
	if err != nil {
		return err
	}

	// Check the Class Tables of the Database.
	err = this.checkDbTablesClass(classes)
	if err != nil {
		return err
	}

	// Check the Class Properties Tables of the Database.
	err = this.checkDbTablesClassProperties(classes)
	if err != nil {
		return err
	}

	// Check the Property Tables of each Class in the Database.
	err = this.checkDbTablesClassesProperty(classes)
	if err != nil {
		return err
	}

	return nil
}

// Checks the Class Table of the Database.
func (this Storage) checkDbTableClass(
	aClass class.Class,
) error {

	var classTableName string
	var columnsSettings []database.DatabaseTableColumnSettings
	var err error
	var ok bool

	// Table Name.
	classTableName = strings.ToLower(aClass.GetName())

	// Check Table's Existence.
	ok, err = this.TableExists(classTableName)
	if err != nil {
		return err
	}
	if ok == false {
		err = fmt.Errorf(
			ErrFormatMissingTable,
			classTableName,
		)
		return err
	}

	// Check Table's Columns.
	columnsSettings, err = this.ShowColumns(classTableName)
	if err != nil {
		return err
	}
	err = this.verifyColumnsForTableClass(columnsSettings)
	if err != nil {
		return err
	}

	return nil
}

// Checks the Classes Table of the Database.
func (this Storage) checkDbTableClasses() error {

	var classesTableName string
	var columnsSettings []database.DatabaseTableColumnSettings
	var err error
	var ok bool

	// Preparations.
	classesTableName = strings.ToLower(this.settings.Table().Classes())

	// Check Table's Existence.
	ok, err = this.TableExists(classesTableName)
	if err != nil {
		return err
	}
	if ok == false {
		err = fmt.Errorf(
			ErrFormatMissingTable,
			classesTableName,
		)
		return err
	}

	// Check Table's Columns.
	columnsSettings, err = this.ShowColumns(classesTableName)
	if err != nil {
		return err
	}
	err = this.verifyColumnsForTableClasses(columnsSettings)
	if err != nil {
		return err
	}

	return nil
}

// Checks the Properties Table of a Class of the Database.
func (this Storage) checkDbTableClassProperties(
	aClass class.Class,
) error {

	var className string
	var classPropertiesTableName string
	var columnsSettings []database.DatabaseTableColumnSettings
	var err error
	var ok bool

	// Table Name.
	className = strings.ToLower(aClass.GetName())
	classPropertiesTableName =
		this.ComposeTableNameClassProperties(className)

	// Check Table's Existence.
	ok, err = this.TableExists(classPropertiesTableName)
	if err != nil {
		return err
	}
	if ok == false {
		err = fmt.Errorf(
			ErrFormatMissingTable,
			classPropertiesTableName,
		)
		return err
	}

	// Check Table's Columns.
	columnsSettings, err = this.ShowColumns(classPropertiesTableName)
	if err != nil {
		return err
	}
	err = this.verifyColumnsForTableClassProperties(columnsSettings)
	if err != nil {
		return err
	}

	return nil
}

// Checks the Class Tables of the Database.
func (this Storage) checkDbTablesClass(
	classes map[uint]class.Class,
) error {

	var err error

	for _, aClass := range classes {

		err = this.checkDbTableClass(aClass)
		if err != nil {
			return err
		}
	}

	return nil
}

// Checks the Property Tables of each Class in the Database.
func (this Storage) checkDbTablesClassesProperty(
	classes map[uint]class.Class,
) error {

	var err error

	for _, aClass := range classes {

		err = this.checkDbTablesClassProperty(aClass)
		if err != nil {
			return err
		}
	}

	return nil
}

// Checks the Properties Table of each Class of the Database.
func (this Storage) checkDbTablesClassProperties(
	classes map[uint]class.Class,
) error {

	var err error

	for _, aClass := range classes {

		err = this.checkDbTableClassProperties(aClass)
		if err != nil {
			return err
		}
	}

	return nil
}

// Checks the Property Tables of a Class in the Database.
func (this Storage) checkDbTablesClassProperty(
	aClass class.Class,
) error {

	var className string
	var classProperties map[uint]property.Property
	var classPropertyName string
	var classPropertyTableName string
	var columnsSettings []database.DatabaseTableColumnSettings
	var dbColumnTypePropertyValue string
	var err error
	var ok bool

	className = strings.ToLower(aClass.GetName())
	classProperties = aClass.GetProperties()
	for _, aClassProperty := range classProperties {

		// Preparations.
		classPropertyName = aClassProperty.GetName()
		classPropertyTableName = this.ComposeTableNameClassProperty(
			className,
			classPropertyName,
		)
		dbColumnTypePropertyValue = aClassProperty.GetKind().DbType

		// Check Table's Existence.
		ok, err = this.TableExists(classPropertyTableName)
		if err != nil {
			return err
		}
		if ok == false {
			err = fmt.Errorf(
				ErrFormatMissingTable,
				classPropertyTableName,
			)
			return err
		}

		// Check Table's Columns.
		columnsSettings, err = this.ShowColumns(classPropertyTableName)
		if err != nil {
			return err
		}
		err = this.verifyColumnsForTableClassProperty(
			columnsSettings,
			dbColumnTypePropertyValue,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

// Composes a Name for a Class Table.
func (this Storage) ComposeTableNameClass(
	className string,
) string {

	var result string

	result = className

	return result
}

// Composes a Name for a Class Properties Table.
func (this Storage) ComposeTableNameClassProperties(
	className string,
) string {

	var result string

	result = className +
		tables.NameSeparator +
		this.settings.Table().PropertyTypes()

	return result
}

// Composes a Name for a Class Property Table.
func (this Storage) ComposeTableNameClassProperty(
	className string,
	propertyName string,
) string {

	var result string

	result = className +
		tables.NameSeparator +
		this.settings.Table().PropertyPrefix() +
		tables.NameSeparator +
		propertyName

	return result
}
