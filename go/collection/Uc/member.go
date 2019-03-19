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

// member.go.

package uc

// Methods to check Entities' Membership.

import (
	databaseCollection "github.com/legacy-vault/library/go/collection/Database/Collection"
	"github.com/legacy-vault/library/go/collection/Uc/Collection/Class"
	"github.com/legacy-vault/library/go/collection/Uc/Collection/Class/Object"
	classProperty "github.com/legacy-vault/library/go/collection/Uc/Collection/Class/Property"
)

// Checks whether a Class Object read from the Database is truly a Member
// of a Class.
func (this Uc) membershipClassObject(
	aClass class.Class,
	dbClassObject databaseCollection.DatabaseClassObject,
) bool {

	var classId uint

	// Preparation.
	classId = aClass.GetId()

	// Verification.
	if dbClassObject.ClassId != classId {
		return false
	}

	return true
}

// Checks whether a Class Object Property read from the Database is truly
// a Member of a Class.
func (this Uc) membershipClassObjectProperty(
	aClass class.Class,
	aClassObject object.Object,
	aClassProperty classProperty.Property,
	dbClassObjectProperty databaseCollection.DatabaseClassObjectProperty,
) bool {

	var classId uint
	var classObjectId uint
	var classPropertyId uint

	// Preparation.
	classId = aClass.GetId()
	classObjectId = aClassObject.GetId()
	classPropertyId = aClassProperty.GetId()

	// Verify the Class ID of an Class Object Property.
	if dbClassObjectProperty.ClassId != classId {
		return false
	}

	// Verify the Object ID of an Class Object Property.
	if dbClassObjectProperty.ObjectId != classObjectId {
		return false
	}

	// Verify the Property ID of an Class Object Property.
	if dbClassObjectProperty.PropertyId != classPropertyId {
		return false
	}

	return true
}

// Checks whether a Class Property read from the Database is truly a Member
// of a Class.
func (this Uc) membershipClassProperty(
	aClass class.Class,
	dbClassProperty databaseCollection.DatabaseClassProperty,
) bool {

	var classId uint

	// Preparation.
	classId = aClass.GetId()

	// Verification.
	if dbClassProperty.ClassId != classId {
		return false
	}

	return true
}
