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

func CheckIndex(key, prefix string, index int) bool {
	b := []byte(key + strconv.Itoa(index))
	h := fmt.Sprintf("%x", md5.Sum(b))

	if strings.HasPrefix(h, prefix) {
		return true
	}

	return false
}

func FindCoin(key, prefix string) (int, error) {
	max := math.MaxInt32

	for index := 0; index < max; index++ {
		if CheckIndex(key, prefix, index) {
			return index, nil
		}
	}

	err := fmt.Errorf("no coin found below %d", max)
	return 0, err
}

func Process(s string) (int, int, error) {
	i5, err := FindCoin(s, "00000")
	if err != nil {
		return 0, 0, err
	}

	i6, err := FindCoin(s, "000000")
	if err != nil {
		return 0, 0, err
	}

	return i5, i6, nil
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

	i5, i6, err := Process(s)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Printf("index5: %d\nindex6: %d\n", i5, i6)
	return 0
}

func main() {
	os.Exit(run())
}
