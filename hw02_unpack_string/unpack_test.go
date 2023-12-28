package hw02unpackstring

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		// uncomment if task with asterisk completed
		// {input: `qwe\4\5`, expected: `qwe45`},
		// {input: `qwe\45`, expected: `qwe44444`},
		// {input: `qwe\\5`, expected: `qwe\\\\\`},
		// {input: `qwe\\\3`, expected: `qwe\3`},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestIsValid(t *testing.T) {
	invalidStrings := []string{"3abc", "45", "aaa10b"}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			res := IsValid(tc)
			require.Equal(t, false, res)
		})
	}
}

func TestPrepare(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "a4b1c2d5e1"},
		{input: "abccd", expected: "a1b1c1c1d1"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "a1a1a0b1"},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			res := Prepare(tc.input)
			require.Equal(t, tc.expected, res)
		})
	}
}

func TestFinallyUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4b1c2d5e1", expected: "aaaabccddddde"},
		{input: "a1b1c1c1d1", expected: "abccd"},
		{input: "", expected: ""},
		{input: "a1a1a0b1", expected: "aab"},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			res := FinallyUnpack(tc.input)
			require.Equal(t, tc.expected, res)
		})
	}
}
