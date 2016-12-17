package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestSortFreqMap(t *testing.T) {
	tests := []struct {
		in  freq
		out pairList
	}{
		{
			freq{
				'a': 5, 'b': 3, 'x': 1, 'y': 1, 'z': 1,
			},
			pairList{
				{'a', 5}, {'b', 3}, {'x', 1}, {'y', 1}, {'z', 1},
			},
		},
		{
			freq{
				'a': 1, 'b': 1, 'c': 1, 'd': 1, 'e': 1, 'f': 1, 'g': 1, 'h': 1,
			},
			pairList{
				{'a', 1}, {'b', 1}, {'c', 1}, {'d', 1}, {'e', 1}, {'f', 1}, {'g', 1}, {'h', 1},
			},
		},
		{
			freq{
				'a': 2, 'e': 1, 'l': 1, 'm': 1, 'n': 1, 'o': 3, 'r': 2, 't': 1,
			},
			pairList{
				{'o', 3}, {'a', 2}, {'r', 2}, {'e', 1}, {'l', 1}, {'m', 1}, {'n', 1}, {'t', 1},
			},
		},
		{
			freq{
				'a': 2, 'e': 1, 'l': 3, 'm': 1, 'o': 3, 'r': 2, 't': 2, 'y': 1,
			},
			pairList{
				{'l', 3}, {'o', 3}, {'a', 2}, {'r', 2}, {'t', 2}, {'e', 1}, {'m', 1}, {'y', 1},
			},
		},
	}

	for _, tt := range tests {
		f := sortFreqMap(tt.in)
		if !reflect.DeepEqual(f, tt.out) {
			t.Errorf("sortFreqMap(%v) = %v, want %v", tt.in, f, tt.out)
		}
	}
}

func TestProcess(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{
			"eedadn\ndrvtee\neandsr\nraavrd\natevrs\ntsrnev\nsdttsa\nrasrtv\nnssdts\nntnada\nsvetve\ntesnvt\nvntsnd\nvrdear\ndvrsen\nenarar",
			"easter",
		},
	}

	for _, tt := range tests {
		m, err := process(strings.NewReader(tt.in))
		if err != nil {
			t.Errorf("process(%q) = error %s, want %q", tt.in, err, tt.out)
		} else if m != tt.out {
			t.Errorf("process(%q) = %q, want %q", tt.in, m, tt.out)
		}
	}
}
