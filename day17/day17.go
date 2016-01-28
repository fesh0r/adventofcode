package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func process(f io.Reader, cap int) (int, error) {
	containers := make([]int, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		v, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return 0, err
		}
		containers = append(containers, v)
	}

	cnt := uint(len(containers))

	var valid int
	for i := uint(0); i < (1 << cnt); i++ {
		var cur int
		for j := uint(0); j < cnt; j++ {
			if (i & (1 << j)) != 0 {
				cur += containers[j]
			}
		}
		if cur == cap {
			valid++
		}
	}

	return valid, nil
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

	valid, err := process(f, 150)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Printf("valid: %d\n", valid)

	return 0
}

func main() {
	os.Exit(run())
}
