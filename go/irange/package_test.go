// package_test.go.

package irange

// A Range of Integer Values :: Internal Test.

// An internal Test of the Package itself.

import (
	"fmt"
	"testing"
)

const ErrorReportPostfix = " Error"

func Test_AreEqual(
	t *testing.T,
) {
	const ErrMsg = "AreEqual" + ErrorReportPostfix

	var irA Range
	var irB Range

	// Prepare common Data.
	irA.isFilled = true
	irA.left = 456
	irA.right = 123

	// #1. Test equal Ranges.
	if AreEqual(irA, irA) != true {
		t.Error(ErrMsg)
		t.FailNow()
	}

	// #2. Test Ranges with different 'isFilled' Field.
	irB.isFilled = !irA.isFilled
	irB.left = irA.left
	irB.right = irA.right
	if AreEqual(irA, irB) != false {
		t.Error(ErrMsg)
		t.FailNow()
	}

	// #3. Test Ranges with different 'left' Field.
	irB.isFilled = irA.isFilled
	irB.left = irA.left + 1
	irB.right = irA.right
	if AreEqual(irA, irB) != false {
		t.Error(ErrMsg)
		t.FailNow()
	}

	// #4. Test Ranges with different 'right' Field.
	irB.isFilled = irA.isFilled
	irB.left = irA.left
	irB.right = irA.right + 1
	if AreEqual(irA, irB) != false {
		t.Error(ErrMsg)
		t.FailNow()
	}
}

func Test_Intersection(
	t *testing.T,
) {
	const ErrMsgPrefix = "Intersection" + ErrorReportPostfix

	const TestsCount = 15

	type TestWorkspace struct {

		// Test's Configuration.
		RangeA          Range
		RangeB          Range
		ResultExpected  Range
		ErrorIsExpected bool

		// Test's Runtime Parameters.
		ResultGot Range
		ErrorGot  error
		Report    error
	}

	var test TestWorkspace
	var tests []TestWorkspace

	tests = make([]TestWorkspace, 0, TestsCount)

	// Test #1.
	test = TestWorkspace{
		RangeA:          New(10, 11),
		RangeB:          New(20, 25),
		ResultExpected:  NewEmpty(),
		ErrorIsExpected: false,
	}
	tests = append(tests, test)

	// Test #2.
	test = TestWorkspace{
		RangeA:          New(10, 20),
		RangeB:          New(20, 25),
		ResultExpected:  New(20, 20),
		ErrorIsExpected: false,
	}
	tests = append(tests, test)

	// Test #3.
	test = TestWorkspace{
		RangeA:          New(10, 21),
		RangeB:          New(20, 25),
		ResultExpected:  New(20, 21),
		ErrorIsExpected: false,
	}
	tests = append(tests, test)

	// Test #4.
	test = TestWorkspace{
		RangeA:          New(10, 25),
		RangeB:          New(20, 25),
		ResultExpected:  New(20, 25),
		ErrorIsExpected: false,
	}
	tests = append(tests, test)

	// Test #5.
	test = TestWorkspace{
		RangeA:          New(10, 29),
		RangeB:          New(20, 25),
		ResultExpected:  New(20, 25),
		ErrorIsExpected: false,
	}
	tests = append(tests, test)

	// Test #6.
	test = TestWorkspace{
		RangeA:          New(30, 31),
		RangeB:          New(30, 32),
		ResultExpected:  New(30, 31),
		ErrorIsExpected: false,
	}
	tests = append(tests, test)

	// Test #7.
	test = TestWorkspace{
		RangeA:          New(50, 51),
		RangeB:          New(40, 41),
		ResultExpected:  NewEmpty(),
		ErrorIsExpected: false,
	}
	tests = append(tests, test)

	// Test #8.
	test = TestWorkspace{
		RangeA:          New(50, 51),
		RangeB:          New(40, 50),
		ResultExpected:  New(50, 50),
		ErrorIsExpected: false,
	}
	tests = append(tests, test)

	// Test #9.
	test = TestWorkspace{
		RangeA:          New(50, 55),
		RangeB:          New(40, 52),
		ResultExpected:  New(50, 52),
		ErrorIsExpected: false,
	}
	tests = append(tests, test)

	// Test #10.
	test = TestWorkspace{
		RangeA:          New(50, 55),
		RangeB:          New(40, 55),
		ResultExpected:  New(50, 55),
		ErrorIsExpected: false,
	}
	tests = append(tests, test)

	// Test #11.
	test = TestWorkspace{
		RangeA:          New(50, 55),
		RangeB:          New(40, 59),
		ResultExpected:  New(50, 55),
		ErrorIsExpected: false,
	}
	tests = append(tests, test)

	// Tests for a non-valid Range...

	// Test #12.
	test = TestWorkspace{
		RangeA:          New(101, 102).Rotated(),
		RangeB:          New(10, 20),
		ResultExpected:  NewEmpty(),
		ErrorIsExpected: true,
	}
	tests = append(tests, test)

	// Test #13.
	test = TestWorkspace{
		RangeA:          New(10, 20),
		RangeB:          New(101, 102).Rotated(),
		ResultExpected:  NewEmpty(),
		ErrorIsExpected: true,
	}
	tests = append(tests, test)

	for testIdx, _ := range tests {
		tests[testIdx].ResultGot, tests[testIdx].ErrorGot = Intersection(
			tests[testIdx].RangeA,
			tests[testIdx].RangeB,
		)

		// Check for expected Error.
		if tests[testIdx].ErrorIsExpected {
			if tests[testIdx].ErrorGot == nil {
				tests[testIdx].Report = fmt.Errorf(
					ErrMsgPrefix+
						"Error was expected, but none has been received. "+
						"Test #%v.",
					testIdx+1,
				)
			}
		} else {
			if tests[testIdx].ErrorGot != nil {
				tests[testIdx].Report = fmt.Errorf(
					ErrMsgPrefix+
						"Error was not expected, but has been received. "+
						"Test #%v.",
					testIdx+1,
				)
			}
		}
		if tests[testIdx].Report != nil {
			t.Error(tests[testIdx].Report)
			t.FailNow()
		}

		// Check for expected Result.
		if !AreEqual(tests[testIdx].ResultGot, tests[testIdx].ResultExpected) {
			tests[testIdx].Report = fmt.Errorf(
				ErrMsgPrefix+
					"Result Value Mismatch. "+
					"Expected: %v, "+
					"Received: %v, "+
					"Test #%v.",
				tests[testIdx].ResultExpected,
				tests[testIdx].ResultGot,
				testIdx+1,
			)
		}
		if tests[testIdx].Report != nil {
			t.Error(tests[testIdx].Report)
			t.FailNow()
		}
	}
}

