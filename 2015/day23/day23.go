package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

//go:generate stringer -type=opcodeType
type opcodeType int

const (
	opHlf opcodeType = iota
	opTpl
	opInc
	opJmp
	opJie
	opJio
)

var opLookup = map[string]opcodeType{
	"hlf": opHlf,
	"tpl": opTpl,
	"inc": opInc,
	"jmp": opJmp,
	"jie": opJie,
	"jio": opJio,
}

//go:generate stringer -type=registerType
type registerType int

const (
	regA registerType = iota
	regB
)

var regLookup = map[string]registerType{
	"a": regA,
	"b": regB,
}

type operation struct {
	op  opcodeType
	reg registerType
	off int
}

type machine struct {
	inst []operation
	regs []int
	pc   int
}

var lineRegex = regexp.MustCompile("^(hlf|tpl|inc|jmp|jie|jio) ([a-z])?(, )?([+-][0-9]+)?$")

func (o operation) String() string {
	return fmt.Sprintf("{%v %v %d}", o.op, o.reg, o.off)
}

func parseLine(s string) (operation, error) {
	errInvalid := fmt.Errorf("invalid operation %q", s)

	m := lineRegex.FindStringSubmatch(s)
	if m == nil {
		return operation{}, errInvalid
	}

	var ok bool

	var op operation

	op.op, ok = opLookup[m[1]]
	if !ok {
		err := fmt.Errorf("invalid opcode %q in %q", m[1], s)
		return operation{}, err
	}

	if op.op != opJmp {
		op.reg, ok = regLookup[m[2]]
		if !ok {
			err := fmt.Errorf("invalid register %q in %q", m[2], s)
			return operation{}, err
		}
	} else {
		if m[2] != "" {
			err := fmt.Errorf("register given for jmp %q in %q", m[2], s)
			return operation{}, err
		}
	}

	if op.op == opJmp || op.op == opJie || op.op == opJio {
		if op.op == opJmp {
			if m[3] != "" {
				err := fmt.Errorf("comma given for jmp %q in %q", m[3], s)
				return operation{}, err
			}
		} else {
			if m[3] != ", " {
				err := fmt.Errorf("comma not given for jmp %q in %q", m[3], s)
				return operation{}, err
			}
		}
		if m[4] == "" {
			err := fmt.Errorf("offset not given for jmp %q in %q", m[4], s)
			return operation{}, err
		}

		var err error
		op.off, err = strconv.Atoi(m[4])
		if err != nil {
			return operation{}, err
		}
	} else {
		if m[3] != "" {
			err := fmt.Errorf("comma given for no-jmp %q in %q", m[3], s)
			return operation{}, err
		}
		if m[4] != "" {
			err := fmt.Errorf("offset given for non-jmp %q in %q", m[4], s)
			return operation{}, err
		}
	}

	return op, nil
}

func (m *machine) run() {
	var nextPc int

	for {
		op := m.inst[m.pc]
		nextPc = m.pc + 1
		switch op.op {
		case opHlf:
			m.regs[op.reg] /= 2
		case opTpl:
			m.regs[op.reg] *= 3
		case opInc:
			m.regs[op.reg] += 1
		case opJmp:
			nextPc = m.pc + op.off
		case opJie:
			if m.regs[op.reg]%2 == 0 {
				nextPc = m.pc + op.off
			}
		case opJio:
			if m.regs[op.reg] == 1 {
				nextPc = m.pc + op.off
			}
		}
		//fmt.Println(m.pc, nextPc, m.inst[m.pc], m.regs)
		if nextPc < len(m.inst) {
			m.pc = nextPc
		} else {
			break
		}
	}
}

func process(f io.Reader) (int, int, error) {
	inst := []operation{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		o, err := parseLine(s)
		if err != nil {
			return 0, 0, nil
		}
		inst = append(inst, o)
	}

	machine1 := machine{inst, make([]int, 2), 0}
	machine1.run()

	regs := make([]int, 2)
	regs[regA] = 1
	machine2 := machine{inst, regs, 0}
	machine2.run()

	return machine1.regs[regB], machine2.regs[regB], nil
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

	rb, rb2, err := process(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Printf("register_b: %d\nregister_b2: %d\n", rb, rb2)
	return 0
}

func main() {
	os.Exit(run())
}
