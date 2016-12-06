package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

var lineRegexp = regexp.MustCompile(
	`^(.+): capacity ([0-9-]+), durability ([0-9-]+), flavor ([0-9-]+), texture ([0-9-]+), calories ([0-9-]+)$`)

func parseLine(s string) (string, int, int, int, int, int, error) {
	errFormat := fmt.Errorf("invalid line %q", s)

	m := lineRegexp.FindStringSubmatch(s)
	if m == nil {
		return "", 0, 0, 0, 0, 0, errFormat
	}

	name := m[1]

	capacity, err := strconv.Atoi(m[2])
	if err != nil {
		return "", 0, 0, 0, 0, 0, errFormat
	}

	durability, err := strconv.Atoi(m[3])
	if err != nil {
		return "", 0, 0, 0, 0, 0, errFormat
	}

	flavor, err := strconv.Atoi(m[4])
	if err != nil {
		return "", 0, 0, 0, 0, 0, errFormat
	}

	texture, err := strconv.Atoi(m[5])
	if err != nil {
		return "", 0, 0, 0, 0, 0, errFormat
	}

	calories, err := strconv.Atoi(m[6])
	if err != nil {
		return "", 0, 0, 0, 0, 0, errFormat
	}

	return name, capacity, durability, flavor, texture, calories, nil
}

func combinations(n int, r int) <-chan []int {
	c := make(chan []int)
	go func() {
		defer close(c)

		indices := make([]int, r)
		for i := range indices {
			indices[i] = 0
		}

		out := make([]int, len(indices))
		copy(out, indices)
		c <- out

		for n > 0 {
			i := r - 1
			for ; i >= 0; i-- {
				if indices[i] != n-1 {
					break
				}
			}

			if i < 0 {
				break
			}

			index := indices[i]
			for j := i; j < r; j++ {
				indices[j] = index + 1
			}

			out := make([]int, len(indices))
			copy(out, indices)
			c <- out
		}
	}()

	return c
}

type ingredient struct {
	name       string
	capacity   int
	durability int
	flavour    int
	texture    int
	calories   int
}

func process(f io.Reader, n int, calories int) (int, int, error) {
	var ingredients []ingredient

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()

		name, capacity, durability, flavor, texture, calories, err := parseLine(s)
		if err != nil {
			return 0, 0, err
		}

		ingredients = append(ingredients, ingredient{name, capacity, durability, flavor, texture, calories})
	}

	var highest int
	var highestC int
	for c := range combinations(len(ingredients), n) {
		var t ingredient
		var cur int
		for _, i := range c {
			t.capacity += ingredients[i].capacity
			t.durability += ingredients[i].durability
			t.flavour += ingredients[i].flavour
			t.texture += ingredients[i].texture
			t.calories += ingredients[i].calories
		}
		if t.capacity < 0 {
			t.capacity = 0
		}
		if t.durability < 0 {
			t.durability = 0
		}
		if t.flavour < 0 {
			t.flavour = 0
		}
		if t.texture < 0 {
			t.texture = 0
		}
		cur = t.capacity * t.durability * t.flavour * t.texture
		if cur > highest {
			highest = cur
		}
		if t.calories == calories && cur > highestC {
			highestC = cur
		}
	}

	return highest, highestC, nil
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

	score, scoreC, err := process(f, 100, 500)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Printf("score: %d\nscoreC: %d\n", score, scoreC)

	return 0
}

func main() {
	os.Exit(run())
}
