package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func checkIndex(key string, index int) (bool, string) {
	b := []byte(key + strconv.Itoa(index))
	h := fmt.Sprintf("%x", md5.Sum(b))

	if strings.HasPrefix(h, "00000") {
		return true, h[5:6]
	}

	return false, ""
}

func findPassword(key string) (string, error) {
	max := math.MaxInt32

	var current_pos int
	var code string

	for index := 0; index < max; index++ {
		f, c := checkIndex(key, index)
		if f {
			code += c
			current_pos++

			if current_pos == 8 {
				return code, nil
			}
		}
	}

	err := fmt.Errorf("no code found below %d", max)
	return "", err
}

func checkIndex2(key string, index int) (bool, int, byte) {
	b := []byte(key + strconv.Itoa(index))
	h := fmt.Sprintf("%x", md5.Sum(b))

	if strings.HasPrefix(h, "00000") {
		pos, err := strconv.ParseInt(h[5:6], 16, 0)
		if err == nil {
			if pos < 8 {
				return true, int(pos), h[6]
			}
		}
	}

	return false, 0, 0
}

func findPassword2(key string) (string, error) {
	max := math.MaxInt32

	var done int
	code := make([]byte, 8)

	for index := 0; index < max; index++ {
		f, p, c := checkIndex2(key, index)
		if f {
			if code[p] == 0 {
				code[p] = c
				done++

				if done == 8 {
					return string(code), nil
				}
			}
		}
	}

	err := fmt.Errorf("no code found below %d", max)
	return "", err
}

func process(s string) (string, string, error) {
	c1, err := findPassword(s)
	if err != nil {
		return "", "", err
	}

	c2, err := findPassword2(s)
	if err != nil {
		return "", "", err
	}

	return c1, c2, nil
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

	c1, c2, err := process(s)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Printf("code1: %s\ncode2: %s\n", c1, c2)
	return 0
}

func main() {
	os.Exit(run())
}
