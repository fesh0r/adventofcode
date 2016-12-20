package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func process(s string) (int, error) {
	p, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	v := make([]int, p / 10)

	for i := 1; i < len(v); i++ {
		for j := i; j < len(v); j += i {
			v[j] += i * 10
		}
	}

	res := -1

	for i := 1; i < len(v); i++ {
		if v[i] >= p {
			res = i
			break
		}
	}

	return res, nil
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

	house, err := process(s)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Printf("house: %d\n", house)
	return 0
}

func main() {
	os.Exit(run())
}
