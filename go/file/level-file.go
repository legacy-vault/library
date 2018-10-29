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

// level-file.go.

// File :: Usual Functions with Level Restrictions.

package file

import (
	"io/ioutil"
	"os"
	"path"
)

// Lists all Files inside the Directory.
// Level Restrictions are used.
// Criteria Type is 'Less or Equal'.
func listFilesLRLE(
	folderPath string,
	criterionType byte,
	criterionValue uint64,
	currentLevel uint64,
) []string {

	var err error
	var filePath string
	var files []string
	var item os.FileInfo
	var items []os.FileInfo
	var subFiles []string
	var subPath string

	// Prepare Data.
	files = []string{}

	// Read one Directory.
	items, err = ioutil.ReadDir(folderPath)
	if err != nil {
		return nil
	}

	// Check Items.
	for _, item = range items {

		// Directory?
		if item.IsDir() {

			// Directory.

			// Analyze the Criterion.
			if currentLevel >= criterionValue {
				continue
			}

			// Check Sub-Levels.
			subPath = path.Join(folderPath, item.Name())
			subFiles = listFilesLRLE(
				subPath,
				criterionType,
				criterionValue,
				currentLevel+1,
			)
			files = append(files, subFiles...)

		} else {

			// Not a Directory.

			// Analyze the Criterion.
			if currentLevel > criterionValue {
				continue
			}

			filePath = path.Join(folderPath, item.Name())
			files = append(files, filePath)
		}
	}

	return files
}

// Lists all Files inside the Directory.
// Level Restrictions are used.
// Criteria Type is 'Equal'.
func listFilesLREQ(
	folderPath string,
	criterionType byte,
	criterionValue uint64,
	currentLevel uint64,
) []string {

	var err error
	var filePath string
	var files []string
	var item os.FileInfo
	var items []os.FileInfo
	var subFiles []string
	var subPath string

	// Prepare Data.
	files = []string{}

	// Read one Directory.
	items, err = ioutil.ReadDir(folderPath)
	if err != nil {
		return nil
	}

	// Check Items.
	for _, item = range items {

		// Directory?
		if item.IsDir() {

			// Directory.

			// Analyze the Criterion.
			if currentLevel >= criterionValue {
				continue
			}

			// Check Sub-Levels.
			subPath = path.Join(folderPath, item.Name())
			subFiles = listFilesLREQ(
				subPath,
				criterionType,
				criterionValue,
				currentLevel+1,
			)
			files = append(files, subFiles...)

		} else {

			// Not a Directory.

			// Analyze the Criterion.
			if currentLevel != criterionValue {
				continue
			}

			filePath = path.Join(folderPath, item.Name())
			files = append(files, filePath)
		}
	}

	return files
}

// Lists all Files inside the Directory.
// Level Restrictions are used.
// Criteria Type is 'Greater or Equal'.
func listFilesLRGE(
	folderPath string,
	criterionType byte,
	criterionValue uint64,
	currentLevel uint64,
) []string {

	var err error
	var filePath string
	var files []string
	var item os.FileInfo
	var items []os.FileInfo
	var subFiles []string
	var subPath string

	// Prepare Data.
	files = []string{}

	// Read one Directory.
	items, err = ioutil.ReadDir(folderPath)
	if err != nil {
		return nil
	}

	// Check Items.
	for _, item = range items {

		// Directory?
		if item.IsDir() {

			// Directory.

			// Check Sub-Levels.
			subPath = path.Join(folderPath, item.Name())
			subFiles = listFilesLRGE(
				subPath,
				criterionType,
				criterionValue,
				currentLevel+1,
			)
			files = append(files, subFiles...)

		} else {

			// Not a Directory.

			// Analyze the Criterion.
			if currentLevel < criterionValue {
				continue
			}

			filePath = path.Join(folderPath, item.Name())
			files = append(files, filePath)
		}
	}

	return files
}
