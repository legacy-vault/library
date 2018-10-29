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

// helper.go.

// Helper-Functions.

package file

import (
	"os"
)

// Counts the Number of Sub-Folders (Sub-Items of the 'Folder' Type).
func SubFoldersCount(items []os.FileInfo) int {

	var subFoldersCount int
	var item os.FileInfo

	for _, item = range items {

		// Directory?
		if !item.IsDir() {
			continue
		}

		subFoldersCount++
	}

	return subFoldersCount
}

// Checks whether the Criterion Type is valid.
func criterionTypeIsValid(criterionType byte) bool {

	if criterionType == LevelRestrictionsCriterionLessOrEqual {
		return true
	}

	if criterionType == LevelRestrictionsCriterionEqual {
		return true
	}

	if criterionType == LevelRestrictionsCriterionGreaterOrEqual {
		return true
	}

	return false
}
