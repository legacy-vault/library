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

// File and Directory :: Usual Functions.

// Last Update Time: 2018-10-30.

package file

import (
	"io/ioutil"
	"os"
	"path"
)

// Lists all Files inside the Directory.
func ListFiles(folderPath string, goSubLevels bool) []string {

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
			if goSubLevels {
				// Check Sub-Levels.
				subPath = path.Join(folderPath, item.Name())
				subFiles = ListFiles(subPath, goSubLevels)
				files = append(files, subFiles...)
			}

		} else {

			// Not a Directory.
			filePath = path.Join(folderPath, item.Name())
			files = append(files, filePath)
		}
	}

	return files
}

// Lists all Directories inside the Directory.
func ListFolders(folderPath string, goSubLevels bool) []string {

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
		folders = append(folders, itemPath)

		// Directory.
		if goSubLevels {
			// Check Sub-Levels.
			subFolders = ListFolders(itemPath, goSubLevels)
			folders = append(folders, subFolders...)
		}
	}

	return folders
}

// Lists all Files and Folders inside the Directory.
func ListFilesAndFolders(folderPath string, goSubLevels bool) []string {

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

		// Add to List.
		itemPath = path.Join(folderPath, item.Name())
		filesAndFolders = append(filesAndFolders, itemPath)

		// Directory?
		if item.IsDir() {

			// Directory.
			if goSubLevels {
				// Check Sub-Levels.
				subPath = path.Join(folderPath, item.Name())
				subFilesAndFolders = ListFilesAndFolders(subPath, goSubLevels)
				filesAndFolders = append(filesAndFolders, subFilesAndFolders...)
			}

		}
	}

	return filesAndFolders
}

// Lists Files inside the Directory and filters them by their Extension.
// Only Files with allowed Extensions will be returned.
// Extensions must be provided without '.' Dot Symbol.
func ListFilesExtAllowed(
	folderPath string,
	goSubLevels bool,
	extensionsAllowed []string,
) []string {

	// Fool Check.
	if len(extensionsAllowed) == 0 {
		return nil
	}

	// Prepare Data.
	for i, ext := range extensionsAllowed {
		extensionsAllowed[i] = "." + ext
	}

	return listFilesExtAllowed(
		folderPath,
		goSubLevels,
		extensionsAllowed,
	)
}

// Lists Files inside the Directory and filters them by their Extension.
// Only Files with allowed Extensions will be returned.
// Extensions must be provided in Golang Format (with '.' Dot Symbol).
func listFilesExtAllowed(
	folderPath string,
	goSubLevels bool,
	extensionsAllowed []string,
) []string {

	var err error
	var ext string
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
			if goSubLevels {
				// Check Sub-Levels.
				subPath = path.Join(folderPath, item.Name())
				subFiles = listFilesExtAllowed(
					subPath,
					goSubLevels,
					extensionsAllowed,
				)
				files = append(files, subFiles...)
			}

		} else {

			// Not a Directory.

			// Check Extension.
			ext = path.Ext(item.Name())
			if !existsIn(ext, extensionsAllowed) {
				continue
			}
			filePath = path.Join(folderPath, item.Name())
			files = append(files, filePath)
		}
	}

	return files
}

// Lists Files inside the Directory and filters them by their Extension.
// Only Files with not-forbidden Extensions will be returned.
// Extensions must be provided without '.' Dot Symbol.
func ListFilesExtForbidden(
	folderPath string,
	goSubLevels bool,
	extensionsForbidden []string,
) []string {

	// Fool Check.
	if len(extensionsForbidden) == 0 {
		return nil
	}

	// Prepare Data.
	for i, ext := range extensionsForbidden {
		extensionsForbidden[i] = "." + ext
	}

	return listFilesExtForbidden(
		folderPath,
		goSubLevels,
		extensionsForbidden,
	)
}

// Lists Files inside the Directory and filters them by their Extension.
// Only Files with not-forbidden Extensions will be returned.
// Extensions must be provided in Golang Format (with '.' Dot Symbol).
func listFilesExtForbidden(
	folderPath string,
	goSubLevels bool,
	extensionsForbidden []string,
) []string {

	var err error
	var ext string
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
			if goSubLevels {
				// Check Sub-Levels.
				subPath = path.Join(folderPath, item.Name())
				subFiles = listFilesExtForbidden(
					subPath,
					goSubLevels,
					extensionsForbidden,
				)
				files = append(files, subFiles...)
			}

		} else {

			// Not a Directory.

			// Check Extension.
			ext = path.Ext(item.Name())
			if existsIn(ext, extensionsForbidden) {
				continue
			}
			filePath = path.Join(folderPath, item.Name())
			files = append(files, filePath)
		}
	}

	return files
}
