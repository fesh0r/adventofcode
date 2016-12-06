package main

import (
	"strings"
	"testing"
)

func TestMove(t *testing.T) {
	tests := []struct {
		inPos Position
		inDir rune
		out   Position
	}{
		{Position{0, 0}, 'U', Position{0, 0}},
		{Position{0, 0}, 'D', Position{0, 1}},
		{Position{1, 1}, 'R', Position{2, 1}},
	}

	for _, tt := range tests {
		pos := tt.inPos
		err := pos.Move(tt.inDir)
		if err != nil {
			t.Errorf("%v.Move(%q) = error %s, want %v", tt.inPos, tt.inDir, err, tt.out)
		} else if pos.X != tt.out.X || pos.Y != tt.out.Y {
			t.Errorf("%v.Move(%q) = %v, want %v", tt.inPos, tt.inDir, pos, tt.out)
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
		var pos Position
		err := pos.Move(tt)
		if err == nil {
			t.Errorf("pos.Move(%q) = %v, want error", tt, pos)
		}
	}
}

func TestCode(t *testing.T) {
	tests := []struct {
		in  Position
		out string
	}{
		{Position{0, 0}, "1"},
		{Position{2, 1}, "6"},
		{Position{1, 1}, "5"},
	}

	for _, tt := range tests {
		s, err := tt.in.Code()
		if err != nil {
			t.Errorf("%v.Code() = error %s, want %q", tt.in, err, tt.out)
		} else if s != tt.out {
			t.Errorf("%v.Code() = %q, want %q", tt.in, s, tt.out)
		}
	}
}

func TestCodeError(t *testing.T) {
	tests := []Position{
		{-1, 0},
		{3, 1},
	}

	for _, tt := range tests {
		s, err := tt.Code()
		if err == nil {
			t.Errorf("%v.Code() = %q, want error", tt, s)
		}
	}
}

func TestProcess(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"U", "2"},
		{"DDDU\nR", "56"},
		{"ULL\nRRDDD\nLURDL\nUUUUD", "1985"},
	}

	for _, tt := range tests {
		c, err := Process(strings.NewReader(tt.in))
		if err != nil {
			t.Errorf("Process(%q) = error %s, want %q", tt.in, err, tt.out)
		} else if c != tt.out {
			t.Errorf("Process(%q) = %q, want %q", tt.in, c, tt.out)
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
		c, err := Process(strings.NewReader(tt))
		if err == nil {
			t.Errorf("Process(%q) = %q, want error", tt, c)
		}
	}
}
