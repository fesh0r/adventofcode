package main

import (
	"bufio"
	"fmt"
	"os"
)

func hasVowels(s string) bool {
	var cnt int
	for _, c := range s {
		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			cnt++
			if cnt >= 3 {
				return true
			}
		}
	}

	return false
}

func hasRepeated(s string) bool {
	r := []rune(s)
	for i := 0; i < len(r)-1; i++ {
		if r[i+1] == r[i] {
			return true
		}
	}

	return false
}

func hasNoBad(s string) bool {
	r := []rune(s)
	for i := 0; i < len(r)-1; i++ {
		switch string(r[i : i+2]) {
		case "ab", "cd", "pq", "xy":
			return false
		}
	}

	return true
}

func hasRepeatedPair(s string) bool {
	r := []rune(s)
	for i := 0; i < len(r)-3; i++ {
		for j := i + 2; j < len(r)-1; j++ {
			if r[j] == r[i] && r[j+1] == r[i+1] {
				return true
			}
		}
	}

	return false
}

func hasRepeatWithGap(s string) bool {
	r := []rune(s)
	for i := 0; i < len(r)-2; i++ {
		if r[i] == r[i+2] {
			return true
		}
	}

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
