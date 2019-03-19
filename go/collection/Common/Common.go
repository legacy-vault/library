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

// Common.go.

package common

// Common Things used in all Packages.

import (
	"reflect"
)

const ErrorReporter = "Common"

const ErrMethodOwnerDoesNotExist = "Method Owner does not exist"
const ErrMethodNotImplemented = "Method is not implemented"

const ErrFormatUpdateIdMismatch = "The provided ID " +
	"does not match the provided Entity's ID ('%v' vs '%v')."

// Converts an empty Interface containing an Array of Bytes taken from a
// Response from an SQL Server into a real Array of Bytes.
func SqlResponseInterfaceToBytes(
	ifc interface{},
) ([]byte, error) {

	var result []byte

	result = reflect.ValueOf(ifc).Bytes()

	return result, nil
}
