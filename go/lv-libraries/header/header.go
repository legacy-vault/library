package header

// Common Message Header Names for HTTP and Mail Protocols.

// Notes.
//
// This Packages contains all IANA registered Header Names which have an RFC Reference as of 2019-05-26.
// For more Information visit this URL:
// https://www.iana.org/assignments/message-headers/message-headers.xml
// This List may contain some old deprecated and obsolete Header Names.

// Delimiters.
const CommaSpace = ", "

// MakeListOfHeaders Function composes a List of Headers delimited by Comma and
// Space.
func MakeListOfHeaders(
	headers []string,
) string {

	var result string

	iLast := len(headers) - 1
	for i, header := range headers {
		if i != iLast {
			result = result + header + CommaSpace
		} else {
			result = result + header
		}
	}

	return result
}
