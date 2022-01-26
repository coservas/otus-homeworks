package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	var result strings.Builder
	var prev rune

	for i, cur := range input {
		isDigit := unicode.IsDigit(cur)
		isDigitPrev := unicode.IsDigit(prev)

		if (i == 0 && isDigit) || (isDigitPrev && isDigit) {
			return "", ErrInvalidString
		}

		if !isDigit {
			result.WriteRune(cur)
			prev = cur
			continue
		}

		if cur == '0' {
			tempString := strings.TrimSuffix(result.String(), string(prev))
			result.Reset()
			result.WriteString(tempString)
			continue
		}

		for j := rune(1); j < cur-'0'; j++ {
			result.WriteRune(prev)
		}

		prev = cur
	}

	return result.String(), nil
}
