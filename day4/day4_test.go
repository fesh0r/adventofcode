package main

import (
	"fmt"
	"testing"
)

func TestFindCoin(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{"abcdef", 609043},
		{"pqrstuv", 1048970},
	}

	for _, tt := range tests {
		i, err := findCoin(tt.in)
		if err != nil {
			t.Errorf("findCoin(%q) expected: %d, got error: %s", tt.in, tt.out, err)
		} else if i != tt.out {
			t.Errorf("findCoin(%q) expected: %d, got: %d", tt.in, tt.out, i)
		}
	}
}

func ExampleFindCoin() {
	s := "abcdef"
	i, _ := findCoin(s)
	fmt.Printf("%q => %d\n", s, i)
	// Output: "abcdef" => 609043
}
