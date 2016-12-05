package main

import "testing"

func TestParse(t *testing.T) {
	tests := []struct {
		in              string
		outDir, outDist int
	}{
		{"R2", 1, 2},
		{"L185", -1, 185},
	}

	for _, tt := range tests {
		dir, dist, err := parse(tt.in)
		if err != nil {
			t.Errorf("parse(%q) = error %s, want %d, %d", tt.in, err, tt.outDir, tt.outDist)
		} else if dir != tt.outDir || dist != tt.outDist {
			t.Errorf("parse(%q) = %d, %d, want %d, %d", tt.in, dir, dist, tt.outDir, tt.outDist)
		}
	}
}

func TestParseError(t *testing.T) {
	tests := []string{
		"",
		"]",
		"☃",
		"L",
		"R-1",
	}

	for _, tt := range tests {
		dir, dist, err := parse(tt)
		if err == nil {
			t.Errorf("parse(%q) = %d, %d, want error", tt, dir, dist)
		}
	}
}

func TestDistance(t *testing.T) {
	tests := []struct {
		in  position
		out int
	}{
		{position{0, 0}, 0},
		{position{1, 1}, 2},
		{position{2, -185}, 187},
	}

	for _, tt := range tests {
		dist := distance(tt.in)
		if dist != tt.out {
			t.Errorf("distance(%q) = %d, want %d", tt.in, dist, tt.out)
		}
	}
}

func TestProcess(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{"R2, L3", 5},
		{"R2, R2, R2", 2},
		{"R5, L5, R5, R3", 12},
	}

	for _, tt := range tests {
		d, err := process(tt.in)
		if err != nil {
			t.Errorf("process(%q) = error %s, want %d", tt.in, err, tt.out)
		} else if d != tt.out {
			t.Errorf("process(%q) = %d, want %d", tt.in, d, tt.out)
		}
	}
}

func TestProcessError(t *testing.T) {
	tests := []string{
		"",
		"]",
		"☃",
		"L",
		"R-1",
		"R2 L3",
		"R2, R2, ",
	}

	for _, tt := range tests {
		d, err := process(tt)
		if err == nil {
			t.Errorf("process(%q) = %d, want error", tt, d)
		}
	}
}
