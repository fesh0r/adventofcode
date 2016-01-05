package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

//go:generate stringer -type=opType
type opType int

const (
	opSet opType = iota
	opNot
	opAnd
	opOr
	opLShift
	opRShift
)

var opLookup = map[string]opType{
	"":       opSet,
	"NOT":    opNot,
	"AND":    opAnd,
	"OR":     opOr,
	"LSHIFT": opLShift,
	"RSHIFT": opRShift,
}

type operation struct {
	p1, p2 string
	op     opType
	w1, w2 *operation
	value  uint16
	valid  bool
}

func (o operation) String() string {
	return fmt.Sprintf("{%q %v %q}", o.p1, o.op, o.p2)
}

func (o *operation) getValue() (uint16, error) {
	if o.valid {
		return o.value, nil
	}

	var wire1, wire2 uint16
	var err error

	if o.w1 != nil {
		wire1, err = o.w1.getValue()
		if err != nil {
			return 0, err
		}
	} else if o.p1 != "" {
		v, _ := strconv.ParseUint(o.p1, 10, 16)
		wire1 = uint16(v)
	}
	if o.w2 != nil {
		wire2, err = o.w2.getValue()
		if err != nil {
			return 0, err
		}
	} else if o.p2 != "" {
		v, _ := strconv.ParseUint(o.p2, 10, 16)
		wire2 = uint16(v)
	}

	var result uint16

	switch o.op {
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

	o.value = result
	o.valid = true

	if !o.valid {
		err = fmt.Errorf("value not found")
		return 0, err
	}

	return o.value, nil
}

type circuit map[string]*operation

func (c circuit) resolveWires() error {
	for k, v := range c {
		if v.p1 != "" {
			if !checkUint16(v.p1) {
				l1, ok := c[v.p1]
				if !ok {
					return fmt.Errorf("invalid operand %q for %q", v.p1, k)
				}
				v.w1 = l1
			}
		}
		if v.p2 != "" {
			if !checkUint16(v.p2) {
				l2, ok := c[v.p2]
				if !ok {
					return fmt.Errorf("invalid operand %q for %q", v.p2, k)
				}
				v.w2 = l2
			}
		}
	}

	return nil
}

var lineRegexp = regexp.MustCompile(
	"^(?:(?:([a-z]+)|([0-9]+)) )?(?:(NOT|OR|AND|LSHIFT|RSHIFT) )?(?:([a-z]+)|([0-9]+)) -> ([a-z]+)$")

func checkUint16(s string) bool {
	_, err := strconv.ParseUint(s, 10, 16)
	return err == nil
}

func parseLine(s string) (operation, string, error) {
	err := fmt.Errorf("invalid instruction %q", s)

	m := lineRegexp.FindStringSubmatch(s)
	if m == nil {
		return operation{}, "", err
	}

	op, ok := opLookup[m[3]]
	if !ok {
		return operation{}, "", err
	}

	var p1, p2 string
	switch op {
	case opSet, opNot:
		if m[1] != "" || m[2] != "" {
			return operation{}, "", err
		}
		if m[5] != "" {
			if !checkUint16(m[5]) {
				return operation{}, "", err
			}
			p1 = m[5]
		} else {
			p1 = m[4]
		}
	case opAnd, opOr:
		if m[1] == "" && m[2] == "" {
			return operation{}, "", err
		}
		if m[2] != "" {
			if !checkUint16(m[2]) {
				return operation{}, "", err
			}
			p1 = m[2]
		} else {
			p1 = m[1]
		}
		if m[5] != "" {
			if !checkUint16(m[5]) {
				return operation{}, "", err
			}
			p2 = m[5]
		} else {
			p2 = m[4]
		}
	case opLShift, opRShift:
		if m[1] == "" && m[2] == "" {
			return operation{}, "", err
		}
		if m[2] != "" {
			if !checkUint16(m[2]) {
				return operation{}, "", err
			}
			p1 = m[2]
		} else {
			p1 = m[1]
		}
		if m[5] == "" {
			return operation{}, "", err
		}
		if !checkUint16(m[5]) {
			return operation{}, "", err
		}
		p2 = m[5]
	}

	lineOp := operation{op: op, p1: p1, p2: p2}
	lineTo := m[6]

	return lineOp, lineTo, nil
}

func process(f io.Reader) (uint16, error) {
	c := make(circuit)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()

		op, to, err := parseLine(s)
		if err != nil {
			return 0, err
		}

		if _, ok := c[to]; ok {
			err := fmt.Errorf("wire %q already set", to)
			return 0, err
		}

		c[to] = &op
	}

	err := c.resolveWires()
	if err != nil {
		return 0, err
	}

	if _, ok := c["a"]; !ok {
		err := fmt.Errorf("wire \"a\" not found")
		return 0, err
	}

	value, err := c["a"].getValue()
	if err != nil {
		return 0, err
	}

	return value, nil
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

	value, err := process(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Printf("a: %d\n", value)

	return 0
}

func main() {
	os.Exit(run())
}
