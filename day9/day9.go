package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

var lineRegexp = regexp.MustCompile("^(.+) to (.+) = ([0-9]+)$")

func parseLine(s string) (string, string, int, error) {
	errFormat := fmt.Errorf("invalid line %q", s)

	m := lineRegexp.FindStringSubmatch(s)
	if m == nil {
		return "", "", 0, errFormat
	}

	from, to := m[1], m[2]

	distance, err := strconv.Atoi(m[3])
	if err != nil {
		return "", "", 0, errFormat
	}

	return from, to, distance, nil
}

func findIndex(l []string, s string) (int, bool) {
	for k, v := range l {
		if v == s {
			return k, true
		}
	}

	return 0, false
}

func permutations(n int) <-chan []int {
	c := make(chan []int)
	go func() {
		defer close(c)

		indices := make([]int, n)
		for i := range indices {
			indices[i] = i
		}

		cycles := make([]int, n)
		for i := range cycles {
			cycles[i] = n - i
		}

		out := make([]int, len(indices))
		copy(out, indices)
		c <- out

		for n > 0 {
			i := n - 1
			for ; i >= 0; i-- {
				cycles[i] -= 1
				if cycles[i] == 0 {
					index := indices[i]
					for j := i; j < n-1; j++ {
						indices[j] = indices[j+1]
					}
					indices[n-1] = index
					cycles[i] = n - i
				} else {
					j := cycles[i]
					indices[i], indices[n-j] = indices[n-j], indices[i]

					out := make([]int, len(indices))
					copy(out, indices)
					c <- out
					break
				}
			}

			if i < 0 {
				break
			}
		}
	}()

	return c
}

type route struct {
	from int
	to   int
}

func process(f io.Reader) (int, int, error) {
	locations := make([]string, 0, 20)
	seenLocations := make(map[string]bool)
	distances := make(map[route]int)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()

		from, to, distance, err := parseLine(s)
		if err != nil {
			return 0, 0, err
		}

		if !seenLocations[from] {
			locations = append(locations, from)
			seenLocations[from] = true
		}
		fromIndex, _ := findIndex(locations, from)

		if !seenLocations[to] {
			locations = append(locations, to)
			seenLocations[to] = true
		}
		toIndex, _ := findIndex(locations, to)

		if _, found := distances[route{fromIndex, toIndex}]; found {
			err := fmt.Errorf("duplicate distance %q, %q in line %q", from, to, s)
			return 0, 0, err
		}

		distances[route{fromIndex, toIndex}] = distance
		distances[route{toIndex, fromIndex}] = distance
	}

	var lowest, highest int
	first := true
	for c := range permutations(len(locations)) {
		cur := 0
		for j := 0; j < len(c)-1; j++ {
			d, found := distances[route{c[j], c[j+1]}]
			if !found {
				err := fmt.Errorf("unknown distance %q, %q", locations[c[j]], locations[c[j+1]])
				return 0, 0, err
			}
			cur += d
		}
		if first || cur < lowest {
			lowest = cur
		}
		if first || cur > highest {
			highest = cur
		}
		first = false
	}

	return lowest, highest, nil
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

	lowest, highest, err := process(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Printf("lowest: %d\nhighest: %d\n", lowest, highest)

	return 0
}

func main() {
	os.Exit(run())
}
