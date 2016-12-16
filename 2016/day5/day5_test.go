package main

import "testing"

func TestCheckIndex(t *testing.T) {
	tests := []struct {
		in    string
		index int
		out   bool
		outC  string
	}{
		{
			"abc", 3231928,
			false, "",
		},
		{
			"abc", 3231929,
			true, "1",
		},
		{
			"abc", 5017308,
			true, "8",
		},
		{
			"abc", 5278568,
			true, "f",
		},
	}

	for _, tt := range tests {
		f, c := checkIndex(tt.in, tt.index)
		if f != tt.out || c != tt.outC {
			t.Errorf("checkIndex(%q,%d) = %t,%q, want %t,%q", tt.in, tt.index, f, c, tt.out, tt.outC)
		}
	}
}

func TestFindPassword(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"abc", "18f47a30"},
	}

	for _, tt := range tests {
		c, err := findPassword(tt.in)
		if err != nil {
			t.Errorf("findPassword(%q) = error %s, want %q", tt.in, err, tt.out)
		} else if c != tt.out {
			t.Errorf("findPassword(%q) = %q, want %q", tt.in, c, tt.out)
		}
	}
}

func TestProcess(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"abc", "18f47a30"},
	}

	for _, tt := range tests {
		c, err := process(tt.in)
		if err != nil {
			t.Errorf("process(%q) = error %s, want %q", tt.in, err, tt.out)
		} else if c != tt.out {
			t.Errorf("process(%q) = %q, want %q", tt.in, c, tt.out)
		}
	}
}
