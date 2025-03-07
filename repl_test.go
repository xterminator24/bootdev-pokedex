package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input:		"  hello   world  ",
			expected: 	[]string{"hello", "world"},
		},
		{
			input:		"",
			expected: 	[]string{},
		},
		{
			input:		"My cool Hat",
			expected: 	[]string{"my", "cool", "hat"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("cleanInput(%q) = %v; want %v", c.input, actual, c.expected)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("cleanInput(%q) = %s; want %s", c.input, word, expectedWord)
			}
		}
	}
}