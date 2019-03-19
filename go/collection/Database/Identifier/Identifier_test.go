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

// Identifier_test.go.

package identifier

// Database Identifier.

import (
	"testing"
)

func Test_DatabaseIdentifier_IsGood(
	t *testing.T,
) {

	var id Identifier

	id.Name = ""
	if id.IsGood(false) != false {
		t.FailNow()
	}

	id.Name = "_"
	if id.IsGood(false) != false {
		t.FailNow()
	}

	id.Name = "_xyz"
	if id.IsGood(false) != false {
		t.FailNow()
	}

	id.Name = "5"
	if id.IsGood(false) != false {
		t.FailNow()
	}

	id.Name = "6abc"
	if id.IsGood(false) != false {
		t.FailNow()
	}

	id.Name = "size_"
	if id.IsGood(false) != false {
		t.FailNow()
	}

	id.Name = "size123_xyz8"
	if id.IsGood(false) != true {
		t.FailNow()
	}

	id.Name = "size123_xyz8"
	if id.IsGood(true) != true {
		t.FailNow()
	}

	id.Name = "_size123_xyz8"
	if id.IsGood(true) != true {
		t.FailNow()
	}

	id.Name = "__size123_xyz8"
	if id.IsGood(true) != true {
		t.FailNow()
	}
}
