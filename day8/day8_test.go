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
		c, m, err := unquotedSize(tt.in)
		if err != nil {
			t.Errorf("unquotedSize(%q) = error %s, want %d, %d", tt.in, err, tt.outCode, tt.outMem)
		} else if c != tt.outCode || m != tt.outMem {
			t.Errorf("unquotedSize(%q) = %d, %d, want %d, %d", tt.in, c, m, tt.outCode, tt.outMem)
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
		c, m, err := unquotedSize(tt)
		if err == nil {
			t.Errorf("unquotedSize(%q) = %d, %d, want error", tt, c, m)
		}
	}
}

func TestProcess(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{`""` + "\n" + `"abc"` + "\n" + `"aaa\"aaa"` + "\n" + `"\x27"` + "\n", 12},
	}

	for _, tt := range tests {
		v, err := process(strings.NewReader(tt.in))
		if err != nil {
			t.Errorf("process(%q) = error %s, want %d", tt.in, err, tt.out)
		} else if v != tt.out {
			t.Errorf("process(%q) = %d, want %d", tt.in, v, tt.out)
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
		v, err := process(strings.NewReader(tt))
		if err == nil {
			t.Errorf("process(%q) = %d, want error", tt, v)
		}
	}
}
