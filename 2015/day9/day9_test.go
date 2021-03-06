package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseLine(t *testing.T) {
	tests := []struct {
		in          string
		outFrom     string
		outTo       string
		outDistance int
	}{
		{"London to Dublin = 464", "London", "Dublin", 464},
		{"London to Belfast = 518", "London", "Belfast", 518},
		{"Dublin to Belfast = 141", "Dublin", "Belfast", 141},
	}

	for _, tt := range tests {
		from, to, distance, err := parseLine(tt.in)
		if err != nil {
			t.Errorf("parseLine(%q) = error %s, want %s, %s, %d", tt.in, err, tt.outFrom, tt.outTo, tt.outDistance)
		} else if from != tt.outFrom || to != tt.outTo || distance != tt.outDistance {
			t.Errorf("parseLine(%q) = %s, %s, %d, want %s, %s, %d",
				tt.in, from, to, distance, tt.outFrom, tt.outTo, tt.outDistance)
		}
	}
}

func TestParseLineError(t *testing.T) {
	tests := []string{
		"",
		"London",
		"London to Dublin",
		"London to Dublin = 0.464",
		"London to Dublin = 464Z",
	}

	for _, tt := range tests {
		from, to, distance, err := parseLine(tt)
		if err == nil {
			t.Errorf("parseLine(%q) = %s, %s, %d, want error", tt, from, to, distance)
		}
	}
}

func TestFindIndex(t *testing.T) {
	tests := []struct {
		in       []string
		inFind   string
		outIndex int
		outFound bool
	}{
		{[]string{"London", "Dublin", "Belfast"}, "Dublin", 1, true},
		{[]string{"London", "Dublin", "Belfast"}, "London", 0, true},
		{[]string{"London", "Dublin", "Belfast"}, "Test", 0, false},
	}

	for _, tt := range tests {
		index, found := findIndex(tt.in, tt.inFind)
		if index != tt.outIndex || found != tt.outFound {
			t.Errorf("findIndex(%q, %q) = %d, %t, want %d, %t",
				tt.in, tt.inFind, index, found, tt.outIndex, tt.outFound)
		}
	}
}

func TestPermutations(t *testing.T) {
	tests := []struct {
		in  int
		out [][]int
	}{
		{1, [][]int{
			{0},
		}},
		{2, [][]int{
			{0, 1},
			{1, 0},
		}},
		{3, [][]int{
			{0, 1, 2},
			{0, 2, 1},
			{1, 0, 2},
			{1, 2, 0},
			{2, 0, 1},
			{2, 1, 0},
		}},
	}

	for _, tt := range tests {
		var out [][]int
		for v := range permutations(tt.in) {
			out = append(out, v)
		}
		if !reflect.DeepEqual(out, tt.out) {
			t.Errorf("permutations(%d) = %v, want %v", tt.in, out, tt.out)
		}
	}
}

func TestProcess(t *testing.T) {
	tests := []struct {
		in   string
		outL int
		outH int
	}{
		{"London to Dublin = 464\nLondon to Belfast = 518\nDublin to Belfast = 141\n", 605, 982},
	}

	for _, tt := range tests {
		lowest, highest, err := process(strings.NewReader(tt.in))
		if err != nil {
			t.Errorf("process(%q) = error %s, want %d, %d", tt.in, err, tt.outL, tt.outH)
		} else if lowest != tt.outL || highest != tt.outH {
			t.Errorf("process(%q) = %d, %d, want %d, %d", tt.in, lowest, highest, tt.outL, tt.outH)
		}
	}
}

func TestProcessError(t *testing.T) {
	tests := []string{
		"London",
		"London to Dublin",
		"London to Dublin = 0.464",
		"London to Dublin = 464Z",
		"London to Dublin = 464\nLondon to Belfast = 518\nDublin to London = 141\n",
		"London to Dublin = 464\nLondon to Belfast = 518\n",
	}

	for _, tt := range tests {
		lowest, highest, err := process(strings.NewReader(tt))
		if err == nil {
			t.Errorf("process(%q) = %d, %d, want error", tt, lowest, highest)
		}
	}
}
