package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "onie gdyby",
			expected: []string{"onie", "gdyby"},
		},
		{
			input:    "Hello, World!",
			expected: []string{"Hello", "World"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("The length of the expected slice differs from the actual slice")
			return
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("the word %s is different than expected %s", word, expectedWord)
				return
			}
		}
	}
}
