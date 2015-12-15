package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var vowels = [...]string{"a", "e", "i", "o", "u"}
var bads = [...]string{"ab", "cd", "pq", "xy"}

func hasVowels(s string) bool {
	var cnt int
	for _, v := range vowels {
		cnt += strings.Count(s, v)

		if cnt >= 3 {
			return true
		}
	}

	return false
}

func hasRepeated(s string) bool {
	var prev rune
	for i, c := range s {
		if i > 0 && prev == c {
			return true
		}
		prev = c
	}

	return false
}

func hasNoBad(s string) bool {
	for _, b := range bads {
		if strings.Count(s, b) > 0 {
			return false
		}
	}

	return true
}

func hasRepeatedPair(s string) bool {
	return false
}

func hasRepeatWithGap(s string) bool {
	return false
}

func checkString(s string) bool {
	return hasVowels(s) && hasRepeated(s) && hasNoBad(s)
}

func checkString2(s string) bool {
	return hasRepeatedPair(s) && hasRepeatWithGap(s)
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
	nice2 := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()

		if checkString(s) {
			nice++
		}

		if checkString2(s) {
			nice2++
		}
	}

	fmt.Printf("nice: %d\nnice2: %d\n", nice, nice2)
	return 0
}

func main() {
	os.Exit(run())
}
