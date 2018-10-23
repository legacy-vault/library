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
// Creation Date:	2018-10-23.
// Web Site Address is an Address in the global Computer Internet Network.
//
//============================================================================//

// record.go.

// Cache Record Functions.

package cache

import "time"

type RecordUID = string

type Record struct {
	UID  RecordUID
	Data interface{}

	lastAccessTime int64

	nextItem     *Record
	previousItem *Record
}

// Updates Record's Data.
func (r *Record) UpdateData(data interface{}) {

	r.Data = data
	r.lastAccessTime = time.Now().Unix()
}

// Updates Record's Last Access Time.
func (r *Record) UpdateLAT() {

	r.lastAccessTime = time.Now().Unix()
}
