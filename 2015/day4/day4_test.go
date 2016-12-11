package main

import (
	"testing"
)

func TestCheckIndex(t *testing.T) {
	tests := []struct {
		in     string
		prefix string
		index  int
		out    bool
	}{
		{"abcdef", "00000", 1, false},
		{"pqrstuv", "00000", 1, false},
		{"abcdef", "00000", 609043, true},
		{"pqrstuv", "00000", 1048970, true},
		{"abcdef", "000000", 6742839, true},
		{"pqrstuv", "000000", 5714438, true},
		{"☃", "00000", 1, false},
		{"☃", "00000", 762997, true},
	}

	for _, tt := range tests {
		b := checkIndex(tt.in, tt.prefix, tt.index)
		if b != tt.out {
			t.Errorf("checkIndex(%q,%q,%d) = %t, want %t", tt.in, tt.prefix, tt.index, b, tt.out)
		}
	}
}

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
			t.Errorf("findCoin(%q,%q) = error %s, want %d", tt.in, tt.prefix, err, tt.out)
		} else if i != tt.out {
			t.Errorf("findCoin(%q,%q) = %d, want %d", tt.in, tt.prefix, i, tt.out)
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
			t.Errorf("findCoin(%q,%q) = error %s, want %d", tt.in, tt.prefix, err, tt.out)
		} else if i != tt.out {
			t.Errorf("findCoin(%q,%q) = %d, want %d", tt.in, tt.prefix, i, tt.out)
		}
	}
}
