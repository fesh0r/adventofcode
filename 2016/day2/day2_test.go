package main

import (
	"strings"
	"testing"
)

func TestMove(t *testing.T) {
	tests := []struct {
		inL    int
		inPos  Position
		inDir  rune
		outPos Position
	}{
		{0, Position{1, 1}, 'U', Position{1, 1}},
		{0, Position{1, 1}, 'D', Position{1, 2}},
		{0, Position{2, 2}, 'R', Position{3, 2}},
	}

	for _, tt := range tests {
		pad := NewPad(tt.inL)
		pad.P = tt.inPos
		err := pad.Move(tt.inDir)
		if err != nil {
			t.Errorf("%v%d.Move(%q) = error %s, want %v", tt.inPos, tt.inL, tt.inDir, err, tt.outPos)
		} else if pad.P.X != tt.outPos.X || pad.P.Y != tt.outPos.Y {
			t.Errorf("%v%d.Move(%q) = %v, want %v", tt.inPos, tt.inL, tt.inDir, pad.P, tt.outPos)
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
		pad := NewPad(0)
		err := pad.Move(tt)
		if err == nil {
			t.Errorf("pos.Move(%q) = %v, want error", tt, pad.P)
		}
	}
}

func TestCode(t *testing.T) {
	tests := []struct {
		inL int
		in  Position
		out string
	}{
		{0, Position{1, 1}, "1"},
		{0, Position{3, 2}, "6"},
		{0, Position{2, 2}, "5"},
	}

	for _, tt := range tests {
		pad := NewPad(tt.inL)
		pad.P = tt.in
		s, err := pad.Code()
		if err != nil {
			t.Errorf("%v%d.Code() = error %s, want %q", tt.in, tt.inL, err, tt.out)
		} else if s != tt.out {
			t.Errorf("%v%d.Code() = %q, want %q", tt.in, tt.inL, s, tt.out)
		}
	}
}

func TestCodeError(t *testing.T) {
	tests := []Position{
		{0, 0},
		{-1, 1},
		{4, 2},
	}

	for _, tt := range tests {
		pad := NewPad(0)
		pad.P = tt
		s, err := pad.Code()
		if err == nil {
			t.Errorf("%v.Code() = %q, want error", tt, s)
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
	}

	for _, tt := range tests {
		c, err := Process(strings.NewReader(tt.in), tt.inL)
		if err != nil {
			t.Errorf("Process(%q,%d) = error %s, want %q", tt.in, tt.inL, err, tt.out)
		} else if c != tt.out {
			t.Errorf("Process(%q,%d) = %q, want %q", tt.in, tt.inL, c, tt.out)
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
		c, err := Process(strings.NewReader(tt), 0)
		if err == nil {
			t.Errorf("Process(%q) = %q, want error", tt, c)
		}
	}
}
