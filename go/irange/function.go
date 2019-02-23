// function.go.

package irange

import "fmt"

// A Range of Integer Values :: Functions.

// A Collection of main Functions used with the Range Class.

// Checks whether two Ranges are equal.
// This Function ignores the Validity Check.
func AreEqual(
	a Range,
	b Range,
) bool {

	if a.isFilled != b.isFilled {

		return false
	}
	if a.left != b.left {

		return false
	}
	if a.right != b.right {

		return false
	}

	return true
}

// Returns an Intersection of two Ranges.
// This Function performs the Validity Check.
func Intersection(
	a Range,
	b Range,
) (Range, error) {

	var err error
	var result Range

	// Check Input Data.
	if !(a.IsValid()) {
		err = fmt.Errorf(
			ErrMsgFormatRangeIsNotValid,
			a,
		)

		return result, err
	}
	if !(b.IsValid()) {
		err = fmt.Errorf(
			ErrMsgFormatRangeIsNotValid,
			b,
		)

		return result, err
	}

	// Analyze.
	if a.left < b.left {

		if a.right < b.left {

			// Empty Result.
			return result, nil

		} else if a.right == b.left {

			// Single Point Result.
			result.left = a.right
			result.right = b.left
			result.isFilled = true

			return result, nil

		} else { // a.right > b.left

			// Compare 'a.right' with 'b.right'.
			if a.right < b.right {

				result.left = b.left
				result.right = a.right
				result.isFilled = true

				return result, nil

			} else if a.right == b.right {

				result.left = b.left
				result.right = a.right
				result.isFilled = true

				return result, nil

			} else { // a.right > b.right

				result.left = b.left
				result.right = b.right
				result.isFilled = true

				return result, nil
			}
		}

	} else if a.left == b.left {

		if b.right <= a.right {

			return b, nil

		} else {

			return a, nil
		}

	} else { // a.left > b.left

		if b.right < a.left {

			// Empty Result.
			return result, nil

		} else if b.right == a.left {

			// Single Point Result.
			result.left = b.right
			result.right = a.left
			result.isFilled = true

			return result, nil

		} else { // b.right > a.left

			// Compare 'b.right' with 'a.right'.
			if b.right < a.right {

				result.left = a.left
				result.right = b.right
				result.isFilled = true

				return result, nil

			} else if b.right == a.right {

				result.left = a.left
				result.right = b.right
				result.isFilled = true

				return result, nil

			} else { // b.right > a.right

				result.left = a.left
				result.right = a.right
				result.isFilled = true

				return result, nil
			}
		}
	}
}

// Creates a new empty Range.
// This Function does not need the Validity Check.
func NewEmpty() Range {

	var result Range

	result.left = NewEmptyLeftCorner
	result.right = NewEmptyRightCorner
	result.isFilled = false

	return result
}

// Creates a new Range.
// This Function performs the Validity Check.
func New(
	left int,
	right int,
) Range {

	var result Range

	result.left = left
	result.right = right
	result.isFilled = true
	result = result.RotatedIfNecessary()

	return result
}
