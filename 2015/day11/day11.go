package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func incRune(r rune) (rune, bool, error) {
	if r < 'a' || r > 'z' {
		err := fmt.Errorf("invalid character %q", r)
		return 0, false, err
	}

	carry := false
	r++
	if r > 'z' {
		r = 'a'
		carry = true
	}

	return r, carry, nil
}

func incString(r []rune) ([]rune, error) {
	for i := len(r) - 1; i >= 0; i-- {
		v, carry, err := incRune(r[i])
		if err != nil {
			return []rune{}, err
		}

		r[i] = v

		if !carry {
			break
		} else if i == 0 {
			r = append([]rune{v}, r...)
		}
	}

	return r, nil
}

func hasNoBad(r []rune) bool {
	for _, c := range r {
		switch c {
		case 'i', 'o', 'l':
			return false
		}
	}

	return true
}

func hasStraight(r []rune) bool {
	for i := 0; i < len(r)-2; i++ {
		if (r[i+1] == (r[i] + 1)) && (r[i+2] == (r[i] + 2)) {
			return true
		}
	}

	return false
}

func hasRepeated(r []rune) bool {
	seen := make(map[rune]bool)

	count := 0
	for i := 0; i < len(r)-1; i++ {
		if r[i+1] == r[i] && !seen[r[i]] {
			count++
			seen[r[i]] = true
		}
	}

	return count > 1
}

func nextPassword(s string) (string, error) {
	r := []rune(s)

	if len(r) > 8 {
		err := fmt.Errorf("invalid initial password %q", s)
		return "", err
	}

	for {
		var err error
		r, err = incString(r)
		if err != nil {
			return "", err
		}

		if len(r) > 8 {
			err := fmt.Errorf("no result found for %q", s)
			return "", err
		}

		if hasNoBad(r) && hasStraight(r) && hasRepeated(r) {
			break
		}
	}

	result := string(r)

	return result, nil
}

func process(s string) (string, string, error) {
	result, err := nextPassword(s)
	if err != nil {
		return "", "", err
	}

	result2, err := nextPassword(result)
	if err != nil {
		return "", "", err
	}

	return result, result2, nil
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

	v, v2, err := process(s)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Printf("result: %q\nresult2: %q\n", v, v2)

	return 0
}

func main() {
	os.Exit(run())
}
