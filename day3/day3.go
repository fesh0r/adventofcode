package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func getHouses(s string) (houses int, err error) {
	var pos [2]int

	h := make(map[[2]int]bool)

	h[pos] = true

	for _, c := range s {
		switch c {
		case '<':
			pos[0]--
		case '>':
			pos[0]++
		case '^':
			pos[1]++
		case 'v':
			pos[1]--
		default:
			err = fmt.Errorf("invalid character %q", c)
			return
		}
		h[pos] = true
	}

	for k := range h {
		if h[k] {
			houses++
		}
	}
	return
}

func getHousesDouble(s string) (houses int, err error) {
	var pos [2][2]int

	h := make(map[[2]int]bool)

	h[pos[0]] = true
	h[pos[1]] = true

	for i, c := range s {
		worker := i % 2
		switch c {
		case '<':
			pos[worker][0]--
		case '>':
			pos[worker][0]++
		case '^':
			pos[worker][1]++
		case 'v':
			pos[worker][1]--
		default:
			err = fmt.Errorf("invalid character %q", c)
			return
		}
		h[pos[worker]] = true
	}

	for k := range h {
		if h[k] {
			houses++
		}
	}
	return
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

	fmt.Printf("houses: %d\n", h)
	fmt.Printf("houses2: %d\n", h2)
	return 0
}

func main() {
	os.Exit(run())
}
