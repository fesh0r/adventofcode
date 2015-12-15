package main

import (
	"fmt"
	"testing"
)

func TestParseChange(t *testing.T) {
	floortests := []struct {
		in  rune
		out int
	}{
		{'(', 1},
		{')', -1},
	}

	for _, tt := range floortests {
		f, err := parseChange(tt.in)
		if err != nil {
			t.Errorf("parseChange(%q) = error %s, want %d", tt.in, err, tt.out)
		} else if f != tt.out {
			t.Errorf("parseChange(%q) = %d, want %d", tt.in, f, tt.out)
		}
	}
}

func TestParseChangeError(t *testing.T) {
	errortests := []rune{
		' ',
		']',
		'☃',
	}

	for _, tt := range errortests {
		f, err := parseChange(tt)
		if err == nil {
			t.Errorf("parseChange(%q) = %d, want error", tt, f)
		}
	}
}

func TestGetFloor(t *testing.T) {
	floortests := []struct {
		in  string
		out int
	}{
		{"", 0},
		{"(())", 0},
		{"()()", 0},
		{"(((", 3},
		{"(()(()(", 3},
		{"))(((((", 3},
		{"())", -1},
		{"))(", -1},
		{")))", -3},
		{")())())", -3},
	}

	for _, tt := range floortests {
		f, err := getFloor(tt.in)
		if err != nil {
			t.Errorf("getFloor(%q) = error %s, want %d", tt.in, err, tt.out)
		} else if f != tt.out {
			t.Errorf("getFloor(%q) = %d, want %d", tt.in, f, tt.out)
		}
	}
}

func TestGetFloorError(t *testing.T) {
	errortests := []string{
		" ",
		"(())]",
		"((☃))",
	}

	for _, tt := range errortests {
		f, err := getFloor(tt)
		if err == nil {
			t.Errorf("getFloor(%q) = %d, want error", tt, f)
		}
	}
}

func TestGetBasement(t *testing.T) {
	basementtests := []struct {
		in  string
		out int
	}{
		{"", 0},
		{")", 1},
		{"()())", 5},
	}

	for _, tt := range basementtests {
		f, err := getBasement(tt.in)
		if err != nil {
			t.Errorf("getBasement(%q) = error %s, want %d", tt.in, err, tt.out)
		} else if f != tt.out {
			t.Errorf("getBasement(%q) = %d, want %d", tt.in, f, tt.out)
		}
	}
}

func TestGetBasementError(t *testing.T) {
	errortests := []string{
		" ",
		"(())]",
		"((☃))",
	}

	for _, tt := range errortests {
		f, err := getBasement(tt)
		if err == nil {
			t.Errorf("getBasement(%q) = %d, want error", tt, f)
		}
	}
}

func ExampleGetFloor() {
	s := "())"
	f, _ := getFloor(s)
	fmt.Printf("%q => %d\n", s, f)
	// Output: "())" => -1
}

func ExampleGetBasement() {
	s := "())"
	f, _ := getBasement(s)
	fmt.Printf("%q => %d\n", s, f)
	// Output: "())" => 3
}
