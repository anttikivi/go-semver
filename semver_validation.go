package semver

import (
	"slices"
	"strings"
)

// IsValid reports whether s is a valid semantic version string.
// The version may have a 'v' prefix.
func IsValid(ver string) bool {
	return isValid(ver)
}

// IsValidPrefix reports whether s is a valid semantic version string.
// It allows the version to have either one of the given prefixes or a 'v'
// prefix.
func IsValidPrefix(ver string, p ...string) bool {
	return isValid(ver, p...)
}

func isValid(ver string, prefixes ...string) bool { //nolint:cyclop,funlen,gocognit,gocyclo // not really too complex
	if ver == "" {
		return false
	}

	pos := strings.IndexFunc(ver, func(r rune) bool { return '0' <= r && r <= '9' })
	if pos == -1 {
		// The version contains digits so it cannot be valid.
		return false
	}

	// The position is other than the first index so the string has a prefix.
	// We need to check whether it is one of the valid prefixes.
	if pos != 0 {
		prefix := ver[:pos]
		if !slices.Contains(prefixes, prefix) && prefix != "v" {
			return false
		}
	}

	length := len(ver)

	// Check the major and minor number.
	// Both of them should start at the next position and end in a dot so we can
	// just repeat this loop twice.
	for range 2 {
		start := pos
		zero := ver[pos] == '0'

		// Check that every number before the next dot is a digit.
		for ; pos < length && ver[pos] != '.'; pos++ {
			if ver[pos] < '0' || ver[pos] > '9' {
				return false
			}
		}

		if pos-start < 1 && zero {
			return false
		}

		// We cannot be at the end yet.
		if pos >= length {
			return false
		}

		// Every character was a digit and we reached a dot before the end of the
		// string so let's hop over the dot and repeat the process for the next
		// number.
		pos++
	}

	// Next check the patch number.
	// Otherwise the check is the same as for major and minor but it can end in
	// a hyphen or a plus.
	start := pos
	zero := ver[pos] == '0'

	// Check that every number before the next dot is a digit.
	for ; pos < length && ver[pos] != '-' && ver[pos] != '+'; pos++ {
		if ver[pos] < '0' || ver[pos] > '9' {
			return false
		}
	}

	if pos-start < 1 && zero {
		return false
	}

	// If the major, minor, and patch were checked successfully and we are at
	// the end, the version is valid.
	if pos >= length {
		return true
	}

	// Check the pre-release identifiers.
	if ver[pos] == '-' { //nolint:nestif // not too complex
		// Skip the hyphen.
		pos++

		num := true
		zero = false
		currentLen := 0

		for ; pos < length && ver[pos] != '+'; pos++ {
			b := ver[pos]
			// If the character is a dot, start a new identifier.
			if b == '.' {
				// If the identifier with a leading zero is a number longer than
				// one character, the version is invalid.
				if zero && num && currentLen > 1 {
					return false
				}

				num = true
				zero = false
				currentLen = 0
				pos++
				// Empty identifier is invalid.
				if b = ver[pos]; b == '+' || pos >= length {
					return false
				}
			}

			if b == '0' && currentLen == 0 {
				zero = true
			}

			// If the identifier is still a number but we encounter a non-digit
			// character, the identifier is no longer a number.
			if num && ('A' <= b && b <= 'Z') || ('a' <= b && b <= 'z') || b == '-' {
				num = false
			}

			// Otherwise just check that the character is valid.
			if ('A' > b || b > 'Z') && ('a' > b || b > 'z') && ('0' > b || b > '9') && b != '-' {
				return false
			}

			currentLen++
		}

		// If the identifier with a leading zero is a number longer than
		// one character, the version is invalid.
		if zero && num && currentLen > 1 {
			return false
		}

		if currentLen == 0 {
			return false
		}
	}

	if pos >= length {
		return true
	}

	if ver[pos] == '+' {
		pos++
		for ; pos < length; pos++ {
			b := ver[pos]
			if ('A' > b || b > 'Z') && ('a' > b || b > 'z') && ('0' > b || b > '9') && b != '-' && b != '.' {
				return false
			}
		}
	}

	return true
}
