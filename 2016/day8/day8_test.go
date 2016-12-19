package main

import (
	"reflect"
	"strings"
	"testing"
)

func getHoDisplay() *display {
	return &display{
		[][]bool{
			{true, false, true, false, true, true, true},
			{true, true, true, false, true, false, true},
			{true, false, true, false, true, true, true},
		},
		7, 3}
}

func TestNewDisplay(t *testing.T) {
	tests := []struct {
		inX, inY int
		out      display
	}{
		{
			2, 2,
			display{
				[][]bool{
					{false, false},
					{false, false},
				},
				2, 2,
			},
		},
		{
			7, 3,
			display{
				[][]bool{
					{false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false},
				},
				7, 3,
			},
		},
	}

	for _, tt := range tests {
		d := newDisplay(tt.inX, tt.inY)
		if !reflect.DeepEqual(d, &tt.out) {
			t.Errorf("newDisplay(%d, %d) = %#v, want %#v", tt.inX, tt.inY, d, tt.out)
		}
	}
}

func TestDisplay_String(t *testing.T) {
	tests := []struct {
		in  display
		out string
	}{
		{
			display{
				[][]bool{
					{true, false},
					{true, false},
				},
				2, 2,
			},
			"#.\n#.",
		},
		{
			*getHoDisplay(),
			"#.#.###\n###.#.#\n#.#.###",
		},
	}

	for _, tt := range tests {
		out := tt.in.String()
		if out != tt.out {
			t.Errorf("display.String() = %q, want %q", out, tt.out)
		}
	}
}

func TestDisplay_on(t *testing.T) {
	tests := []struct {
		inX, inY int
		out      bool
	}{
		{0, 0, true},
		{1, 0, false},
		{1, 1, true},
		{3, 0, false},
		{6, 2, true},
	}

	for _, tt := range tests {
		d := getHoDisplay()
		out := d.on(tt.inX, tt.inY)
		if out != tt.out {
			t.Errorf("display.on(%d,%d) = %v, want %v", tt.inX, tt.inY, out, tt.out)
		}
	}
}

func TestDisplay_onCount(t *testing.T) {
	tests := []struct {
		in  display
		out int
	}{
		{
			display{
				[][]bool{
					{true, false},
					{true, false},
				},
				2, 2,
			},
			2,
		},
		{
			*getHoDisplay(),
			15,
		},
	}

	for _, tt := range tests {
		out := tt.in.onCount()
		if out != tt.out {
			t.Errorf("display.onCount() = %d, want %d", out, tt.out)
		}
	}
}

func TestDisplay_rect(t *testing.T) {
	tests := []struct {
		inD      display
		inW, inH int
		out      string
	}{
		{
			*getHoDisplay(), 1, 1,
			"#.#.###\n###.#.#\n#.#.###",
		},
		{
			*getHoDisplay(), 4, 2,
			"#######\n#####.#\n#.#.###",
		},
		{
			*newDisplay(7, 3), 3, 2,
			"###....\n###....\n.......",
		},
	}

	for _, tt := range tests {
		d := tt.inD
		d.rect(tt.inW, tt.inH)
		out := d.String()
		if out != tt.out {
			t.Errorf("display.rect(%d,%d) = %q, want %q", tt.inW, tt.inH, out, tt.out)
		}
	}
}

func TestDisplay_rotateRow(t *testing.T) {
	tests := []struct {
		inD      display
		inY, inC int
		out      string
	}{
		{
			display{
				[][]bool{
					{true, false},
					{false, false},
				},
				2, 2,
			}, 0, 1,
			".#\n..",
		},
		{
			display{
				[][]bool{
					{true, false, true, false, false, false, false},
					{true, true, true, false, false, false, false},
					{false, true, false, false, false, false, false},
				},
				7, 3,
			}, 0, 4,
			"....#.#\n###....\n.#.....",
		},
		{
			*getHoDisplay(), 0, 1,
			"##.#.##\n###.#.#\n#.#.###",
		},
		{
			*getHoDisplay(), 2, 4,
			"#.#.###\n###.#.#\n.####.#",
		},
	}

	for _, tt := range tests {
		d := tt.inD
		d.rotateRow(tt.inY, tt.inC)
		out := d.String()
		if out != tt.out {
			t.Errorf("display.rotateRow(%d,%d) = %q, want %q", tt.inY, tt.inC, out, tt.out)
		}
	}
}

func TestDisplay_rotateColumn(t *testing.T) {
	tests := []struct {
		inD      display
		inX, inC int
		out      string
	}{
		{
			display{
				[][]bool{
					{true, false},
					{false, false},
				},
				2, 2,
			}, 0, 1,
			"..\n#.",
		},
		{
			display{
				[][]bool{
					{true, true, true, false, false, false, false},
					{true, true, true, false, false, false, false},
					{false, false, false, false, false, false, false},
				},
				7, 3,
			}, 1, 1,
			"#.#....\n###....\n.#.....",
		},
		{
			display{
				[][]bool{
					{false, false, false, false, true, false, true},
					{true, true, true, false, false, false, false},
					{false, true, false, false, false, false, false},
				},
				7, 3,
			}, 1, 1,
			".#..#.#\n#.#....\n.#.....",
		},
		{
			*getHoDisplay(), 0, 1,
			"#.#.###\n###.#.#\n#.#.###",
		},
		{
			*getHoDisplay(), 1, 4,
			"#.#.###\n#.#.#.#\n###.###",
		},
	}

	for _, tt := range tests {
		d := tt.inD
		d.rotateColumn(tt.inX, tt.inC)
		out := d.String()
		if out != tt.out {
			t.Errorf("display.rotateRow(%d,%d) = \n%v\n, want \n%v\n", tt.inX, tt.inC, out, tt.out)
		}
	}
}

func TestProcess(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{
			"rect 3x2\nrotate column x=1 by 1\nrotate row y=0 by 4\nrotate column x=1 by 1",
			6,
		},
	}

	for _, tt := range tests {
		c, err := process(strings.NewReader(tt.in), 7, 3)
		if err != nil {
			t.Errorf("process(%q) = error %s, want %d", tt.in, err, tt.out)
		} else if c != tt.out {
			t.Errorf("process(%q) = %d, want %d", tt.in, c, tt.out)
		}
	}
}
