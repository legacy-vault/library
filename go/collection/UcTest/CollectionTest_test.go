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

// CollectionTest_test.go.

package main

import (
	"testing"

	"github.com/kr/pretty"

	dbClass "github.com/legacy-vault/library/go/collection/Database/Class/Value"
	databaseSettings "github.com/legacy-vault/library/go/collection/Database/Settings"
	"github.com/legacy-vault/library/go/collection/Test"
	"github.com/legacy-vault/library/go/collection/Uc"
	"github.com/legacy-vault/library/go/collection/Uc/Collection/Class"
	"github.com/legacy-vault/library/go/collection/Uc/Collection/Class/Object"
	objectProperty "github.com/legacy-vault/library/go/collection/Uc/Collection/Class/Object/Property"
	classProperty "github.com/legacy-vault/library/go/collection/Uc/Collection/Class/Property"
	"github.com/legacy-vault/library/go/collection/Uc/Collection/Class/Property/KindSettings"
	storageSettings "github.com/legacy-vault/library/go/collection/Uc/Storage/Settings"
)

const ErrErrorWasExpectedButNoneReceived = "Error was expected, " +
	"but none has been received"

// A Test to create a primitive Collection.
func Test_Collection_Create(t *testing.T) {

	// Working Entities.
	var aClassA class.Class
	var aClassAPropertyHeight classProperty.Property
	var aClassAPropertyHeightKindSettings kindsettings.KindSettings
	var aClassAPropertyWidth classProperty.Property
	var aClassAPropertyWidthKindSettings kindsettings.KindSettings
	var aClassAPropertyFilepath classProperty.Property
	var aClassAPropertyFilepathKindSettings kindsettings.KindSettings
	var aClassAObjectX object.Object
	var aClassAObjectXPropertyFilepath objectProperty.Property

	var aClassB class.Class
	var aClassBPropertyColor classProperty.Property
	var aClassBPropertyColorKindSettings kindsettings.KindSettings
	var aClassBObjectY object.Object
	var aClassAObjectYPropertyColor objectProperty.Property

	// Other Variables.
	var err error
	var theTest *test.Test
	var ucDbSettings databaseSettings.Settings
	var ucSettings storageSettings.Settings
	var userCollection *uc.Uc

	theTest = test.New(t)

	// Prepare the Settings for a Collection.
	ucDbSettings = databaseSettings.Settings{
		Database:   "test",
		HostName:   "localhost",
		Password:   "test",
		PortNumber: 3306,
		Type:       dbClass.MysqlAlias,
		UserName:   "test",

		Journal: true,
	}
	ucSettings, err = storageSettings.NewSettings(
		ucDbSettings,
		"myClasses", // Empty => Default.
		"",          // Empty => Default.
		"",          // Empty => Default.
		"",          // Empty => Default.
	)
	theTest.CheckError(err)

	// Connect to a Collection's Database.
	userCollection, err = uc.New(
		"My Collection",
		ucSettings,
	)
	theTest.CheckError(err)

	// Create an empty Collection in the Database.
	err = userCollection.CreateDbCollection()
	theTest.CheckError(err)

	// Create some primitive Entities in the Collection...

	// 1. Add a Class.
	aClassA, err = userCollection.AddClass("image")
	theTest.CheckError(err)
	pretty.Println(aClassA)

	// 2. Add a Class Property 'Width'.
	aClassAPropertyWidthKindSettings, err =
		kindsettings.NewWithDbType("INT(10) UNSIGNED NOT NULL")
	theTest.CheckError(err)
	aClassAPropertyWidth, err = userCollection.AddClassProperty(
		aClassA.GetId(),
		0, // Empty => Automatic Generation.
		"Width",
		"Image Width in Pixels",
		aClassAPropertyWidthKindSettings,
	)
	theTest.CheckError(err)
	pretty.Println(aClassAPropertyWidth)

	// 3. Add a Class Property 'Height'.
	aClassAPropertyHeightKindSettings = aClassAPropertyWidthKindSettings
	theTest.CheckError(err)
	aClassAPropertyHeight, err = userCollection.AddClassProperty(
		aClassA.GetId(),
		0, // Empty => Automatic Generation.
		"Height",
		"Image Height in Pixels",
		aClassAPropertyHeightKindSettings,
	)
	theTest.CheckError(err)
	pretty.Println(aClassAPropertyHeight)

	// 4. Add a Class Property 'FilePath'.
	aClassAPropertyFilepathKindSettings, err =
		kindsettings.NewWithDbType("TEXT NOT NULL")
	theTest.CheckError(err)
	aClassAPropertyFilepath, err = userCollection.AddClassProperty(
		aClassA.GetId(),
		0, // Empty => Automatic Generation.
		"FilePath",
		"An absolute Path to the Image's File",
		aClassAPropertyFilepathKindSettings,
	)
	theTest.CheckError(err)
	pretty.Println(aClassAPropertyFilepath)

	// 5. Add another Class.
	aClassB, err = userCollection.AddClass("planet")
	theTest.CheckError(err)
	pretty.Println(aClassB)

	// 6. Add a Class Property 'Color'.
	aClassBPropertyColorKindSettings, err =
		kindsettings.NewWithDbType("VARCHAR(255) NOT NULL")
	theTest.CheckError(err)
	aClassBPropertyColor, err = userCollection.AddClassProperty(
		aClassB.GetId(),
		0, // Empty => Automatic Generation.
		"Color",
		"Planet's Color",
		aClassBPropertyColorKindSettings,
	)
	theTest.CheckError(err)
	pretty.Println(aClassBPropertyColor)

	// 7. Add an Object of the 'Image' Class.
	aClassAObjectX, err = userCollection.AddClassObject(
		aClassA.GetId(),
		0, // Empty => Automatic Generation.
	)
	pretty.Println(aClassAObjectX)

	// 8. Add a 'FilePath' Property to the previously created Object.
	aClassAObjectXPropertyFilepath, err =
		userCollection.AddClassObjectProperty(
			aClassA.GetId(),
			aClassAObjectX.GetId(),
			aClassAPropertyFilepath.GetId(),
			"c:/Images/cat.png",
		)
	theTest.CheckError(err)
	pretty.Println(aClassAObjectXPropertyFilepath)

	// 9. Add an Object of the 'Planet' Class.
	aClassBObjectY, err = userCollection.AddClassObject(
		aClassB.GetId(),
		0, // Empty => Automatic Generation.
	)
	pretty.Println(aClassBObjectY)

	// 10. Add a 'Color' Property to the previously created Object.
	aClassAObjectYPropertyColor, err =
		userCollection.AddClassObjectProperty(
			aClassB.GetId(),
			aClassBObjectY.GetId(),
			aClassBPropertyColor.GetId(),
			"Yellow",
		)
	theTest.CheckError(err)
	pretty.Println(aClassAObjectYPropertyColor)

	// Close the Collection.
	err = userCollection.Close()
	theTest.CheckError(err)
}

