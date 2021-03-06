package main

import (
	"testing"
)

func TestParseDirection(t *testing.T) {
	tests := []struct {
		in   rune
		outX int
		outY int
	}{
		{'>', 1, 0},
		{'<', -1, 0},
		{'^', 0, 1},
		{'v', 0, -1},
	}

	for _, tt := range tests {
		x, y, err := parseDirection(tt.in)
		if err != nil {
			t.Errorf("parseDirection(%q) = error %s, want %d, %d", tt.in, err, tt.outX, tt.outY)
		} else if x != tt.outX || y != tt.outY {
			t.Errorf("parseDirection(%q) = %d, %d, want %d, %d", tt.in, x, y, tt.outX, tt.outY)
		}
	}
}

func TestParseDirectionError(t *testing.T) {
	tests := []rune{
		' ',
		'A',
		'☃',
	}

	for _, tt := range tests {
		x, y, err := parseDirection(tt)
		if err == nil {
			t.Errorf("parseDirection(%q) = %d, %d, want error", tt, x, y)
		}
	}
}

func TestGetHouses(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{"", 1},
		{">", 2},
		{"^>v<", 4},
		{"^v^v^v^v^v", 2},
	}

	for _, tt := range tests {
		h, err := getHouses(tt.in)
		if err != nil {
			t.Errorf("getHouses(%q) = error %s, want %d", tt.in, err, tt.out)
		} else if h != tt.out {
			t.Errorf("getHouses(%q) = %d, want %d", tt.in, h, tt.out)
		}
	}
}

func TestGetHousesError(t *testing.T) {
	tests := []string{
		" ",
		"A",
		"☃",
	}

	for _, tt := range tests {
		h, err := getHouses(tt)
		if err == nil {
			t.Errorf("getHouses(%q) = %d, want error", tt, h)
		}
	}
}

func TestGetHousesDouble(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{"", 1},
		{"^v", 3},
		{"^>v<", 3},
		{"^v^v^v^v^v", 11},
	}

	for _, tt := range tests {
		h, err := getHousesDouble(tt.in)
		if err != nil {
			t.Errorf("getHousesDouble(%q) = error %s, want %d", tt.in, err, tt.out)
		} else if h != tt.out {
			t.Errorf("getHousesDouble(%q) = %d, want %d", tt.in, h, tt.out)
		}
	}
}

func TestGetHousesDoubleError(t *testing.T) {
	tests := []string{
		" ",
		"A",
		"☃",
	}

	for _, tt := range tests {
		h, err := getHousesDouble(tt)
		if err == nil {
			t.Errorf("getHousesDouble(%q) = %d, want error", tt, h)
		}
	}
}

func TestProcess(t *testing.T) {
	tests := []struct {
		in   string
		out  int
		out2 int
	}{
		{"", 1, 1},
		{">", 2, 2},
		{"^>v<", 4, 3},
		{"^v^v^v^v^v", 2, 11},
	}

	for _, tt := range tests {
		h, h2, err := process(tt.in)
		if err != nil {
			t.Errorf("process(%q) = error %s, want %d, %d", tt.in, err, tt.out, tt.out2)
		} else if h != tt.out || h2 != tt.out2 {
			t.Errorf("process(%q) = %d, %d, want %d, %d", tt.in, h, h2, tt.out, tt.out2)
		}
	}
}

func TestProcessError(t *testing.T) {
	tests := []string{
		" ",
		"A",
		"☃",
	}

	for _, tt := range tests {
		h, h2, err := process(tt)
		if err == nil {
			t.Errorf("process(%q) = %d, %d, want error", tt, h, h2)
		}
	}
}
