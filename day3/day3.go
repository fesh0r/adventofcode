package main

import (
	"fmt"
	"os"
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

func run() int {
	s := ">"
	h, err := getHouses(s)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Printf("houses: %d\n", h)
	return 0
}

func main() {
	os.Exit(run())
}
