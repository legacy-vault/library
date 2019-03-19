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

// Errorz.go.

package errorz

// Functions for Errors Processing.

import "errors"

const TextError = "Error"
const TextDelimiter = " "

const CombineSeparator = ";" + TextDelimiter
const ReportSeparator = TextDelimiter + ":" + TextDelimiter
const ReportBorderLeft = "[" + TextDelimiter
const ReportBorderRight = TextDelimiter + "]"
const ReporterBorderLeft = "'"
const ReporterBorderRight = "'"

// Combines two Errors.
func Combine(
	a error,
	b error,
) error {

	var result error

	result = errors.New(a.Error() + CombineSeparator + b.Error())

	return result
}

// Composes an Error Report.
func Report(
	reporter string,
	report error,
) error {

	var msg string
	var result error

	msg = ReportBorderLeft +
		ReporterBorderLeft + reporter + ReporterBorderRight +
		TextDelimiter + TextError + ReportSeparator +
		report.Error() +
		ReportBorderRight

	result = errors.New(msg)

	return result
}
