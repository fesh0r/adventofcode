package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type group struct {
	c  []int
	qe int64
}

type groupList []group

func combinationsQE(n int, r int, weights []int, target int) <-chan group {
	c := make(chan group)
	go func() {
		defer close(c)

		indices := make([]int, r)
		for i := range indices {
			indices[i] = i
		}

		cw := 0
		qe := int64(1)
		for _, w := range indices {
			cw += weights[w]
			qe *= int64(weights[w])
		}
		if cw == target {
			out := make([]int, len(indices))
			copy(out, indices)
			c <- group{out, qe}
		}

		for n > 0 {
			i := r - 1
			for ; i >= 0; i-- {
				if indices[i] != i+n-r {
					break
				}
			}
			if i < 0 {
				break
			}

			indices[i]++
			for j := i + 1; j < r; j++ {
				indices[j] = indices[j-1] + 1
			}

			cw = 0
			qe = 1
			for _, w := range indices {
				cw += weights[w]
				qe *= int64(weights[w])
			}
			if cw == target {
				out := make([]int, len(indices))
				copy(out, indices)
				c <- group{out, qe}
			}
		}
	}()

	return c
}

func findLowest(weights []int, groups int) int64 {
	totalWeight := 0
	for _, v := range weights {
		totalWeight += v
	}

	targetWeight := totalWeight / groups

	for l := 1; l < len(weights)-(groups-2); l++ {
		for c := range combinationsQE(len(weights), l, weights, targetWeight) {
			return c.qe
		}
	}

	return 0
}

func process(f io.Reader) (int64, error) {
	var weights []int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		if s != "" {
			w, err := strconv.Atoi(s)
			if err != nil {
				return 0, err
			}
			weights = append(weights, w)
		}
	}

	qe3 := findLowest(weights, 3)

	return qe3, nil
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

	qe3, err := process(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Printf("qe3: %d\n", qe3)
	return 0
}

func main() {
	os.Exit(run())
}
