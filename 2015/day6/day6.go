package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

//go:generate stringer -type=Instruction
type Instruction int

const (
	turnOn Instruction = iota
	turnOff
	toggle
)

type Coordinates struct {
	XStart int
	YStart int
	XEnd   int
	YEnd   int
}

var lineRegexp = regexp.MustCompile("^(turn on|turn off|toggle) (\\d+),(\\d+) through (\\d+),(\\d+)$")

func MakeLights(xSize, ySize uint) [][]bool {
	lights := make([][]bool, ySize)
	allLights := make([]bool, xSize*ySize)
	for i := range lights {
		lights[i], allLights = allLights[:xSize], allLights[xSize:]
	}

	return lights
}

func MakeLights2(xSize, ySize uint) [][]int {
	lights := make([][]int, ySize)
	allLights := make([]int, xSize*ySize)
	for i := range lights {
		lights[i], allLights = allLights[:xSize], allLights[xSize:]
	}

	return lights
}

func ParseLine(s string) (Instruction, Coordinates, error) {
	var err error

	m := lineRegexp.FindStringSubmatch(s)
	if m == nil {
		err = fmt.Errorf("invalid instruction %q", s)
		return 0, Coordinates{}, err
	}

	var inst Instruction
	switch m[1] {
	case "turn on":
		inst = turnOn
	case "turn off":
		inst = turnOff
	case "toggle":
		inst = toggle
	}

	var coords Coordinates
	coords.XStart, err = strconv.Atoi(m[2])
	if err != nil {
		return 0, Coordinates{}, err
	}
	coords.YStart, err = strconv.Atoi(m[3])
	if err != nil {
		return 0, Coordinates{}, err
	}
	coords.XEnd, err = strconv.Atoi(m[4])
	if err != nil {
		return 0, Coordinates{}, err
	}
	coords.YEnd, err = strconv.Atoi(m[5])
	if err != nil {
		return 0, Coordinates{}, err
	}

	if coords.XStart > coords.XEnd || coords.YStart > coords.YEnd {
		err := fmt.Errorf("invalid coordinates %d", coords)
		return 0, Coordinates{}, err
	}

	return inst, coords, nil
}

func ProcessLine(l [][]bool, s string) error {
	inst, coords, err := ParseLine(s)
	if err != nil {
		return err
	}

	if coords.YStart > len(l) || coords.YEnd > len(l) ||
		coords.XStart > len(l[0]) || coords.XEnd > len(l[0]) {
		err := fmt.Errorf("invalid coordinates %d", coords)
		return err
	}

	for y := coords.YStart; y <= coords.YEnd; y++ {
		for x := coords.XStart; x <= coords.XEnd; x++ {
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

func ProcessLine2(l [][]int, s string) error {
	inst, coords, err := ParseLine(s)
	if err != nil {
		return err
	}

	if coords.YStart > len(l) || coords.YEnd > len(l) ||
		coords.XStart > len(l[0]) || coords.XEnd > len(l[0]) {
		err := fmt.Errorf("invalid coordinates %d", coords)
		return err
	}

	for y := coords.YStart; y <= coords.YEnd; y++ {
		for x := coords.XStart; x <= coords.XEnd; x++ {
			switch inst {
			case turnOn:
				l[x][y] += 1
			case turnOff:
				if l[x][y] > 0 {
					l[x][y] -= 1
				}
			case toggle:
				l[x][y] += 2
			}
		}
	}

	return nil
}

func Process(f io.Reader) (int, int, error) {
	lights := MakeLights(1000, 1000)
	lights2 := MakeLights2(1000, 1000)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var err error

		s := scanner.Text()

		err = ProcessLine(lights, s)
		if err != nil {
			return 0, 0, err
		}

		err = ProcessLine2(lights2, s)
		if err != nil {
			return 0, 0, err
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

	var brightness int
	for _, row := range lights2 {
		for _, l := range row {
			brightness += l
		}
	}

	return lightCount, brightness, nil
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

	lightCount, brightness, err := Process(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Printf("lights: %d\n", lightCount)
	fmt.Printf("brightness: %d\n", brightness)

	return 0
}

func main() {
	os.Exit(run())
}
