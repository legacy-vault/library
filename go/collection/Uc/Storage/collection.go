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

// collection.go.

package storage

// Low-Level Methods to work with Collection's Parts in the Database.
// These Methods compose Queries and run them on a Database Server.

import (
	"fmt"
	"strings"

	databaseCollection "github.com/legacy-vault/library/go/collection/Database/Collection"
	"github.com/legacy-vault/library/go/collection/Errorz"
	"github.com/legacy-vault/library/go/collection/Uc/Collection/Class"
	"github.com/legacy-vault/library/go/collection/Uc/Collection/Class/Object"
	objectProperty "github.com/legacy-vault/library/go/collection/Uc/Collection/Class/Object/Property"
	classProperty "github.com/legacy-vault/library/go/collection/Uc/Collection/Class/Property"
)

// Adds the Class to the Classes Table in the Database.
func (this *Storage) AddClass(
	aClass class.Class,
) error {

	var className string
	var err error
	var errJournal error
	var msg string
	var query string

	// Query: Class Addition.
	query = this.ComposeQueryAddClass(aClass)
	err = this.SimpleQuery(query)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Report.
	className = strings.ToLower(aClass.GetName())
	msg = fmt.Sprintf(
		MsgFormatClassRegistered,
		className,
	)
	errJournal = this.AddJournalRecord(msg)
	if errJournal != nil {
		errJournal = errorz.Report(ErrorReporter, errJournal)
		return errJournal
	}

	return nil
}

// Adds a Class Object to the Class Table in the Database.
func (this *Storage) AddClassObject(
	aClass class.Class,
	aClassObject object.Object,
) error {

	var className string
	var err error
	var errJournal error
	var msg string
	var query string

	// Query: Add a Class Object into Class Table.
	query = this.ComposeQueryAddClassObject(aClass, aClassObject)
	err = this.SimpleQuery(query)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Report.
	className = strings.ToLower(aClass.GetName())
	msg = fmt.Sprintf(
		MsgFormatClassObjectRegistered,
		aClassObject.GetId(),
		className,
	)
	errJournal = this.AddJournalRecord(msg)
	if errJournal != nil {
		errJournal = errorz.Report(ErrorReporter, errJournal)
		return errJournal
	}

	return nil
}

// Adds a Class Object Property to the Class Property Table in the Database.
func (this *Storage) AddClassObjectProperty(
	aClass class.Class,
	aClassObject object.Object,
	aClassObjectProperty objectProperty.Property,
) error {

	var className string
	var classPropertyName string
	var err error
	var errJournal error
	var msg string
	var query string

	// Parameters for a Query.

	// Query: Add a Class Object Property into Class Property Table.
	query, classPropertyName, err = this.ComposeQueryAddClassObjectProperty(
		aClass,
		aClassObject,
		aClassObjectProperty,
	)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}
	err = this.SimpleQuery(query)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Report.
	className = strings.ToLower(aClass.GetName())
	msg = fmt.Sprintf(
		MsgFormatClassObjectPropertyRegistered,
		className,
		classPropertyName,
		aClassObject.GetId(),
	)
	errJournal = this.AddJournalRecord(msg)
	if errJournal != nil {
		errJournal = errorz.Report(ErrorReporter, errJournal)
		return errJournal
	}

	return nil
}

// Adds a Class Property to the Class Properties Table in the Database.
func (this *Storage) AddClassProperty(
	aClass class.Class,
	aClassProperty classProperty.Property,
) error {

	var className string
	var classPropertyName string
	var err error
	var errJournal error
	var msg string
	var query string

	// Query: Class Property Addition.
	query = this.ComposeQueryAddClassProperty(aClass, aClassProperty)
	err = this.SimpleQuery(query)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Report.
	className = strings.ToLower(aClass.GetName())
	classPropertyName = strings.ToLower(aClassProperty.GetName())
	msg = fmt.Sprintf(
		MsgFormatClassPropertyRegistered,
		className,
		classPropertyName,
	)
	errJournal = this.AddJournalRecord(msg)
	if errJournal != nil {
		errJournal = errorz.Report(ErrorReporter, errJournal)
		return errJournal
	}

	return nil
}

