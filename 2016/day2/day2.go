package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
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

func (pos *position) move(dir rune) error {
	var err error

	change, ok := direction[dir]
	if !ok {
		err = fmt.Errorf("invalid direction %q", dir)
		return err
	}

	pos.x += change.x
	if pos.x < 0 {
		pos.x = 0
	}
	if pos.x > 2 {
		pos.x = 2
	}

	pos.y += change.y
	if pos.y < 0 {
		pos.y = 0
	}
	if pos.y > 2 {
		pos.y = 2
	}

	return nil
}

func (pos *position) code() (string, error) {
	var err error

	var code string

	if pos.x < 0 || pos.x > 2 || pos.y < 0 || pos.y > 2 {
		err = fmt.Errorf("invalid position %v", pos)
		return "", err
	}

	code = strconv.Itoa(pos.x + pos.y*3 + 1)

	return code, nil
}

func process(f io.Reader) (string, error) {
	var err error

	var pos position
	var code string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()

		for _, d := range s {
			err = pos.move(d)
			if err != nil {
				return "", err
			}
		}

		var curCode string
		curCode, err = pos.code()
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

	code, err := process(f)
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