func Test_New(
	t *testing.T,
) {
	const ErrMsg = "New" + ErrorReportPostfix

	var irCreated Range
	var left int
	var right int

	// Prepare Data.
	left = 10
	right = 20

	// Ensure that the Test is serious.
	if left == right {
		t.Error(ErrMsg)
		t.FailNow()
	}

	// #1. Normal Test.
	irCreated = New(left, right)
	if irCreated.isFilled != true {
		t.Error(ErrMsg)
		t.FailNow()
	}
	if irCreated.left != left {
		t.Error(ErrMsg)
		t.FailNow()
	}
	if irCreated.right != right {
		t.Error(ErrMsg)
		t.FailNow()
	}

	// #2. Rotated Test.
	irCreated = New(right, left)
	if irCreated.isFilled != true {
		t.Error(ErrMsg)
		t.FailNow()
	}
	if irCreated.left != left {
		t.Error(ErrMsg)
		t.FailNow()
	}
	if irCreated.right != right {
		t.Error(ErrMsg)
		t.FailNow()
	}
}

func Test_NewEmpty(
	t *testing.T,
) {
	const ErrMsg = "NewEmpty" + ErrorReportPostfix

	var irCreated Range

	irCreated = NewEmpty()

	if irCreated.isFilled != false {
		t.Error(ErrMsg)
		t.FailNow()
	}
	if irCreated.left != NewEmptyLeftCorner {
		t.Error(ErrMsg)
		t.FailNow()
	}
	if irCreated.right != NewEmptyRightCorner {
		t.Error(ErrMsg)
		t.FailNow()
	}
}

func Test_IsEmpty(
	t *testing.T,
) {
	const ErrMsg = "IsEmpty" + ErrorReportPostfix

	var irEmpty Range
	var irNotEmpty Range

	// Prepare Data.
	irEmpty = Range{
		left:     123,
		right:    456,
		isFilled: false,
	}
	irNotEmpty = Range{
		left:     999,
		right:    888,
		isFilled: true,
	}

	// #1.
	if irEmpty.IsEmpty() != true {
		t.Error(ErrMsg)
		t.FailNow()
	}

	// #2.
	if irNotEmpty.IsEmpty() != false {
		t.Error(ErrMsg)
		t.FailNow()
	}
}

