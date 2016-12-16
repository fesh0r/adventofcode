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
			t.Errorf("checkIndex(%q, %d) = %t, %q, want %t, %q", tt.in, tt.index, f, c, tt.out, tt.outC)
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

func TestCheckIndex2(t *testing.T) {
	tests := []struct {
		in    string
		index int
		out   bool
		outP  int
		outC  byte
	}{
		{
			"abc", 3231928,
			false, 0, 0,
		},
		{
			"abc", 3231929,
			true, 1, byte('5'),
		},
		{
			"abc", 5017308,
			false, 0, 0,
		},
		{
			"abc", 5357525,
			true, 4, byte('e'),
		},
	}

	for _, tt := range tests {
		f, p, c := checkIndex2(tt.in, tt.index)
		if f != tt.out || c != tt.outC {
			t.Errorf("checkIndex2(%q, %d) = %t, %d, %q, want %t, %d, %q",
				tt.in, tt.index, f, p, c, tt.out, tt.outP, tt.outC)
		}
	}
}

func TestProcess(t *testing.T) {
	tests := []struct {
		in   string
		out1 string
		out2 string
	}{
		{"abc", "18f47a30", "05ace8e3"},
	}

	for _, tt := range tests {
		c1, c2, err := process(tt.in)
		if err != nil {
			t.Errorf("process(%q) = error %s, want %q, %q", tt.in, err, tt.out1, tt.out2)
		} else if c1 != tt.out1 || c2 != tt.out2 {
			t.Errorf("process(%q) = %q, %q, want %q, %q", tt.in, c1, c2, tt.out1, tt.out2)
		}
	}
}
