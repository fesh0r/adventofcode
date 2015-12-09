package main

import "fmt"

func getFloor(s string) (floor int, err error) {
	for _, c := range s {
		switch c {
		case '(':
			floor++
		case ')':
			floor--
		default:
			err = fmt.Errorf("Character '%c' invalid", c)
			return
		}
	}
	return
}

func main() {
	s := "())"
	f, err := getFloor(s)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("spec: '%s'\nfloor: %d\n", s, f)
	}
}
