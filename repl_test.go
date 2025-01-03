package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, c := range cases {
		got := cleanInput(c.input)
		// Check the length of the actual slice
		if len(got) != len(c.expected) {
			t.Errorf("cleanInput(%q) == %q, want %q", c.input, got, c.expected)
		}
		for i := range got {
			word := got[i]
			if word != c.expected[i] {
				t.Errorf("cleanInput(%q) == %q, want %q", c.input, got, c.expected)
			}
		}
	}
}
