package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

var lineRegex = regexp.MustCompile("^([A-Za-z]+) => ([A-Za-z]+)$")
var moleculeRegex = regexp.MustCompile("[A-Z][a-z]?")

type replace struct {
	in, out string
}

func process(f io.Reader) (int, error) {
	var replacements []replace
	var molecule []string

	results := make(map[string]bool)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		m := lineRegex.FindStringSubmatch(s)
		if m != nil {
			replacements = append(replacements, replace{m[1], m[2]})
		} else if len(s) > 0 {
			molecule = moleculeRegex.FindAllString(s, -1)
		}
	}

	for k, v := range molecule {
		for _, r := range replacements {
			if r.in == v {
				newMolecule := make([]string, len(molecule))
				copy(newMolecule, molecule)
				newMolecule[k] = r.out
				newMoleculeStr := strings.Join(newMolecule, "")
				results[newMoleculeStr] = true
			}
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
