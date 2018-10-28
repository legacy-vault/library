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

// file.go.

// File Functions.

package bencode

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"time"
)

// Parses an input File into an Interface.
// Also stores some additional Data, all packed into an Object.
func ParseFile(filePath string) (*DecodedObject, error) {

	var bufioReader *bufio.Reader
	var err error
	var file *os.File
	var fileContents []byte
	var ifc interface{}
	var obj *DecodedObject
	var ok bool
	var reader io.Reader

	// Open the File and prepare a Stream Reader.
	file, err = os.Open(filePath)
	if err != nil {
		return nil, err
	}
	reader = bufio.NewReader(file)
	bufioReader = bufio.NewReader(reader)

	// Parse the File encoded with 'bencode' Encoding into an Object.
	ifc, err = getBencodedValue(bufioReader)
	if err != nil {
		return nil, err
	}

	// Get File Contents.
	_, err = file.Seek(0, 0)
	if err != nil {
		return nil, err
	}
	fileContents, err = ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	// Close the File.
	err = file.Close()
	if err != nil {
		return nil, err
	}

	// Prepare Result.
	obj = &DecodedObject{
		FilePath:        filePath,
		DecodeTimestamp: time.Now().Unix(),
		SourceData:      fileContents,
		DecodedObject:   ifc,
	}

	// Perform a Self-Check.
	ok = obj.SelfCheck()
	if !ok {
		err = ErrSelfCheck
		return nil, err
	}

	// Calculate BTIH.
	err = obj.CalculateBTIH()
	if err != nil {
		return nil, err
	}

	return obj, nil
}
