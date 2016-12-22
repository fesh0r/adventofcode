package main

import (
	"strings"
	"testing"
)

func TestProcess(t *testing.T) {
	tests := []struct {
		in  string
		out int64
	}{
		{"1\n2\n3\n4\n5\n7\n8\n9\n10\n11\n", 99},
	}

	for _, tt := range tests {
		qe, err := process(strings.NewReader(tt.in))
		if err != nil {
			t.Errorf("process(%q) = error %s, want %d", tt.in, err, tt.out)
		} else if qe != tt.out {
			t.Errorf("process(%q) = %d, want %d", tt.in, qe, tt.out)
		}
	}
}
