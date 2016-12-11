package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseSize(t *testing.T) {
	tests := []struct {
		in  string
		out []int
	}{
		{"2x3x4", []int{2, 3, 4}},
		{"1x1x10", []int{1, 1, 10}},
	}

	for _, tt := range tests {
		w, err := parseSize(tt.in)
		if err != nil {
			t.Errorf("parseSize(%q) = error %s, want %d", tt.in, err, tt.out)
		} else if !reflect.DeepEqual(w, tt.out) {
			t.Errorf("parseSize(%q) = %d, want %d", tt.in, w, tt.out)
		}
	}
}

func TestParseSizeError(t *testing.T) {
	tests := []string{
		" ",
		"a",
		"2",
		"1x2",
		"1x2x3x4",
		"2x3y2",
		"1x2x3☃",
	}

	for _, tt := range tests {
		w, err := parseSize(tt)
		if err == nil {
			t.Errorf("parseSize(%q) = %d, want error", tt, w)
		}
	}
}

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
			t.Errorf("getWrapping(%q) = error %s, want %d", tt.in, err, tt.out)
		} else if w != tt.out {
			t.Errorf("getWrapping(%q) = %d, want %d", tt.in, w, tt.out)
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
		"1x2x3☃",
	}

	for _, tt := range tests {
		w, err := getWrapping(tt)
		if err == nil {
			t.Errorf("getWrapping(%q) = %d, want error", tt, w)
		}
	}
}

func TestGetRibbon(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{"2x3x4", 34},
		{"1x1x10", 14},
	}

	for _, tt := range tests {
		r, err := getRibbon(tt.in)
		if err != nil {
			t.Errorf("getRibbon(%q) = error %s, want %d", tt.in, err, tt.out)
		} else if r != tt.out {
			t.Errorf("getRibbon(%q) = %d, want %d", tt.in, r, tt.out)
		}
	}
}

func TestGetRibbonError(t *testing.T) {
	tests := []string{
		" ",
		"a",
		"2",
		"1x2",
		"1x2x3x4",
		"2x3y2",
		"1x2x3☃",
	}

	for _, tt := range tests {
		r, err := getRibbon(tt)
		if err == nil {
			t.Errorf("getRibbon(%q) = %d, want error", tt, r)
		}
	}
}

func TestProcess(t *testing.T) {
	tests := []struct {
		in   string
		outW int
		outR int
	}{
		{"2x3x4\n1x1x10\n", 101, 48},
	}

	for _, tt := range tests {
		w, r, err := process(strings.NewReader(tt.in))
		if err != nil {
			t.Errorf("process(%q) = error %s, want %d, %d", tt.in, err, tt.outW, tt.outR)
		} else if w != tt.outW || r != tt.outR {
			t.Errorf("process(%q) = %d, %d, want %d, %d", tt.in, w, r, tt.outW, tt.outR)
		}
	}
}

func TestProcessError(t *testing.T) {
	tests := []string{
		" ",
		"a",
		"2",
		"1x2",
		"1x2x3x4",
		"2x3y2",
		"1x2x3☃",
	}

	for _, tt := range tests {
		w, r, err := process(strings.NewReader(tt))
		if err == nil {
			t.Errorf("process(%q) = %d, %d, want error", tt, w, r)
		}
	}
}
