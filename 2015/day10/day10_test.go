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
			t.Errorf("repeatLookSay(%q, %d) = error %s, want %s", tt.in, tt.inC, err, tt.out)
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
		in   string
		out  int
		out2 int
	}{
		{"1", 82350, 1166642},
	}

	for _, tt := range tests {
		v, v2, err := process(tt.in)
		if err != nil {
			t.Errorf("process(%q) = error %s, want %d, %d", tt.in, err, tt.out, tt.out2)
		} else if v != tt.out || v2 != tt.out2 {
			t.Errorf("process(%q) = %d, %d, want %d, %d", tt.in, v, v2, tt.out, tt.out2)
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
		v, v2, err := process(tt)
		if err == nil {
			t.Errorf("process(%q) = %d, %d, want error", tt, v, v2)
		}
	}
}
