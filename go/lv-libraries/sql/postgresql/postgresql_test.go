// +build test

package postgresql

import (
	"testing"
)

func Test_MakePostgresqlDsn(t *testing.T) {

	var dsnExpected string
	var dsnReceived string

	// Test #1. Full String.
	dsnExpected = "postgresql://vasya:pwd@localhost:1234/vasyadb?sslmode=allow"
	dsnReceived = MakePostgresqlDsn(
		"localhost",
		"1234",
		"vasyadb",
		"vasya",
		"pwd",
		"sslmode=allow",
	)
	if dsnExpected != dsnReceived {
		t.Error("Full String")
		t.FailNow()
	}

	// Test #2. No Password.
	dsnExpected = "postgresql://vasya@localhost:1234/vasyadb?sslmode=allow"
	dsnReceived = MakePostgresqlDsn(
		"localhost",
		"1234",
		"vasyadb",
		"vasya",
		"",
		"sslmode=allow",
	)
	if dsnExpected != dsnReceived {
		t.Error("No Password")
		t.FailNow()
	}

	// Test #3. No Username.
	dsnExpected = "postgresql://localhost:1234/vasyadb?sslmode=allow"
	dsnReceived = MakePostgresqlDsn(
		"localhost",
		"1234",
		"vasyadb",
		"",
		"password-not-used",
		"sslmode=allow",
	)
	if dsnExpected != dsnReceived {
		t.Error("No Username")
		t.FailNow()
	}

	// Test #4. No Database.
	dsnExpected = "postgresql://localhost:1234?sslmode=allow"
	dsnReceived = MakePostgresqlDsn(
		"localhost",
		"1234",
		"",
		"",
		"password-not-used",
		"sslmode=allow",
	)
	if dsnExpected != dsnReceived {
		t.Error("No Database")
		t.FailNow()
	}

	// Test #5. No Parameters.
	dsnExpected = "postgresql://localhost:1234"
	dsnReceived = MakePostgresqlDsn(
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
