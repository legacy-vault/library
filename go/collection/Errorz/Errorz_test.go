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

// Errorz_test.go.

package errorz

// Functions for Errors Processing.

import (
	"errors"
	"fmt"
	"testing"
)

func Test_Combine(t *testing.T) {

	var errA error
	var errB error
	var errDouble error
	var ok bool

	errA = errors.New("a")
	errB = errors.New("b")

	errDouble = Combine(errA, errB)
	fmt.Println(errDouble)
	ok = (errDouble.Error() == "a"+CombineSeparator+"b")
	if !ok {
		t.FailNow()
	}
}

func Test_Report(t *testing.T) {

	var errFromReporter error
	var errReport error
	var ok bool
	var reporter string

	reporter = "Service A"
	errFromReporter = errors.New("Something has gone wild")
	errReport = Report(reporter, errFromReporter)

	fmt.Println(errReport)
	ok = (errReport.Error() == "[ 'Service A' Error : Something has gone wild ]")
	if !ok {
		t.FailNow()
	}
}
