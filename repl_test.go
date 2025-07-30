package main

import (
	"reflect"
	"testing"
)

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
			input:    "a\tstring\nwith mixed\t\twhitespace",
			expected: []string{"a", "string", "with", "mixed", "whitespace"},
		},
		{
			input:    "Words   with    extra   spaces",
			expected: []string{"Words", "with", "extra", "spaces"},
		},
		{
			input:    "This is a sample sentence",
			expected: []string{"This", "is", "a", "sample", "sentence"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("Expected %v, but got %v", c.expected, actual)
		}
	}
}
