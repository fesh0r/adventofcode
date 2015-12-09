package main

import "fmt"

func getFloor(s string) (floor int) {
	for _, c := range s {
		switch c {
		case '(':
			floor++
		case ')':
			floor--
		}
	}
	return
}

func main() {
	s := "())"
	f := getFloor(s)
	fmt.Printf("spec: '%s'\nfloor: %d\n", s, f)
}
