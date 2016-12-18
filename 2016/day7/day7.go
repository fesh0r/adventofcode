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

type abaMap map[string]bool

func getAbas(s string) abaMap {
	result := make(abaMap)
	for i := 0; i < len(s)-2; i++ {
		if s[i+1] != s[i] && s[i+2] == s[i] {
			result[s[i:i+3]] = true
		}
	}

	return result
}

func checkLine(s string) (bool, bool) {
	var end int
	var inHypernet, goodAbba, badAbba, hasSsl bool
	var abas, babs abaMap

	abas = make(abaMap)
	babs = make(abaMap)

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
			for k := range getAbas(s[:end]) {
				if inHypernet {
					babs[k] = true
				} else {
					abas[k] = true
				}
			}
		}

		if end == len(s) {
			break
		}
		s = s[end+1:]

		inHypernet = !inHypernet
	}

	for k := range abas {
		test := k[1:2] + k[0:1] + k[1:2]
		if babs[test] {
			hasSsl = true
			break
		}
	}

	return goodAbba && !badAbba, hasSsl
}

func process(f io.Reader) (int, int, error) {
	var tls, ssl int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()

		hasTls, hasSsl := checkLine(s)
		if hasTls {
			tls++
		}
		if hasSsl {
			ssl++
		}
	}

	return tls, ssl, nil
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

	tls, ssl, err := process(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Printf("tls: %d\nssl: %d\n", tls, ssl)
	return 0
}

func main() {
	os.Exit(run())
}
