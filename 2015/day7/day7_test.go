package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseLine(t *testing.T) {
	tests := []struct {
		in    string
		outOp Operation
		outTo string
	}{
		{"123 -> x", Operation{Op: opSet, P1: "123"}, "x"},
		{"x AND y -> z", Operation{Op: opAnd, P1: "x", P2: "y"}, "z"},
		{"p LSHIFT 2 -> q", Operation{Op: opLShift, P1: "p", P2: "2"}, "q"},
		{"NOT e -> f", Operation{Op: opNot, P1: "e"}, "f"},
		{"1 OR g -> h", Operation{Op: opOr, P1: "1", P2: "g"}, "h"},
		{"h -> i", Operation{Op: opSet, P1: "h"}, "i"},
	}

	for _, tt := range tests {
		op, to, err := ParseLine(tt.in)
		if err != nil {
			t.Errorf("ParseLine(%q) = error %s, want %v, %q", tt.in, err, tt.outOp, tt.outTo)
		} else if !reflect.DeepEqual(op, tt.outOp) || to != tt.outTo {
			t.Errorf("ParseLine(%q) = %v, %q, want %v, %q", tt.in, op, to, tt.outOp, tt.outTo)
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
		op, to, err := ParseLine(tt)
		if err == nil {
			t.Errorf("ParseLine(%q) = %v, %q, want error", tt, op, to)
		}
	}
}

func TestProcess(t *testing.T) {
	tests := []struct {
		in   string
		out  uint16
		out2 uint16
	}{
		{
			"123 -> x\n456 -> b\nx AND b -> d\nx OR b -> e\nx LSHIFT 2 -> f\nb RSHIFT 2 -> a\nNOT x -> h\nNOT b -> i\n",
			114,
			28,
		},
	}
	for _, tt := range tests {
		value, value2, err := Process(strings.NewReader(tt.in))
		if err != nil {
			t.Errorf("Process(%q) = error %s, want %d, %d", tt.in, err, tt.out, tt.out2)
		} else if value != tt.out || value2 != tt.out2 {
			t.Errorf("Process(%q) = %d, %d, want %d, %d", tt.in, value, value2, tt.out, tt.out2)
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
		value, value2, err := Process(strings.NewReader(tt))
		if err == nil {
			t.Errorf("Process(%q) = %d, %d, want error", tt, value, value2)
		}
	}
}
