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

// Uc.go.

package uc

// User's Collection.
// An accessible Wrapper of the Collection and its Storage.

import (
	"xxx/Uc/Collection"
	"xxx/Uc/Storage"
)

const ErrorReporter = "Uc"

const ErrNotInitialized = "Not Initialized"
const ErrNotOpened = "Not opened"
const ErrOpened = "Already opened"
const ErrNotClosed = "Not closed"

const ErrFormatPropertyClassId = "Property's Class ID Mismatch. " +
	"Expected: '%v', received: '%v'. " +
	"ClassId='%v' PropertyId='%v'."
const ErrFormatObjectClassId = "Object's  Class ID Mismatch. " +
	"Expected: '%v', received: '%v'. " +
	"ClassId='%v' ObjectId='%v'."
const ErrFormatObjectPropertyClassId = "Object Property's Class ID Mismatch. " +
	"Expected: '%v', received: '%v'. " +
	"ClassId='%v' PropertyId='%v ObjectId='%v'."
const ErrFormatObjectPropertyPropertyId = "Object Property's Property ID Mismatch. " +
	"Expected: '%v', received: '%v'. " +
	"ClassId='%v' PropertyId='%v ObjectId='%v'."
const ErrFormatObjectPropertyObjectId = "Object Property's Object ID Mismatch. " +
	"Expected: '%v', received: '%v'. " +
	"ClassId='%v' PropertyId='%v ObjectId='%v'."
const ErrFormatClassPropertyMembership = "Class Property Membership Error"
const ErrFormatClassObjectMembership = "Class Object Membership Error"
const ErrFormatClassObjectPropertyMembership = "Class Object Property " +
	"Membership Error"
const ErrPropertyKindSettings = "Property Kind Settings Error"
const ErrPropertyId = "Property ID Error"

type Uc struct {

	// Main Entities.
	collection collection.Collection
	storage    storage.Storage

	// Parameters & Settings.
	initializationIsDone bool
	isOpen               bool
}
