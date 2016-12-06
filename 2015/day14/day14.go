package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

var lineRegexp = regexp.MustCompile(
	`^(.+) can fly ([0-9]+) km/s for ([0-9]+) seconds, but then must rest for ([0-9]+) seconds\.$`)

func ParseLine(s string) (string, int, int, int, error) {
	errFormat := fmt.Errorf("invalid line %q", s)

	m := lineRegexp.FindStringSubmatch(s)
	if m == nil {
		return "", 0, 0, 0, errFormat
	}

	name := m[1]

	speed, err := strconv.Atoi(m[2])
	if err != nil {
		return "", 0, 0, 0, errFormat
	}

	fly, err := strconv.Atoi(m[3])
	if err != nil {
		return "", 0, 0, 0, errFormat
	}

	rest, err := strconv.Atoi(m[4])
	if err != nil {
		return "", 0, 0, 0, errFormat
	}

	return name, speed, fly, rest, nil
}

type Deer struct {
	Name     string
	Speed    int
	Fly      int
	Rest     int
	Distance int
	Flying   bool
	Next     int
	Points   int
}

func Process(f io.Reader, endTime int) (int, int, error) {
	var allDeer []Deer

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()

		name, speed, fly, rest, err := ParseLine(s)
		if err != nil {
			return 0, 0, err
		}

		allDeer = append(allDeer, Deer{Name: name, Speed: speed, Fly: fly, Rest: rest})
	}

	for t := 0; t < endTime; t++ {
		for k := range allDeer {
			d := &allDeer[k]

			if d.Next == t {
				if d.Flying {
					d.Flying = false
					d.Next += d.Rest
				} else {
					d.Flying = true
					d.Next += d.Fly
				}
			}

			if d.Flying {
				d.Distance += d.Speed
			}
		}

		var curMax int
		for i, d := range allDeer {
			if i == 0 || d.Distance > curMax {
				curMax = d.Distance
			}
		}

		for i := range allDeer {
			if allDeer[i].Distance == curMax {
				allDeer[i].Points++
			}
		}
	}

	var maxDistance, maxPoints int
	for _, d := range allDeer {
		if d.Distance > maxDistance {
			maxDistance = d.Distance
		}
		if d.Points > maxPoints {
			maxPoints = d.Points
		}
	}

	return maxDistance, maxPoints, nil
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

	distance, points, err := Process(f, 2503)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Printf("distance: %d\npoints: %d\n", distance, points)

	return 0
}

func main() {
	os.Exit(run())
}
