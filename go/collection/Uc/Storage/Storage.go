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

// Storage.

import (
	"github.com/legacy-vault/library/go/collection/Database"
	"github.com/legacy-vault/library/go/collection/Uc/Storage/Settings"
)

const ErrorReporter = "Storage"

// Error Messages.
const ErrFormatColumnsCount = "Columns Count Mismatch. " +
	"Expected: '%v', received: '%v'."
const ErrFormatColumnName = "Column Name Mismatch. " +
	"Expected: '%v', received: '%v'."
const ErrFormatColumnType = "Column Type Mismatch. " +
	"Expected: '%v', received: '%v'."
const ErrFormatMissingTable = "Missing Table '%v'"

// Creation Messages.
const MsgClassesTableCreated = "Classes Table " +
	"has been created."
const MsgFormatClassTableCreated = "Class '%v' Table " +
	"has been created."
const MsgFormatClassPropertiesTableCreated = "Class '%v' Properties Table " +
	"has been created."
const MsgFormatClassPropertyTableCreated = "Class '%v' Property '%v' Table " +
	"has been created."

// Registration Messages.
const MsgFormatClassRegistered = "Class '%v' " +
	"has been registered."
const MsgFormatClassObjectRegistered = "Object with ID '%v' " +
	"has been registered in the Class '%v'"
const MsgFormatClassPropertyRegistered = "Class '%v' Property '%v' " +
	"has been registered."
const MsgFormatClassObjectPropertyRegistered = "Class '%v' Property '%v' " +
	"has been registered for an Object with ID='%v'."

const ErrConfiguration = "Configuration Error"
const ErrConnection = "Connection Error"
const ErrDisconnection = "Disconnection Error"

type Storage struct {
	database *database.Database
	settings settings.Settings
}
