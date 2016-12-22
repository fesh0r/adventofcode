package main

import "testing"

func TestParseInput(t *testing.T) {
	tests := []struct {
		in  string
		out character
	}{
		{
			"Hit Points: 1\nDamage: 2\nArmor: 3\n",
			character{1, 2, 3},
		},
	}

	for _, tt := range tests {
		c, err := parseInput(tt.in)
		if err != nil {
			t.Errorf("parseInput(%q) = error %s, want %v", tt.in, err, tt.out)
		} else if c != tt.out {
			t.Errorf("parseInput(%q) = %v, want %v", tt.in, c, tt.out)
		}
	}
}

func TestFight(t *testing.T) {
	tests := []struct {
		inP, inB character
		out      bool
	}{
		{
			character{8, 5, 5}, character{12, 7, 2},
			true,
		},
	}

	for _, tt := range tests {
		w := fight(tt.inP, tt.inB)
		if w != tt.out {
			t.Errorf("fight(%v, %v) = %v, want %v", tt.inP, tt.inB, w, tt.out)
		}
	}
}

func TestProcess(t *testing.T) {
	tests := []struct {
		in        string
		inH       int
		out, out2 int
	}{
		{
			"Hit Points: 12\nDamage: 7\nArmor: 2\n", 10,
			65, 148,
		},
		{
			"Hit Points: 12\nDamage: 7\nArmor: 2\n", 100,
			8, -1,
		},
		{
			"Hit Points: 12\nDamage: 7\nArmor: 2\n", 5,
			116, 255,
		},
	}

	for _, tt := range tests {
		c, c2, err := process(tt.in, tt.inH)
		if err != nil {
			t.Errorf("process(%q, %d) = error %s, want %d, %d", tt.in, tt.inH, err, tt.out, tt.out2)
		} else if c != tt.out || c2 != tt.out2 {
			t.Errorf("process(%q, %d) = %d, %d, want %d, %d", tt.in, tt.inH, c, c2, tt.out, tt.out2)
		}
	}
}
