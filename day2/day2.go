package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func parseSize(s string) ([]int, error) {
	var err error

	r, err := regexp.Compile("^(\\d+)x(\\d+)x(\\d+)$")
	if err != nil {
		return nil, err
	}

	m := r.FindStringSubmatch(s)
	if m == nil {
		err = fmt.Errorf("invalid size string %q", s)
		return nil, err
	}

	size := make([]int, 3)
	for i := 0; i < 3; i++ {
		size[i], err = strconv.Atoi(m[i+1])
		if err != nil {
			return nil, err
		}
	}
	sort.Ints(size)

	return size, nil
}

func getWrapping(s string) (int, error) {
	l, err := parseSize(s)
	if err != nil {
		return 0, err
	}

	area := 2*l[0]*l[1] + 2*l[0]*l[2] + 2*l[1]*l[2] + l[0]*l[1]

	return area, nil
}

func getRibbon(s string) (int, error) {
	l, err := parseSize(s)
	if err != nil {
		return 0, err
	}

	length := 2*l[0] + 2*l[1] + l[0]*l[1]*l[2]

	return length, nil
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

	area := 0
	ribbon := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		size := scanner.Text()

		a, err := getWrapping(size)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
		area += a

		r, err := getRibbon(size)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
		ribbon += r
	}

	fmt.Printf("area: %d\nribbon: %d\n", area, ribbon)
	return 0
}

func main() {
	os.Exit(run())
}
