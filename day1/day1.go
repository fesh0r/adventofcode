package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func parseChange(c rune) (int, error) {
	var change int

	switch c {
	case '(':
		change = 1
	case ')':
		change = -1
	default:
		err := fmt.Errorf("invalid character %q", c)
		return 0, err
	}

	return change, nil
}

func getFloor(s string) (int, error) {
	var floor int

	for _, c := range s {
		change, err := parseChange(c)
		if err != nil {
			return 0, err
		}

		floor += change
	}

	return floor, nil
}

func getBasement(s string) (int, error) {
	var floor int

	for i, c := range s {
		change, err := parseChange(c)
		if err != nil {
			return 0, err
		}

		floor += change
		if floor < 0 {
			index := i + 1
			return index, nil
		}
	}

	return 0, nil
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
