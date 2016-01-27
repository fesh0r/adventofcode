package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseLine(t *testing.T) {
	tests := []struct {
		in         string
		outIndex   int
		outAttribs attributes
	}{
		{"Sue 119: goldfish: 6, perfumes: 3, children: 1", 119, attributes{
			"goldfish": 6,
			"children": 1,
			"perfumes": 3,
		}},
	}

	for _, tt := range tests {
		index, attribs, err := parseLine(tt.in)
		if err != nil {
			t.Errorf("parseLine(%q) = error %s, want %d, %v", tt.in, err, tt.outIndex, tt.outAttribs)
		} else if index != tt.outIndex || !reflect.DeepEqual(attribs, tt.outAttribs) {
			t.Errorf("parseLine(%q) = %d, %v, want %d, %v", tt.in, index, attribs, tt.outIndex, tt.outAttribs)
		}
	}
}

func TestParseLineError(t *testing.T) {
	tests := []string{
		"",
		"Sue",
		"Sue A: goldfish: 6, perfumes: 3, children: 1",
		"Sue 119: goldfish: A, perfumes: 3, children: 1",
		"Sue 119: fish: 6, perfumes: 3, children: 1",
	}

	for _, tt := range tests {
		index, attribs, err := parseLine(tt)
		if err == nil {
			t.Errorf("parseLine(%q) = %d, %v, want error", tt, index, attribs)
		}
	}
}

func TestProcess(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{"Sue 24: akitas: 5, goldfish: 6, vizslas: 6\nSue 40: vizslas: 0, cats: 7, akitas: 0\n", 40},
	}

	for _, tt := range tests {
		aunt, err := process(strings.NewReader(tt.in))
		if err != nil {
			t.Errorf("process(%q) = error %s, want %d", tt.in, err, tt.out)
		} else if aunt != tt.out {
			t.Errorf("process(%q) = %d, want %d", tt.in, aunt, tt.out)
		}
	}
}

func TestProcessError(t *testing.T) {
	tests := []string{
		"Sue",
		"Sue A: goldfish: 6, perfumes: 3, children: 1",
		"Sue 119: goldfish: A, perfumes: 3, children: 1",
		"Sue 119: fish: 6, perfumes: 3, children: 1",
	}

	for _, tt := range tests {
		aunt, err := process(strings.NewReader(tt))
		if err == nil {
			t.Errorf("process(%q) = %d, want error", tt, aunt)
		}
	}
}
