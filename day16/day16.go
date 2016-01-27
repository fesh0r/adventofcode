package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

type attributes map[string]int

var sample = attributes{
	"children":    3,
	"cats":        7,
	"samoyeds":    2,
	"pomeranians": 3,
	"akitas":      0,
	"vizslas":     0,
	"goldfish":    5,
	"trees":       3,
	"cars":        2,
	"perfumes":    1,
}

var lineRegexp = regexp.MustCompile(`^Sue ([0-9]+): (.+)$`)
var splitRegexp = regexp.MustCompile(`, `)
var attributeRegexp = regexp.MustCompile(`([a-z]+): ([0-9]+)`)

func parseLine(s string) (int, attributes, error) {
	errFormat := fmt.Errorf("invalid line %q", s)

	ml := lineRegexp.FindStringSubmatch(s)
	if ml == nil {
		return 0, nil, errFormat
	}

	index, err := strconv.Atoi(ml[1])
	if err != nil {
		return 0, nil, errFormat
	}

	attribs := make(attributes)
	sa := splitRegexp.Split(ml[2], -1)
	if sa == nil {
		return 0, nil, errFormat
	}

	for _, a := range sa {
		ma := attributeRegexp.FindStringSubmatch(a)
		if ma == nil {
			return 0, nil, errFormat
		}

		if _, ok := sample[ma[1]]; !ok {
			return 0, nil, errFormat
		}

		v, err := strconv.Atoi(ma[2])
		if err != nil {
			return 0, nil, errFormat
		}

		attribs[ma[1]] = v
	}

	return index, attribs, nil
}

func process(f io.Reader) (int, error) {
	var highest, highestIndex int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()

		index, attribs, err := parseLine(s)
		if err != nil {
			return 0, err
		}

		var cur int
		for k, v := range attribs {
			if sample[k] == v {
				cur++
			}
		}
		if cur > highest {
			highest = cur
			highestIndex = index
		}
	}

	return highestIndex, nil
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

	aunt, err := process(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Printf("aunt: %d\n", aunt)

	return 0
}

func main() {
	os.Exit(run())
}
