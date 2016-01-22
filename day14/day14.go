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

func parseLine(s string) (string, int, int, int, error) {
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

type deer struct {
	name     string
	speed    int
	fly      int
	rest     int
	distance int
	flying   bool
	next     int
	points   int
}

func process(f io.Reader, endTime int) (int, int, error) {
	var allDeer []deer

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()

		name, speed, fly, rest, err := parseLine(s)
		if err != nil {
			return 0, 0, err
		}

		allDeer = append(allDeer, deer{name: name, speed: speed, fly: fly, rest: rest})
	}

	for t := 0; t < endTime; t++ {
		for k := range allDeer {
			d := &allDeer[k]

			if d.next == t {
				if d.flying {
					d.flying = false
					d.next += d.rest
				} else {
					d.flying = true
					d.next += d.fly
				}
			}

			if d.flying {
				d.distance += d.speed
			}
		}

		var curMax int
		for i, d := range allDeer {
			if i == 0 || d.distance > curMax {
				curMax = d.distance
			}
		}

		for i := range allDeer {
			if allDeer[i].distance == curMax {
				allDeer[i].points++
			}
		}
	}

	var maxDistance, maxPoints int
	for i, d := range allDeer {
		if i == 0 || d.distance > maxDistance {
			maxDistance = d.distance
		}
		if i == 0 || d.points > maxPoints {
			maxPoints = d.points
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

	distance, points, err := process(f, 2503)
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
