package main

import (
	"reflect"
	"testing"
)

func getGliderBoard() *board {
	return &board{
		[][]bool{
			{false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false},
			{false, false, false, true, false, false, false},
			{false, false, false, false, true, false, false},
			{false, false, true, true, true, false, false},
			{false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false},
		}, 5, 5}
}

func getGliderLife() *life {
	return &life{getGliderBoard(), newBoard(5, 5), 5, 5}
}

func getGliderString() string {
	return ".....\n..#..\n...#.\n.###.\n....."
}

func TestNewBoard(t *testing.T) {
	tests := []struct {
		inX, inY int
		out      board
	}{
		{2, 2, board{
			[][]bool{
				{false, false, false, false},
				{false, false, false, false},
				{false, false, false, false},
				{false, false, false, false},
			}, 2, 2},
		},
		{1, 4, board{
			[][]bool{
				{false, false, false},
				{false, false, false},
				{false, false, false},
				{false, false, false},
				{false, false, false},
				{false, false, false},
			}, 1, 4},
		},
	}

	for _, tt := range tests {
		board := newBoard(tt.inX, tt.inY)
		if !reflect.DeepEqual(board, &tt.out) {
			t.Errorf("newBoard(%d,%d) = %#v, want %#v", tt.inX, tt.inY, board, tt.out)
		}
	}
}

func TestBoard_String(t *testing.T) {
	tests := []struct {
		in  board
		out string
	}{
		{
			board{
				[][]bool{
					{false, false, false},
					{false, true, false},
					{false, false, false},
				}, 1, 1},
			"#",
		},
		{
			*getGliderBoard(),
			".....\n..#..\n...#.\n.###.\n.....",
		},
	}

	for _, tt := range tests {
		out := tt.in.String()
		if out != tt.out {
			t.Errorf("board.String() = %q, want %q", out, tt.out)
		}
	}
}

func TestBoard_On(t *testing.T) {
	tests := []struct {
		inX, inY int
		out      bool
	}{
		{1, 1, false},
		{3, 2, true},
		{3, 3, true},
	}

	for _, tt := range tests {
		b := getGliderBoard()
		out := b.on(tt.inX, tt.inY)
		if out != tt.out {
			t.Errorf("board.on(%d,%d) = %v, want %v", tt.inX, tt.inY, out, tt.out)
		}
	}
}

func TestBoard_Set(t *testing.T) {
	tests := []struct {
		inX, inY int
		inOn     bool
		out      bool
	}{
		{0, 0, false, false},
		{2, 1, false, false},
		{2, 2, true, true},
		{3, 2, true, true},
	}

	for _, tt := range tests {
		b := getGliderBoard()
		b.set(tt.inX, tt.inY, tt.inOn)
		out := b.on(tt.inX, tt.inY)
		if out != tt.out {
			t.Errorf("board.set(%d,%d,%v) = %v, want %v", tt.inX, tt.inY, tt.inOn, out, tt.out)
		}
	}
}

func TestBoard_Next(t *testing.T) {
	tests := []struct {
		inX, inY int
		out      bool
	}{
		{0, 0, false},
		{2, 1, false},
		{2, 2, false},
		{3, 2, true},
	}

	for _, tt := range tests {
		b := getGliderBoard()
		out := b.next(tt.inX, tt.inY)
		if out != tt.out {
			t.Errorf("board.next(%d,%d) = %v, want %v", tt.inX, tt.inY, out, tt.out)
		}
	}
}

func TestNewLife(t *testing.T) {
	tests := []struct {
		in  string
		out life
	}{
		{
			".....\n..#..\n...#.\n.###.\n.....",
			*getGliderLife(),
		},
	}

	for _, tt := range tests {
		life := newLife(tt.in)
		if !reflect.DeepEqual(life, &tt.out) {
			t.Errorf("newLife(%q) = %#v, want %#v", tt.in, life.a, tt.out.a)
		}
	}
}

func TestLife_String(t *testing.T) {
	tests := []struct {
		in  life
		out string
	}{
		{
			*getGliderLife(),
			getGliderString(),
		},
	}

	for _, tt := range tests {
		out := tt.in.String()
		if out != tt.out {
			t.Errorf("life.String() = %q, want %q", out, tt.out)
		}
	}
}

func TestLife_Next(t *testing.T) {
	tests := []struct {
		in  life
		out string
	}{
		{
			*getGliderLife(),
			".....\n.....\n.#.#.\n..##.\n..#..",
		},
		{
			*newLife("...\n###\n..."),
			".#.\n.#.\n.#.",
		},
		{
			*newLife(".#.#.#\n...##.\n#....#\n..#...\n#.#..#\n####.."),
			"..##..\n..##.#\n...##.\n......\n#.....\n#.##..",
		},
		{
			*newLife("..##..\n..##.#\n...##.\n......\n#.....\n#.##.."),
			"..###.\n......\n..###.\n......\n.#....\n.#....",
		},
		{
			*newLife("..###.\n......\n..###.\n......\n.#....\n.#...."),
			"...#..\n......\n...#..\n..##..\n......\n......",
		},
		{
			*newLife("...#..\n......\n...#..\n..##..\n......\n......"),
			"......\n......\n..##..\n..##..\n......\n......",
		},
	}

	for _, tt := range tests {
		tt.in.next()
		out := tt.in.String()
		if out != tt.out {
			t.Errorf("life.next() = %q, want %q", out, tt.out)
		}
	}
}

func TestLife_On(t *testing.T) {
	tests := []struct {
		in  life
		out int
	}{
		{
			*getGliderLife(),
			5,
		},
		{
			*newLife(".#.#.#\n...##.\n#....#\n..#...\n#.#..#\n####.."),
			15,
		},
		{
			*newLife("......\n......\n..##..\n..##..\n......\n......"),
			4,
		},
	}

	for _, tt := range tests {
		out := tt.in.on()
		if out != tt.out {
			t.Errorf("life.on() = %d, want %d", out, tt.out)
		}
	}
}

func TestLife_Fixed(t *testing.T) {
	tests := []struct {
		in  life
		out string
	}{
		{
			*getGliderLife(),
			"#...#\n..#..\n...#.\n.###.\n#...#",
		},
		{
			*newLife(".#.#.#\n...##.\n#....#\n..#...\n#.#..#\n####.."),
			"##.#.#\n...##.\n#....#\n..#...\n#.#..#\n####.#",
		},
		{
			*newLife("......\n......\n..##..\n..##..\n......\n......"),
			"#....#\n......\n..##..\n..##..\n......\n#....#",
		},
	}

	for _, tt := range tests {
		tt.in.fixed()
		out := tt.in.String()
		if out != tt.out {
			t.Errorf("life.on() = %q, want %q", out, tt.out)
		}
	}
}

func TestProcess(t *testing.T) {
	tests := []struct {
		in      string
		inGen   int
		inFixed bool
		out     int
	}{
		{
			getGliderString(), 4, false,
			5,
		},
		{
			".#.#.#\n...##.\n#....#\n..#...\n#.#..#\n####..", 4, false,
			4,
		},
		{
			getGliderString(), 4, true,
			11,
		},
		{
			".#.#.#\n...##.\n#....#\n..#...\n#.#..#\n####..", 5, true,
			17,
		},
	}

	for _, tt := range tests {
		out, _ := process(tt.in, tt.inGen, tt.inFixed)
		if out != tt.out {
			t.Errorf("process(%d,%v) = %d, want %d", tt.inGen, tt.inFixed, out, tt.out)
		}
	}
}