func Test_IsEqualTo(
	t *testing.T,
) {
	const ErrMsg = "IsEqualTo" + ErrorReportPostfix

	var irA Range
	var irB Range

	// Prepare common Data.
	irA.isFilled = true
	irA.left = 456
	irA.right = 123

	// #1. Test equal Ranges.
	if irA.IsEqualTo(irA) != true {
		t.Error(ErrMsg)
		t.FailNow()
	}

	// #2. Test Ranges with different 'isFilled' Field.
	irB.isFilled = !irA.isFilled
	irB.left = irA.left
	irB.right = irA.right
	if irA.IsEqualTo(irB) != false {
		t.Error(ErrMsg)
		t.FailNow()
	}
	if irB.IsEqualTo(irA) != false {
		t.Error(ErrMsg)
		t.FailNow()
	}

	// #3. Test Ranges with different 'left' Field.
	irB.isFilled = irA.isFilled
	irB.left = irA.left + 1
	irB.right = irA.right
	if irA.IsEqualTo(irB) != false {
		t.Error(ErrMsg)
		t.FailNow()
	}
	if irB.IsEqualTo(irA) != false {
		t.Error(ErrMsg)
		t.FailNow()
	}

	// #4. Test Ranges with different 'right' Field.
	irB.isFilled = irA.isFilled
	irB.left = irA.left
	irB.right = irA.right + 1
	if irA.IsEqualTo(irB) != false {
		t.Error(ErrMsg)
		t.FailNow()
	}
	if irB.IsEqualTo(irA) != false {
		t.Error(ErrMsg)
		t.FailNow()
	}
}

func Test_IsNotEmpty(
	t *testing.T,
) {
	const ErrMsg = "IsNotEmpty" + ErrorReportPostfix

	var irEmpty Range
	var irNotEmpty Range

	// Prepare Data.
	irEmpty = Range{
		left:     123,
		right:    456,
		isFilled: false,
	}
	irNotEmpty = Range{
		left:     999,
		right:    888,
		isFilled: true,
	}

	// #1.
	if irEmpty.IsNotEmpty() != false {
		t.Error(ErrMsg)
		t.FailNow()
	}

	// #2.
	if irNotEmpty.IsNotEmpty() != true {
		t.Error(ErrMsg)
		t.FailNow()
	}
}

func Test_IsValid(
	t *testing.T,
) {
	const ErrMsg = "IsValid" + ErrorReportPostfix

	var irValidEmpty Range
	var irValidNonEmpty Range
	var irNotValid Range

	// Prepare Data.
	irValidEmpty = Range{
		left:     123,
		right:    456,
		isFilled: false,
	}
	irValidNonEmpty = Range{
		left:     123,
		right:    456,
		isFilled: true,
	}
	irNotValid = Range{
		left:     999,
		right:    888,
		isFilled: true,
	}

	// #1.
	if irValidEmpty.IsValid() != true {
		t.Error(ErrMsg)
		t.FailNow()
	}

	// #2.
	if irValidNonEmpty.IsValid() != true {
		t.Error(ErrMsg)
		t.FailNow()
	}

	// #3.
	if irNotValid.IsValid() != false {
		t.Error(ErrMsg)
		t.FailNow()
	}
}

func Test_Left(
	t *testing.T,
) {
	const ErrMsg = "Left" + ErrorReportPostfix

	var err error
	var irNotValid Range
	var irEmpty Range
	var irNormal Range
	var left int
	var leftExpected int
	var right int

	// Prepare Data.
	left = 14
	right = left + 1
	irNotValid = Range{
		left:     right,
		right:    left,
		isFilled: true,
	}
	irEmpty = Range{
		isFilled: false,
	}
	irNormal = Range{
		left:     left,
		right:    right,
		isFilled: true,
	}
	leftExpected = left

	// #1.
	left, err = irNotValid.Left()
	if err == nil {
		t.Error(ErrMsg)
		t.FailNow()
	}

	// #2.
	left, err = irEmpty.Left()
	if err == nil {
		t.Error(ErrMsg)
		t.FailNow()
	}

	// #3.
	left, err = irNormal.Left()
	if err != nil {
		t.Error(ErrMsg)
		t.FailNow()
	}
	if left != leftExpected {
		t.Error(ErrMsg)
		t.FailNow()
	}
}

