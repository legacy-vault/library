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

// close.go.

package uc

// User Collection Methods for Finalization.

import (
	"errors"

	"github.com/legacy-vault/library/go/collection/Common"
	"github.com/legacy-vault/library/go/collection/Errorz"
)

// Closes the User's Collection.
func (this *Uc) Close() error {

	var err error

	// Fool's Check.
	if this == nil {
		err = errors.New(common.ErrMethodOwnerDoesNotExist)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Check Status.
	if this.isInitialized() != true {
		err = errors.New(ErrNotInitialized)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	// Close.
	err = this.close()
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	return nil
}

// Closes the User's Collection.
func (this *Uc) close() error {

	var err error

	// Disconnect.
	err = this.disconnect()
	if err != nil {
		return err
	}

	// Set Status.
	this.isOpen = false

	// Check Status.
	if this.IsOpened() != false {
		err = errors.New(ErrNotClosed)
		return err
	}

	return nil
}

// Disconnects from the User's Collection's Database.
func (this *Uc) disconnect() error {

	var err error

	// Disconnect.
	err = this.storage.Disconnect()
	if err != nil {
		return err
	}

	return nil
}
