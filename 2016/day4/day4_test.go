package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseRoom(t *testing.T) {
	tests := []struct {
		in  string
		out room
	}{
		{
			"aaaaa-bbb-z-y-x-123[abxyz]",
			room{"aaaaa-bbb-z-y-x", 123, "abxyz"},
		},
		{
			"a-b-c-d-e-f-g-h-987[abcde]",
			room{"a-b-c-d-e-f-g-h", 987, "abcde"},
		},
		{
			"not-a-real-room-404[oarel]",
			room{"not-a-real-room", 404, "oarel"},
		},
		{
			"totally-real-room-200[decoy]",
			room{"totally-real-room", 200, "decoy"},
		},
	}

	for _, tt := range tests {
		s, err := parseRoom(tt.in)
		if err != nil {
			t.Errorf("parseRoom(%q) = error %s, want %v", tt.in, err, tt.out)
		} else if s != tt.out {
			t.Errorf("parseRoom(%q) = %v, want %v", tt.in, s, tt.out)
		}
	}
}

func TestParseRoomError(t *testing.T) {
	tests := []string{
		"",
		"aaaaa",
		"a-b-c-d-e-f-g-h-987[abcde",
		"not-a-real-room-aa[oarel]",
		"totally-real-room-200[decoy1]",
		"totally-real-room-200[deco]",
	}

	for _, tt := range tests {
		s, err := parseRoom(tt)
		if err == nil {
			t.Errorf("parseRoom(%q) = %v, want error", tt, s)
		}
	}
}

func TestGetFreq(t *testing.T) {
	tests := []struct {
		in  string
		out freq
	}{
		{
			"aaaaabbbzyx",
			freq{
				'a': 5, 'b': 3, 'x': 1, 'y': 1, 'z': 1,
			},
		},
		{
			"abcdefgh",
			freq{
				'a': 1, 'b': 1, 'c': 1, 'd': 1, 'e': 1, 'f': 1, 'g': 1, 'h': 1,
			},
		},
		{
			"notarealroom",
			freq{
				'a': 2, 'e': 1, 'l': 1, 'm': 1, 'n': 1, 'o': 3, 'r': 2, 't': 1,
			},
		},
		{
			"totallyrealroom",
			freq{
				'a': 2, 'e': 1, 'l': 3, 'm': 1, 'o': 3, 'r': 2, 't': 2, 'y': 1,
			},
		},
	}

	for _, tt := range tests {
		f := getFreq(tt.in)
		if !reflect.DeepEqual(f, tt.out) {
			t.Errorf("getFreq(%q) = %v, want %v", tt.in, f, tt.out)
		}
	}
}

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

func TestGetCode(t *testing.T) {
	tests := []struct {
		in  room
		out string
	}{
		{
			room{"aaaaa-bbb-z-y-x", 123, "abxyz"},
			"abxyz",
		},
		{
			room{"a-b-c-d-e-f-g-h", 987, "abcde"},
			"abcde",
		},
		{
			room{"not-a-real-room", 404, "oarel"},
			"oarel",
		},
		{
			room{"totally-real-room", 200, "decoy"},
			"loart",
		},
	}

	for _, tt := range tests {
		c := getCode(tt.in)
		if c != tt.out {
			t.Errorf("getCode(%v) = %v, want %v", tt.in, c, tt.out)
		}
	}
}

func TestCheckRoom(t *testing.T) {
	tests := []struct {
		in  room
		out bool
	}{
		{
			room{"aaaaa-bbb-z-y-x", 123, "abxyz"},
			true,
		},
		{
			room{"a-b-c-d-e-f-g-h", 987, "abcde"},
			true,
		},
		{
			room{"not-a-real-room", 404, "oarel"},
			true,
		},
		{
			room{"totally-real-room", 200, "decoy"},
			false,
		},
	}

	for _, tt := range tests {
		ok := checkRoom(tt.in)
		if ok != tt.out {
			t.Errorf("checkRoom(%v) = %v, want %v", tt.in, ok, tt.out)
		}
	}
}

func TestProcess(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{
			"aaaaa-bbb-z-y-x-123[abxyz]\na-b-c-d-e-f-g-h-987[abcde]\nnot-a-real-room-404[oarel]\ntotally-real-room-200[decoy]",
			1514,
		},
	}

	for _, tt := range tests {
		s, err := process(strings.NewReader(tt.in))
		if err != nil {
			t.Errorf("process(%q) = error %s, want %d", tt.in, err, tt.out)
		} else if s != tt.out {
			t.Errorf("process(%q) = %d, want %d", tt.in, s, tt.out)
		}
	}
}

func TestProcessError(t *testing.T) {
	tests := []string{
		"aaaaa",
		"a-b-c-d-e-f-g-h-987[abcde",
		"not-a-real-room-aa[oarel]",
		"totally-real-room-200[decoy1]",
	}

	for _, tt := range tests {
		s, err := process(strings.NewReader(tt))
		if err == nil {
			t.Errorf("process(%q) = %d, want error", tt, s)
		}
	}
}
