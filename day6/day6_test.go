package main

import (
	"reflect"
	"testing"
)

func TestParseLine(t *testing.T) {
	tests := []struct {
		in      string
		outInst instruction
		out     coordinates
	}{
		{"turn on 0,0 through 999,999", turnOn, coordinates{0, 0, 999, 999}},
		{"toggle 0,0 through 999,0", toggle, coordinates{0, 0, 999, 0}},
		{"turn off 499,499 through 500,500", turnOff, coordinates{499, 499, 500, 500}},
		{"turn off 499,499 through 500,500", turnOff, coordinates{499, 499, 500, 500}},
		{"turn off 499,499 through 500,498", turnOff, coordinates{499, 499, 500, 498}},
	}

	for _, tt := range tests {
		i, c, err := parseLine(tt.in)
		if err != nil {
			t.Errorf("parseLine(%q) = error %s, want %v, %d", tt.in, err, tt.outInst, tt.out)
		} else if i != tt.outInst || !reflect.DeepEqual(c, tt.out) {
			t.Errorf("parseLine(%q) = %v, %d, want %v, %d", tt.in, i, c, tt.outInst, tt.out)
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
		i, c, err := parseLine(tt)
		if err == nil {
			t.Errorf("parseLine(%q) = %v, %d, want error", tt, i, c)
		}
	}
}
