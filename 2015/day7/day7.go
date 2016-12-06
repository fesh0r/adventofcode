package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

//go:generate stringer -type=OpType
type OpType int

const (
	opSet OpType = iota
	opNot
	opAnd
	opOr
	opLShift
	opRShift
)

var opLookup = map[string]OpType{
	"":       opSet,
	"NOT":    opNot,
	"AND":    opAnd,
	"OR":     opOr,
	"LSHIFT": opLShift,
	"RSHIFT": opRShift,
}

type Operation struct {
	P1, P2 string
	Op     OpType
	W1, W2 *Operation
	Value  uint16
	Valid  bool
}

func (o Operation) String() string {
	return fmt.Sprintf("{%q %v %q}", o.P1, o.Op, o.P2)
}

func (o *Operation) GetValue() (uint16, error) {
	if o.Valid {
		return o.Value, nil
	}

	var wire1, wire2 uint16
	var err error

	if o.W1 != nil {
		wire1, err = o.W1.GetValue()
		if err != nil {
			return 0, err
		}
	} else if o.P1 != "" {
		v, _ := strconv.ParseUint(o.P1, 10, 16)
		wire1 = uint16(v)
	}
	if o.W2 != nil {
		wire2, err = o.W2.GetValue()
		if err != nil {
			return 0, err
		}
	} else if o.P2 != "" {
		v, _ := strconv.ParseUint(o.P2, 10, 16)
		wire2 = uint16(v)
	}

	var result uint16

	switch o.Op {
	case opSet:
		result = wire1
	case opNot:
		result = ^wire1
	case opAnd:
		result = wire1 & wire2
	case opOr:
		result = wire1 | wire2
	case opLShift:
		result = wire1 << wire2
	case opRShift:
		result = wire1 >> wire2
	}

	o.Value = result
	o.Valid = true

	if !o.Valid {
		err = fmt.Errorf("value not found")
		return 0, err
	}

	return o.Value, nil
}

type Circuit map[string]*Operation

func (c Circuit) ResolveWires() error {
	for k, v := range c {
		if v.P1 != "" {
			if !CheckUint16(v.P1) {
				l1, ok := c[v.P1]
				if !ok {
					return fmt.Errorf("invalid operand %q for %q", v.P1, k)
				}
				v.W1 = l1
			}
		}
		if v.P2 != "" {
			if !CheckUint16(v.P2) {
				l2, ok := c[v.P2]
				if !ok {
					return fmt.Errorf("invalid operand %q for %q", v.P2, k)
				}
				v.W2 = l2
			}
		}
	}

	return nil
}

func (c Circuit) Reset() {
	for _, v := range c {
		v.Value = 0
		v.Valid = false
	}
}

var lineRegexp = regexp.MustCompile(
	"^(?:(?:([a-z]+)|([0-9]+)) )?(?:(NOT|OR|AND|LSHIFT|RSHIFT) )?(?:([a-z]+)|([0-9]+)) -> ([a-z]+)$")

func CheckUint16(s string) bool {
	_, err := strconv.ParseUint(s, 10, 16)
	return err == nil
}

func ParseLine(s string) (Operation, string, error) {
	err := fmt.Errorf("invalid instruction %q", s)

	m := lineRegexp.FindStringSubmatch(s)
	if m == nil {
		return Operation{}, "", err
	}

	op, ok := opLookup[m[3]]
	if !ok {
		return Operation{}, "", err
	}

	var p1, p2 string
	switch op {
	case opSet, opNot:
		if m[1] != "" || m[2] != "" {
			return Operation{}, "", err
		}
		if m[5] != "" {
			if !CheckUint16(m[5]) {
				return Operation{}, "", err
			}
			p1 = m[5]
		} else {
			p1 = m[4]
		}
	case opAnd, opOr:
		if m[1] == "" && m[2] == "" {
			return Operation{}, "", err
		}
		if m[2] != "" {
			if !CheckUint16(m[2]) {
				return Operation{}, "", err
			}
			p1 = m[2]
		} else {
			p1 = m[1]
		}
		if m[5] != "" {
			if !CheckUint16(m[5]) {
				return Operation{}, "", err
			}
			p2 = m[5]
		} else {
			p2 = m[4]
		}
	case opLShift, opRShift:
		if m[1] == "" && m[2] == "" {
			return Operation{}, "", err
		}
		if m[2] != "" {
			if !CheckUint16(m[2]) {
				return Operation{}, "", err
			}
			p1 = m[2]
		} else {
			p1 = m[1]
		}
		if m[5] == "" {
			return Operation{}, "", err
		}
		if !CheckUint16(m[5]) {
			return Operation{}, "", err
		}
		p2 = m[5]
	}

	lineOp := Operation{Op: op, P1: p1, P2: p2}
	lineTo := m[6]

	return lineOp, lineTo, nil
}

func Process(f io.Reader) (uint16, uint16, error) {
	c := make(Circuit)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()

		op, to, err := ParseLine(s)
		if err != nil {
			return 0, 0, err
		}

		if _, ok := c[to]; ok {
			err := fmt.Errorf("wire %q already set", to)
			return 0, 0, err
		}

		c[to] = &op
	}

	err := c.ResolveWires()
	if err != nil {
		return 0, 0, err
	}

	if _, ok := c["a"]; !ok {
		err := fmt.Errorf("wire \"a\" not found")
		return 0, 0, err
	}

	value, err := c["a"].GetValue()
	if err != nil {
		return 0, 0, err
	}

	c.Reset()

	if _, ok := c["b"]; !ok {
		err := fmt.Errorf("wire \"b\" not found")
		return 0, 0, err
	}

	c["b"].Value = value
	c["b"].Valid = true

	value2, err := c["a"].GetValue()
	if err != nil {
		return 0, 0, err
	}

	return value, value2, nil
}

func run() int {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "%s filename\n", os.Args[0])
		return 1
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	defer f.Close()

	value, value2, err := Process(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Printf("a: %d\na2: %d\n", value, value2)

	return 0
}

func main() {
	os.Exit(run())
}
