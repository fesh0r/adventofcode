package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseLine(t *testing.T) {
	tests := []struct {
		in            string
		outName       string
		outCapacity   int
		outDurability int
		outFlavor     int
		outTexture    int
		outCalories   int
	}{
		{"Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8", "Butterscotch", -1, -2, 6, 3, 8},
		{"Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3", "Cinnamon", 2, 3, -2, -1, 3},
	}

	for _, tt := range tests {
		name, capacity, durability, flavor, texture, calories, err := parseLine(tt.in)
		if err != nil {
			t.Errorf("parseLine(%q) = error %s, want %s, %d, %d, %d, %d, %d",
				tt.in, err, tt.outName, tt.outCapacity, tt.outDurability, tt.outFlavor, tt.outTexture, tt.outCalories)
		} else if name != tt.outName || capacity != tt.outCapacity || durability != tt.outDurability ||
			flavor != tt.outFlavor || texture != tt.outTexture || calories != tt.outCalories {
			t.Errorf("parseLine(%q) = %s, %d, %d, %d, %d, %d, want %s, %d, %d, %d, %d, %d",
				tt.in, name, capacity, durability, flavor, texture, calories, tt.outName, tt.outCapacity,
				tt.outDurability, tt.outFlavor, tt.outTexture, tt.outCalories)
		}
	}
}

func TestParseLineError(t *testing.T) {
	tests := []string{
		"",
		"Butterscotch",
		"Butterscotch: capacity -1.0, durability -2, flavor 6, texture 3, calories 8",
		"Butterscotch: capacity -1A, durability -2, flavor 6, texture 3, calories 8",
		"Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories AZ",
		"Butterscotch: capacity 0-1, durability -2, flavor 6, texture 3, calories 8",
	}

	for _, tt := range tests {
		name, capacity, durability, flavor, texture, calories, err := parseLine(tt)
		if err == nil {
			t.Errorf("parseLine(%q) = %s, %d, %d, %d, %d, %d, want error",
				tt, name, capacity, durability, flavor, texture, calories)
		}
	}
}

func TestCombinations(t *testing.T) {
	tests := []struct {
		inN int
		inR int
		out [][]int
	}{
		{1, 1, [][]int{
			{0},
		}},
		{2, 1, [][]int{
			{0}, {1},
		}},
		{2, 2, [][]int{
			{0, 0}, {0, 1}, {1, 1},
		}},
		{4, 2, [][]int{
			{0, 0}, {0, 1}, {0, 2}, {0, 3}, {1, 1}, {1, 2}, {1, 3}, {2, 2}, {2, 3}, {3, 3},
		}},
	}

	for _, tt := range tests {
		var out [][]int
		for v := range combinations(tt.inN, tt.inR) {
			out = append(out, v)
		}
		if !reflect.DeepEqual(out, tt.out) {
			t.Errorf("permutations(%d, %d) = %v, want %v", tt.inN, tt.inR, out, tt.out)
		}
	}
}

func TestProcess(t *testing.T) {
	tests := []struct {
		in        string
		inN       int
		inC       int
		outScore  int
		outScoreC int
	}{
		{"Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8\nCinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3",
			100, 500, 62842880, 57600000},
	}

	for _, tt := range tests {
		score, scoreC, err := process(strings.NewReader(tt.in), tt.inN, tt.inC)
		if err != nil {
			t.Errorf("process(%q, %d) = error %s, want %d, %d", tt.in, tt.inN, err, tt.outScore, tt.outScoreC)
		} else if score != tt.outScore || scoreC != tt.outScoreC {
			t.Errorf("process(%q, %d) = %d, %d, want %d, %d", tt.in, tt.inN, score, scoreC, tt.outScore, tt.outScoreC)
		}
	}
}

func TestProcessError(t *testing.T) {
	tests := []struct {
		in  string
		inN int
		inC int
	}{
		{"Butterscotch", 100, 500},
		{"Butterscotch: capacity -1.0, durability -2, flavor 6, texture 3, calories 8", 100, 500},
		{"Butterscotch: capacity -1A, durability -2, flavor 6, texture 3, calories 8", 100, 500},
		{"Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories AZ", 100, 500},
		{"Butterscotch: capacity 0-1, durability -2, flavor 6, texture 3, calories 8", 100, 500},
	}

	for _, tt := range tests {
		score, scoreC, err := process(strings.NewReader(tt.in), tt.inN, tt.inC)
		if err == nil {
			t.Errorf("process(%q, %d) = %d, %d, want error", tt.in, tt.inN, score, scoreC)
		}
	}
}
