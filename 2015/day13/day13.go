package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

var lineRegexp = regexp.MustCompile(`^(.+) would (lose|gain) ([0-9]+) happiness units by sitting next to (.+)\.$`)

func parseLine(s string) (string, string, int, error) {
	errFormat := fmt.Errorf("invalid line %q", s)

	m := lineRegexp.FindStringSubmatch(s)
	if m == nil {
		return "", "", 0, errFormat
	}

	from, to := m[1], m[4]

	happy, err := strconv.Atoi(m[3])
	if err != nil {
		return "", "", 0, errFormat
	}
	if m[2] == "lose" {
		happy = -happy
	}

	return from, to, happy, nil
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

type seating struct {
	from int
	to   int
}

func process(f io.Reader, addSelf bool) (int, error) {
	people := make([]string, 0, 20)
	seenPeople := make(map[string]bool)
	happiness := make(map[seating]int)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()

		from, to, happy, err := parseLine(s)
		if err != nil {
			return 0, err
		}

		if !seenPeople[from] {
			people = append(people, from)
			seenPeople[from] = true
		}
		fromIndex, _ := findIndex(people, from)

		if !seenPeople[to] {
			people = append(people, to)
			seenPeople[to] = true
		}
		toIndex, _ := findIndex(people, to)

		if _, found := happiness[seating{fromIndex, toIndex}]; found {
			err := fmt.Errorf("duplicate happiness %q, %q in line %q", from, to, s)
			return 0, err
		}

		happiness[seating{fromIndex, toIndex}] = happy
	}

	if addSelf {
		from := "Me"
		people = append(people, from)
		seenPeople[from] = true
		fromIndex, _ := findIndex(people, from)

		for i := range people {
			if i != fromIndex {
				happiness[seating{fromIndex, i}] = 0
				happiness[seating{i, fromIndex}] = 0
			}
		}
	}

	var highest int
	for c := range permutations(len(people)) {
		cur := 0
		for j := 0; j < len(c); j++ {
			k := j + 1
			if k >= len(c) {
				k = 0
			}
			h, found := happiness[seating{c[j], c[k]}]
			if !found {
				err := fmt.Errorf("unknown happiness %q, %q", people[c[j]], people[c[k]])
				return 0, err
			}
			h2, found := happiness[seating{c[k], c[j]}]
			if !found {
				err := fmt.Errorf("unknown happiness %q, %q", people[c[k]], people[c[j]])
				return 0, err
			}
			cur += h + h2
		}
		if cur > highest {
			highest = cur
		}
	}

	return highest, nil
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

	highest, err := process(f, false)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	_, err = f.Seek(0, 0)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	highest2, err := process(f, true)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Printf("highest: %d\nhighest2: %d\n", highest, highest2)

	return 0
}

func main() {
	os.Exit(run())
}