func Test_Right(
	t *testing.T,
) {
	const ErrMsg = "Right" + ErrorReportPostfix

	var err error
	var irNotValid Range
	var irEmpty Range
	var irNormal Range
	var left int
	var right int
	var rightExpected int

	// Prepare Data.
	left = 14
	right = left + 1
	irNotValid = Range{
		left:     right,
		right:    left,
		isFilled: true,
	}
	irEmpty = Range{
		isFilled: false,
	}
	irNormal = Range{
		left:     left,
		right:    right,
		isFilled: true,
	}
	rightExpected = right

	// #1.
	right, err = irNotValid.Right()
	if err == nil {
		t.Error(ErrMsg)
		t.FailNow()
	}

	// #2.
	right, err = irEmpty.Right()
	if err == nil {
		t.Error(ErrMsg)
		t.FailNow()
	}

	// #3.
	right, err = irNormal.Right()
	if err != nil {
		t.Error(ErrMsg)
		t.FailNow()
	}
	if right != rightExpected {
		t.Error(ErrMsg)
		t.FailNow()
	}
}

func Test_Rotated(
	t *testing.T,
) {
	const ErrMsg = "Rotated" + ErrorReportPostfix

	var first Range
	var irRotated Range
	var left int
	var right int
	var newLeftExpected int
	var newRightExpected int

	// Prepare Data.
	left = 10
	right = 20
	first = Range{
		left:     left,
		right:    right,
		isFilled: false,
	}
	newLeftExpected = right
	newRightExpected = left
	irRotated = first.Rotated()

	// #1.
	if irRotated.left != newLeftExpected {
		t.Error(ErrMsg)
		t.FailNow()
	}
	if irRotated.right != newRightExpected {
		t.Error(ErrMsg)
		t.FailNow()
	}
	if irRotated.isFilled != first.isFilled {
		t.Error(ErrMsg)
		t.FailNow()
	}

	// #2.
	if !AreEqual(first, irRotated.Rotated()) {
		t.Error(ErrMsg)
		t.FailNow()
	}
}

func Test_RotatedIfNecessary(
	t *testing.T,
) {

	const ErrMsg = "RotatedIfNecessary" + ErrorReportPostfix

	var irNotValid Range
	var irRotatedIfNecessary Range
	var irValid Range
	var left int
	var right int
	var newFilledExpected bool
	var newLeftExpected int
	var newRightExpected int

	// Prepare Data.
	left = 10
	right = 20
	irValid = Range{
		left:     left,
		right:    right,
		isFilled: true,
	}
	irNotValid = Range{
		left:     right,
		right:    left,
		isFilled: true,
	}
	newLeftExpected = left
	newRightExpected = right
	newFilledExpected = irValid.isFilled

	// #1.
	irRotatedIfNecessary = irValid.RotatedIfNecessary()
	if irRotatedIfNecessary.left != newLeftExpected {
		t.Error(ErrMsg)
		t.FailNow()
	}
	if irRotatedIfNecessary.right != newRightExpected {
		t.Error(ErrMsg)
		t.FailNow()
	}
	if irRotatedIfNecessary.isFilled != newFilledExpected {
		t.Error(ErrMsg)
		t.FailNow()
	}

	// #2.
	irRotatedIfNecessary = irNotValid.RotatedIfNecessary()
	if irRotatedIfNecessary.left != newLeftExpected {
		t.Error(ErrMsg)
		t.FailNow()
	}
	if irRotatedIfNecessary.right != newRightExpected {
		t.Error(ErrMsg)
		t.FailNow()
	}
	if irRotatedIfNecessary.isFilled != newFilledExpected {
		t.Error(ErrMsg)
		t.FailNow()
	}
}

func Test_Size(
	t *testing.T,
) {
	const ErrMsg = "Size" + ErrorReportPostfix

	var delta int
	var err error
	var irEmpty Range
	var irNotValid Range
	var irNormal Range
	var left int
	var right int
	var size int
	var sizeExpected int

	// Prepare Data.
	left = 10
	delta = 2
	right = left + delta
	sizeExpected = delta + 1
	irNotValid = Range{
		left:     right,
		right:    left,
		isFilled: true,
	}
	irEmpty = Range{
		isFilled: false,
	}
	irNormal = Range{
		left:     left,
		right:    right,
		isFilled: true,
	}

	// #1.
	size, err = irNotValid.Size()
	if err == nil {
		t.Error(ErrMsg)
		t.FailNow()
	}

	// #2.
	size, err = irEmpty.Size()
	if err != nil {
		t.Error(ErrMsg)
		t.FailNow()
	}
	if size != 0 {
		t.Error(ErrMsg)
		t.FailNow()
	}

	// #3.
	size, err = irNormal.Size()
	if err != nil {
		t.Error(ErrMsg)
		t.FailNow()
	}
	if size != sizeExpected {
		t.Error(ErrMsg)
		t.FailNow()
	}
}
