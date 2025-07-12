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
			input: "   hello world   ",
			expected: []string{"hello", "world"},
		},
		{
			input: "hello world   ",
			expected: []string{"hello", "world"},
		},
		{
			input: "   hello world",
			expected: []string{"hello", "world"},
		},
		{
			input: "hello    world",
			expected: []string{"hello", "world"},
		},
		{
			input: "   THIS is A Message    To ALL my BabY   girLs  ",
			expected: []string{"this","is","a","message","to","all","my","baby","girls"},
		},
	}

	// Checks each test case written above
	for _, c := range cases {
		// Will check to see if both the expected and made slices are the same size
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Error: unexpected amount of items in slice: %d", len(actual))
			continue
		}
		// Will check to see if each slice has the same words
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Error: mismatched word found in slice")
			}
		}
	}
}