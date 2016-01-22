package main

import (
	"strings"
	"testing"
)

func TestParseLine(t *testing.T) {
	tests := []struct {
		in       string
		outName  string
		outSpeed int
		outFly   int
		outRest  int
	}{
		{"Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.", "Comet", 14, 10, 127},
		{"Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.", "Dancer", 16, 11, 162},
	}

	for _, tt := range tests {
		name, speed, fly, rest, err := parseLine(tt.in)
		if err != nil {
			t.Errorf("parseLine(%q) = error %s, want %s, %d, %d, %d",
				tt.in, err, tt.outName, tt.outSpeed, tt.outFly, tt.outRest)
		} else if name != tt.outName || speed != tt.outSpeed || fly != tt.outFly || rest != tt.outRest {
			t.Errorf("parseLine(%q) = %s, %d, %d, %d, want %s, %d, %d, %d",
				tt.in, name, speed, fly, rest, tt.outName, tt.outSpeed, tt.outFly, tt.outRest)
		}
	}
}

func TestParseLineError(t *testing.T) {
	tests := []string{
		"",
		"Comet",
		"Comet can fly 14Z km/s for 10 seconds, but then must rest for 127 seconds.",
		"Comet can fly 14 km/s for 10.0 seconds, but then must rest for 127 seconds.",
		"Comet can fly 14 km/s for 10.0 seconds, but then must rest for ABC seconds.",
	}

	for _, tt := range tests {
		name, speed, fly, rest, err := parseLine(tt)
		if err == nil {
			t.Errorf("parseLine(%q) = %s, %d, %d, %d, want error", tt, name, speed, fly, rest)
		}
	}
}

func TestProcess(t *testing.T) {
	tests := []struct {
		in          string
		inTime      int
		outDistance int
		outPoints   int
	}{
		{"Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.\nDancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.\n",
			1000, 1120, 689},
	}

	for _, tt := range tests {
		distance, points, err := process(strings.NewReader(tt.in), tt.inTime)
		if err != nil {
			t.Errorf("process(%q, %d) = error %s, want %d, %d", tt.in, tt.inTime, err, tt.outDistance, tt.outPoints)
		} else if distance != tt.outDistance || points != tt.outPoints {
			t.Errorf("process(%q, %d) = %d, %d want %d, %d",
				tt.in, tt.inTime, distance, points, tt.outDistance, tt.outPoints)
		}
	}
}

func TestProcessError(t *testing.T) {
	tests := []struct {
		in     string
		inTime int
	}{
		{"Comet", 1000},
		{"Comet can fly 14Z km/s for 10 seconds, but then must rest for 127 seconds.", 1000},
		{"Comet can fly 14 km/s for 10.0 seconds, but then must rest for 127 seconds.", 1000},
		{"Comet can fly 14 km/s for 10.0 seconds, but then must rest for ABC seconds.", 1000},
	}

	for _, tt := range tests {
		distance, points, err := process(strings.NewReader(tt.in), tt.inTime)
		if err == nil {
			t.Errorf("process(%q, %d) = %d, %d, want error", tt.in, tt.inTime, distance, points)
		}
	}
}
