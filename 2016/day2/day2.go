package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Position struct {
	X, Y int
}

var direction = map[rune]Position{
	'U': {0, -1},
	'R': {1, 0},
	'D': {0, 1},
	'L': {-1, 0},
}

type Layout struct {
	L    [][]string
	S    Position
	W, H int
}

var layout = []Layout{
	{
		[][]string{
			[]string{" ", " ", " ", " ", " "},
			[]string{" ", "1", "2", "3", " "},
			[]string{" ", "4", "5", "6", " "},
			[]string{" ", "7", "8", "9", " "},
			[]string{" ", " ", " ", " ", " "},
		},
		Position{2, 2},
		4, 4,
	},
}

type Pad struct {
	P Position
	Layout
}

func NewPad(l int) Pad {
	return Pad{layout[l].S, layout[l]}
}

func (pad *Pad) Move(dir rune) error {
	var err error

	change, ok := direction[dir]
	if !ok {
		err = fmt.Errorf("invalid direction %q", dir)
		return err
	}

	var newPos Position
	newPos.X = pad.P.X + change.X
	newPos.Y = pad.P.Y + change.Y

	if pad.L[newPos.Y][newPos.X] != " " {
		pad.P = newPos
	}

	return nil
}

func (pad *Pad) Code() (string, error) {
	var err error

	var code string

	if pad.P.X < 0 || pad.P.X > pad.W || pad.P.Y < 0 || pad.P.Y > pad.H {
		err = fmt.Errorf("invalid position %v", pad.P)
		return "", err
	}

	code = pad.L[pad.P.Y][pad.P.X]
	if code == " " {
		err = fmt.Errorf("invalid position %v", pad.P)
		return "", err
	}

	return code, nil
}

func Process(f io.Reader, l int) (string, error) {
	var err error

	pad := NewPad(l)
	var code string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()

		for _, d := range s {
			err = pad.Move(d)
			if err != nil {
				return "", err
			}
		}

		var curCode string
		curCode, err = pad.Code()
		if err != nil {
			return "", err
		}

		code += curCode
	}

	return code, nil
}

func run() int {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "%s filename\n", os.Args[0])
		return 1
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	defer f.Close()

	code, err := Process(f, 0)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Printf("code: %s\n", code)
	return 0
}

func main() {
	os.Exit(run())
}
