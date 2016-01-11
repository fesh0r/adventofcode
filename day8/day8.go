package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func unquotedSize(s string) (int, int, error) {
	unquoted, err := strconv.Unquote(s)
	if err != nil {
		return 0, 0, err
	}

	sizeCode := len(s)
	sizeMem := len(unquoted)

	return sizeCode, sizeMem, nil
}

func process(f io.Reader) (int, error) {
	var size int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()

		c, m, err := unquotedSize(s)
		if err != nil {
			return 0, err
		}

		size += c - m
	}

	return size, nil
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

	v, err := process(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Printf("v: %d\n", v)

	return 0
}

func main() {
	os.Exit(run())
}
