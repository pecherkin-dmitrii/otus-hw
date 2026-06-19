package hw02unpackstring

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

var ErrInvalidString = errors.New("invalid string")

var startsWithDigits = regexp.MustCompile(`^\d+`)
var containsNumber = regexp.MustCompile(`\d{2,}`)

func Unpack(str string) (string, error) {
	switch {
	case utf8.RuneCountInString(str) == 0:
		return "", nil
	case startsWithDigits.MatchString(str) || containsNumber.MatchString(str):
		return "", ErrInvalidString
	default:
		builder := strings.Builder{}
		runeSlice := []rune(str)
		for i, ch := range runeSlice {
			if _, err := strconv.Atoi(string(ch)); err == nil {
				continue
			} else if i+1 == len(runeSlice) {
				builder.WriteRune(ch)
			} else if digit, err := strconv.Atoi(string(runeSlice[i+1])); err != nil {
				builder.WriteRune(ch)
				continue
			} else {
				builder.WriteString(strings.Repeat(string(ch), digit))
			}
		}
		return builder.String(), nil
	}
}
