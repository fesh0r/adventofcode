package main

import (
	"bufio"
	"fmt"
	"os"
)

func hasVowels(s string) bool {
	return false
}

func hasRepeated(s string) bool {
	return false
}

func hasNoBad(s string) bool {
	return false
}

func checkString(s string) bool {
	return hasVowels(s) && hasRepeated(s) && hasNoBad(s)
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

	nice := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()

		if checkString(s) {
			nice++
		}
	}

	fmt.Printf("nice: %d\n", nice)
	return 0
}

func main() {
	os.Exit(run())
}
