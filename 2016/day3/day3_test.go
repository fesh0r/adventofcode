package main

import (
	"strings"
	"testing"
)

func TestProcess(t *testing.T) {
	tests := []struct {
		in        string
		out, out2 int
	}{
		{"5 10 25", 0, 0},
		{"5 10 14", 1, 0},
		{"5 10 14\n5 10 13\n5 10 25", 2, 3},
		{"    5 10  14\n5      10 13\n 5  10   25\n", 2, 3},
		{"101 301 501\n102 302 502\n103 303 503\n201 401 601\n202 402 602\n203 403 603", 3, 6},
	}

	for _, tt := range tests {
		l, l2, err := process(strings.NewReader(tt.in))
		if err != nil {
			t.Errorf("process(%q) = error %s, want %d, %d", tt.in, err, tt.out, tt.out2)
		} else if l != tt.out || l2 != tt.out2 {
			t.Errorf("process(%q) = %d, %d, want %d, %d", tt.in, l, l2, tt.out, tt.out2)
		}
	}
}

func TestProcessError(t *testing.T) {
	tests := []string{
		" ",
		"u",
		"0",
		"☃",
		"103 104",
		"103 104  105\n ☃",
		"103 104    105\nA",
	}

	for _, tt := range tests {
		l, l2, err := process(strings.NewReader(tt))
		if err == nil {
			t.Errorf("process(%q) = %d, %d, want error", tt, l, l2)
		}
	}
}
