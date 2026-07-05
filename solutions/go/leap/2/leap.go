// Package leap should have a package comment that summarizes what it's about.
package leap

// IsLeapYear should have a comment documenting it.
func IsLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	}

	if year%4 == 0 && year%100 != 0 {
		return true
	}

	return false
}
