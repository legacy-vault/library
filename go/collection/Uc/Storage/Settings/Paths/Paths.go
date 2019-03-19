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

// Paths.go.

package paths

// Collection Settings Paths.

import (
	"errors"

	"github.com/legacy-vault/library/go/collection/Common"
	"github.com/legacy-vault/library/go/collection/Errorz"
)

const ErrorReporter = "Paths"

// Path.
const FolderSqlCommandTemplatesDefault = "SQL"

// File.
const (
	FileSqlCommandTemplateAddClass                   = "addClass.sql"
	FileSqlCommandTemplateAddClassObject             = "addClassObject.sql"
	FileSqlCommandTemplateAddClassObjectProperty     = "addClassObjectProperty.sql"
	FileSqlCommandTemplateAddClassProperty           = "addClassProperty.sql"
	FileSqlCommandTemplateCreateClassesTable         = "createClassesTable.sql"
	FileSqlCommandTemplateCreateClassTable           = "createClassTable.sql"
	FileSqlCommandTemplateCreateClassPropertiesTable = "createClassPropertiesTable.sql"
	FileSqlCommandTemplateCreateClassPropertyTable   = "createClassPropertyTable.sql"
	FileSqlCommandTemplateReadClasses                = "readClasses.sql"
	FileSqlCommandTemplateReadClassProperties        = "readClassProperties.sql"
	FileSqlCommandTemplateReadClassObjects           = "readClassObjects.sql"
	FileSqlCommandTemplateReadClassObjectProperty    = "readClassObjectProperty.sql"
)

// Database Table Settings: Names.
type Paths struct {

	// Path to a Folder with SQL Command Templates.
	// Usually, 'SQL'.
	sqlCommandTemplates string
}

// Returns the 'sqlCommandTemplates' Field.
func (this Paths) GetSqlCommandTemplates() string {

	return this.sqlCommandTemplates
}

// Sets all the Fields.
func (this *Paths) SetAll(
	sqlCommandTemplates string,
) error {

	var err error

	// Fool's Check.
	if this == nil {
		err = errors.New(common.ErrMethodOwnerDoesNotExist)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// 1. SQL Command Templates Folder Path.
	if len(sqlCommandTemplates) == 0 {
		this.sqlCommandTemplates = FolderSqlCommandTemplatesDefault
	} else {
		this.sqlCommandTemplates = sqlCommandTemplates
	}

	return nil
}
