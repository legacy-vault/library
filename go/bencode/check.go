//============================================================================//
//
// Copyright © 2018 by McArcher.
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
//============================================================================//
//
// Web Site:		'https://github.com/legacy-vault'.
// Author:			McArcher.
// Creation Date:	2018-10-29.
// Web Site Address is an Address in the global Computer Internet Network.
//
//============================================================================//

// check.go.

// Self-Check.

// Last Update Time: 2018-10-29.

package bencode

import "bytes"

// A simple Self-Check.
// Encodes the decoded Data and compares with the Source.
func (do *DecodedObject) SelfCheck() bool {

	var baEncoded []byte
	var checkResult int
	var err error

	// Encode decoded Data.
	baEncoded, err = Encode(do.DecodedObject)
	if err != nil {
		return false
	}

	// Compare encoded decoded Data with original Data.
	checkResult = bytes.Compare(baEncoded, do.SourceData)
	if checkResult != 0 {
		return false
	}

	do.SelfChecked = true

	return true
}
