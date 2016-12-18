package main

import (
	"strings"
	"testing"
)

func TestHasAbba(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{"abba", true},
		{"mnop", false},
		{"qrst", false},
		{"abcd", false},
		{"bddb", true},
		{"xyyx", true},
		{"aaaa", false},
		{"qwer", false},
		{"tyui", false},
		{"ioxxoj", true},
		{"asdfgh", false},
		{"zxcvbn", false},
	}

	for _, tt := range tests {
		r := hasAbba(tt.in)
		if r != tt.out {
			t.Errorf("hasAbba(%q) = %v, want %v", tt.in, r, tt.out)
		}
	}
}

func TestCheckLine(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{
			"abba[mnop]qrst",
			true,
		},
		{
			"abcd[bddb]xyyx",
			false,
		},
		{
			"aaaa[qwer]tyui",
			false,
		},
		{
			"ioxxoj[asdfgh]zxcvbn",
			true,
		},
		{
			"abba[mnop]qrstabcd[bddb]xyyx",
			false,
		},
		{
			"aaaa[qwer]tyuiioxxoj[asdfgh]zxcvbn",
			true,
		},
		{
			"abba[mnop]",
			true,
		},
		{
			"[mnop]qrst",
			false,
		},
	}

	for _, tt := range tests {
		c := checkLine(tt.in)
		if c != tt.out {
			t.Errorf("checkLine(%q) = %v, want %v", tt.in, c, tt.out)
		}
	}
}

func TestProcess(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{
			"abba[mnop]qrst\nabcd[bddb]xyyx\naaaa[qwer]tyui\nioxxoj[asdfgh]zxcvbn",
			2,
		},
		{
			"abba[mnop]qrstabcd[bddb]xyyx\naaaa[qwer]tyuiioxxoj[asdfgh]zxcvbn",
			1,
		},
	}

	for _, tt := range tests {
		c, err := process(strings.NewReader(tt.in))
		if err != nil {
			t.Errorf("process(%q) = error %s, want %d", tt.in, err, tt.out)
		} else if c != tt.out {
			t.Errorf("process(%q) = %d, want %d", tt.in, c, tt.out)
		}
	}
}
