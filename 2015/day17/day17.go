package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func process(f io.Reader, cap int) (int, int, error) {
	containers := make([]int, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		v, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return 0, 0, err
		}
		containers = append(containers, v)
	}

	cnt := uint(len(containers))

	var valid int
	used := make([]int, cnt)
	for i := uint(0); i < (1 << cnt); i++ {
		var cur, filled int
		for j := uint(0); j < cnt; j++ {
			if (i & (1 << j)) != 0 {
				cur += containers[j]
				filled++
			}
		}
		if cur == cap {
			valid++
			used[filled]++
		}
	}

	var validMin int
	for _, v := range used {
		if v > 0 {
			validMin = v
			break
		}
	}

	return valid, validMin, nil
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

	valid, validMin, err := process(f, 150)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Printf("valid: %d\nvalidMin: %d\n", valid, validMin)

	return 0
}

func main() {
	os.Exit(run())
}
