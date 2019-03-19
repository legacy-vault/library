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

// Strings_test.go.

package strings

// Functions for Strings Processing.

import "testing"

func Test_All(t *testing.T) {

	var array []string

	array = []string{}
	if ArrayContainsString(array, "x") != false {
		t.FailNow()
	}

	array = []string{
		"aaa",
		"bbb",
	}
	if ArrayContainsString(array, "ccc") != false {
		t.FailNow()
	}

	array = []string{
		"aaa",
		"bbb",
		"ccc",
	}
	if ArrayContainsString(array, "bbb") != true {
		t.FailNow()
	}
	if IndexOfString("bbb", array) != 1 {
		t.FailNow()
	}
	if IndexOfString("zzz", array) != IndexNotFound {
		t.FailNow()
	}
}
