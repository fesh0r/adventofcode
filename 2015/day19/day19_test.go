package main

import (
	"strings"
	"testing"
)

func TestProcess(t *testing.T) {
	tests := []struct {
		in        string
		out, out2 int
	}{
		{
			"e => H\ne => O\nH => HO\nH => OH\nO => HH\n\nHOH\n",
			4, 3,
		},
		//{
		//	"e => H\ne => O\nH => HO\nH => OH\nO => HH\n\nHOHOHO\n",
		//	7, 6,
		//},
	}

	for _, tt := range tests {
		c, c2, err := process(strings.NewReader(tt.in))
		if err != nil {
			t.Errorf("process(%q) = error %s, want %d, %d", tt.in, err, tt.out, tt.out2)
		} else if c != tt.out {
			t.Errorf("process(%q) = %d, %d, want %d, %d", tt.in, c, c2, tt.out, tt.out2)
		}
	}
}
