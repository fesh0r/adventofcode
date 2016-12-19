package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

var lineRegex = regexp.MustCompile("^([A-Za-z]+) => ([A-Za-z]+)$")

type replace struct {
	in, out  string
	inRegexp *regexp.Regexp
}

func process(f io.Reader) (int, error) {
	var replacements []replace
	var molecule string

	results := make(map[string]bool)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		m := lineRegex.FindStringSubmatch(s)
		if m != nil {
			replacements = append(replacements, replace{m[1], m[2], regexp.MustCompile(m[1])})
		} else if len(s) > 0 {
			molecule = s
		}
	}

	for _, r := range replacements {
		m := r.inRegexp.FindAllStringIndex(molecule, -1)
		for _, v := range m {
			newMolecule := molecule[:v[0]] + r.out + molecule[v[1]:]
			results[newMolecule] = true
		}
	}

	return len(results), nil
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

	c, err := process(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Printf("count: %d\n", c)
	return 0
}

func main() {
	os.Exit(run())
}
