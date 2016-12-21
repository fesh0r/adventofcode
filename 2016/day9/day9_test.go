package main

import "testing"

func TestExpand(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"ADVENT", "ADVENT"},
		{"A(1x5)BC", "ABBBBBC"},
		{"(3x3)XYZ", "XYZXYZXYZ"},
		{"A(2x2)BCD(2x2)EFG", "ABCBCDEFEFG"},
		{"(6x1)(1x3)A", "(1x3)A"},
		{"X(8x2)(3x3)ABCY", "X(3x3)ABC(3x3)ABCY"},
	}

	for _, tt := range tests {
		o, err := expand(tt.in)
		if err != nil {
			t.Errorf("expand(%q) = error %s, want %q", tt.in, err, tt.out)
		} else if o != tt.out {
			t.Errorf("expand(%q) = %q, want %q", tt.in, o, tt.out)
		}
	}
}

func TestProcess(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{"ADVENT", 6},
		{"A(1x5)BC", 7},
		{"(3x3)XYZ", 9},
		{"A(2x2)BCD(2x2)EFG", 11},
		{"(6x1)(1x3)A", 6},
		{"X(8x2)(3x3)ABCY", 18},
	}

	for _, tt := range tests {
		l, err := process(tt.in)
		if err != nil {
			t.Errorf("process(%q) = error %s, want %d", tt.in, err, tt.out)
		} else if l != tt.out {
			t.Errorf("process(%q) = %d, want %d", tt.in, l, tt.out)
		}
	}
}

func TestProcess2(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{"(3x3)XYZ", 9},
		{"X(8x2)(3x3)ABCY", 20},
		{"(27x12)(20x12)(13x14)(7x10)(1x12)A", 241920},
		{"(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN", 445},
	}

	for _, tt := range tests {
		l, err := process2(tt.in)
		if err != nil {
			t.Errorf("process2(%q) = error %s, want %d", tt.in, err, tt.out)
		} else if l != tt.out {
			t.Errorf("process2(%q) = %d, want %d", tt.in, l, tt.out)
		}
	}
}
