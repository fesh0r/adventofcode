package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func NextLookSay(s string) string {
	b := []rune(s)

	result := make([]rune, 0, len(b)*2)

	for i, j := 0, 0; i < len(b); i = j {
		j = i + 1
		for j < len(b) && b[j] == b[i] {
			j++
		}
		result = append(result, []rune(strconv.Itoa(j-i))...)
		result = append(result, b[i])
	}

	s = string(result)

	return s
}

var digitsRegexp = regexp.MustCompile("^[0-9]+$")

func RepeatLookSay(s string, c int) (string, error) {
	if !digitsRegexp.MatchString(s) {
		err := fmt.Errorf("invalid input %q", s)
		return "", err
	}

	for i := 0; i < c; i++ {
		s = NextLookSay(s)
	}

	return s, nil
}

func Process(s string) (int, int, error) {
	r, err := RepeatLookSay(s, 40)
	if err != nil {
		return 0, 0, err
	}

	l := len(r)

	r2, err := RepeatLookSay(s, 50)
	if err != nil {
		return 0, 0, err
	}

	l2 := len(r2)

	return l, l2, nil
}

func run() int {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "%s filename\n", os.Args[0])
		return 1
	}

	b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	s := strings.TrimSpace(string(b))

	v, v2, err := Process(s)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Printf("result: %d\nresult2: %d\n", v, v2)

	return 0
}

func main() {
	os.Exit(run())
}