// Creates a Class Table in the Database.
func (this *Storage) CreateTableClass(
	aClass class.Class,
) error {

	var className string
	var err error
	var errJournal error
	var msg string
	var query string

	// Query: Class Table Creation.
	query = this.ComposeQueryCreateTableClass(aClass)
	err = this.SimpleQuery(query)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Report.
	className = strings.ToLower(aClass.GetName())
	msg = fmt.Sprintf(
		MsgFormatClassTableCreated,
		className,
	)
	errJournal = this.AddJournalRecord(msg)
	if errJournal != nil {
		errJournal = errorz.Report(ErrorReporter, errJournal)
		return errJournal
	}

	return nil
}

// Creates a Classes Table in the Database.
func (this *Storage) CreateTableClasses() error {

	var err error
	var errJournal error
	var msg string
	var query string

	// Query: Classes Table Creation.
	query = this.ComposeQueryCreateTableClasses()
	err = this.SimpleQuery(query)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Report.
	msg = MsgClassesTableCreated
	errJournal = this.AddJournalRecord(msg)
	if errJournal != nil {
		errJournal = errorz.Report(ErrorReporter, errJournal)
		return errJournal
	}

	return nil
}

// Creates a Properties Table for a Class in the Database.
func (this *Storage) CreateTableClassProperties(
	aClass class.Class,
) error {

	var className string
	var err error
	var errJournal error
	var msg string
	var query string

	// Query: Class Properties Table Creation.
	query = this.ComposeQueryCreateTableClassProperties(aClass)
	err = this.SimpleQuery(query)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Report.
	className = strings.ToLower(aClass.GetName())
	msg = fmt.Sprintf(
		MsgFormatClassPropertiesTableCreated,
		className,
	)
	errJournal = this.AddJournalRecord(msg)
	if errJournal != nil {
		errJournal = errorz.Report(ErrorReporter, errJournal)
		return errJournal
	}

	return nil
}

// Creates a Class Property Table in the Database.
func (this *Storage) CreateTableClassProperty(
	aClass class.Class,
	aClassProperty classProperty.Property,
) error {

	var className string
	var classPropertyName string
	var err error
	var errJournal error
	var msg string
	var query string

	// Query: Class Property Table Creation.
	query = this.ComposeQueryCreateTableClassProperty(aClass, aClassProperty)
	err = this.SimpleQuery(query)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Report.
	className = strings.ToLower(aClass.GetName())
	classPropertyName = strings.ToLower(aClassProperty.GetName())
	msg = fmt.Sprintf(
		MsgFormatClassPropertyTableCreated,
		className,
		classPropertyName,
	)
	errJournal = this.AddJournalRecord(msg)
	if errJournal != nil {
		errJournal = errorz.Report(ErrorReporter, errJournal)
		return errJournal
	}

	return nil
}

// Reads the Classes List from a Database.
func (this *Storage) ReadClasses() ([]databaseCollection.DatabaseClass, error) {

	var err error
	var query string
	var result []databaseCollection.DatabaseClass

	// Query: Read Classes Table.
	query = this.ComposeQueryReadTableClasses()
	result, err = this.QueryClasses(query)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return result, err
	}

	return result, err
}

// Reads the Class Object Property for a specific Class Object from a Database.
func (this *Storage) ReadClassObjectProperty(
	aClass class.Class,
	aClassProperty classProperty.Property,
	aClassObject object.Object,
) (databaseCollection.DatabaseClassObjectProperty, error) {

	var err error
	var query string
	var result databaseCollection.DatabaseClassObjectProperty

	// Query: Read the Class Properties Table.
	query = this.ComposeQueryReadTableClassObjectProperty(
		aClass,
		aClassProperty,
		aClassObject,
	)
	result, err = this.QueryClassObjectProperty(query)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return result, err
	}

	return result, err
}

// Reads the Class Objects List for a specific Class from a Database.
func (this *Storage) ReadClassObjects(
	aClass class.Class,
) ([]databaseCollection.DatabaseClassObject, error) {

	var err error
	var query string
	var result []databaseCollection.DatabaseClassObject

	// Query: Read the Class Table.
	query = this.ComposeQueryReadTableClass(aClass)
	result, err = this.QueryClassObjects(query)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return result, err
	}

	return result, err
}

// Reads the Class Properties List for a specific Class from a Database.
func (this *Storage) ReadClassProperties(
	aClass class.Class,
) ([]databaseCollection.DatabaseClassProperty, error) {

	var err error
	var query string
	var result []databaseCollection.DatabaseClassProperty

	// Query: Read the Class Properties Table.
	query = this.ComposeQueryReadTableClassProperties(aClass)
	result, err = this.QueryClassProperties(query)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return result, err
	}

	return result, err
}
