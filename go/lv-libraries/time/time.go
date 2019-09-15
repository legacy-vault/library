package time

import (
	"fmt"
	"time"
)

var emptyTime time.Time

func NewTimeStringRFC3339(
	year uint,
	month uint,
	day uint,
	hour uint,
	minute uint,
	second uint,
) string {

	var result string

	// Year Numbers less then 100 are considered to be 20xx.
	if year < 100 {
		year = year + 2000
	}

	yearStr := fmt.Sprintf("%04d", year)
	monthStr := fmt.Sprintf("%02d", month)
	dayStr := fmt.Sprintf("%02d", day)
	hourStr := fmt.Sprintf("%02d", hour)
	minuteStr := fmt.Sprintf("%02d", minute)
	secondStr := fmt.Sprintf("%02d", second)

	result = fmt.Sprintf(
		"%s-%s-%sT%s:%s:%sZ",
		yearStr,
		monthStr,
		dayStr,
		hourStr,
		minuteStr,
		secondStr,
	)

	return result
}

func Minimum(
	a time.Time,
	b time.Time,
) time.Time {
	if a.After(b) {
		return b
	}
	return a
}

func Maximum(
	a time.Time,
	b time.Time,
) time.Time {
	if a.After(b) {
		return a
	}
	return b
}

func IsEmpty(
	t time.Time,
) bool {
	if t == emptyTime {
		return true
	}
	return false
}
