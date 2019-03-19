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

// query.go.

package storage

// Storage Query Composing.

import (
	"fmt"
	"github.com/legacy-vault/library/go/collection/Database/StringLiteral"
	"github.com/legacy-vault/library/go/collection/Uc/Collection/Class"
	"github.com/legacy-vault/library/go/collection/Uc/Collection/Class/Object"
	objectProperty "github.com/legacy-vault/library/go/collection/Uc/Collection/Class/Object/Property"
	classProperty "github.com/legacy-vault/library/go/collection/Uc/Collection/Class/Property"
	"github.com/legacy-vault/library/go/collection/Uc/Collection/Class/Property/KindSettings"
	"strings"
)

// Composes a Query to add a Class.
func (this Storage) ComposeQueryAddClass(
	aClass class.Class,
) string {

	var classId uint
	var className string
	var classesTableName string
	var result string
	var template string

	template = this.settings.SqlCommandTemplate().GetAddClass()
	classesTableName = this.settings.Table().Classes()
	classId = aClass.GetId()
	className = strings.ToLower(aClass.GetName())

	result = fmt.Sprintf(
		template,
		classesTableName,
		classId,
		className,
	)

	return result
}

// Composes a Query to add a Class Object.
func (this Storage) ComposeQueryAddClassObject(
	aClass class.Class,
	aClassObject object.Object,
) string {

	var classId uint
	var className string
	var classObjectId uint
	var classTableName string
	var result string
	var template string

	template = this.settings.SqlCommandTemplate().GetAddClassObject()
	className = strings.ToLower(aClass.GetName())
	classTableName = this.ComposeTableNameClass(className)
	classObjectId = aClassObject.GetId()
	classId = aClass.GetId()

	result = fmt.Sprintf(
		template,
		classTableName,
		classObjectId,
		classId,
	)

	return result
}

// Composes a Query to add a Class Object Property.
// Returns a Query and a Class Property Name.
func (this Storage) ComposeQueryAddClassObjectProperty(
	aClass class.Class,
	aClassObject object.Object,
	aClassObjectProperty objectProperty.Property,
) (string, string, error) {

	var aClassProperty classProperty.Property
	var classId uint
	var className string
	var classObjectId uint
	var classObjectPropertyId uint
	var classObjectPropertyValueFinal string // Value with single Quotes if needed.
	var classPropertyKind kindsettings.KindSettings
	var classPropertyName string
	var classPropertyTableName string
	var err error
	var query string
	var template string

	template = this.settings.SqlCommandTemplate().GetAddClassObjectProperty()
	className = strings.ToLower(aClass.GetName())
	classId = aClass.GetId()
	classObjectId = aClassObject.GetId()
	classObjectPropertyId = aClassObjectProperty.GetId()
	aClassProperty, err = aClass.GetPropertyById(classObjectPropertyId)
	if err != nil {
		return "", "", err
	}
	classPropertyName = strings.ToLower(aClassProperty.GetName())
	classPropertyTableName = this.ComposeTableNameClassProperty(
		className,
		classPropertyName,
	)
	classPropertyKind = aClassProperty.GetKind()
	classObjectPropertyValueFinal, err =
		aClassObjectProperty.GetValueString(classPropertyKind)
	if err != nil {
		return "", "", err
	}

	query = fmt.Sprintf(
		template,
		classPropertyTableName,
		classObjectPropertyId,
		classObjectId,
		classId,
		classObjectPropertyValueFinal,
	)

	return query, classPropertyName, nil
}

// Composes a Query to add a Class Property.
func (this Storage) ComposeQueryAddClassProperty(
	aClass class.Class,
	aClassProperty classProperty.Property,
) string {

	var aPropertyDescription string
	var aPropertyName string
	var classId uint
	var className string
	var classObjectPropertyTypeString string
	var classObjectPropertyId uint
	var classPropertiesTableName string
	var result string
	var template string

	template = this.settings.SqlCommandTemplate().GetAddClassProperty()
	className = strings.ToLower(aClass.GetName())
	classPropertiesTableName = this.ComposeTableNameClassProperties(className)
	classObjectPropertyId = aClassProperty.GetId()
	aPropertyName = strings.ToLower(aClassProperty.GetName())
	classObjectPropertyTypeString = aClassProperty.GetType().DbType
	classId = aClass.GetId()

	// Property Description escaped.
	aPropertyDescription = aClassProperty.GetDescription()
	sl := stringliteral.StringLiteral{Value: aPropertyDescription}
	aPropertyDescription = sl.Escaped()

	result = fmt.Sprintf(
		template,
		classPropertiesTableName,
		classObjectPropertyId,
		aPropertyName,
		aPropertyDescription,
		classObjectPropertyTypeString,
		classId,
	)

	return result
}

