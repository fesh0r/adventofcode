package main

import (
	"strings"
	"testing"
)

func TestProcess(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{
			"H => HO\nH => OH\nO => HH\n\nHOH\n",
			4,
		},
		{
			"H => HO\nH => OH\nO => HH\n\nHOHOHO\n",
			7,
		},
	}

	for _, tt := range tests {
		c, err := process(strings.NewReader(tt.in))
		if err != nil {
			t.Errorf("process(%q) = error %s, want %d", tt.in, err, tt.out)
		} else if c != tt.out {
			t.Errorf("process(%q) = %d, want %d", tt.in, c, tt.out)
		}
	}
}
