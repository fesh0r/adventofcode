package main

import (
	"strings"
	"testing"
)

func TestProcess(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{"5 10 25", 0},
		{"5 10 14", 1},
		{"5 10 14\n5 10 13\n5 10 25", 2},
		{"    5 10  14\n5      10 13\n 5  10   25\n", 2},
	}

	for _, tt := range tests {
		l, err := process(strings.NewReader(tt.in))
		if err != nil {
			t.Errorf("process(%q) = error %s, want %d", tt.in, err, tt.out)
		} else if l != tt.out {
			t.Errorf("process(%q) = %d, want %d", tt.in, l, tt.out)
		}
	}
}

func TestProcessError(t *testing.T) {
	tests := []string{
		" ",
		"u",
		"0",
		"☃",
		"103 104",
		"103 104  105\n ☃",
		"103 104    105\nA",
	}

	for _, tt := range tests {
		l, err := process(strings.NewReader(tt))
		if err == nil {
			t.Errorf("process(%q) = %d, want error", tt, l)
		}
	}
}
