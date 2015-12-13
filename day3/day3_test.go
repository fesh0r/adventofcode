package main

import (
	"fmt"
	"testing"
)

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
			t.Errorf("getHouses(%q) expected: %d, got error: %s", tt.in, tt.out, err)
		} else if h != tt.out {
			t.Errorf("getHouses(%q) expected: %d, got: %d", tt.in, tt.out, h)
		}
	}
}

func TestGetHousesError(t *testing.T) {
	tests := []string{
		" ",
		"A",
	}
	for _, tt := range tests {
		h, err := getHouses(tt)
		if err == nil {
			t.Errorf("getHouses(%q) expected error, got: %d", tt, h)
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
			t.Errorf("getHousesDouble(%q) expected: %d, got error: %s", tt.in, tt.out, err)
		} else if h != tt.out {
			t.Errorf("getHousesDouble(%q) expected: %d, got: %d", tt.in, tt.out, h)
		}
	}
}

func TestGetHousesDoubleError(t *testing.T) {
	tests := []string{
		" ",
		"A",
	}
	for _, tt := range tests {
		h, err := getHousesDouble(tt)
		if err == nil {
			t.Errorf("getHousesDouble(%q) expected error, got: %d", tt, h)
		}
	}
}

func ExampleGetHouses() {
	s := "^>v<"
	h, _ := getHouses(s)
	fmt.Printf("%q => %d\n", s, h)
	// Output: "^>v<" => 4
}

func ExampleGetHousesDouble() {
	s := "^>v<"
	h, _ := getHousesDouble(s)
	fmt.Printf("%q => %d\n", s, h)
	// Output: "^>v<" => 3
}
