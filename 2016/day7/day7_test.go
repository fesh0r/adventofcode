package main

import (
	"reflect"
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

func TestGetAbas(t *testing.T) {
	tests := []struct {
		in  string
		out abaMap
	}{
		{
			"aba",
			abaMap{"aba": true},
		},
		{
			"aaa",
			abaMap{},
		},
		{
			"zazbz",
			abaMap{"zaz": true, "zbz": true},
		},
	}

	for _, tt := range tests {
		r := getAbas(tt.in)
		if !reflect.DeepEqual(r, tt.out) {
			t.Errorf("hasAbba(%q) = %v, want %v", tt.in, r, tt.out)
		}
	}
}

func TestCheckLine(t *testing.T) {
	tests := []struct {
		in        string
		out, out2 bool
	}{
		{
			"abba[mnop]qrst",
			true, false,
		},
		{
			"abcd[bddb]xyyx",
			false, false,
		},
		{
			"aaaa[qwer]tyui",
			false, false,
		},
		{
			"ioxxoj[asdfgh]zxcvbn",
			true, false,
		},
		{
			"abba[mnop]qrstabcd[bddb]xyyx",
			false, false,
		},
		{
			"aaaa[qwer]tyuiioxxoj[asdfgh]zxcvbn",
			true, false,
		},
		{
			"abba[mnop]",
			true, false,
		},
		{
			"[mnop]qrst",
			false, false,
		},
		{
			"aba[bab]xyz",
			false, true,
		},
		{
			"xyx[xyx]xyx",
			false, false,
		},
		{
			"aaa[kek]eke",
			false, true,
		},
		{
			"zazbz[bzb]cdb",
			false, true,
		},
	}

	for _, tt := range tests {
		c, c2 := checkLine(tt.in)
		if c != tt.out || c2 != tt.out2 {
			t.Errorf("checkLine(%q) = %v, %v, want %v, %v", tt.in, c, c2, tt.out, tt.out2)
		}
	}
}

func TestProcess(t *testing.T) {
	tests := []struct {
		in        string
		out, out2 int
	}{
		{
			"abba[mnop]qrst\nabcd[bddb]xyyx\naaaa[qwer]tyui\nioxxoj[asdfgh]zxcvbn",
			2, 0,
		},
		{
			"abba[mnop]qrstabcd[bddb]xyyx\naaaa[qwer]tyuiioxxoj[asdfgh]zxcvbn",
			1, 0,
		},
		{
			"aba[bab]xyz\nxyx[xyx]xyx\naaa[kek]eke\nzazbz[bzb]cdb",
			0, 3,
		},
	}

	for _, tt := range tests {
		c, c2, err := process(strings.NewReader(tt.in))
		if err != nil {
			t.Errorf("process(%q) = error %s, want %d, %d", tt.in, err, tt.out, tt.out2)
		} else if c != tt.out {
			t.Errorf("process(%q) = %d, %d, want %d, %d", tt.in, c, c2, tt.out, tt.out2)
		}
	}
}
