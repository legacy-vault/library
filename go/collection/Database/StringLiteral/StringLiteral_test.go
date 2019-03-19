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

// StringLiteral.go.

package stringliteral

// String Literal.

import (
	"fmt"
	"testing"
)

func Test_All(t *testing.T) {

	var sl StringLiteral

	sl = StringLiteral{
		Value: "Hello World",
	}
	fmt.Println(sl.Escaped())
	fmt.Println(sl.EscapedAndSingleQuoted())

	sl = StringLiteral{
		Value: `He said: "Hello!"`,
	}
	fmt.Println(sl.Escaped())
	fmt.Println(sl.EscapedAndSingleQuoted())

	sl = StringLiteral{
		Value: `a'b''c'''d`,
	}
	fmt.Println(sl.Escaped())
	fmt.Println(sl.EscapedAndSingleQuoted())

	sl = StringLiteral{
		Value: `c:\Temp\cat.png`,
	}
	fmt.Println(sl.Escaped())
	fmt.Println(sl.EscapedAndSingleQuoted())

	sl = StringLiteral{
		Value: `/etc/fstab`,
	}
	fmt.Println(sl.Escaped())
	fmt.Println(sl.EscapedAndSingleQuoted())

	sl = StringLiteral{
		Value: `100%, %%%, %%, \%`,
	}
	fmt.Println(sl.Escaped())
	fmt.Println(sl.EscapedAndSingleQuoted())
}
