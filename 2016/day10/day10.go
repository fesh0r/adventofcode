package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

var lineRegex = regexp.MustCompile("value ([0-9]+) goes to bot ([0-9]+)|bot ([0-9]+) gives low to (bot|output) ([0-9]+) and high to (bot|output) ([0-9]+)")

type bot struct {
	values    []int
	lowDest   string
	lowIndex  int
	highDest  string
	highIndex int
}

type factory struct {
	b map[int]bot
	o map[int]int
}

func newFactory() factory {
	var f factory
	f.b = make(map[int]bot)
	f.o = make(map[int]int)

	return f
}

func (f factory) addLine(s string) (int, error) {
	startIndex := -1
	m := lineRegex.FindStringSubmatch(s)
	if m != nil {
		if m[1] != "" {
			index, err := strconv.Atoi(m[2])
			if err != nil {
				return 0, err
			}
			value, err := strconv.Atoi(m[1])
			if err != nil {
				return 0, err
			}

			bot := f.b[index]
			bot.values = append(bot.values, value)
			f.b[index] = bot
			if len(bot.values) > 1 {
				startIndex = index
			}
		} else {
			index, err := strconv.Atoi(m[3])
			if err != nil {
				return 0, err
			}

			bot := f.b[index]

			bot.lowDest = m[4]
			bot.lowIndex, err = strconv.Atoi(m[5])
			if err != nil {
				return 0, err
			}
			bot.highDest = m[6]
			bot.highIndex, err = strconv.Atoi(m[7])
			if err != nil {
				return 0, err
			}

			f.b[index] = bot
		}
	}

	return startIndex, nil
}

func (f factory) propagate(index int) int {
	bot := f.b[index]

	if len(bot.values) != 2 {
		return -1
	}

	var low, high int
	if bot.values[0] < bot.values[1] {
		low, high = bot.values[0], bot.values[1]
	} else {
		low, high = bot.values[1], bot.values[0]
	}
	bot.values = nil
	f.b[index] = bot

	compareIndex := -1
	if low == 17 && high == 61 {
		compareIndex = index
	}

	if bot.lowDest == "bot" {
		botLow := f.b[bot.lowIndex]
		botLow.values = append(botLow.values, low)
		f.b[bot.lowIndex] = botLow
	}
	if bot.highDest == "bot" {
		botHigh := f.b[bot.highIndex]
		botHigh.values = append(botHigh.values, high)
		f.b[bot.highIndex] = botHigh
	}

	if bot.lowDest == "bot" {
		lowClear := f.propagate(bot.lowIndex)
		if lowClear != -1 {
			compareIndex = lowClear
		}
	} else {
		f.o[bot.lowIndex] = low
	}
	if bot.highDest == "bot" {
		highClear := f.propagate(bot.highIndex)
		if highClear != -1 {
			compareIndex = highClear
		}
	} else {
		f.o[bot.highIndex] = high
	}

	return compareIndex
}

func process(f io.Reader) (int, error) {
	fac := newFactory()

	var startIndex int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		i, err := fac.addLine(s)
		if err != nil {
			return 0, err
		}
		if i != -1 {
			startIndex = i
		}
	}

	clearIndex := fac.propagate(startIndex)

	return clearIndex, nil
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

	b, err := process(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Printf("bot: %d\n", b)
	return 0
}

func main() {
	os.Exit(run())
}
