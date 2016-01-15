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
		d, _ := parse(tt.in)
		v, err := walkVar(d)
		sort.Ints(v)
		sort.Ints(tt.out)
		if err != nil {
			t.Errorf("walkVar(`%v`) = error %s, want `%v`", tt.in, err, tt.out)
		} else if !reflect.DeepEqual(v, tt.out) {
			t.Errorf("walkVar(`%v`) = `%v`, want `%v`", tt.in, v, tt.out)
		}
	}
}

func TestProcess(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{`[1,2,3]`, 6},
		{`[[[3]]]`, 3},
		{`{"a":{"b":4},"c":-1}`, 3},
		{`{"a":[-1,1]}`, 0},
		{`[-1,{"a":1}]`, 0},
		{`[]`, 0},
		{`{}`, 0},
	}

	for _, tt := range tests {
		v, err := process(tt.in)
		if err != nil {
			t.Errorf("process(%#q) = error %s, want %d", tt.in, err, tt.out)
		} else if v != tt.out {
			t.Errorf("process(%#q) = %d, want %d", tt.in, v, tt.out)
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
		v, err := process(tt)
		if err == nil {
			t.Errorf("process(%#q) = %d, want error", tt, v)
		}
	}
}
