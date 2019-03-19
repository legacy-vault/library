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

// database.go.

package storage

// Storage Database Methods.
// Wrappers for the Database Methods.

import (
	"xxx/Database"
	databaseCollection "xxx/Database/Collection"
)

func (this Storage) QueryClasses(
	query string,
) ([]databaseCollection.DatabaseClass, error) {

	return this.database.QueryClasses(query)
}

func (this Storage) QueryClassObjectProperty(
	query string,
) (databaseCollection.DatabaseClassObjectProperty, error) {

	return this.database.QueryClassObjectProperty(query)
}

func (this Storage) QueryClassObjects(
	query string,
) ([]databaseCollection.DatabaseClassObject, error) {

	return this.database.QueryClassObjects(query)
}

func (this Storage) QueryClassProperties(
	query string,
) ([]databaseCollection.DatabaseClassProperty, error) {

	return this.database.QueryClassProperties(query)
}

func (this Storage) ShowColumns(
	tableName string,
) ([]database.DatabaseTableColumnSettings, error) {

	return this.database.ShowColumns(tableName)
}

func (this Storage) SimpleQuery(
	query string,
) error {

	return this.database.SimpleQuery(query)
}

func (this Storage) TableExists(
	tableName string,
) (bool, error) {

	return this.database.TableExists(tableName)
}
