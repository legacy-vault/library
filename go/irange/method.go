// method.go.

package irange

import "fmt"

// A Range of Integer Values :: Methods.

// A Collection of main Methods used with the Range Class.

// Checks whether a Range is empty or not.
// This Function ignores the Validity Check.
func (this Range) IsEmpty() bool {

	return !(this.isFilled)
}

// Checks whether a Range equals to another Range.
// This Function ignores the Validity Check.
func (this Range) IsEqualTo(
	that Range,
) bool {

	return AreEqual(this, that)
}

// Checks whether a Range is empty or not.
// This Function ignores the Validity Check.
func (this Range) IsNotEmpty() bool {

	return this.isFilled
}

// Checks the Integrity of a Range.
// This Function performs the Validity Check.
func (this Range) IsValid() bool {

	if this.IsEmpty() {

		return true
	}

	if this.left > this.right {

		return false
	}

	return true
}

// Returns the left Corner Item of the Range.
// This Function performs the Validity Check.
func (this Range) Left() (int, error) {

	var err error

	if !(this.IsValid()) {
		err = fmt.Errorf(
			ErrMsgFormatRangeIsNotValid,
			this,
		)

		return ItemOnError, err
	}

	if this.IsEmpty() {
		err = fmt.Errorf(
			ErrMsgRangeIsEmpty,
			this,
		)

		return ItemOnError, err
	}

	return this.left, nil
}

// Returns the right Corner Item of the Range.
// This Function performs the Validity Check.
func (this Range) Right() (int, error) {

	var err error

	if !(this.IsValid()) {
		err = fmt.Errorf(
			ErrMsgFormatRangeIsNotValid,
			this,
		)

		return ItemOnError, err
	}

	if this.IsEmpty() {
		err = fmt.Errorf(
			ErrMsgRangeIsEmpty,
			this,
		)
		return ItemOnError, err
	}

	return this.right, nil
}

// Returns a Range with Corner Items swapped.
// This Function ignores the Validity Check.
func (this Range) Rotated() Range {

	return Range{
		left:     this.right,
		right:    this.left,
		isFilled: this.isFilled,
	}
}

// Returns a valid Range with Corner Items swapped if necessary.
// This Function performs the Validity Check.
func (this Range) RotatedIfNecessary() Range {

	if this.IsValid() {

		return this
	}

	return this.Rotated()
}

// Returns the Number of Items in the valid Range.
// This Function performs the Validity Check.
func (this Range) Size() (int, error) {

	var err error

	if !(this.IsValid()) {
		err = fmt.Errorf(
			ErrMsgFormatRangeIsNotValid,
			this,
		)

		return SizeOnError, err
	}

	if this.IsEmpty() {

		return 0, nil
	}

	return this.right - this.left + 1, nil
}
