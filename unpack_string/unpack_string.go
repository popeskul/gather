package unpack_string

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

// UnpackString performs primitive string unpacking.
func UnpackString(input string) (string, error) {
	var result strings.Builder
	var prev rune
	escapeMode := false

	for i, r := range input {
		if escapeMode {
			escapeMode = false
			if unicode.IsDigit(r) || r == '\\' {
				result.WriteRune(r)
				prev = r
				continue
			} else {
				return "", errors.New("invalid escape sequence")
			}
		}

		if r == '\\' {
			if i == len(input)-1 {
				return "", errors.New("string ends with a single backslash")
			}
			escapeMode = true
			continue
		}

		if unicode.IsDigit(r) {
			if prev == 0 {
				return "", errors.New("string starts with a digit or has consecutive digits")
			}

			count, err := strconv.Atoi(string(r))
			if err != nil {
				return "", err
			}

			// Let's make sure that the number of repetitions is not negative
			if count > 0 {
				result.WriteString(strings.Repeat(string(prev), count-1))
			} else if count == 0 && result.Len() > 0 {
				// If the number of repetitions is 0, then the previous character is repeated 0 times
				temp := result.String()
				result.Reset()
				result.WriteString(temp[:len(temp)-1])
			}

			prev = 0
			continue
		} else {
			result.WriteRune(r)
		}
		prev = r
	}

	if escapeMode {
		return "", errors.New("string ends with a single backslash")
	}

	return result.String(), nil
}
