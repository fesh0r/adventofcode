package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type position struct {
	x, y int
}

var direction = map[rune]position{
	'U': {0, -1},
	'R': {1, 0},
	'D': {0, 1},
	'L': {-1, 0},
}

type layout struct {
	l    [][]string
	s    position
	w, h int
}

var layouts = []layout{
	{
		[][]string{
			[]string{" ", " ", " ", " ", " "},
			[]string{" ", "1", "2", "3", " "},
			[]string{" ", "4", "5", "6", " "},
			[]string{" ", "7", "8", "9", " "},
			[]string{" ", " ", " ", " ", " "},
		},
		position{2, 2},
		4, 4,
	},
	{
		[][]string{
			[]string{" ", " ", " ", " ", " ", " ", " "},
			[]string{" ", " ", " ", "1", " ", " ", " "},
			[]string{" ", " ", "2", "3", "4", " ", " "},
			[]string{" ", "5", "6", "7", "8", "9", " "},
			[]string{" ", " ", "A", "B", "C", " ", " "},
			[]string{" ", " ", " ", "D", " ", " ", " "},
			[]string{" ", " ", " ", " ", " ", " ", " "},
		},
		position{1, 3},
		7, 7,
	},
}

type pad struct {
	p position
	layout
}

func newPad(l int) pad {
	return pad{layouts[l].s, layouts[l]}
}

func (pad *pad) move(dir rune) error {
	var err error

	change, ok := direction[dir]
	if !ok {
		err = fmt.Errorf("invalid direction %q", dir)
		return err
	}

	var newPos position
	newPos.x = pad.p.x + change.x
	newPos.y = pad.p.y + change.y

	if pad.l[newPos.y][newPos.x] != " " {
		pad.p = newPos
	}

	return nil
}

func (pad *pad) code() (string, error) {
	var err error

	var code string

	if pad.p.x < 0 || pad.p.x > pad.w || pad.p.y < 0 || pad.p.y > pad.h {
		err = fmt.Errorf("invalid position %v", pad.p)
		return "", err
	}

	code = pad.l[pad.p.y][pad.p.x]
	if code == " " {
		err = fmt.Errorf("invalid position %v", pad.p)
		return "", err
	}

	return code, nil
}

func process(f io.Reader, l int) (string, error) {
	var err error

	pad := newPad(l)
	var code string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()

		for _, d := range s {
			err = pad.move(d)
			if err != nil {
				return "", err
			}
		}

		var curCode string
		curCode, err = pad.code()
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

	code, err := process(f, 0)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	_, err = f.Seek(0, 0)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	code2, err := process(f, 1)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Printf("code: %s\ncode2: %s\n", code, code2)
	return 0
}

func main() {
	os.Exit(run())
}
