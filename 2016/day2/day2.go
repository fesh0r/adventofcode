package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
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

func (pos *Position) Move(dir rune) error {
	var err error

	change, ok := direction[dir]
	if !ok {
		err = fmt.Errorf("invalid direction %q", dir)
		return err
	}

	pos.X += change.X
	if pos.X < 0 {
		pos.X = 0
	}
	if pos.X > 2 {
		pos.X = 2
	}

	pos.Y += change.Y
	if pos.Y < 0 {
		pos.Y = 0
	}
	if pos.Y > 2 {
		pos.Y = 2
	}

	return nil
}

func (pos *Position) Code() (string, error) {
	var err error

	var code string

	if pos.X < 0 || pos.X > 2 || pos.Y < 0 || pos.Y > 2 {
		err = fmt.Errorf("invalid position %v", pos)
		return "", err
	}

	code = strconv.Itoa(pos.X + pos.Y*3 + 1)

	return code, nil
}

func Process(f io.Reader) (string, error) {
	var err error

	pos := Position{1, 1}
	var code string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()

		for _, d := range s {
			err = pos.Move(d)
			if err != nil {
				return "", err
			}
		}

		var curCode string
		curCode, err = pos.Code()
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

	code, err := Process(f)
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
