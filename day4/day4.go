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

func findCoin(key, prefix string) (index int, err error) {
	max := math.MaxInt32
	for index = 0; index < max; index++ {
		b := []byte(key + strconv.Itoa(index))
		h := fmt.Sprintf("%x", md5.Sum(b))
		if strings.HasPrefix(h, prefix) {
			return
		}
	}
	err = fmt.Errorf("no coin found below %d", max)
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

	i, err := findCoin(s, "00000")
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
