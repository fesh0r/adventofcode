package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

var lineRegexp = regexp.MustCompile("^(rect|rotate (?:row|column)) (.+)$")
var rectRegexp = regexp.MustCompile("^([0-9]+)x([0-9]+)$")
var columnRegexp = regexp.MustCompile("^x=([0-9]+) by ([0-9]+)$")
var rowRegexp = regexp.MustCompile("^y=([0-9]+) by ([0-9]+)$")

type display struct {
	d    [][]bool
	w, h int
}

func newDisplay(w, h int) *display {
	d := make([][]bool, h)
	for y := 0; y < h; y++ {
		d[y] = make([]bool, w)
	}

	return &display{d, w, h}
}

func (d *display) String() string {
	var s string

	for y := 0; y < d.h; y++ {
		for x := 0; x < d.w; x++ {
			if d.on(x, y) {
				s += "#"
			} else {
				s += "."
			}
		}
		if y < d.h-1 {
			s += "\n"
		}
	}

	return s
}

func (d *display) on(x, y int) bool {
	return d.d[y][x]
}

func (d *display) onCount() int {
	var pixels int

	for y := 0; y < d.h; y++ {
		for x := 0; x < d.w; x++ {
			if d.d[y][x] {
				pixels++
			}
		}
	}

	return pixels
}

func (d *display) rect(w, h int) {
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			d.d[y][x] = true
		}
	}
}

func (d *display) rotateRow(y, c int) {
	for i := 0; i < c; i++ {
		tmp := d.d[y][d.w-1]
		for x := d.w - 1; x > 0; x-- {
			d.d[y][x] = d.d[y][x-1]
		}
		d.d[y][0] = tmp
	}
}

func (d *display) rotateColumn(x, c int) {
	for i := 0; i < c; i++ {
		tmp := d.d[d.h-1][x]
		for y := d.h - 1; y > 0; y-- {
			d.d[y][x] = d.d[y-1][x]
		}
		d.d[0][x] = tmp
	}
}

func (d *display) apply(s string) error {
	err := fmt.Errorf("invalid line %q", s)

	m := lineRegexp.FindStringSubmatch(s)
	if m == nil {
		return err
	}

	switch m[1] {
	case "rect":
		m2 := rectRegexp.FindStringSubmatch(m[2])
		if m2 == nil {
			return err
		}
		w, err2 := strconv.Atoi(m2[1])
		if err2 != nil {
			return err2
		}
		h, err2 := strconv.Atoi(m2[2])
		if err2 != nil {
			return err2
		}
		d.rect(w, h)
	case "rotate column":
		m2 := columnRegexp.FindStringSubmatch(m[2])
		if m2 == nil {
			return err
		}
		x, err2 := strconv.Atoi(m2[1])
		if err2 != nil {
			return err2
		}
		c, err2 := strconv.Atoi(m2[2])
		if err2 != nil {
			return err2
		}
		d.rotateColumn(x, c)
	case "rotate row":
		m2 := rowRegexp.FindStringSubmatch(m[2])
		if m2 == nil {
			return err
		}
		y, err2 := strconv.Atoi(m2[1])
		if err2 != nil {
			return err2
		}
		c, err2 := strconv.Atoi(m2[2])
		if err2 != nil {
			return err2
		}
		d.rotateRow(y, c)
	}

	return nil
}

func process(f io.Reader, w, h int) (int, string, error) {
	var pixels int

	d := newDisplay(w, h)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()

		err := d.apply(s)
		if err != nil {
			return 0, "", err
		}
	}

	pixels = d.onCount()

	return pixels, d.String(), nil
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

	c, o, err := process(f, 50, 6)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Printf("pixels: %d\noutput:\n%s\n", c, o)
	return 0
}

func main() {
	os.Exit(run())
}
