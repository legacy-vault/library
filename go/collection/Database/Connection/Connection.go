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

// Connection.go.

package connection

// Database Connection.

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/legacy-vault/library/go/collection/Common"
	"github.com/legacy-vault/library/go/collection/Database/Class"
	"github.com/legacy-vault/library/go/collection/Database/Class/Value"
	"github.com/legacy-vault/library/go/collection/Database/Configuration"
	"github.com/legacy-vault/library/go/collection/Errorz"
)

const ErrorReporter = "Connection"

const ErrAlreadyConnected = "Already connected"
const ErrNotConnected = "Not connected"

type Connection struct {

	// Driver-specific Connections.
	dbMySQL *sql.DB

	// Flags.
	isConnected bool
}

// Closes the Connection.
func (this *Connection) Close(
	databaseClass value.Value,
) error {

	var err error

	// Fool's Check.
	if this == nil {
		err = errors.New(common.ErrMethodOwnerDoesNotExist)
		err = errorz.Report(ErrorReporter, err)
		return err
	}
	if this.IsEstablished() == false {
		err = fmt.Errorf(
			ErrNotConnected,
		)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Action.
	switch databaseClass {

	case value.Mysql:

		// Close the MySQL Connection.
		err = this.closeMysql()
		if err != nil {
			err = errorz.Report(ErrorReporter, err)
			return err
		}

	default:

		// Database Class in unsupported.
		err = fmt.Errorf(
			class.ErrFormatClassValueAliasUnknown,
			databaseClass,
		)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Finalization.
	this.isConnected = false

	return nil
}

// Closes the Connection.
// MySQL Method Variant.
func (this *Connection) closeMysql() error {

	var err error

	err = this.dbMySQL.Close()
	if err != nil {
		return err
	}

	return nil
}

// Establishes the Connection.
func (this *Connection) Establish(
	databaseClass value.Value,
	configuration configuration.Configuration,
) error {

	var err error

	// Fool's Check.
	if this == nil {
		err = errors.New(common.ErrMethodOwnerDoesNotExist)
		err = errorz.Report(ErrorReporter, err)
		return err
	}
	if this.IsEstablished() == true {
		err = fmt.Errorf(
			ErrAlreadyConnected,
		)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Action.
	switch databaseClass {

	case value.Mysql:

		// Establish the MySQL Connection.
		err = this.establishMysql(configuration)
		if err != nil {
			err = errorz.Report(ErrorReporter, err)
			return err
		}

	default:

		// Database Class in unsupported.
		err = fmt.Errorf(
			class.ErrFormatClassValueAliasUnknown,
			databaseClass,
		)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Finalization.
	this.isConnected = true

	return nil
}

// Establishes the MySQL Connection.
// MySQL Method Variant.
func (this *Connection) establishMysql(
	configuration configuration.Configuration,
) error {

	var err error

	// Establish the MySQL Connection.
	this.dbMySQL, err = sql.Open(
		"mysql",
		configuration.GetDsn().GetValue(),
	)
	if err != nil {
		return err
	}
	err = this.dbMySQL.Ping()
	if err != nil {
		return err
	}

	return nil
}

// Returns the 'connected' State.
func (this Connection) IsEstablished() bool {

	return this.isConnected
}

// Runs a Query.
// It is the Caller's Responsibility to close the Rows Object.
func (this Connection) Query(
	databaseClass value.Value,
	query string,
	args ...interface{},
) (*sql.Rows, error) {

	var err error
	var rows *sql.Rows

	// Action.
	switch databaseClass {

	case value.Mysql:

		// MySQL Query Execution.
		rows, err = this.queryMysql(query, args...)
		if err != nil {
			err = errorz.Report(ErrorReporter, err)
			return rows, err
		}

	default:

		// Database Class in unsupported.
		err = fmt.Errorf(
			class.ErrFormatClassValueAliasUnknown,
			databaseClass,
		)
		err = errorz.Report(ErrorReporter, err)
		return nil, err
	}

	// Finalization.
	// None.
	return rows, nil
}

// Runs a Query.
// It is the Caller's Responsibility to close the Rows Object.
// MySQL Method Variant.
func (this Connection) queryMysql(
	query string,
	args ...interface{},
) (*sql.Rows, error) {

	return this.dbMySQL.Query(query, args...)
}
