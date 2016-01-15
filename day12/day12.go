package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func parse(s string) (interface{}, error) {
	var d interface{}
	err := json.Unmarshal([]byte(s), &d)
	if err != nil {
		return 0, err
	}

	return d, err
}

func walkVarFilter(d interface{}, doSkip bool, skip string) ([]int, error) {
	r := make([]int, 0)

	switch d.(type) {
	case float64:
		r = append(r, int(d.(float64)))
	case []interface{}:
		for _, v := range d.([]interface{}) {
			ir, err := walkVarFilter(v, doSkip, skip)
			if err != nil {
				return nil, err
			}
			r = append(r, ir...)
		}
	case map[string]interface{}:
		var cr []int
		var skipCur bool
		for _, v := range d.(map[string]interface{}) {
			if doSkip && v == skip {
				skipCur = true
				break
			}
			ir, err := walkVarFilter(v, doSkip, skip)
			if err != nil {
				return nil, err
			}
			cr = append(cr, ir...)
		}
		if !skipCur {
			r = append(r, cr...)
		}
	case string:
		// ignore
	default:
		err := fmt.Errorf("unhandled type %T", d)
		return nil, err
	}

	return r, nil
}

func walkVar(d interface{}) ([]int, error) {
	v, err := walkVarFilter(d, false, "")
	if err != nil {
		return nil, err
	}

	return v, nil
}

func walkVarSkip(d interface{}, skip string) ([]int, error) {
	v, err := walkVarFilter(d, true, skip)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func process(s string) (int, int, error) {
	d, err := parse(s)
	if err != nil {
		return 0, 0, err
	}

	r, err := walkVar(d)
	if err != nil {
		return 0, 0, err
	}

	var sum int
	for _, v := range r {
		sum += v
	}

	r2, err := walkVarSkip(d, "red")
	if err != nil {
		return 0, 0, err
	}

	var sum2 int
	for _, v := range r2 {
		sum2 += v
	}

	return sum, sum2, nil
}

func run() int {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "%s filename\n", os.Args[0])
		return 1
	}

	b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	s := strings.TrimSpace(string(b))

	v, v2, err := process(s)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Printf("result: %d\nresult2: %d\n", v, v2)

	return 0
}

func main() {
	os.Exit(run())
}
