package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(source string) (string, error) {
	if !IsValid(source) {
		return source, ErrInvalidString
	}

	return FinallyUnpack(Prepare(source)), nil
}

func FinallyUnpack(prepared string) string {
	result := ""
	prevLetter := ""
	for i, v := range prepared {
		if i%2 == 1 {
			count, err := strconv.Atoi(string(v))
			if err != nil {
				panic("Cant parse int")
			}
			result += strings.Repeat(prevLetter, count)
		} else {
			prevLetter = string(v)
		}
	}

	return result
}

func Prepare(source string) string {
	isPrevLetter := false
	preparedString := ""
	for i, v := range source {
		if isPrevLetter && unicode.IsLetter(v) {
			preparedString += "1" + string(v)
		} else {
			preparedString += string(v)
		}

		if unicode.IsLetter(v) {
			isPrevLetter = true
		} else {
			isPrevLetter = false
		}

		if i == (len(source)-1) && unicode.IsLetter(v) {
			preparedString += "1"
		}
	}

	return preparedString
}

func IsValid(source string) bool {
	isPrevDigit := false
	for i, v := range source {
		if i == 0 && unicode.IsDigit(v) {
			return false
		}

		if unicode.IsDigit(v) && isPrevDigit {
			return false
		}

		if unicode.IsDigit(v) && !isPrevDigit {
			isPrevDigit = true
		}

		if unicode.IsLetter(v) {
			isPrevDigit = false
		}
	}

	return true
}
