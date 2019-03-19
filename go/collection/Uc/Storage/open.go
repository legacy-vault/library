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

// Storage.go.

package storage

// Storage Initialization.

import (
	"errors"
	"io/ioutil"
	"path/filepath"

	"xxx/Database"
	"xxx/Errorz"
	"xxx/Uc/Storage/Settings"
	"xxx/Uc/Storage/Settings/Paths"
	"xxx/Uc/Storage/Settings/Tables"
)

// Sets the Settings
func (this *Storage) Configure(
	settings settings.Settings,
) error {

	var err error
	var tableSettings tables.Tables

	this.settings = settings

	// Check the Settings.
	tableSettings = this.settings.Table()
	err = tableSettings.Check()
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Initialize the Database.
	this.database, err = database.New(
		this.settings.Database().Type,
	)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	err = this.database.Configure(
		this.settings.Database(),
	)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}
	if this.database.IsConfigured() != true {
		err = errors.New(ErrConfiguration)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	return nil
}

// Connects to the Database.
func (this *Storage) Connect() error {

	var err error

	// Connect.
	err = this.database.Connect()
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}
	if this.database.IsConnected() != true {
		err = errors.New(ErrConnection)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	return nil
}

// Gets the Settings
func (this Storage) GetSettings() settings.Settings {

	return this.settings
}

// Loads SQL Command Templates.
func (this *Storage) LoadSqlCommandTemplates() error {

	var addClass string
	var addClassObject string
	var addClassObjectProperty string
	var addClassProperty string
	var ba []byte
	var createClassesTable string
	var createClassPropertiesTable string
	var createClassPropertyTable string
	var createClassTable string
	var err error
	var path string
	var readClasses string
	var readClassProperties string
	var readClassObjects string
	var readClassObjectProperty string

	// 'addCollectionClass.sql' File.
	path = filepath.Join(
		this.settings.Path().GetSqlCommandTemplates(),
		paths.FileSqlCommandTemplateAddClass,
	)
	ba, err = ioutil.ReadFile(path)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}
	addClass = string(ba)

	// 'createDbTableClasses.sql' File.
	path = filepath.Join(
		this.settings.Path().GetSqlCommandTemplates(),
		paths.FileSqlCommandTemplateCreateClassesTable,
	)
	ba, err = ioutil.ReadFile(path)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}
	createClassesTable = string(ba)

	// 'createClassTable.sql' File.
	path = filepath.Join(
		this.settings.Path().GetSqlCommandTemplates(),
		paths.FileSqlCommandTemplateCreateClassTable,
	)
	ba, err = ioutil.ReadFile(path)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}
	createClassTable = string(ba)

	// 'createClassPropertiesTable.sql' File.
	path = filepath.Join(
		this.settings.Path().GetSqlCommandTemplates(),
		paths.FileSqlCommandTemplateCreateClassPropertiesTable,
	)
	ba, err = ioutil.ReadFile(path)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}
	createClassPropertiesTable = string(ba)

	// 'createClassPropertyTable.sql' File.
	path = filepath.Join(
		this.settings.Path().GetSqlCommandTemplates(),
		paths.FileSqlCommandTemplateCreateClassPropertyTable,
	)
	ba, err = ioutil.ReadFile(path)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}
	createClassPropertyTable = string(ba)

	// 'createCollectionClassProperty.sql' File.
	path = filepath.Join(
		this.settings.Path().GetSqlCommandTemplates(),
		paths.FileSqlCommandTemplateAddClassProperty,
	)
	ba, err = ioutil.ReadFile(path)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}
	addClassProperty = string(ba)

	// 'createClassObject.sql' File.
	path = filepath.Join(
		this.settings.Path().GetSqlCommandTemplates(),
		paths.FileSqlCommandTemplateAddClassObject,
	)
	ba, err = ioutil.ReadFile(path)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}
	addClassObject = string(ba)

	// 'createClassObject.sql' File.
	path = filepath.Join(
		this.settings.Path().GetSqlCommandTemplates(),
		paths.FileSqlCommandTemplateAddClassObjectProperty,
	)
	ba, err = ioutil.ReadFile(path)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}
	addClassObjectProperty = string(ba)

	// 'readClasses.sql' File.
	path = filepath.Join(
		this.settings.Path().GetSqlCommandTemplates(),
		paths.FileSqlCommandTemplateReadClasses,
	)
	ba, err = ioutil.ReadFile(path)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}
	readClasses = string(ba)

	// 'readClassProperties.sql' File.
	path = filepath.Join(
		this.settings.Path().GetSqlCommandTemplates(),
		paths.FileSqlCommandTemplateReadClassProperties,
	)
	ba, err = ioutil.ReadFile(path)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}
	readClassProperties = string(ba)

	// 'readClassObjects.sql' File.
	path = filepath.Join(
		this.settings.Path().GetSqlCommandTemplates(),
		paths.FileSqlCommandTemplateReadClassObjects,
	)
	ba, err = ioutil.ReadFile(path)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}
	readClassObjects = string(ba)

	// 'readClassObjectProperty.sql' File.
	path = filepath.Join(
		this.settings.Path().GetSqlCommandTemplates(),
		paths.FileSqlCommandTemplateReadClassObjectProperty,
	)
	ba, err = ioutil.ReadFile(path)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}
	readClassObjectProperty = string(ba)

	// createDbClassesAndProperties Templates.
	err = this.settings.UpdateSqlCommandTemplate(
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
