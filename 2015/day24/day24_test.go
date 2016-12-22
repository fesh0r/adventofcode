package main

import (
	"strings"
	"testing"
)

func TestProcess(t *testing.T) {
	tests := []struct {
		in         string
		out3, out4 int64
	}{
		{"1\n2\n3\n4\n5\n7\n8\n9\n10\n11\n", 99, 44},
	}

	for _, tt := range tests {
		qe3, qe4, err := process(strings.NewReader(tt.in))
		if err != nil {
			t.Errorf("process(%q) = error %s, want %d, %d", tt.in, err, tt.out3, tt.out4)
		} else if qe3 != tt.out3 || qe4 != tt.out4 {
			t.Errorf("process(%q) = %d, %d, want %d, %d", tt.in, qe3, qe4, tt.out3, tt.out4)
		}
	}
}
