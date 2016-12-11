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

func process(f io.Reader) (int, int, error) {
	var count, count2, index int

	l2 := make([][]int, 3)
	for i := range l2 {
		l2[i] = make([]int, 3)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		ls := strings.Fields(s)

		if len(ls) != 3 {
			err := fmt.Errorf("invalid line %q", s)
			return 0, 0, err
		}

		l := make([]int, 3)
		for k, v := range ls {
			val, err := strconv.Atoi(strings.TrimSpace(v))
			if err != nil {
				return 0, 0, err
			}
			l[k] = val
			l2[k][index] = val
		}

		sort.Ints(l)

		if l[2] < l[0]+l[1] {
			count++
		}

		index++
		if index == 3 {
			for i := 0; i < 3; i++ {
				sort.Ints(l2[i])

				if l2[i][2] < l2[i][0]+l2[i][1] {
					count2++
				}
			}
			index = 0
		}
	}

	return count, count2, nil
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

	count, count2, err := process(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Printf("count: %d\ncount2: %d\n", count, count2)
	return 0
}

func main() {
	os.Exit(run())
}