// A Test to open an existing Collection.
func Test_Collection_OpenClose(t *testing.T) {

	var aTest *test.Test
	var classes map[uint]class.Class
	var err error
	var ucDbSettings databaseSettings.Settings
	var ucSettings storageSettings.Settings
	var userCollection *uc.Uc

	aTest = test.New(t)

	// Prepare the Settings for a Collection.
	ucDbSettings = databaseSettings.Settings{
		Database:   "test",
		HostName:   "localhost",
		Password:   "test",
		PortNumber: 3306,
		Type:       dbClass.MysqlAlias,
		UserName:   "test",

		Journal: true,
	}
	ucSettings, err = storageSettings.NewSettings(
		ucDbSettings,
		"myClasses", // Empty => Default.
		"",          // Empty => Default.
		"",          // Empty => Default.
		"",          // Empty => Default.
	)
	aTest.CheckError(err)

	// Connect to a Collection's Database.
	userCollection, err = uc.New(
		"My Collection",
		ucSettings,
	)
	aTest.CheckError(err)

	// Open a Collection from Database.
	err = userCollection.Open()
	aTest.CheckError(err)
	if userCollection.IsOpened() != true {
		aTest.Stop("IsOpened Error")
	}

	// Get Classes.
	classes, err = userCollection.GetClasses()
	aTest.CheckError(err)
	pretty.Println(classes)

	// Close the Collection.
	err = userCollection.Close()
	aTest.CheckError(err)
}
