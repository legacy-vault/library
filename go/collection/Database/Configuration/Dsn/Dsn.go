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

// Dsn.go.

package dsn

// Database Configuration Data Source Name.

import (
	"errors"
	"fmt"
	"strconv"

	"xxx/Common"
	"xxx/Errorz"
)

const ErrorReporter = "Dsn"

const DriverMysqlConnectionType = "tcp"
const FormatMysql = "%s:%s@%s(%s:%s)/%s%s"
const FormatMysqlParametersPrefix = "?"

type Dsn struct {
	database string
	host     string
	isDone   bool

	// List of Parameters, without leading '?' Symbol.
	// Format: "param1=value1&...&paramN=valueN".
	parameters string

	password string
	portStr  string
	portInt  uint16
	username string
	value    string
}

// Creates a DSN for the MySQL Database Driver.
func (this Dsn) createDSNForMySQL() string {

	var parametersFull string
	var result string

	if len(this.parameters) > 0 {
		parametersFull = FormatMysqlParametersPrefix + this.parameters
	} else {
		parametersFull = ""
	}

	result = fmt.Sprintf(
		FormatMysql,
		this.username,
		this.password,
		DriverMysqlConnectionType,
		this.host,
		this.portStr,
		this.database,
		parametersFull,
	)

	return result
}

// Returns the Host Name.
func (this Dsn) GetHost() string {

	return this.host
}

// Returns the Port as Number.
func (this Dsn) GetPortAsNumber() int {

	return int(this.portInt)
}

// Returns the Port as String.
func (this Dsn) GetPortAsString() string {

	return this.portStr
}

// Returns the User Name.
func (this Dsn) GetUsername() string {

	return this.username
}

// Returns the total DSN Value as String.
func (this Dsn) GetValue() string {

	return this.value
}

// Returns the 'configured' State.
func (this Dsn) IsConfigured() bool {

	return this.isDone
}

// Sets all the Fields of a DSN.
func (this *Dsn) SetAll(
	hostName string,
	portNumber uint16,
	userName string,
	password string,
	database string,
	parameters string,
) error {

	var err error

	// Fool's Check.
	if this == nil {
		err = errors.New(common.ErrMethodOwnerDoesNotExist)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	this.host = hostName
	this.portInt = portNumber
	this.portStr = strconv.FormatUint(uint64(portNumber), 10)
	this.username = userName
	this.password = password
	this.database = database
	this.parameters = parameters

	this.value = this.createDSNForMySQL()
	this.isDone = true

	return nil
}
