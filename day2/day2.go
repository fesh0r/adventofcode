package main

import (
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
	s := "2x3x4"
	w, err := getWrapping(s)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Printf("%q => %d\n", s, w)
	return 0
}

func main() {
	os.Exit(run())
}
