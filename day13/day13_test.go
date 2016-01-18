package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseLine(t *testing.T) {
	tests := []struct {
		in       string
		outFrom  string
		outTo    string
		outHappy int
	}{
		{"Alice would gain 54 happiness units by sitting next to Bob.", "Alice", "Bob", 54},
		{"Alice would lose 79 happiness units by sitting next to Carol.", "Alice", "Carol", -79},
		{"David would gain 41 happiness units by sitting next to Carol.", "David", "Carol", 41},
	}

	for _, tt := range tests {
		from, to, happy, err := parseLine(tt.in)
		if err != nil {
			t.Errorf("parseLine(%q) = error %s, want %s, %s, %d", tt.in, err, tt.outFrom, tt.outTo, tt.outHappy)
		} else if from != tt.outFrom || to != tt.outTo || happy != tt.outHappy {
			t.Errorf("parseLine(%q) = %s, %s, %d, want %s, %s, %d",
				tt.in, from, to, happy, tt.outFrom, tt.outTo, tt.outHappy)
		}
	}
}

func TestParseLineError(t *testing.T) {
	tests := []string{
		"",
		"Alice",
		"Alice would lose 79 happiness units",
		"Alice would lose -2 happiness units by sitting next to David.",
		"Alice would lose 2.0 happiness units by sitting next to David.",
		"Alice would lose 2Z happiness units by sitting next to David.",
	}

	for _, tt := range tests {
		from, to, happy, err := parseLine(tt)
		if err == nil {
			t.Errorf("parseLine(%q) = %s, %s, %d, want error", tt, from, to, happy)
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
		{[]string{"Alice", "Bob", "Carol"}, "Bob", 1, true},
		{[]string{"Alice", "Bob", "Carol"}, "Alice", 0, true},
		{[]string{"Alice", "Bob", "Carol"}, "Test", 0, false},
	}

	for _, tt := range tests {
		index, found := findIndex(tt.in, tt.inFind)
		if index != tt.outIndex || found != tt.outFound {
			t.Errorf("findIndex(%q, %q) = %d, %t, want %d, %t",
				tt.in, tt.inFind, index, found, tt.outIndex, tt.outFound)
		}
	}
}

func TestAppendCopy(t *testing.T) {
	tests := []struct {
		in    [][]int
		inAdd []int
		out   [][]int
	}{
		{[][]int{}, []int{0}, [][]int{{0}}},
		{[][]int{{0}}, []int{0}, [][]int{{0}, {0}}},
		{[][]int{{0}}, []int{1}, [][]int{{0}, {1}}},
		{[][]int{{0}, {1}}, []int{1}, [][]int{{0}, {1}, {1}}},
		{[][]int{{0, 1}, {1, 2}}, []int{2, 1}, [][]int{{0, 1}, {1, 2}, {2, 1}}},
	}

	for _, tt := range tests {
		out := appendCopy(tt.in, tt.inAdd)
		if !reflect.DeepEqual(out, tt.out) {
			t.Errorf("appendCopy(%v, %v) = %v, want %v", tt.in, tt.inAdd, out, tt.out)
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
		out := permutations(tt.in)
		if !reflect.DeepEqual(out, tt.out) {
			t.Errorf("permutations(%d) = %v, want %v", tt.in, out, tt.out)
		}
	}
}

func TestProcess(t *testing.T) {
	tests := []struct {
		in        string
		inAddSelf bool
		out       int
	}{
		{"Alice would gain 54 happiness units by sitting next to Bob.\nAlice would lose 79 happiness units by sitting next to Carol.\nAlice would lose 2 happiness units by sitting next to David.\nBob would gain 83 happiness units by sitting next to Alice.\nBob would lose 7 happiness units by sitting next to Carol.\nBob would lose 63 happiness units by sitting next to David.\nCarol would lose 62 happiness units by sitting next to Alice.\nCarol would gain 60 happiness units by sitting next to Bob.\nCarol would gain 55 happiness units by sitting next to David.\nDavid would gain 46 happiness units by sitting next to Alice.\nDavid would lose 7 happiness units by sitting next to Bob.\nDavid would gain 41 happiness units by sitting next to Carol.\n", false, 330},
		{"Alice would gain 54 happiness units by sitting next to Bob.\nAlice would lose 79 happiness units by sitting next to Carol.\nAlice would lose 2 happiness units by sitting next to David.\nBob would gain 83 happiness units by sitting next to Alice.\nBob would lose 7 happiness units by sitting next to Carol.\nBob would lose 63 happiness units by sitting next to David.\nCarol would lose 62 happiness units by sitting next to Alice.\nCarol would gain 60 happiness units by sitting next to Bob.\nCarol would gain 55 happiness units by sitting next to David.\nDavid would gain 46 happiness units by sitting next to Alice.\nDavid would lose 7 happiness units by sitting next to Bob.\nDavid would gain 41 happiness units by sitting next to Carol.\n", true, 286},
	}

	for _, tt := range tests {
		highest, err := process(strings.NewReader(tt.in), tt.inAddSelf)
		if err != nil {
			t.Errorf("process(%q, %t) = error %s, want %d", tt.in, tt.inAddSelf, err, tt.out)
		} else if highest != tt.out {
			t.Errorf("process(%q, %t) = %d, want %d", tt.in, tt.inAddSelf, highest, tt.out)
		}
	}
}

func TestProcessError(t *testing.T) {
	tests := []struct {
		in        string
		inAddSelf bool
	}{
		{"Alice", false},
		{"Alice would lose 79 happiness units", false},
		{"Alice would gain 54 happiness units by sitting next to Bob.\nAlice would lose 79 happiness units by sitting next to Bob.\n", false},
		{"Alice would gain 54 happiness units by sitting next to Bob.\nAlice would lose 79 happiness units by sitting next to Carol.\n", false},
		{"Alice would gain 54 happiness units by sitting next to Bob.\nAlice would lose 79 happiness units by sitting next to Carol.\n", true},
	}

	for _, tt := range tests {
		highest, err := process(strings.NewReader(tt.in), tt.inAddSelf)
		if err == nil {
			t.Errorf("process(%q, %t) = %d, want error", tt.in, tt.inAddSelf, highest)
		}
	}
}
