package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestWalkVar(t *testing.T) {
	tests := []struct {
		in  string
		out []int
	}{
		{`[1,2,3]`, []int{1, 2, 3}},
		{`[[[3]]]`, []int{3}},
		{`{"a":{"b":4},"c":-1,"d":23,"e":2,"f":5,"g":6}`, []int{4, -1, 23, 2, 5, 6}},
		{`{"a":[-1,1]}`, []int{-1, 1}},
		{`[-1,{"a":1}]`, []int{-1, 1}},
		{`[]`, []int{}},
		{`{}`, []int{}},
	}

	for _, tt := range tests {
		d, _ := Parse(tt.in)
		v, err := WalkVar(d)
		sort.Ints(v)
		sort.Ints(tt.out)
		if err != nil {
			t.Errorf("WalkVar(`%v`) = error %s, want `%v`", tt.in, err, tt.out)
		} else if !reflect.DeepEqual(v, tt.out) {
			t.Errorf("WalkVar(`%v`) = `%v`, want `%v`", tt.in, v, tt.out)
		}
	}
}

func TestWalkVarSkip(t *testing.T) {
	tests := []struct {
		in   string
		skip string
		out  []int
	}{
		{`[1,2,3]`, "red", []int{1, 2, 3}},
		{`[[[3]]]`, "red", []int{3}},
		{`{"a":{"b":4},"c":-1,"d":23,"e":2,"f":5,"g":6}`, "red", []int{4, -1, 23, 2, 5, 6}},
		{`{"a":[-1,1]}`, "red", []int{-1, 1}},
		{`[-1,{"a":1}]`, "red", []int{-1, 1}},
		{`[]`, "red", []int{}},
		{`{}`, "red", []int{}},
		{`[1,{"c":"red","b":2},3]`, "red", []int{1, 3}},
		{`{"d":"red","e":[1,2,3,4],"f":5}`, "red", []int{}},
		{`[1,"red",5]`, "red", []int{1, 5}},
		{`{"d":"red","e":[1,2,3,4],"f":5}`, "blue", []int{1, 2, 3, 4, 5}},
	}

	for _, tt := range tests {
		d, _ := Parse(tt.in)
		v, err := WalkVarSkip(d, tt.skip)
		sort.Ints(v)
		sort.Ints(tt.out)
		if err != nil {
			t.Errorf("WalkVarSkip(`%v`) = error %s, want `%v`", tt.in, err, tt.out)
		} else if !reflect.DeepEqual(v, tt.out) {
			t.Errorf("WalkVarSkip(`%v`) = `%v`, want `%v`", tt.in, v, tt.out)
		}
	}
}

func TestProcess(t *testing.T) {
	tests := []struct {
		in   string
		out  int
		out2 int
	}{
		{`[1,2,3]`, 6, 6},
		{`[[[3]]]`, 3, 3},
		{`{"a":{"b":4},"c":-1}`, 3, 3},
		{`{"a":[-1,1]}`, 0, 0},
		{`[-1,{"a":1}]`, 0, 0},
		{`[]`, 0, 0},
		{`{}`, 0, 0},
		{`[1,{"c":"red","b":2},3]`, 6, 4},
		{`{"d":"red","e":[1,2,3,4],"f":5}`, 15, 0},
		{`[1,"red",5]`, 6, 6},
	}

	for _, tt := range tests {
		v, v2, err := Process(tt.in)
		if err != nil {
			t.Errorf("Process(%#q) = error %s, want %d, %d", tt.in, err, tt.out, tt.out2)
		} else if v != tt.out || v2 != tt.out2 {
			t.Errorf("Process(%#q) = %d, %d, want %d, %d", tt.in, v, v2, tt.out, tt.out2)
		}
	}
}

func TestProcessError(t *testing.T) {
	tests := []string{
		``,
		`{`,
		`"{`,
	}

	for _, tt := range tests {
		v, v2, err := Process(tt)
		if err == nil {
			t.Errorf("Process(%#q) = %d, %d, want error", tt, v, v2)
		}
	}
}
