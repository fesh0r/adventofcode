package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func findCoin(key string) (index int, err error) {
	return
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

	i, err := findCoin(s)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Printf("index: %d\n", i)
	return 0
}

func main() {
	os.Exit(run())
}
