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

// SqlTemplates.go.

package sqlTemplates

// Collection Settings SQL Command Templates.

import (
	"errors"

	"github.com/legacy-vault/library/go/collection/Common"
	"github.com/legacy-vault/library/go/collection/Errorz"
)

const ErrorReporter = "SqlTemplates"

// SQL Command Templates.
type SqlTemplates struct {

	// Commands for Table Creation.

	// Addition.
	addClass               string
	addClassObject         string
	addClassProperty       string
	addClassObjectProperty string

	// Creation.
	createClassesTable         string
	createClassPropertiesTable string
	createClassPropertyTable   string
	createClassTable           string

	// Reading.
	readClasses             string
	readClassProperties     string
	readClassObjects        string
	readClassObjectProperty string
}

// Returns the 'addClass' Field.
func (this SqlTemplates) GetAddClass() string {

	return this.addClass
}

// Returns the 'addClassObject' Field.
func (this SqlTemplates) GetAddClassObject() string {

	return this.addClassObject
}

// Returns the 'addClassObjectProperty' Field.
func (this SqlTemplates) GetAddClassObjectProperty() string {

	return this.addClassObjectProperty
}

// Returns the 'addClassProperty' Field.
func (this SqlTemplates) GetAddClassProperty() string {

	return this.addClassProperty
}

// Returns the 'createClassesTable' Field.
func (this SqlTemplates) GetCreateClassesTable() string {

	return this.createClassesTable
}

// Returns the 'createClassPropertiesTable' Field.
func (this SqlTemplates) GetCreateClassPropertiesTable() string {

	return this.createClassPropertiesTable
}

// Returns the 'createClassPropertyTable' Field.
func (this SqlTemplates) GetCreateClassPropertyTable() string {

	return this.createClassPropertyTable
}

// Returns the 'createClassTable' Field.
func (this SqlTemplates) GetCreateClassTable() string {

	return this.createClassTable
}

// Returns the 'readClasses' Field.
func (this SqlTemplates) GetReadClasses() string {

	return this.readClasses
}

// Returns the 'readClassObjectProperty' Field.
func (this SqlTemplates) GetReadClassObjectProperty() string {

	return this.readClassObjectProperty
}

// Returns the 'readClassObjects' Field.
func (this SqlTemplates) GetReadClassObjects() string {

	return this.readClassObjects
}

// Returns the 'readClassProperties' Field.
func (this SqlTemplates) GetReadClassProperties() string {

	return this.readClassProperties
}

// Sets all the Fields.
func (this *SqlTemplates) SetAll(
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

	this.addClass = addClass
	this.addClassObject = addClassObject
	this.addClassObjectProperty = addClassObjectProperty
	this.addClassProperty = addClassProperty

	this.createClassesTable = createClassesTable
	this.createClassTable = createClassTable
	this.createClassPropertiesTable = createClassPropertiesTable
	this.createClassPropertyTable = createClassPropertyTable

	this.readClasses = readClasses
	this.readClassProperties = readClassProperties
	this.readClassObjects = readClassObjects
	this.readClassObjectProperty = readClassObjectProperty

	return nil
}
