package main

import (
	"strings"
	"testing"
)

func TestProcess(t *testing.T) {
	tests := []struct {
		in   string
		inC  int
		out  int
		outM int
	}{
		{"20\n15\n10\n5\n5\n", 25, 4, 3},
	}

	for _, tt := range tests {
		valid, validMin, err := process(strings.NewReader(tt.in), tt.inC)
		if err != nil {
			t.Errorf("process(%q, %d) = error %s, want %d, %d", tt.in, tt.inC, err, tt.out, tt.outM)
		} else if valid != tt.out || validMin != tt.outM {
			t.Errorf("process(%q, %d) = %d, %d, want %d, %d", tt.in, tt.inC, valid, validMin, tt.out, tt.outM)
		}
	}
}

func TestProcessError(t *testing.T) {
	tests := []struct {
		in  string
		inC int
	}{
		{"A", 1},
		{"42\n23\nZ\n", 1},
	}

	for _, tt := range tests {
		valid, validMin, err := process(strings.NewReader(tt.in), tt.inC)
		if err == nil {
			t.Errorf("process(%q, %d) = %d, %d, want error", tt.in, tt.inC, valid, validMin)
		}
	}
}
