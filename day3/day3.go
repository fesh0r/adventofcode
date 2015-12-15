package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func parseDirection(c rune) (int, int, error) {
	var x, y int

	switch c {
	case '<':
		x = -1
	case '>':
		x = 1
	case '^':
		y = 1
	case 'v':
		y = -1
	default:
		err := fmt.Errorf("invalid character %q", c)
		return 0, 0, err
	}

	return x, y, nil
}

func getHouses(s string) (int, error) {
	var pos [2]int

	h := make(map[[2]int]bool)

	h[pos] = true

	for _, c := range s {
		x, y, err := parseDirection(c)
		if err != nil {
			return 0, err
		}

		pos[0] += x
		pos[1] += y
		h[pos] = true
	}

	var houses int
	for k := range h {
		if h[k] {
			houses++
		}
	}

	return houses, nil
}

func getHousesDouble(s string) (int, error) {
	var pos [2][2]int

	h := make(map[[2]int]bool)

	h[pos[0]] = true
	h[pos[1]] = true

	for i, c := range s {
		x, y, err := parseDirection(c)
		if err != nil {
			return 0, err
		}

		worker := i % 2
		pos[worker][0] += x
		pos[worker][1] += y
		h[pos[worker]] = true
	}

	var houses int
	for k := range h {
		if h[k] {
			houses++
		}
	}

	return houses, nil
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

	h, err := getHouses(s)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	h2, err := getHousesDouble(s)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Printf("houses: %d\nhouses2: %d\n", h, h2)
	return 0
}

func main() {
	os.Exit(run())
}
