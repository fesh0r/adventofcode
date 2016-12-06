package main

import (
	"fmt"
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
		x, y, err := ParseDirection(tt.in)
		if err != nil {
			t.Errorf("ParseDirection(%q) = error %s, want %d, %d", tt.in, err, tt.outX, tt.outY)
		} else if x != tt.outX || y != tt.outY {
			t.Errorf("ParseDirection(%q) = %d, %d, want %d, %d", tt.in, x, y, tt.outX, tt.outY)
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
		x, y, err := ParseDirection(tt)
		if err == nil {
			t.Errorf("ParseDirection(%q) = %d, %d, want error", tt, x, y)
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
		h, err := GetHouses(tt.in)
		if err != nil {
			t.Errorf("GetHouses(%q) = error %s, want %d", tt.in, err, tt.out)
		} else if h != tt.out {
			t.Errorf("GetHouses(%q) = %d, want %d", tt.in, h, tt.out)
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
		h, err := GetHouses(tt)
		if err == nil {
			t.Errorf("GetHouses(%q) = %d, want error", tt, h)
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
		h, err := GetHousesDouble(tt.in)
		if err != nil {
			t.Errorf("GetHousesDouble(%q) = error %s, want %d", tt.in, err, tt.out)
		} else if h != tt.out {
			t.Errorf("GetHousesDouble(%q) = %d, want %d", tt.in, h, tt.out)
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
		h, err := GetHousesDouble(tt)
		if err == nil {
			t.Errorf("GetHousesDouble(%q) = %d, want error", tt, h)
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
		h, h2, err := Process(tt.in)
		if err != nil {
			t.Errorf("Process(%q) = error %s, want %d, %d", tt.in, err, tt.out, tt.out2)
		} else if h != tt.out || h2 != tt.out2 {
			t.Errorf("Process(%q) = %d, %d, want %d, %d", tt.in, h, h2, tt.out, tt.out2)
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
		h, h2, err := Process(tt)
		if err == nil {
			t.Errorf("Process(%q) = %d, %d, want error", tt, h, h2)
		}
	}
}

func ExampleGetHouses() {
	s := "^>v<"
	h, _ := GetHouses(s)
	fmt.Printf("%q => %d\n", s, h)
	// Output: "^>v<" => 4
}

func ExampleGetHousesDouble() {
	s := "^>v<"
	h, _ := GetHousesDouble(s)
	fmt.Printf("%q => %d\n", s, h)
	// Output: "^>v<" => 3
}

func ExampleProcess() {
	s := "^>v<"
	h, h2, _ := Process(s)
	fmt.Printf("%q => %d, %d\n", s, h, h2)
	// Output: "^>v<" => 4, 3
}
