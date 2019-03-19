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

// Configuration.go.

package configuration

// Database Configuration.

import (
	"errors"
	"fmt"
	"time"

	"xxx/Common"
	"xxx/Database/Configuration/Dsn"
	"xxx/Errorz"
)

const ErrorReporter = "Configuration"

const ErrNotConfiguredDsn = "DSN Configuration is not done"

type Configuration struct {

	// Data Source Name.
	dsn dsn.Dsn

	// 'Configuration is Done' Flag.
	isDone bool

	// Journal writing 'enabled' Flag.
	journalIsOn bool

	// Last Set Time, Time of the last Settings Change.
	lst time.Time
}

// Finalizes the Configuration.
func (this *Configuration) Finalize() error {

	var err error

	// Fool's Check.
	if this == nil {
		err = errors.New(common.ErrMethodOwnerDoesNotExist)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	if this.dsn.IsConfigured() != true {
		err = fmt.Errorf(
			ErrNotConfiguredDsn,
		)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	this.isDone = true
	this.lst = time.Now()

	return nil
}

// Returns the 'dsn' Field.
func (this Configuration) GetDsn() dsn.Dsn {

	return this.dsn
}

// Returns the 'journalIsOn' State.
func (this Configuration) GetJournalIsOn() bool {

	return this.journalIsOn
}

// Returns the 'done' State.
func (this Configuration) IsSet() bool {

	return this.isDone
}

// Sets the 'journalIsOn' State.
func (this *Configuration) SetJournalIsOn(
	journalIsOn bool,
) error {

	var err error

	// Fool's Check.
	if this == nil {
		err = errors.New(common.ErrMethodOwnerDoesNotExist)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	this.journalIsOn = journalIsOn

	return nil
}

// Updates the 'dsn' Field.
func (this *Configuration) UpdateDsn(
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

	// Set the DSN.
	err = this.dsn.SetAll(
		hostName,
		portNumber,
		userName,
		password,
		database,
		parameters,
	)
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	return nil
}
