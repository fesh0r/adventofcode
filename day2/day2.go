package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func getWrapping(s string) (area int, err error) {
	r, err := regexp.Compile("^(\\d+)x(\\d+)x(\\d+)$")
	if err != nil {
		return
	}

	m := r.FindStringSubmatch(s)
	if m == nil {
		err = fmt.Errorf("invalid size string %q", s)
		return
	}

	l := make([]int, 3)
	for i := 0; i < 3; i++ {
		l[i], err = strconv.Atoi(m[i+1])
		if err != nil {
			return
		}
	}

	sort.Ints(l)

	area = 2*l[0]*l[1] + 2*l[0]*l[2] + 2*l[1]*l[2] + l[0]*l[1]

	return
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

	area := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		a, err := getWrapping(scanner.Text())
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
		area += a
	}

	fmt.Printf("area => %d\n", area)
	return 0
}

func main() {
	os.Exit(run())
}
