package main

import (
	"fmt"
	"testing"
)

func TestGetWrapping(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{"2x3x4", 58},
		{"1x1x10", 43},
	}

	for _, tt := range tests {
		w, err := getWrapping(tt.in)
		if err != nil {
			t.Errorf("getWrapping(%q) expected: %d, got error: %s", tt.in, tt.out, err)
		} else if w != tt.out {
			t.Errorf("getWrapping(%q) expected: %d, got: %d", tt.in, tt.out, w)
		}
	}
}

func TestGetWrappingError(t *testing.T) {
	tests := []string{
		" ",
		"a",
		"2",
		"1x2",
		"1x2x3x4",
		"2x3y2",
	}
	for _, tt := range tests {
		w, err := getWrapping(tt)
		if err == nil {
			t.Errorf("getWrapping(%q) expected error, got: %d", tt, w)
		}
	}
}

func ExampleGetWrapping() {
	s := "2x3x4"
	w, _ := getWrapping(s)
	fmt.Printf("%q => %d\n", s, w)
	// Output: "2x3x4" => 58
}
