package main

import (
	"fmt"
	"testing"
)

func TestHasVowels(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{"", false},
		{"xyz", false},
		{"aei", true},
		{"xazegov", true},
		{"aeiouaeiouaeiou", true},
		{"ugknbfddgicrmopn", true},
		{"aaa", true},
		{"jchzalrnumimnmhp", true},
		{"haegwjzuvuyypxyu", true},
		{"dvszwmarrgswjxmb", false},
	}

	for _, tt := range tests {
		i := hasVowels(tt.in)
		if i != tt.out {
			t.Errorf("hasVowels(%q) = %t, want %t", tt.in, i, tt.out)
		}
	}
}

func TestHasRepeated(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{"", false},
		{"xyz", false},
		{"xx", true},
		{"abcdde", true},
		{"aabbccdd", true},
		{"ugknbfddgicrmopn", true},
		{"aaa", true},
		{"jchzalrnumimnmhp", false},
		{"haegwjzuvuyypxyu", true},
		{"dvszwmarrgswjxmb", true},
	}

	for _, tt := range tests {
		i := hasRepeated(tt.in)
		if i != tt.out {
			t.Errorf("hasRepeated(%q) = %t, want %t", tt.in, i, tt.out)
		}
	}
}

func TestHasNoBad(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{"", true},
		{"xyz", false},
		{"abcde", false},
		{"aabcde", false},
		{"ugknbfddgicrmopn", true},
		{"aaa", true},
		{"jchzalrnumimnmhp", true},
		{"haegwjzuvuyypxyu", false},
		{"dvszwmarrgswjxmb", true},
	}

	for _, tt := range tests {
		i := hasNoBad(tt.in)
		if i != tt.out {
			t.Errorf("hasNoBad(%q) = %t, want %t", tt.in, i, tt.out)
		}
	}
}

func TestCheckString(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{"ugknbfddgicrmopn", true},
		{"aaa", true},
		{"jchzalrnumimnmhp", false},
		{"haegwjzuvuyypxyu", false},
		{"dvszwmarrgswjxmb", false},
	}

	for _, tt := range tests {
		i := checkString(tt.in)
		if i != tt.out {
			t.Errorf("checkString(%q) = %t, want %t", tt.in, i, tt.out)
		}
	}
}

func ExampleCheckString() {
	s := "ugknbfddgicrmopn"
	n := checkString(s)
	fmt.Printf("%q => %t\n", s, n)
	// Output: "ugknbfddgicrmopn" => true
}
