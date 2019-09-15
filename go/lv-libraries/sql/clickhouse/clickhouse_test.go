// +build test

package clickhouse

import (
	"testing"
)

func Test_MakeClickhouseDsn(t *testing.T) {

	var dsnExpected string
	var dsnReceived string

	// Test #1. Full String.
	dsnExpected = "http://vasya:pwd@localhost:1234/vasyadb?xyz=123"
	dsnReceived = MakeClickhouseDsn(
		"http",
		"localhost",
		"1234",
		"vasyadb",
		"vasya",
		"pwd",
		"xyz=123",
	)
	if dsnExpected != dsnReceived {
		t.Error("Full String")
		t.FailNow()
	}

	// Test #2. No Password.
	dsnExpected = "http://vasya@localhost:1234/vasyadb?xyz=123"
	dsnReceived = MakeClickhouseDsn(
		"http",
		"localhost",
		"1234",
		"vasyadb",
		"vasya",
		"",
		"xyz=123",
	)
	if dsnExpected != dsnReceived {
		t.Error("No Password")
		t.FailNow()
	}

	// Test #3. No Username.
	dsnExpected = "http://localhost:1234/vasyadb?xyz=123"
	dsnReceived = MakeClickhouseDsn(
		"http",
		"localhost",
		"1234",
		"vasyadb",
		"",
		"password-not-used",
		"xyz=123",
	)
	if dsnExpected != dsnReceived {
		t.Error("No Username")
		t.FailNow()
	}

	// Test #4. No Database.
	dsnExpected = "http://localhost:1234?xyz=123"
	dsnReceived = MakeClickhouseDsn(
		"http",
		"localhost",
		"1234",
		"",
		"",
		"password-not-used",
		"xyz=123",
	)
	if dsnExpected != dsnReceived {
		t.Error("No Database")
		t.FailNow()
	}

	// Test #5. No Parameters.
	dsnExpected = "http://localhost:1234"
	dsnReceived = MakeClickhouseDsn(
		"http",
		"localhost",
		"1234",
		"",
		"",
		"password-not-used",
		"",
	)
	if dsnExpected != dsnReceived {
		t.Error("No Parameters")
		t.FailNow()
	}
}
