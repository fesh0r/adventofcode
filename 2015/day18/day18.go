package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Board struct {
	b    [][]bool
	w, h int
}

func NewBoard(w, h int) *Board {
	b := make([][]bool, h+2)

	for i := range b {
		b[i] = make([]bool, w+2)
	}

	return &Board{b, w, h}
}

func (b *Board) String() string {
	var s string

	for y := 0; y < b.h; y++ {
		for x := 0; x < b.w; x++ {
			if b.On(x, y) {
				s += "#"
			} else {
				s += "."
			}
		}
		if y < b.h-1 {
			s += "\n"
		}
	}

	return s
}

func (b *Board) On(x, y int) bool {
	return b.b[y+1][x+1]
}

func (b *Board) Set(x, y int, alive bool) {
	b.b[y+1][x+1] = alive
}

func (b *Board) Next(x, y int) bool {
	c := 0
	for ny := -1; ny <= 1; ny++ {
		for nx := -1; nx <= 1; nx++ {
			if (nx != 0 || ny != 0) && b.On(x+nx, y+ny) {
				c++
			}
		}
	}

	return c == 3 || c == 2 && b.On(x, y)
}

type Life struct {
	a, b *Board
	w, h int
}

func NewLife(s string) *Life {
	ss := strings.Split(s, "\n")

	w := len(ss[0])
	h := len(ss)
	a := NewBoard(w, h)

	for y, row := range ss {
		for x, cell := range row {
			a.Set(x, y, cell == '#')
		}
	}

	return &Life{a, NewBoard(w, h), w, h}
}

func (l *Life) String() string {
	return l.a.String()
}

func (l *Life) Next() {
	for y := 0; y < l.h; y++ {
		for x := 0; x < l.w; x++ {
			l.b.Set(x, y, l.a.Next(x, y))
		}
	}

	l.a, l.b = l.b, l.a
}

func (l *Life) On() int {
	on := 0
	for y := 0; y < l.h; y++ {
		for x := 0; x < l.w; x++ {
			if l.a.On(x, y) {
				on++
			}
		}
	}
	return on
}

func (l *Life) Fixed() {
	l.a.Set(0, 0, true)
	l.a.Set(l.w-1, 0, true)
	l.a.Set(0, l.h-1, true)
	l.a.Set(l.w-1, l.h-1, true)
}

func process(s string, g int, fixed bool) (int, error) {
	l := NewLife(s)

	if fixed {
		l.Fixed()
	}

	for i := 0; i < g; i++ {
		l.Next()
		if fixed {
			l.Fixed()
		}
	}

	on := l.On()

	return on, nil
}

func run() int {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "%s filename\n", os.Args[0])
		return 1
	}

	b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	s := strings.TrimSpace(string(b))

	l, err := process(s, 100, false)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	l2, err := process(s, 100, true)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Printf("lights: %d\nlights2: %d\n", l, l2)

	return 0
}

func main() {
	os.Exit(run())
}
