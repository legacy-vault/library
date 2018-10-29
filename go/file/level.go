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

// usual.go.

// File and Directory :: Usual Functions with Level Restrictions.

package file

// Critreria Type of Level Restrictions.
const LevelRestrictionsCriterionLessOrEqual byte = 1
const LevelRestrictionsCriterionEqual byte = 2
const LevelRestrictionsCriterionGreaterOrEqual byte = 4

// Lists all Files inside the Directory.
// Level Restrictions are used.
func ListFilesLR(
	folderPath string,
	criterionType byte,
	criterionValue uint64,
) []string {

	// Check the Criterion.
	if !criterionTypeIsValid(criterionType) {
		return nil
	}

	return listFilesLR(
		folderPath,
		criterionType,
		criterionValue,
		1,
	)
}

// Lists all Files inside the Directory.
// Level Restrictions are used.
func listFilesLR(
	folderPath string,
	criterionType byte,
	criterionValue uint64,
	currentLevel uint64,
) []string {

	if criterionType == LevelRestrictionsCriterionLessOrEqual {
		return listFilesLRLE(
			folderPath,
			criterionType,
			criterionValue,
			currentLevel,
		)
	}

	if criterionType == LevelRestrictionsCriterionEqual {
		return listFilesLREQ(
			folderPath,
			criterionType,
			criterionValue,
			currentLevel,
		)
	}

	if criterionType == LevelRestrictionsCriterionGreaterOrEqual {
		return listFilesLRGE(
			folderPath,
			criterionType,
			criterionValue,
			currentLevel,
		)
	}

	return nil
}

// Lists all Folders inside the Directory.
// Level Restrictions are used.
func ListFoldersLR(
	folderPath string,
	criterionType byte,
	criterionValue uint64,
) []string {

	// Check the Criterion.
	if !criterionTypeIsValid(criterionType) {
		return nil
	}

	return listFoldersLR(
		folderPath,
		criterionType,
		criterionValue,
		1,
	)
}

// Lists all Folders inside the Directory.
// Level Restrictions are used.
func listFoldersLR(
	folderPath string,
	criterionType byte,
	criterionValue uint64,
	currentLevel uint64,
) []string {

	if criterionType == LevelRestrictionsCriterionLessOrEqual {
		return listFoldersLRLE(
			folderPath,
			criterionType,
			criterionValue,
			currentLevel,
		)
	}

	if criterionType == LevelRestrictionsCriterionEqual {
		return listFoldersLREQ(
			folderPath,
			criterionType,
			criterionValue,
			currentLevel,
		)
	}

	if criterionType == LevelRestrictionsCriterionGreaterOrEqual {
		return listFoldersLRGE(
			folderPath,
			criterionType,
			criterionValue,
			currentLevel,
		)
	}

	return nil
}

// Lists all Files and Folders inside the Directory.
// Level Restrictions are used.
func ListFilesAndFoldersLR(
	folderPath string,
	criterionType byte,
	criterionValue uint64,
) []string {

	// Check the Criterion.
	if !criterionTypeIsValid(criterionType) {
		return nil
	}

	return listFilesAndFoldersLR(
		folderPath,
		criterionType,
		criterionValue,
		1,
	)
}

// Lists all Files inside the Directory.
// Level Restrictions are used.
func listFilesAndFoldersLR(
	folderPath string,
	criterionType byte,
	criterionValue uint64,
	currentLevel uint64,
) []string {

	if criterionType == LevelRestrictionsCriterionLessOrEqual {
		return listFilesAndFoldersLRLE(
			folderPath,
			criterionType,
			criterionValue,
			currentLevel,
		)
	}

	if criterionType == LevelRestrictionsCriterionEqual {
		return listFilesAndFoldersLREQ(
			folderPath,
			criterionType,
			criterionValue,
			currentLevel,
		)
	}

	if criterionType == LevelRestrictionsCriterionGreaterOrEqual {
		return listFilesAndFoldersLRGE(
			folderPath,
			criterionType,
			criterionValue,
			currentLevel,
		)
	}

	return nil
}