// Composes a Query to create a Class Table.
func (this Storage) ComposeQueryCreateTableClass(
	aClass class.Class,
) string {

	var classesTableName string
	var className string
	var classTableName string
	var result string
	var template string

	template = this.settings.SqlCommandTemplate().GetCreateClassTable()
	classesTableName = this.GetSettings().Table().Classes()
	className = strings.ToLower(aClass.GetName())
	classTableName = this.ComposeTableNameClass(className)

	result = fmt.Sprintf(
		template,
		classTableName,
		classTableName,
		classesTableName,
		classesTableName,
	)

	return result
}

// Composes a Query to create a Classes Table.
func (this Storage) ComposeQueryCreateTableClasses() string {

	var result string
	var template string

	template = this.settings.SqlCommandTemplate().GetCreateClassesTable()
	result = fmt.Sprintf(
		template,
		this.settings.Table().Classes(),
	)

	return result
}

// Composes a Query to create a Class Properties Table.
func (this Storage) ComposeQueryCreateTableClassProperties(
	aClass class.Class,
) string {

	var className string
	var classPropertiesTableName string
	var classesTableName string
	var result string
	var template string

	template =
		this.settings.SqlCommandTemplate().GetCreateClassPropertiesTable()
	classesTableName = this.settings.Table().Classes()
	className = strings.ToLower(aClass.GetName())
	classPropertiesTableName = this.ComposeTableNameClassProperties(
		className,
	)

	result = fmt.Sprintf(
		template,
		classPropertiesTableName,
		classPropertiesTableName,
		classesTableName,
		classesTableName,
	)

	return result
}

// Composes a Query to create a Class Property Table.
func (this Storage) ComposeQueryCreateTableClassProperty(
	aClass class.Class,
	aClassProperty classProperty.Property,
) string {

	var classesTableName string
	var className string
	var classPropertiesTableName string
	var classPropertyName string
	var classPropertyTableName string
	var classPropertyTypeString string
	var query string
	var template string

	template = this.settings.SqlCommandTemplate().GetCreateClassPropertyTable()
	classPropertyName = strings.ToLower(aClassProperty.GetName())
	className = strings.ToLower(aClass.GetName())
	classPropertiesTableName = this.ComposeTableNameClassProperties(
		className,
	)
	classPropertyTableName = this.ComposeTableNameClassProperty(
		className,
		classPropertyName,
	)
	classPropertyTypeString = aClassProperty.GetType().DbType
	classesTableName = this.settings.Table().Classes()

	query = fmt.Sprintf(
		template,
		classPropertyTableName,
		classPropertyTypeString,
		classPropertyTableName,
		classPropertiesTableName,
		classPropertiesTableName,
		classPropertyTableName,
		className,
		className,
		classPropertyTableName,
		classesTableName,
		classesTableName,
	)

	return query
}

// Composes a Query to read the Class Objects Table.
func (this Storage) ComposeQueryReadTableClass(
	aClass class.Class,
) string {

	var className string
	var classTableName string
	var result string
	var template string

	template = this.settings.SqlCommandTemplate().GetReadClassObjects()
	className = strings.ToLower(aClass.GetName())
	classTableName = this.ComposeTableNameClass(className)

	result = fmt.Sprintf(
		template,
		classTableName,
	)

	return result
}

// Composes a Query to read the Classes Table.
func (this Storage) ComposeQueryReadTableClasses() string {

	var classesTableName string
	var result string
	var template string

	template = this.settings.SqlCommandTemplate().GetReadClasses()
	classesTableName = this.settings.Table().Classes()

	result = fmt.Sprintf(
		template,
		classesTableName,
	)

	return result
}

// Composes a Query to read a Class Object Property.
func (this Storage) ComposeQueryReadTableClassObjectProperty(
	aClass class.Class,
	aClassProperty classProperty.Property,
	aClassObject object.Object,
) string {

	var className string
	var classObjectId uint
	var classObjectPropertyTableName string
	var classPropertyName string
	var result string
	var template string

	template = this.settings.SqlCommandTemplate().GetReadClassObjectProperty()
	className = aClass.GetName()
	classPropertyName = aClassProperty.GetName()
	classObjectPropertyTableName = this.ComposeTableNameClassProperty(
		className,
		classPropertyName,
	)
	classObjectId = aClassObject.GetId()

	result = fmt.Sprintf(
		template,
		classObjectPropertyTableName,
		classObjectId,
	)

	return result
}

// Composes a Query to read the Class Properties Table.
func (this Storage) ComposeQueryReadTableClassProperties(
	aClass class.Class,
) string {

	var className string
	var classPropertiesTableName string
	var result string
	var template string

	template = this.settings.SqlCommandTemplate().GetReadClassProperties()
	className = strings.ToLower(aClass.GetName())
	classPropertiesTableName = this.ComposeTableNameClassProperties(className)

	result = fmt.Sprintf(
		template,
		classPropertiesTableName,
	)

	return result
}
