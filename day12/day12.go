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

func walkVar(d interface{}) ([]int, error) {
	r := make([]int, 0)

	switch d.(type) {
	case float64:
		r = append(r, int(d.(float64)))
	case []interface{}:
		for _, v := range d.([]interface{}) {
			ir, err := walkVar(v)
			if err != nil {
				return nil, err
			}
			r = append(r, ir...)
		}
	case map[string]interface{}:
		for _, v := range d.(map[string]interface{}) {
			ir, err := walkVar(v)
			if err != nil {
				return nil, err
			}
			r = append(r, ir...)
		}
	case string:
		// ignore
	default:
		err := fmt.Errorf("unhandled type %T", d)
		return nil, err
	}

	return r, nil
}

func process(s string) (int, error) {
	d, err := parse(s)
	if err != nil {
		return 0, err
	}

	r, err := walkVar(d)
	if err != nil {
		return 0, err
	}

	var sum int
	for _, v := range r {
		sum += v
	}

	return sum, nil
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

	v, err := process(s)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Printf("result: %d\n", v)

	return 0
}

func main() {
	os.Exit(run())
}
