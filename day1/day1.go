package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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

func getBasement(s string) (index int, err error) {
	var floor int
	for i, c := range s {
		switch c {
		case '(':
			floor++
		case ')':
			floor--
		default:
			err = fmt.Errorf("invalid character %q", c)
			return
		}
		if floor < 0 {
			index = i + 1
			return
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

	f, err := getFloor(s)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	i, err := getBasement(s)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Printf("floor: %d\nbasement: %d\n", f, i)
	return 0
}

func main() {
	os.Exit(run())
}
