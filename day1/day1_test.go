package main

import (
	"fmt"
	"testing"
)

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
			t.Errorf("getFloor(%q) expected: %d, got error: %s", tt.in, tt.out, err)
		} else if f != tt.out {
			t.Errorf("getFloor(%q) expected: %d, got: %d", tt.in, tt.out, f)
		}
	}
}

func TestGetFloorError(t *testing.T) {
	errortests := []string{
		" ",
		"(())]",
	}
	for _, tt := range errortests {
		f, err := getFloor(tt)
		if err == nil {
			t.Errorf("getFloor(%q) expected error, got: %d", tt, f)
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
			t.Errorf("getBasement(%q) expected: %d, got error: %s", tt.in, tt.out, err)
		} else if f != tt.out {
			t.Errorf("getBasement(%q) expected: %d, got: %d", tt.in, tt.out, f)
		}
	}
}

func TestGetBasementError(t *testing.T) {
	errortests := []string{
		" ",
		"(())]",
	}
	for _, tt := range errortests {
		f, err := getBasement(tt)
		if err == nil {
			t.Errorf("getBasement(%q) expected error, got: %d", tt, f)
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
