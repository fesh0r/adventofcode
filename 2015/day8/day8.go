package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func UnquotedSize(s string) (int, int, error) {
	unquoted, err := strconv.Unquote(s)
	if err != nil {
		return 0, 0, err
	}

	sizeCode := len(s)
	sizeMem := len(unquoted)

	return sizeCode, sizeMem, nil
}

func QuotedSize(s string) (int, int) {
	quoted := strconv.Quote(s)

	sizeCode := len(quoted)
	sizeMem := len(s)

	return sizeCode, sizeMem
}

func Process(f io.Reader) (int, int, error) {
	var size, size2 int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()

		c, m, err := UnquotedSize(s)
		if err != nil {
			return 0, 0, err
		}

		c2, _ := QuotedSize(s)

		size += c - m
		size2 += c2 - c
	}

	return size, size2, nil
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

	v, v2, err := Process(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Printf("v: %d\nv2: %d\n", v, v2)

	return 0
}

func main() {
	os.Exit(run())
}
