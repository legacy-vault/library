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

// level-file&folder.go.

// File and Folder :: Usual Functions with Level Restrictions.

package file

import (
	"io/ioutil"
	"os"
	"path"
)

// Lists all Files and Folders inside the Directory.
// Level Restrictions are used.
// Criteria Type is 'Less or Equal'.
func listFilesAndFoldersLRLE(
	folderPath string,
	criterionType byte,
	criterionValue uint64,
	currentLevel uint64,
) []string {

	var err error
	var filesAndFolders []string
	var item os.FileInfo
	var itemPath string
	var items []os.FileInfo
	var subFilesAndFolders []string
	var subPath string

	// Prepare Data.
	filesAndFolders = []string{}

	// Read one Directory.
	items, err = ioutil.ReadDir(folderPath)
	if err != nil {
		return nil
	}

	// Check Items.
	for _, item = range items {

		itemPath = path.Join(folderPath, item.Name())

		// Analyze the Criterion.
		if currentLevel <= criterionValue {

			// Add to List.
			filesAndFolders = append(filesAndFolders, itemPath)
		}

		// Directory?
		if item.IsDir() {

			// Directory.

			// Analyze the Criterion.
			if currentLevel >= criterionValue {
				continue
			}

			// Check Sub-Levels.
			subPath = path.Join(folderPath, item.Name())
			subFilesAndFolders = listFilesAndFoldersLRLE(
				subPath,
				criterionType,
				criterionValue,
				currentLevel+1,
			)
			filesAndFolders = append(filesAndFolders, subFilesAndFolders...)
		}
	}

	return filesAndFolders
}

// Lists all Files and Folders inside the Directory.
// Level Restrictions are used.
// Criteria Type is 'Equal'.
func listFilesAndFoldersLREQ(
	folderPath string,
	criterionType byte,
	criterionValue uint64,
	currentLevel uint64,
) []string {

	var err error
	var filesAndFolders []string
	var item os.FileInfo
	var itemPath string
	var items []os.FileInfo
	var subFilesAndFolders []string
	var subPath string

	// Prepare Data.
	filesAndFolders = []string{}

	// Read one Directory.
	items, err = ioutil.ReadDir(folderPath)
	if err != nil {
		return nil
	}

	// Check Items.
	for _, item = range items {

		itemPath = path.Join(folderPath, item.Name())

		// Analyze the Criterion.
		if currentLevel == criterionValue {

			// Add to List.
			filesAndFolders = append(filesAndFolders, itemPath)
		}

		// Directory?
		if item.IsDir() {

			// Directory.

			// Analyze the Criterion.
			if currentLevel >= criterionValue {
				continue
			}

			// Check Sub-Levels.
			subPath = path.Join(folderPath, item.Name())
			subFilesAndFolders = listFilesAndFoldersLREQ(
				subPath,
				criterionType,
				criterionValue,
				currentLevel+1,
			)
			filesAndFolders = append(filesAndFolders, subFilesAndFolders...)
		}
	}

	return filesAndFolders
}

// Lists all Files and Folders inside the Directory.
// Level Restrictions are used.
// Criteria Type is 'Greater or Equal'.
func listFilesAndFoldersLRGE(
	folderPath string,
	criterionType byte,
	criterionValue uint64,
	currentLevel uint64,
) []string {

	var err error
	var filesAndFolders []string
	var item os.FileInfo
	var itemPath string
	var items []os.FileInfo
	var subFilesAndFolders []string
	var subPath string

	// Prepare Data.
	filesAndFolders = []string{}

	// Read one Directory.
	items, err = ioutil.ReadDir(folderPath)
	if err != nil {
		return nil
	}

	// Check Items.
	for _, item = range items {

		itemPath = path.Join(folderPath, item.Name())

		// Analyze the Criterion.
		if currentLevel >= criterionValue {

			// Add to List.
			filesAndFolders = append(filesAndFolders, itemPath)
		}

		// Directory?
		if item.IsDir() {

			// Directory.

			// Check Sub-Levels.
			subPath = path.Join(folderPath, item.Name())
			subFilesAndFolders = listFilesAndFoldersLRGE(
				subPath,
				criterionType,
				criterionValue,
				currentLevel+1,
			)
			filesAndFolders = append(filesAndFolders, subFilesAndFolders...)
		}
	}

	return filesAndFolders
}
