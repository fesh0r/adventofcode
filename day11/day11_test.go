package main

import "testing"

func TestIncRune(t *testing.T) {
	tests := []struct {
		in   rune
		out  rune
		outC bool
	}{
		{'a', 'b', false},
		{'m', 'n', false},
		{'z', 'a', true},
	}

	for _, tt := range tests {
		v, c, err := incRune(tt.in)
		if err != nil {
			t.Errorf("incRune(%q) = error %s, want %q, %t", tt.in, err, tt.out, tt.outC)
		} else if v != tt.out || c != tt.outC {
			t.Errorf("incRune(%q) = %q, %t, want %q, %t", tt.in, v, c, tt.out, tt.outC)
		}
	}
}

func TestIncRuneError(t *testing.T) {
	tests := []rune{
		0,
		'A',
		'☃',
	}

	for _, tt := range tests {
		v, c, err := incRune(tt)
		if err == nil {
			t.Errorf("incRune(%q) = %q, %t, want error", tt, v, c)
		}
	}
}

func TestIncString(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"a", "b"},
		{"m", "n"},
		{"z", "aa"},
		{"xyzzzzzz", "xzaaaaaa"},
	}

	for _, tt := range tests {
		r, err := incString([]rune(tt.in))
		v := string(r)
		if err != nil {
			t.Errorf("incString(%q) = error %s, want %q", tt.in, err, tt.out)
		} else if v != tt.out {
			t.Errorf("incString(%q) = %q, want %q", tt.in, v, tt.out)
		}
	}
}

func TestIncStringError(t *testing.T) {
	tests := []string{
		"\x00",
		"A",
		"☃",
		"abc\x00",
	}

	for _, tt := range tests {
		r, err := incString([]rune(tt))
		v := string(r)
		if err == nil {
			t.Errorf("incString(%q) = %q, want error", tt, v)
		}
	}
}

func TestHasNoBad(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{"a", true},
		{"m", true},
		{"xyzzzzzz", true},
		{"o", false},
		{"i", false},
		{"l", false},
		{"oiler", false},
	}

	for _, tt := range tests {
		v := hasNoBad([]rune(tt.in))
		if v != tt.out {
			t.Errorf("hasNoBad(%q) = %t, want %t", tt.in, v, tt.out)
		}
	}
}

func TestHasStraight(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{"a", false},
		{"ab", false},
		{"abc", true},
		{"abbc", false},
		{"hijklmmn", true},
		{"abbcemno", true},
	}

	for _, tt := range tests {
		v := hasStraight([]rune(tt.in))
		if v != tt.out {
			t.Errorf("hasStraight(%q) = %t, want %t", tt.in, v, tt.out)
		}
	}
}

func TestHasRepeated(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{"a", false},
		{"aa", false},
		{"aaa", false},
		{"aabaa", false},
		{"aabb", true},
		{"abbcegjk", false},
		{"abbceffg", true},
	}

	for _, tt := range tests {
		v := hasRepeated([]rune(tt.in))
		if v != tt.out {
			t.Errorf("hasRepeated(%q) = %t, want %t", tt.in, v, tt.out)
		}
	}
}

func TestNextPassword(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"abcdefgh", "abcdffaa"},
		{"ghijklmn", "ghjaabcc"},
	}

	for _, tt := range tests {
		v, err := nextPassword(tt.in)
		if err != nil {
			t.Errorf("nextPassword(%q) = error %s, want %q", tt.in, err, tt.out)
		} else if v != tt.out {
			t.Errorf("nextPassword(%q) = %q, want %q", tt.in, v, tt.out)
		}
	}
}

func TestNextPasswordError(t *testing.T) {
	tests := []string{
		"\x00",
		"A",
		"☃",
		"abc\x00",
		"aaaaaaaaa",
		"zzzzzzzz",
	}

	for _, tt := range tests {
		v, err := nextPassword(tt)
		if err == nil {
			t.Errorf("nextPassword(%q) = %q, want error", tt, v)
		}
	}
}

func TestProcess(t *testing.T) {
	tests := []struct {
		in  string
		out string
		out2 string
	}{
		{"abcdefgh", "abcdffaa", "abcdffbb"},
		{"ghijklmn", "ghjaabcc", "ghjbbcdd"},
	}

	for _, tt := range tests {
		v, v2, err := process(tt.in)
		if err != nil {
			t.Errorf("process(%q) = error %s, want %q, %q", tt.in, err, tt.out, tt.out2)
		} else if v != tt.out || v2 != tt.out2 {
			t.Errorf("process(%q) = %q, %q, want %q, %q", tt.in, v, v2, tt.out, tt.out2)
		}
	}
}

func TestProcessError(t *testing.T) {
	tests := []string{
		"\x00",
		"A",
		"☃",
		"abc\x00",
		"aaaaaaaaa",
		"zzzzzzzz",
	}

	for _, tt := range tests {
		v, v2, err := process(tt)
		if err == nil {
			t.Errorf("process(%q) = %q, %q, want error", tt, v, v2)
		}
	}
}
