package main

import (
	"testing"
)

func TestProcess(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{"150", 8},
		{"1000", 48},
	}

	for _, tt := range tests {
		c, err := process(tt.in)
		if err != nil {
			t.Errorf("process(%q) = error %s, want %d", tt.in, err, tt.out)
		} else if c != tt.out {
			t.Errorf("process(%q) = %d, want %d", tt.in, c, tt.out)
		}
	}
}

func TestProcess2(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{"150", 8},
		{"1000", 36},
	}

	for _, tt := range tests {
		c, err := process2(tt.in)
		if err != nil {
			t.Errorf("process2(%q) = error %s, want %d", tt.in, err, tt.out)
		} else if c != tt.out {
			t.Errorf("process2(%q) = %d, want %d", tt.in, c, tt.out)
		}
	}
}
