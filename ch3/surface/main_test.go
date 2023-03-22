package main

import (
	"fmt"
	"testing"
)

func TestIsDataValid(t *testing.T) {
	var z float64
	cases := []struct {
		input  []float64
		expect bool
	}{
		{[]float64{0.1}, true},
		{[]float64{0.0}, true},
		{[]float64{0.0, 1 / z}, false},
		{[]float64{0.0, -1 / z}, false},
		{[]float64{0.0, z / z}, false},
	}

	for _, c := range cases {
		result := isDataValid(c.input...)
		if result != c.expect {
			t.Fatalf("Result of [%v] is %v, expect %v\n", c.input, result, c.expect)
		}
		// t.Logf("Result of [%v] is %v, expect %v\n", c.input, result, c.expect)
	}
}

func TestGetColor(t *testing.T) {
	fmt.Printf("color: %v from %v\n", getColor(0.9), 0.9)
}
