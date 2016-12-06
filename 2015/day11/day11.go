package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func IncRune(r rune) (rune, bool, error) {
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

func IncString(r []rune) ([]rune, error) {
	for i := len(r) - 1; i >= 0; i-- {
		v, carry, err := IncRune(r[i])
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

func HasNoBad(r []rune) bool {
	for _, c := range r {
		switch c {
		case 'i', 'o', 'l':
			return false
		}
	}

	return true
}

func HasStraight(r []rune) bool {
	for i := 0; i < len(r)-2; i++ {
		if (r[i+1] == (r[i] + 1)) && (r[i+2] == (r[i] + 2)) {
			return true
		}
	}

	return false
}

func HasRepeated(r []rune) bool {
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

func NextPassword(s string) (string, error) {
	r := []rune(s)

	if len(r) > 8 {
		err := fmt.Errorf("invalid initial password %q", s)
		return "", err
	}

	for {
		var err error
		r, err = IncString(r)
		if err != nil {
			return "", err
		}

		if len(r) > 8 {
			err := fmt.Errorf("no result found for %q", s)
			return "", err
		}

		if HasNoBad(r) && HasStraight(r) && HasRepeated(r) {
			break
		}
	}

	result := string(r)

	return result, nil
}

func Process(s string) (string, string, error) {
	result, err := NextPassword(s)
	if err != nil {
		return "", "", err
	}

	result2, err := NextPassword(result)
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

	v, v2, err := Process(s)
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
