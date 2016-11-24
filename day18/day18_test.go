package main

import (
	"reflect"
	"testing"
)

func getGliderBoard() *Board {
	return &Board{
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

func getGliderLife() *Life {
	return &Life{getGliderBoard(), NewBoard(5, 5), 5, 5}
}

func getGliderString() string {
	return ".....\n..#..\n...#.\n.###.\n....."
}

func TestNewBoard(t *testing.T) {
	tests := []struct {
		inX, inY int
		out      Board
	}{
		{2, 2, Board{
			[][]bool{
				{false, false, false, false},
				{false, false, false, false},
				{false, false, false, false},
				{false, false, false, false},
			}, 2, 2},
		},
		{1, 4, Board{
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
		board := NewBoard(tt.inX, tt.inY)
		if !reflect.DeepEqual(board, &tt.out) {
			t.Errorf("NewBoard(%d,%d) = %#v, want %#v", tt.inX, tt.inY, board, tt.out)
		}
	}
}

func TestBoard_String(t *testing.T) {
	tests := []struct {
		in  Board
		out string
	}{
		{
			Board{
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
			t.Errorf("Board.String() = %q, want %q", out, tt.out)
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
		out := b.On(tt.inX, tt.inY)
		if out != tt.out {
			t.Errorf("Board.On(%d,%d) = %v, want %v", tt.inX, tt.inY, out, tt.out)
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
		b.Set(tt.inX, tt.inY, tt.inOn)
		out := b.On(tt.inX, tt.inY)
		if out != tt.out {
			t.Errorf("Board.Set(%d,%d,%v) = %v, want %v", tt.inX, tt.inY, tt.inOn, out, tt.out)
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
		out := b.Next(tt.inX, tt.inY)
		if out != tt.out {
			t.Errorf("Board.Next(%d,%d) = %v, want %v", tt.inX, tt.inY, out, tt.out)
		}
	}
}

func TestNewLife(t *testing.T) {
	tests := []struct {
		in  string
		out Life
	}{
		{
			".....\n..#..\n...#.\n.###.\n.....",
			*getGliderLife(),
		},
	}

	for _, tt := range tests {
		life := NewLife(tt.in)
		if !reflect.DeepEqual(life, &tt.out) {
			t.Errorf("NewLife(%q) = %#v, want %#v", tt.in, life.a, tt.out.a)
		}
	}
}

func TestLife_String(t *testing.T) {
	tests := []struct {
		in  Life
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
			t.Errorf("Life.String() = %q, want %q", out, tt.out)
		}
	}
}

func TestLife_Next(t *testing.T) {
	tests := []struct {
		in  Life
		out string
	}{
		{
			*getGliderLife(),
			".....\n.....\n.#.#.\n..##.\n..#..",
		},
		{
			*NewLife("...\n###\n..."),
			".#.\n.#.\n.#.",
		},
		{
			*NewLife(".#.#.#\n...##.\n#....#\n..#...\n#.#..#\n####.."),
			"..##..\n..##.#\n...##.\n......\n#.....\n#.##..",
		},
		{
			*NewLife("..##..\n..##.#\n...##.\n......\n#.....\n#.##.."),
			"..###.\n......\n..###.\n......\n.#....\n.#....",
		},
		{
			*NewLife("..###.\n......\n..###.\n......\n.#....\n.#...."),
			"...#..\n......\n...#..\n..##..\n......\n......",
		},
		{
			*NewLife("...#..\n......\n...#..\n..##..\n......\n......"),
			"......\n......\n..##..\n..##..\n......\n......",
		},
	}

	for _, tt := range tests {
		tt.in.Next()
		out := tt.in.String()
		if out != tt.out {
			t.Errorf("Life.Next() = %q, want %q", out, tt.out)
		}
	}
}

func TestLife_On(t *testing.T) {
	tests := []struct {
		in  Life
		out int
	}{
		{
			*getGliderLife(),
			5,
		},
		{
			*NewLife(".#.#.#\n...##.\n#....#\n..#...\n#.#..#\n####.."),
			15,
		},
		{
			*NewLife("......\n......\n..##..\n..##..\n......\n......"),
			4,
		},
	}

	for _, tt := range tests {
		out := tt.in.On()
		if out != tt.out {
			t.Errorf("Life.On() = %d, want %d", out, tt.out)
		}
	}
}

func TestLife_Fixed(t *testing.T) {
	tests := []struct {
		in  Life
		out string
	}{
		{
			*getGliderLife(),
			"#...#\n..#..\n...#.\n.###.\n#...#",
		},
		{
			*NewLife(".#.#.#\n...##.\n#....#\n..#...\n#.#..#\n####.."),
			"##.#.#\n...##.\n#....#\n..#...\n#.#..#\n####.#",
		},
		{
			*NewLife("......\n......\n..##..\n..##..\n......\n......"),
			"#....#\n......\n..##..\n..##..\n......\n#....#",
		},
	}

	for _, tt := range tests {
		tt.in.Fixed()
		out := tt.in.String()
		if out != tt.out {
			t.Errorf("Life.On() = %q, want %q", out, tt.out)
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
			t.Errorf("process(%d,%d) = %d, want %d", tt.inGen, tt.inFixed, out, tt.out)
		}
	}
}
