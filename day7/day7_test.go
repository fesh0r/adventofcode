package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseLine(t *testing.T) {
	tests := []struct {
		in    string
		outOp operation
		outTo string
	}{
		{"123 -> x", operation{op: opSet, p1: "123"}, "x"},
		{"x AND y -> z", operation{op: opAnd, p1: "x", p2: "y"}, "z"},
		{"p LSHIFT 2 -> q", operation{op: opLShift, p1: "p", p2: "2"}, "q"},
		{"NOT e -> f", operation{op: opNot, p1: "e"}, "f"},
		{"1 OR g -> h", operation{op: opOr, p1: "1", p2: "g"}, "h"},
		{"h -> i", operation{op: opSet, p1: "h"}, "i"},
	}

	for _, tt := range tests {
		op, to, err := parseLine(tt.in)
		if err != nil {
			t.Errorf("parseLine(%q) = error %s, want %v, %q", tt.in, err, tt.outOp, tt.outTo)
		} else if !reflect.DeepEqual(op, tt.outOp) || to != tt.outTo {
			t.Errorf("parseLine(%q) = %v, %q, want %v, %q", tt.in, op, to, tt.outOp, tt.outTo)
		}
	}
}

func TestParseLineError(t *testing.T) {
	tests := []string{
		"",
		"->",
		"-> i",
		"a b -> i",
		"a NOT e -> f",
		"p LSHIFT a -> q",
		"70000 -> i",
		"100000 OR a -> g",
		"b AND 1000000 -> g",
	}

	for _, tt := range tests {
		op, to, err := parseLine(tt)
		if err == nil {
			t.Errorf("parseLine(%q) = %v, %q, want error", tt, op, to)
		}
	}
}

func TestProcess(t *testing.T) {
	tests := []struct {
		in  string
		out uint16
		out2 uint16
	}{
		{
			"123 -> x\n456 -> b\nx AND b -> d\nx OR b -> e\nx LSHIFT 2 -> f\nb RSHIFT 2 -> a\nNOT x -> h\nNOT b -> i\n",
			114,
			28,
		},
	}
	for _, tt := range tests {
		value, value2, err := process(strings.NewReader(tt.in))
		if err != nil {
			t.Errorf("process(%q) = error %s, want %d, %d", tt.in, err, tt.out, tt.out2)
		} else if value != tt.out || value2 != tt.out2 {
			t.Errorf("process(%q) = %d, %d, want %d, %d", tt.in, value, value2, tt.out, tt.out2)
		}
	}
}

func TestProcessError(t *testing.T) {
	tests := []string{
		"",
		"->",
		"-> i",
		"a b -> i",
		"a NOT e -> f",
		"p LSHIFT a -> q",
		"70000 -> i",
		"100000 OR a -> g",
		"b AND 1000000 -> g",
		"123 -> b\n456 -> c\n",
	}

	for _, tt := range tests {
		value, value2, err := process(strings.NewReader(tt))
		if err == nil {
			t.Errorf("process(%q) = %d, %d, want error", tt, value, value2)
		}
	}
}
