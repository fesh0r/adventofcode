package main

import (
	"fmt"
	"testing"
)

func TestFindCoin(t *testing.T) {
	tests := []struct {
		in     string
		prefix string
		out    int
	}{
		{"abcdef", "00000", 609043},
		{"pqrstuv", "00000", 1048970},
	}

	for _, tt := range tests {
		i, err := findCoin(tt.in, tt.prefix)
		if err != nil {
			t.Errorf("findCoin(%q,%q) expected: %d, got error: %s", tt.in, tt.prefix, tt.out, err)
		} else if i != tt.out {
			t.Errorf("findCoin(%q,%q) expected: %d, got: %d", tt.in, tt.prefix, tt.out, i)
		}
	}
}

func TestFindCoin6(t *testing.T) {
	tests := []struct {
		in     string
		prefix string
		out    int
	}{
		{"abcdef", "000000", 6742839},
		{"pqrstuv", "000000", 5714438},
	}

	for _, tt := range tests {
		i, err := findCoin(tt.in, tt.prefix)
		if err != nil {
			t.Errorf("findCoin(%q,%q) expected: %d, got error: %s", tt.in, tt.prefix, tt.out, err)
		} else if i != tt.out {
			t.Errorf("findCoin(%q,%q) expected: %d, got: %d", tt.in, tt.prefix, tt.out, i)
		}
	}
}

func ExampleFindCoin() {
	s := "abcdef"
	i, _ := findCoin(s, "00000")
	fmt.Printf("%q => %d\n", s, i)
	// Output: "abcdef" => 609043
}
