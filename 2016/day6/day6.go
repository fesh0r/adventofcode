package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

type freq map[rune]int

type pair struct {
	key   rune
	value int
}

type pairList []pair

func (p pairList) Len() int {
	return len(p)
}

func (p pairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p pairList) Less(i, j int) bool {
	if p[i].value != p[j].value {
		return p[i].value > p[j].value
	} else {
		return p[i].key < p[j].key
	}
}

func sortFreqMap(m freq) pairList {
	pl := make(pairList, len(m))

	var i int
	for k, v := range m {
		pl[i] = pair{k, v}
		i++
	}

	sort.Sort(pl)

	return pl
}

func process(f io.Reader) (string, error) {
	var freqs []freq
	var msg string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()

		if freqs == nil {
			freqs = make([]freq, len(s))
			for i := 0; i < len(s); i++ {
				freqs[i] = make(freq)
			}
		}

		for p, l := range s {
			freqs[p][l]++
		}
	}

	for _, f := range freqs {
		pl := sortFreqMap(f)
		msg += string(pl[0].key)
	}

	return msg, nil
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

	msg, err := process(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Printf("msg: %s\n", msg)
	return 0
}

func main() {
	os.Exit(run())
}
