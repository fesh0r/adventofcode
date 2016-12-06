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
		in   string
		out  int
		out2 int
	}{
		{"Sue 24: akitas: 5, goldfish: 6, vizslas: 6\nSue 40: vizslas: 0, cats: 7, akitas: 0\nSue 241: cars: 2, pomeranians: 1, samoyeds: 2\n",
			40, 241},
	}

	for _, tt := range tests {
		aunt, aunt2, err := process(strings.NewReader(tt.in))
		if err != nil {
			t.Errorf("process(%q) = error %s, want %d, %d", tt.in, err, tt.out, tt.out2)
		} else if aunt != tt.out || aunt2 != tt.out2 {
			t.Errorf("process(%q) = %d, %d, want %d, %d", tt.in, aunt, aunt2, tt.out, tt.out2)
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
		aunt, aunt2, err := process(strings.NewReader(tt))
		if err == nil {
			t.Errorf("process(%q) = %d, %d, want error", tt, aunt, aunt2)
		}
	}
}
