package main

import "testing"

func TestNextLookSay(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"1", "11"},
		{"21", "1211"},
		{"111221", "312211"},
	}

	for _, tt := range tests {
		v := nextLookSay(tt.in)
		if v != tt.out {
			t.Errorf("nextLookSay(%q) = %q, want %q", tt.in, v, tt.out)
		}
	}
}

func TestRepeatLookSay(t *testing.T) {
	tests := []struct {
		in  string
		inC int
		out string
	}{
		{"1", 0, "1"},
		{"1", 1, "11"},
		{"1211", 1, "111221"},
		{"1", 5, "312211"},
	}

	for _, tt := range tests {
		v, err := repeatLookSay(tt.in, tt.inC)
		if err != nil {
			t.Errorf("repeatLookSay(%q, %d) = error %s, want %d", tt.in, tt.inC, err, tt.out)
		} else if v != tt.out {
			t.Errorf("repeatLookSay(%q, %d) = %s, want %s", tt.in, tt.inC, v, tt.out)
		}
	}
}

func TestRepeatLookSayError(t *testing.T) {
	tests := []struct {
		in  string
		inC int
	}{
		{"", 0},
		{"A", 1},
		{"123Z", 2},
	}

	for _, tt := range tests {
		v, err := repeatLookSay(tt.in, tt.inC)
		if err == nil {
			t.Errorf("repeatLookSay(%q, %d) = %s, want error", tt.in, tt.inC, v)
		}
	}
}

func TestProcess(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{"1", 82350},
	}

	for _, tt := range tests {
		v, err := process(tt.in)
		if err != nil {
			t.Errorf("process(%q) = error %s, want %d", tt.in, err, tt.out)
		} else if v != tt.out {
			t.Errorf("process(%q) = %d, want %d", tt.in, v, tt.out)
		}
	}
}

func TestProcessError(t *testing.T) {
	tests := []string{
		"",
		"A",
		"123Z",
	}

	for _, tt := range tests {
		v, err := process(tt)
		if err == nil {
			t.Errorf("process(%q) = %d, want error", tt, v)
		}
	}
}
