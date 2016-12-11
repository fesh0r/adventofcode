package main

import (
	"strings"
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

func TestRepeatedPair(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{"", false},
		{"xxx", false},
		{"ababab", true},
		{"qjhvhtzxzqqjkmpb", true},
		{"xxyxx", true},
		{"uurcxstgmygtbstg", true},
		{"ieodomkazucvgmuy", false},
	}

	for _, tt := range tests {
		i := hasRepeatedPair(tt.in)
		if i != tt.out {
			t.Errorf("repeatedPair(%q) = %t, want %t", tt.in, i, tt.out)
		}
	}
}

func TestRepeatWithGap(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{"", false},
		{"xx", false},
		{"xyx", true},
		{"xyyx", false},
		{"qjhvhtzxzqqjkmpb", true},
		{"xxyxx", true},
		{"uurcxstgmygtbstg", false},
		{"ieodomkazucvgmuy", true},
	}

	for _, tt := range tests {
		i := hasRepeatWithGap(tt.in)
		if i != tt.out {
			t.Errorf("hasRepeatWithGap(%q) = %t, want %t", tt.in, i, tt.out)
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

func TestCheckString2(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{"qjhvhtzxzqqjkmpb", true},
		{"xxyxx", true},
		{"uurcxstgmygtbstg", false},
		{"ieodomkazucvgmuy", false},
	}

	for _, tt := range tests {
		i := checkString2(tt.in)
		if i != tt.out {
			t.Errorf("checkString2(%q) = %t, want %t", tt.in, i, tt.out)
		}
	}
}

func TestProcess(t *testing.T) {
	tests := []struct {
		in   string
		out  int
		out2 int
	}{
		{"ugknbfddgicrmopn\nqjhvhtzxzqqjkmpb\n", 1, 1},
	}

	for _, tt := range tests {
		n, n2 := process(strings.NewReader(tt.in))
		if n != tt.out || n2 != tt.out2 {
			t.Errorf("process(%q) = %d, %d, want %d, %d", tt.in, n, n2, tt.out, tt.out2)
		}
	}
}
