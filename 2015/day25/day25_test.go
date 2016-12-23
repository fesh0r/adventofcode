package main

import "testing"

func TestParseInput(t *testing.T) {
	tests := []struct {
		in         string
		outX, outY int
	}{
		{
			"To continue, please consult the code grid in the manual.  Enter the code at row 1, column 1.",
			1, 1,
		},
		{
			"To continue, please consult the code grid in the manual.  Enter the code at row 3, column 4.",
			4, 3,
		},
		{
			"To continue, please consult the code grid in the manual.  Enter the code at row 5, column 6.",
			6, 5,
		},
	}

	for _, tt := range tests {
		x, y, err := parseInput(tt.in)
		if err != nil {
			t.Errorf("parseInput(%q) = error %s, want %d, %d", tt.in, err, tt.outX, tt.outY)
		} else if x != tt.outX {
			t.Errorf("parseInput(%q) = %d, %d, want %d, %d", tt.in, x, y, tt.outX, tt.outY)
		}
	}
}

func TestGetIndex(t *testing.T) {
	tests := []struct {
		inX, inY int
		out      int
	}{
		{
			1, 1,
			1,
		},
		{
			4, 3,
			19,
		},
	}

	for _, tt := range tests {
		i := getIndex(tt.inX, tt.inY)
		if i != tt.out {
			t.Errorf("process(%d, %d) = %d, want %d", tt.inX, tt.inY, i, tt.out)
		}
	}
}

func TestGetCode(t *testing.T) {
	tests := []struct {
		in  int
		out int64
	}{
		{
			1,
			20151125,
		},
		{
			19,
			7981243,
		},
	}

	for _, tt := range tests {
		c := getCode(tt.in)
		if c != tt.out {
			t.Errorf("process(%d) = %d, want %d", tt.in, c, tt.out)
		}
	}
}

func TestProcess(t *testing.T) {
	tests := []struct {
		in  string
		out int64
	}{
		{
			"To continue, please consult the code grid in the manual.  Enter the code at row 1, column 1.",
			20151125,
		},
		{
			"To continue, please consult the code grid in the manual.  Enter the code at row 3, column 4.",
			7981243,
		},
		{
			"To continue, please consult the code grid in the manual.  Enter the code at row 5, column 6.",
			31663883,
		},
	}

	for _, tt := range tests {
		c, err := process(tt.in)
		if err != nil {
			t.Errorf("process(%q) = error %s, want %d", tt.in, err, tt.out)
		} else if c != tt.out {
			t.Errorf("process(%q) = %d, want %d", tt.in, c, tt.out)
		}
	}
}
