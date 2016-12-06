package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseLine(t *testing.T) {
	tests := []struct {
		in      string
		outInst Instruction
		out     Coordinates
	}{
		{"turn on 0,0 through 999,999", turnOn, Coordinates{0, 0, 999, 999}},
		{"toggle 0,0 through 999,0", toggle, Coordinates{0, 0, 999, 0}},
		{"turn off 499,499 through 500,500", turnOff, Coordinates{499, 499, 500, 500}},
	}

	for _, tt := range tests {
		i, c, err := ParseLine(tt.in)
		if err != nil {
			t.Errorf("ParseLine(%q) = error %s, want %v, %d", tt.in, err, tt.outInst, tt.out)
		} else if i != tt.outInst || !reflect.DeepEqual(c, tt.out) {
			t.Errorf("ParseLine(%q) = %v, %d, want %v, %d", tt.in, i, c, tt.outInst, tt.out)
		}
	}
}

func TestParseLineError(t *testing.T) {
	tests := []string{
		"",
		"turn away 0,0 through 999,999",
		"turn on 0,0 through 999,ABCD",
		"turn off 499,499 through 500,10000000000000000000",
		"turn on 499,499 through 500,498",
	}

	for _, tt := range tests {
		i, c, err := ParseLine(tt)
		if err == nil {
			t.Errorf("ParseLine(%q) = %v, %d, want error", tt, i, c)
		}
	}
}

func TestProcess(t *testing.T) {
	tests := []struct {
		in   string
		outC int
		outB int
	}{
		{"turn on 0,0 through 999,999\ntoggle 0,0 through 999,0\n", 999000, 1002000},
		{"toggle 499,499 through 500,500\ntoggle 0,0 through 999,0\n", 1004, 2008},
	}

	for _, tt := range tests {
		c, b, err := Process(strings.NewReader(tt.in))
		if err != nil {
			t.Errorf("Process(%q) = error %s, want %d, %d", tt.in, err, tt.outC, tt.outB)
		} else if c != tt.outC || b != tt.outB {
			t.Errorf("Process(%q) = %v, %d, want %d, %d", tt.in, c, b, tt.outC, tt.outB)
		}
	}
}

func TestProcessError(t *testing.T) {
	tests := []string{
		"turn away 0,0 through 999,999",
		"turn on 0,0 through 999,ABCD",
		"turn off 499,499 through 500,10000000000000000000",
		"turn on 499,499 through 500,498",
	}

	for _, tt := range tests {
		c, b, err := Process(strings.NewReader(tt))
		if err == nil {
			t.Errorf("Process(%q) = %d, %d, want error", tt, c, b)
		}
	}
}
