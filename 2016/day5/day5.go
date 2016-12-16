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

func process(s string) (string, error) {
	c1, err := findPassword(s)
	if err != nil {
		return "", err
	}

	return c1, nil
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

	c1, err := process(s)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Printf("code1: %s\n", c1)
	return 0
}

func main() {
	os.Exit(run())
}
