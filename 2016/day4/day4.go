package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var roomRegexp = regexp.MustCompile("^([a-z-]+)-([0-9]+)\\[([a-z]{5})]$")

type room struct {
	name     string
	sector   int
	checksum string
}

func parseRoom(s string) (room, error) {
	m := roomRegexp.FindStringSubmatch(s)
	if m == nil {
		err := fmt.Errorf("invalid room string %q", s)
		return room{}, err
	}

	var err error
	var r room
	r.name = m[1]
	r.sector, err = strconv.Atoi(m[2])
	if err != nil {
		return room{}, err
	}
	r.checksum = m[3]

	return r, nil
}

type freq map[rune]int

func getFreq(s string) freq {
	freq := make(freq)

	for _, l := range s {
		freq[l]++
	}

	return freq
}

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

func getCode(r room) string {
	f := getFreq(strings.Replace(r.name, "-", "", -1))

	pl := sortFreqMap(f)

	var code string
	for i := 0; i < 5; i++ {
		code += string(pl[i].key)
	}

	return code
}

func checkRoom(r room) bool {
	code := getCode(r)

	return code == r.checksum
}

func process(f io.Reader) (int, error) {
	var sum int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()

		r, err := parseRoom(s)
		if err != nil {
			return 0, err
		}
		if checkRoom(r) {
			sum += r.sector
		}
	}

	return sum, nil
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

	sum, err := process(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Printf("sum: %d\n", sum)
	return 0
}

func main() {
	os.Exit(run())
}
