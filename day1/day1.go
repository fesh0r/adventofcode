package main

import (
	"fmt"
	"os"
)

func getFloor(s string) (floor int, err error) {
	for _, c := range s {
		switch c {
		case '(':
			floor++
		case ')':
			floor--
		default:
			err = fmt.Errorf("invalid character %q", c)
			return
		}
	}
	return
}

func run() int {
	s := "())"

	f, err := getFloor(s)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Printf("spec: %q\nfloor: %d\n", s, f)
	return 0
}

func main() {
	os.Exit(run())
}
