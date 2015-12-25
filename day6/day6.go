package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

//go:generate stringer -type=instruction
type instruction int

const (
	turnOn instruction = iota
	turnOff
	toggle
)

type coordinates struct {
	xStart int
	yStart int
	xEnd   int
	yEnd   int
}

var lineRegexp = regexp.MustCompile("^(turn on|turn off|toggle) (\\d+),(\\d+) through (\\d+),(\\d+)$")

func makeLights(xSize, ySize uint) [][]bool {
	lights := make([][]bool, ySize)
	allLights := make([]bool, xSize*ySize)
	for i := range lights {
		lights[i], allLights = allLights[:xSize], allLights[xSize:]
	}

	return lights
}

func parseLine(s string) (instruction, coordinates, error) {
	var err error

	m := lineRegexp.FindStringSubmatch(s)
	if m == nil {
		err = fmt.Errorf("invalid instruction %q", s)
		return 0, coordinates{}, err
	}

	var inst instruction
	switch m[1] {
	case "turn on":
		inst = turnOn
	case "turn off":
		inst = turnOff
	case "toggle":
		inst = toggle
	}

	var coords coordinates
	coords.xStart, err = strconv.Atoi(m[2])
	if err != nil {
		return 0, coordinates{}, err
	}
	coords.yStart, err = strconv.Atoi(m[3])
	if err != nil {
		return 0, coordinates{}, err
	}
	coords.xEnd, err = strconv.Atoi(m[4])
	if err != nil {
		return 0, coordinates{}, err
	}
	coords.yEnd, err = strconv.Atoi(m[5])
	if err != nil {
		return 0, coordinates{}, err
	}

	if coords.xStart > coords.xEnd || coords.yStart > coords.yEnd {
		err := fmt.Errorf("invalid coordinates %d", coords)
		return 0, coordinates{}, err
	}

	return inst, coords, nil
}

func processLine(l [][]bool, s string) error {
	inst, coords, err := parseLine(s)
	if err != nil {
		return err
	}

	if coords.yStart > len(l) || coords.yEnd > len(l) ||
		coords.xStart > len(l[0]) || coords.xEnd > len(l[0]) {
		err := fmt.Errorf("invalid coordinates %d", coords)
		return err
	}

	for y := coords.yStart; y <= coords.yEnd; y++ {
		for x := coords.xStart; x <= coords.xEnd; x++ {
			switch inst {
			case turnOn:
				l[x][y] = true
			case turnOff:
				l[x][y] = false
			case toggle:
				l[x][y] = !l[x][y]
			}
		}
	}

	return nil
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

	lights := makeLights(1000, 1000)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()

		err := processLine(lights, s)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
	}

	var lightCount int
	for _, row := range lights {
		for _, l := range row {
			if l {
				lightCount++
			}
		}
	}

	fmt.Printf("lights: %d\n", lightCount)

	return 0
}

func main() {
	os.Exit(run())
}
