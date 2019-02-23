// class.go.

package irange

// A Range of Integer Values :: Class Declaration.

// A Range of Integers is either a Range of existing Integers (e.g. [1..5]) or
// an empty Range (an empty Collection of Integers).

////////////////////////////////////////////////////////////////////////////////
//
// N.B.
//
// The Reader of this Package must forgive me my Habit taken from other (more
// mature) Computer Programming Languages which I am used to. Though 'Go'
// Language has no Classes and does not allow Allman's Coding Style, am I a
// Protector of normal human-friendly Coding Style rather than Golang's
// "IDONTCARE" Style.
//
////////////////////////////////////////////////////////////////////////////////

type Range struct {

	// A Flag showing whether the Range is filled (not empty) or not.
	// This Parameter has a stronger Priority than other Parameters when reading
	// a Range.
	isFilled bool

	// Corners Items of the Range.
	// The left Corner has a lesser Value than the Value of the right Corner.
	// The right Corner has a greater Value than the Value of the left Corner.
	left  int
	right int
}
