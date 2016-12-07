package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func process(f io.Reader) (int, error) {
	var err error

	var count int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		ls := strings.Fields(s)

		if len(ls) != 3 {
			err = fmt.Errorf("invalid line %q", s)
			return 0, err
		}

		l := make([]int, 3)
		for k, v := range ls {
			val, err := strconv.Atoi(strings.TrimSpace(v))
			if err != nil {
				return 0, err
			}
			l[k] = val
		}

		sort.Ints(l)

		if l[2] < l[0]+l[1] {
			count++
		}
	}

	return count, nil
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

	count, err := process(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Printf("count: %d\n", count)
	return 0
}

func main() {
	os.Exit(run())
}
