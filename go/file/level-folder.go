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

// level-folder.go.

// Folder :: Usual Functions with Level Restrictions.

package file

import (
	"io/ioutil"
	"os"
	"path"
)

// Lists all Directories inside the Directory.
// Level Restrictions are used.
// Criteria Type is 'Less or Equal'.
func listFoldersLRLE(
	folderPath string,
	criterionType byte,
	criterionValue uint64,
	currentLevel uint64,
) []string {

	var err error
	var folders []string
	var item os.FileInfo
	var itemPath string
	var items []os.FileInfo
	var subFolders []string

	// Prepare Data.
	folders = []string{}

	// Read one Directory.
	items, err = ioutil.ReadDir(folderPath)
	if err != nil {
		return nil
	}

	// Check Items.
	for _, item = range items {

		// Directory?
		if !item.IsDir() {
			continue
		}

		itemPath = path.Join(folderPath, item.Name())

		// Analyze the Criterion.
		if currentLevel <= criterionValue {

			// Save Directory.
			folders = append(folders, itemPath)
		}

		// Analyze the Criterion.
		if currentLevel >= criterionValue {
			continue
		}

		// Check Sub-Levels.
		subFolders = listFoldersLRLE(
			itemPath,
			criterionType,
			criterionValue,
			currentLevel+1,
		)
		folders = append(folders, subFolders...)
	}

	return folders
}

// Lists all Directories inside the Directory.
// Level Restrictions are used.
// Criteria Type is 'Equal'.
func listFoldersLREQ(
	folderPath string,
	criterionType byte,
	criterionValue uint64,
	currentLevel uint64,
) []string {

	var err error
	var folders []string
	var item os.FileInfo
	var itemPath string
	var items []os.FileInfo
	var subFolders []string

	// Prepare Data.
	folders = []string{}

	// Read one Directory.
	items, err = ioutil.ReadDir(folderPath)
	if err != nil {
		return nil
	}

	// Check Items.
	for _, item = range items {

		// Directory?
		if !item.IsDir() {
			continue
		}

		itemPath = path.Join(folderPath, item.Name())

		// Analyze the Criterion.
		if currentLevel == criterionValue {

			// Save Directory.
			folders = append(folders, itemPath)
		}

		// Analyze the Criterion.
		if currentLevel >= criterionValue {
			continue
		}

		// Check Sub-Levels.
		subFolders = listFoldersLREQ(
			itemPath,
			criterionType,
			criterionValue,
			currentLevel+1,
		)
		folders = append(folders, subFolders...)
	}

	return folders
}

// Lists all Directories inside the Directory.
// Level Restrictions are used.
// Criteria Type is 'Greater or Equal'.
func listFoldersLRGE(
	folderPath string,
	criterionType byte,
	criterionValue uint64,
	currentLevel uint64,
) []string {

	var err error
	var folders []string
	var item os.FileInfo
	var itemPath string
	var items []os.FileInfo
	var subFolders []string

	// Prepare Data.
	folders = []string{}

	// Read one Directory.
	items, err = ioutil.ReadDir(folderPath)
	if err != nil {
		return nil
	}

	// Check Items.
	for _, item = range items {

		// Directory?
		if !item.IsDir() {
			continue
		}

		itemPath = path.Join(folderPath, item.Name())

		// Analyze the Criterion.
		if currentLevel >= criterionValue {

			// Save Directory.
			folders = append(folders, itemPath)
		}

		// Check Sub-Levels.
		subFolders = listFoldersLRGE(
			itemPath,
			criterionType,
			criterionValue,
			currentLevel+1,
		)
		folders = append(folders, subFolders...)
	}

	return folders
}
