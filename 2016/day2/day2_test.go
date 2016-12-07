package main

import (
	"strings"
	"testing"
)

func TestMove(t *testing.T) {
	tests := []struct {
		inL    int
		inPos  position
		inDir  rune
		outPos position
	}{
		{0, position{0, 0}, 'U', position{0, 0}},
		{0, position{0, 0}, 'D', position{0, 1}},
		{0, position{1, 1}, 'R', position{2, 1}},
		{1, position{1, 1}, 'U', position{1, 1}},
		{1, position{1, 1}, 'D', position{1, 2}},
		{1, position{2, 2}, 'R', position{3, 2}},
	}

	for _, tt := range tests {
		pad := newPad(tt.inL)
		pad.p = tt.inPos
		err := pad.move(tt.inDir)
		if err != nil {
			t.Errorf("%v%d.move(%q) = error %s, want %v", tt.inPos, tt.inL, tt.inDir, err, tt.outPos)
		} else if pad.p.x != tt.outPos.x || pad.p.y != tt.outPos.y {
			t.Errorf("%v%d.move(%q) = %v, want %v", tt.inPos, tt.inL, tt.inDir, pad.p, tt.outPos)
		}
	}
}

func TestMoveError(t *testing.T) {
	tests := []rune{
		' ',
		'u',
		'0',
		'☃',
	}

	for _, tt := range tests {
		pad := newPad(0)
		err := pad.move(tt)
		if err == nil {
			t.Errorf("pos.move(%q) = %v, want error", tt, pad.p)
		}
	}
}

func TestCode(t *testing.T) {
	tests := []struct {
		inL int
		in  position
		out string
	}{
		{0, position{0, 0}, "1"},
		{0, position{2, 1}, "6"},
		{0, position{1, 1}, "5"},
		{1, position{1, 1}, "2"},
		{1, position{3, 2}, "8"},
		{1, position{2, 2}, "7"},
	}

	for _, tt := range tests {
		pad := newPad(tt.inL)
		pad.p = tt.in
		s, err := pad.code()
		if err != nil {
			t.Errorf("%v%d.code() = error %s, want %q", tt.in, tt.inL, err, tt.out)
		} else if s != tt.out {
			t.Errorf("%v%d.code() = %q, want %q", tt.in, tt.inL, s, tt.out)
		}
	}
}

func TestCodeError(t *testing.T) {
	tests := []position{
		{-1, -1},
		{-2, 0},
		{3, 1},
	}

	for _, tt := range tests {
		pad := newPad(0)
		pad.p = tt
		s, err := pad.code()
		if err == nil {
			t.Errorf("%v.code() = %q, want error", tt, s)
		}
	}
}

func TestProcess(t *testing.T) {
	tests := []struct {
		inL int
		in  string
		out string
	}{
		{0, "U", "2"},
		{0, "DDDU\nR", "56"},
		{0, "ULL\nRRDDD\nLURDL\nUUUUD", "1985"},
		{1, "ULL\nRRDDD\nLURDL\nUUUUD", "5DB3"},
	}

	for _, tt := range tests {
		c, err := process(strings.NewReader(tt.in), tt.inL)
		if err != nil {
			t.Errorf("process(%q,%d) = error %s, want %q", tt.in, tt.inL, err, tt.out)
		} else if c != tt.out {
			t.Errorf("process(%q,%d) = %q, want %q", tt.in, tt.inL, c, tt.out)
		}
	}
}

func TestProcessError(t *testing.T) {
	tests := []string{
		" ",
		"u",
		"0",
		"☃",
		"LL\nU ",
		"RL\n☃",
		"UDLR\nUr",
	}

	for _, tt := range tests {
		c, err := process(strings.NewReader(tt), 0)
		if err == nil {
			t.Errorf("process(%q) = %q, want error", tt, c)
		}
	}
}
