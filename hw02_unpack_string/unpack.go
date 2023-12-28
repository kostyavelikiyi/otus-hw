package hw02unpackstring

import (
	"errors"
	"strconv"
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
	var result = ""
	var prevLetter = ""
	for i, v := range prepared {
		if i%2 == 1 {
			var count, err = strconv.Atoi(string(v))
			if err != nil {
				panic("Cant parse int")
			}
			for j := 0; j < count; j++ {
				result += prevLetter
			}
		} else {
			prevLetter = string(v)
		}
	}

	return result
}

func Prepare(source string) string {
	var isPrevLetter = false
	var preparedString = ""
	for i, v := range source {
		if isPrevLetter && unicode.IsLetter(v) {
			preparedString += "1" + string(v)
			isPrevLetter = true
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

	var isPrevDigit = false
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
