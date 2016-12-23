package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func parseInput(s string) (int, int, error) {
	var x, y int
	cnt, err := fmt.Sscanf(s, "To continue, please consult the code grid in the manual.  Enter the code at row %d, column %d.\n", &y, &x)
	if err != nil {
		return 0, 0, err
	}
	if cnt != 2 {
		err := fmt.Errorf("invalid input %q", s)
		return 0, 0, err
	}

	return x, y, nil
}

func getIndex(x, y int) int {
	triangle := (x + y - 1) * (x + y) / 2
	index := triangle - y + 1

	return index
}

func getCode(index int) int64 {
	code := int64(20151125)
	for i := 1; i < index; i++ {
		code *= 252533
		code %= 33554393
	}

	return code
}

func process(s string) (int64, error) {
	x, y, err := parseInput(s)
	if err != nil {
		return 0, err
	}

	index := getIndex(x, y)

	code := getCode(index)

	return code, nil
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
	s := string(b)

	code, err := process(s)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Printf("code: %d\n", code)
	return 0
}

func main() {
	os.Exit(run())
}
