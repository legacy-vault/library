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

// edge.go.

// Functions for Edge Files and Edge Directories.

package file

import (
	"io/ioutil"
	"os"
	"path"
)

// From inside the Directory, lists all those Files which sit in Directories
// which have no Sub-Directories (i.e. are Edge Directories in the File System).
func ListEdgeFiles(folderPath string) []string {

	var edgeFiles []string
	var edgeFolder string
	var edgeFolders []string
	var files []string

	// List Edge Folders.
	edgeFolders = ListEdgeFolders(folderPath)
	edgeFiles = []string{}

	// Add Files from each Edge Folder.
	for _, edgeFolder = range edgeFolders {

		files = ListFiles(edgeFolder, true)
		edgeFiles = append(edgeFiles, files...)
	}

	return edgeFiles
}

// From inside the Directory, lists all those Directories
// which have no Sub-Directories (i.e. are Edge Directories in the File System).
func ListEdgeFolders(folderPath string) []string {

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

	// Edge?
	if SubFoldersCount(items) == 0 {
		folders = append(folders, folderPath)
		return folders
	}

	// Check Items.
	for _, item = range items {

		// Directory?
		if !item.IsDir() {
			continue
		}

		itemPath = path.Join(folderPath, item.Name())

		// Check Sub-Levels.
		subFolders = ListEdgeFolders(itemPath)
		folders = append(folders, subFolders...)
	}

	return folders
}
