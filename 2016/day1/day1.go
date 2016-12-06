package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type position struct {
	x, y int
}

var direction = []position{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func parse(s string) (int, int, error) {
	var err error
	var dir, dist int

	if len(s) < 2 {
		err = fmt.Errorf("invalid instruction %q", s)
		return 0, 0, err
	}

	if s[0] == 'L' {
		dir = -1
	} else if s[0] == 'R' {
		dir = 1
	} else {
		err = fmt.Errorf("invalid direction %q", s[0])
		return 0, 0, err
	}

	dist, err = strconv.Atoi(s[1:])
	if err != nil {
		return 0, 0, err
	}

	if dist < 0 {
		err = fmt.Errorf("invalid distance %d", dist)
		return 0, 0, err
	}

	return dir, dist, nil
}

func distance(p position) int {
	var dist int

	if p.x < 0 {
		dist += -p.x
	} else {
		dist += p.x
	}
	if p.y < 0 {
		dist += -p.y
	} else {
		dist += p.y
	}

	return dist
}

func process(s string) (int, int, error) {
	dirs := strings.Split(s, ", ")

	var pos position
	var curDir, dist2 int
	var found2 bool

	visited := make(map[position]int)

	for _, v := range dirs {
		dir, dist, err := parse(v)
		if err != nil {
			return 0, 0, err
		}
		curDir = (curDir + dir + 4) % 4

		for i := 0; i < dist; i++ {
			pos.x += direction[curDir].x
			pos.y += direction[curDir].y

			visited[pos]++
			if !found2 && visited[pos] > 1 {
				dist2 = distance(pos)
				found2 = true
			}
		}
	}

	fullDist := distance(pos)

	return fullDist, dist2, nil
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

	dist, dist2, err := process(s)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Printf("blocks: %d\nblocks2: %d\n", dist, dist2)
	return 0
}

func main() {
	os.Exit(run())
}
