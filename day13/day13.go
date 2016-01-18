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

func appendCopy(a [][]int, b []int) [][]int {
	r := make([]int, len(b))
	copy(r, b)
	a = append(a, r)
	return a
}

func permutations(n int) [][]int {
	indices := make([]int, n)
	for i := range indices {
		indices[i] = i
	}

	cycles := make([]int, n)
	for i := range cycles {
		cycles[i] = n - i
	}

	var results [][]int

	results = appendCopy(results, indices)

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

				results = appendCopy(results, indices)
				break
			}
		}

		if i < 0 {
			break
		}
	}

	return results
}

type seating struct {
	from int
	to   int
}

func process(f io.Reader) (int, error) {
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

	p := permutations(len(people))

	var highest int

	for i, c := range p {
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
		if i == 0 || cur > highest {
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

	highest, err := process(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Printf("highest: %d\n", highest)

	return 0
}

func main() {
	os.Exit(run())
}
