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
// Creation Date:	2018-10-26.
// Web Site Address is an Address in the global Computer Internet Network.
//
//============================================================================//

// file.go.

// File Functions.

package sfstorage

import (
	"os"
)

const NewFolderMode = 0755

// Checks whether the speciofied File exists in the Storage.
// Relative Path is the Path relative to the Storage Root Folder.
func FileExists(relPath string) bool {

	var err error

	// Get File Information from O.S.
	_, err = os.Stat(relPath)
	if err != nil {
		// If Error occurs, we consider that
		// the File is not accessible.
		// The Reason is not important for us here now.
		return false
	}

	return true
}

// Creates the Folder.
func CreateFolder(path string) error {

	var err error

	// Try to create the Folder with all its Parents.
	err = os.MkdirAll(path, NewFolderMode)
	if err != nil {
		return err
	}

	return nil
}
