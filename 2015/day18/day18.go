package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type board struct {
	b    [][]bool
	w, h int
}

func newBoard(w, h int) *board {
	b := make([][]bool, h+2)

	for i := range b {
		b[i] = make([]bool, w+2)
	}

	return &board{b, w, h}
}

func (b *board) String() string {
	var s string

	for y := 0; y < b.h; y++ {
		for x := 0; x < b.w; x++ {
			if b.on(x, y) {
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

func (b *board) on(x, y int) bool {
	return b.b[y+1][x+1]
}

func (b *board) set(x, y int, alive bool) {
	b.b[y+1][x+1] = alive
}

func (b *board) next(x, y int) bool {
	var c int
	for ny := -1; ny <= 1; ny++ {
		for nx := -1; nx <= 1; nx++ {
			if (nx != 0 || ny != 0) && b.on(x+nx, y+ny) {
				c++
			}
		}
	}

	return c == 3 || c == 2 && b.on(x, y)
}

type life struct {
	a, b *board
	w, h int
}

func newLife(s string) *life {
	ss := strings.Split(s, "\n")

	w := len(ss[0])
	h := len(ss)
	a := newBoard(w, h)

	for y, row := range ss {
		for x, cell := range row {
			a.set(x, y, cell == '#')
		}
	}

	return &life{a, newBoard(w, h), w, h}
}

func (l *life) String() string {
	return l.a.String()
}

func (l *life) next() {
	for y := 0; y < l.h; y++ {
		for x := 0; x < l.w; x++ {
			l.b.set(x, y, l.a.next(x, y))
		}
	}

	l.a, l.b = l.b, l.a
}

func (l *life) on() int {
	var on int
	for y := 0; y < l.h; y++ {
		for x := 0; x < l.w; x++ {
			if l.a.on(x, y) {
				on++
			}
		}
	}
	return on
}

func (l *life) fixed() {
	l.a.set(0, 0, true)
	l.a.set(l.w-1, 0, true)
	l.a.set(0, l.h-1, true)
	l.a.set(l.w-1, l.h-1, true)
}

func process(s string, g int, fixed bool) (int, error) {
	l := newLife(s)

	if fixed {
		l.fixed()
	}

	for i := 0; i < g; i++ {
		l.next()
		if fixed {
			l.fixed()
		}
	}

	on := l.on()

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
