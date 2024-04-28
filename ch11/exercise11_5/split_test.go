// Exercise 11.5: Extend TestSplit to use a table of inputs and expected outputs.
package split_test

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	var tests = []struct {
		s    string
		sep  string
		want int
	}{
		{"a:b:c", ":", 3},
		{"", ":", 1},
		{":", ":", 2},
		{",:,", ",", 3},
	}
	for _, test := range tests {
		words := strings.Split(test.s, test.sep)
		if got, want := len(words), test.want; got != want {
			t.Errorf("Split(%q, %q) returned %d words, want %d", test.s, test.sep, got, want)
		}

	}
}
