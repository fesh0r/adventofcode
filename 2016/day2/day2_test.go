package main

import (
	"strings"
	"testing"
)

func TestMove(t *testing.T) {
	tests := []struct {
		inPos position
		inDir rune
		out   position
	}{
		{position{0, 0}, 'U', position{0, 0}},
		{position{0, 0}, 'D', position{0, 1}},
		{position{1, 1}, 'R', position{2, 1}},
	}

	for _, tt := range tests {
		pos := tt.inPos
		err := pos.move(tt.inDir)
		if err != nil {
			t.Errorf("%v.move(%q) = error %s, want %v", tt.inPos, tt.inDir, err, tt.out)
		} else if pos.x != tt.out.x || pos.y != tt.out.y {
			t.Errorf("%v.move(%q) = %v, want %v", tt.inPos, tt.inDir, pos, tt.out)
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
		var pos position
		err := pos.move(tt)
		if err == nil {
			t.Errorf("pos.move(%q) = %v, want error", tt, pos)
		}
	}
}

func TestCode(t *testing.T) {
	tests := []struct {
		in  position
		out string
	}{
		{position{0, 0}, "1"},
		{position{2, 1}, "6"},
		{position{1, 1}, "5"},
	}

	for _, tt := range tests {
		s, err := tt.in.code()
		if err != nil {
			t.Errorf("%v.code() = error %s, want %q", tt.in, err, tt.out)
		} else if s != tt.out {
			t.Errorf("%v.code() = %q, want %q", tt.in, s, tt.out)
		}
	}
}

func TestCodeError(t *testing.T) {
	tests := []position{
		{-1, 0},
		{3, 1},
	}

	for _, tt := range tests {
		s, err := tt.code()
		if err == nil {
			t.Errorf("%v.code() = %q, want error", tt, s)
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
		c, err := process(strings.NewReader(tt.in))
		if err != nil {
			t.Errorf("process(%q) = error %s, want %q", tt.in, err, tt.out)
		} else if c != tt.out {
			t.Errorf("process(%q) = %q, want %q", tt.in, c, tt.out)
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
		c, err := process(strings.NewReader(tt))
		if err == nil {
			t.Errorf("process(%q) = %q, want error", tt, c)
		}
	}
}
