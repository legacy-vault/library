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

// Test.go.

package test

// Methods which simplify Error Checking in built-in (Go Language's) Tests.

import (
	"fmt"
	"testing"

	"github.com/legacy-vault/library/go/collection/Common"
)

const ErrIsNotRunning = "Test is not running"

type Test struct {
	t *testing.T
}

// Checks the Test's Error.
// If an Error is set, stops the Test.
func (this *Test) CheckError(
	e error,
) {

	// Fool's Check.
	if this == nil {
		this.t.Error(common.ErrMethodOwnerDoesNotExist)
		this.t.FailNow()
		return
	}
	if !this.IsRunning() {
		this.t.Error(ErrIsNotRunning)
		this.t.FailNow()
		return
	}

	// Check an Error.
	if e != nil {
		this.t.Error(e)
		this.t.FailNow()
	}
}

// Checks whether the Test has been initialized.
func (this Test) IsRunning() bool {

	if this.t == nil {
		return false
	}

	return true
}

// Stops the Test.
func (this *Test) Stop(
	errorMessage string,
) {

	var err error

	// Fool's Check.
	if this == nil {
		this.t.Error(common.ErrMethodOwnerDoesNotExist)
		this.t.FailNow()
		return
	}
	if !this.IsRunning() {
		this.t.Error(ErrIsNotRunning)
		this.t.FailNow()
		return
	}

	// Stop the Test.
	err = fmt.Errorf(errorMessage)
	this.t.Error(err)
	this.t.FailNow()
}

// Initializes the Test.
func New(
	t *testing.T,
) *Test {

	var result *Test

	result = new(Test)
	result.t = t

	return result
}
