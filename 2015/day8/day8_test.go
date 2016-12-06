package main

import (
	"strings"
	"testing"
)

func TestUnquotedSize(t *testing.T) {
	tests := []struct {
		in      string
		outCode int
		outMem  int
	}{
		{`""`, 2, 0},
		{`"abc"`, 5, 3},
		{`"aaa\"aaa"`, 10, 7},
		{`"\x27"`, 6, 1},
	}

	for _, tt := range tests {
		c, m, err := UnquotedSize(tt.in)
		if err != nil {
			t.Errorf("UnquotedSize(%q) = error %s, want %d, %d", tt.in, err, tt.outCode, tt.outMem)
		} else if c != tt.outCode || m != tt.outMem {
			t.Errorf("UnquotedSize(%q) = %d, %d, want %d, %d", tt.in, c, m, tt.outCode, tt.outMem)
		}
	}
}

func TestUnquotedSizeError(t *testing.T) {
	tests := []string{
		``,
		`"`,
		`"""`,
		`"\"`,
		`"\xZX"`,
	}

	for _, tt := range tests {
		c, m, err := UnquotedSize(tt)
		if err == nil {
			t.Errorf("UnquotedSize(%q) = %d, %d, want error", tt, c, m)
		}
	}
}

func TestQuotedSize(t *testing.T) {
	tests := []struct {
		in      string
		outCode int
		outMem  int
	}{
		{`""`, 6, 2},
		{`"abc"`, 9, 5},
		{`"aaa\"aaa"`, 16, 10},
		{`"\x27"`, 11, 6},
	}

	for _, tt := range tests {
		c, m := QuotedSize(tt.in)
		if c != tt.outCode || m != tt.outMem {
			t.Errorf("QuotedSize(%q) = %d, %d, want %d, %d", tt.in, c, m, tt.outCode, tt.outMem)
		}
	}
}

func TestProcess(t *testing.T) {
	tests := []struct {
		in   string
		out  int
		out2 int
	}{
		{`""` + "\n" + `"abc"` + "\n" + `"aaa\"aaa"` + "\n" + `"\x27"` + "\n", 12, 19},
	}

	for _, tt := range tests {
		v, v2, err := Process(strings.NewReader(tt.in))
		if err != nil {
			t.Errorf("Process(%q) = error %s, want %d, %d", tt.in, err, tt.out, tt.out2)
		} else if v != tt.out || v2 != tt.out2 {
			t.Errorf("Process(%q) = %d, %d, want %d, %d", tt.in, v, v2, tt.out, tt.out2)
		}
	}
}

func TestProcessError(t *testing.T) {
	tests := []string{
		`"`,
		`"""`,
		`"\"`,
		`"\xZX"`,
	}

	for _, tt := range tests {
		v, v2, err := Process(strings.NewReader(tt))
		if err == nil {
			t.Errorf("Process(%q) = %d, %d, want error", tt, v, v2)
		}
	}
}
