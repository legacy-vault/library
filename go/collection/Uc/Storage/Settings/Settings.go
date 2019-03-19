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

// Settings.go.

package settings

// Collection Settings.

import (
	"errors"

	"github.com/legacy-vault/library/go/collection/Common"
	"github.com/legacy-vault/library/go/collection/Database/Settings"
	"github.com/legacy-vault/library/go/collection/Errorz"
	"github.com/legacy-vault/library/go/collection/Uc/Storage/Settings/Paths"
	"github.com/legacy-vault/library/go/collection/Uc/Storage/Settings/SqlTemplates"
	"github.com/legacy-vault/library/go/collection/Uc/Storage/Settings/Tables"
)

const ErrorReporter = "Settings"

type Settings struct {

	// Database Settings.
	database settings.Settings

	// Table Settings.
	table tables.Tables

	// Path Settings.
	path paths.Paths

	// SQL Command Templates.
	sqlCommandTemplate sqlTemplates.SqlTemplates
}

// Returns the 'database' Field.
func (this Settings) Database() settings.Settings {

	return this.database
}

// Returns the 'path' Field.
func (this Settings) Path() paths.Paths {

	return this.path
}

// Returns the 'sqlCommandTemplate' Field.
func (this Settings) SqlCommandTemplate() sqlTemplates.SqlTemplates {

	return this.sqlCommandTemplate
}

// Returns the 'table' Field.
func (this Settings) Table() tables.Tables {

	return this.table
}

// Updates the 'sqlCommandTemplate' Field.
func (this *Settings) UpdateSqlCommandTemplate(
	addClass string,
	addClassObject string,
	addClassObjectProperty string,
	addClassProperty string,
	createClassesTable string,
	createClassTable string,
	createClassPropertiesTable string,
	createClassPropertyTable string,
	readClasses string,
	readClassProperties string,
	readClassObjects string,
	readClassObjectProperty string,
) error {

	var err error

	// Fool's Check.
	if this == nil {
		err = errors.New(common.ErrMethodOwnerDoesNotExist)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	err = this.sqlCommandTemplate.SetAll(
		addClass,
		addClassObject,
		addClassObjectProperty,
		addClassProperty,
		createClassesTable,
		createClassTable,
		createClassPropertiesTable,
		createClassPropertyTable,
		readClasses,
		readClassProperties,
		readClassObjects,
		readClassObjectProperty,
	)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	return nil
}

// Creates the Settings.
// Empty Strings are replaced with the default Values.
func NewSettings(
	databaseSettings settings.Settings,
	classesTableName string,
	propertyTypesTableNamePostfix string,
	propertyTablesNamePrefix string,
	sqlCommandTemplatesPath string,
) (Settings, error) {

	var err error
	var result Settings

	result.database = databaseSettings

	err = result.table.SetAll(
		classesTableName,
		propertyTypesTableNamePostfix,
		propertyTablesNamePrefix,
	)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return result, err
	}

	err = result.path.SetAll(
		sqlCommandTemplatesPath,
	)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return result, err
	}

	return result, nil
}
