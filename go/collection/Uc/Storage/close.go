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

// Storage.go.

package storage

// Storage Finalization.

import (
	"errors"

	"github.com/legacy-vault/library/go/collection/Errorz"
)

// Disconnects from the Collection's Database.
func (this *Storage) Disconnect() error {

	var err error

	// Disconnect.
	err = this.database.Disconnect()
	if err != nil {
		err = errorz.Report(ErrorReporter, err)
		return err
	}
	if this.database.IsConnected() != false {
		err = errors.New(ErrDisconnection)
		err = errorz.Report(ErrorReporter, err)
		return err
	}

	return nil
}
