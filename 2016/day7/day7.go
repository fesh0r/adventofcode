package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func hasAbba(s string) bool {
	for i := 0; i < len(s)-3; i++ {
		if s[i+1] != s[i] && s[i+3] == s[i] && s[i+1] == s[i+2] {
			return true
		}
	}

	return false
}

func checkLine(s string) bool {
	var end int
	var inHypernet, goodAbba, badAbba bool

	for {
		if inHypernet {
			end = strings.Index(s, "]")
		} else {
			end = strings.Index(s, "[")
		}
		if end < 0 {
			end = len(s)
		}

		if end > 0 {
			if hasAbba(s[:end]) {
				if inHypernet {
					badAbba = true
				} else {
					goodAbba = true
				}
			}
		}

		if end == len(s) {
			break
		}
		s = s[end+1:]

		inHypernet = !inHypernet
	}

	return goodAbba && !badAbba
}

func process(f io.Reader) (int, error) {
	var tls int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()

		if checkLine(s) {
			tls++
		}
	}

	return tls, nil
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

	tls, err := process(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Printf("tls: %d\n", tls)
	return 0
}

func main() {
	os.Exit(run())
}
