package main

import (
	"strings"
	"testing"
)

func TestParseLine(t *testing.T) {
	tests := []struct {
		in  string
		out operation
	}{
		{
			"inc a",
			operation{opInc, regA, 0},
		},
		{
			"jio a, +2",
			operation{opJio, regA, 2},
		},
		{
			"tpl b",
			operation{opTpl, regB, 0},
		},
	}

	for _, tt := range tests {
		op, err := parseLine(tt.in)
		if err != nil {
			t.Errorf("parseLine(%q) = error %s, want %v", tt.in, err, tt.out)
		} else if op != tt.out {
			t.Errorf("parseLine(%q) = %v, want %v", tt.in, op, tt.out)
		}
	}
}

func TestProcess(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{"inc a\njio a, +2\ntpl a\ninc a\n", 0},
	}

	for _, tt := range tests {
		rb, err := process(strings.NewReader(tt.in))
		if err != nil {
			t.Errorf("process(%q) = error %s, want %d", tt.in, err, tt.out)
		} else if rb != tt.out {
			t.Errorf("process(%q) = %d, want %d", tt.in, rb, tt.out)
		}
	}
}
