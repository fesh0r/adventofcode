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
		dir, dist, err := Parse(tt.in)
		if err != nil {
			t.Errorf("Parse(%q) = error %s, want %d, %d", tt.in, err, tt.outDir, tt.outDist)
		} else if dir != tt.outDir || dist != tt.outDist {
			t.Errorf("Parse(%q) = %d, %d, want %d, %d", tt.in, dir, dist, tt.outDir, tt.outDist)
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
		dir, dist, err := Parse(tt)
		if err == nil {
			t.Errorf("Parse(%q) = %d, %d, want error", tt, dir, dist)
		}
	}
}

func TestDistance(t *testing.T) {
	tests := []struct {
		in  Position
		out int
	}{
		{Position{0, 0}, 0},
		{Position{1, 1}, 2},
		{Position{2, -185}, 187},
	}

	for _, tt := range tests {
		dist := Distance(tt.in)
		if dist != tt.out {
			t.Errorf("Distance(%q) = %d, want %d", tt.in, dist, tt.out)
		}
	}
}

func TestProcess(t *testing.T) {
	tests := []struct {
		in        string
		out, out2 int
	}{
		{"R2, L3", 5, 0},
		{"R2, R2, R2", 2, 0},
		{"R5, L5, R5, R3", 12, 0},
		{"R8, R4, R4, R8", 8, 4},
	}

	for _, tt := range tests {
		d, d2, err := Process(tt.in)
		if err != nil {
			t.Errorf("Process(%q) = error %s, want %d, %d", tt.in, err, tt.out, tt.out2)
		} else if d != tt.out || d2 != tt.out2 {
			t.Errorf("Process(%q) = %d, %d, want %d, %d", tt.in, d, d2, tt.out, tt.out2)
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
		d, d2, err := Process(tt)
		if err == nil {
			t.Errorf("Process(%q) = %d, %d, want error", tt, d, d2)
		}
	}
}
