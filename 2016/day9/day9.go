package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var markerRegex = regexp.MustCompile("\\(([0-9]+)x([0-9]+)\\)")

func expand(s string) (string, error) {
	var out string

	for len(s) > 0 {
		//fmt.Println(s)
		m := markerRegex.FindStringSubmatchIndex(s)
		if m == nil {
			out += s
			s = ""
		} else {
			out += s[:m[0]]
			l, err := strconv.Atoi(s[m[2]:m[3]])
			if err != nil {
				return "", err
			}
			c, err := strconv.Atoi(s[m[4]:m[5]])
			if err != nil {
				return "", err
			}
			out += strings.Repeat(s[m[1]:m[1]+l], c)
			s = s[m[1]+l:]
			//fmt.Printf("%s %#v %d %d\n", out, m, l, c)
		}
	}

	return out, nil
}

func process(s string) (int, error) {
	out, err := expand(s)
	if err != nil {
		return 0, err
	}

	return len(out), nil
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

	l, err := process(s)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Printf("length: %d\n", l)
	return 0
}

func main() {
	os.Exit(run())
}
