package main

import (
	"math"
	"testing"
)

func TestCalculateBMI(t *testing.T) {
	testCases := []struct {
		weight   float64
		height   float64
		expected float64
	}{
		{70, 1.75, 22.86}, 
		{50, 1.6, 19.53},  
		{90, 1.8, 27.78}, 
		{70, 0, 0.0},      
	}

	for _, tc := range testCases {
		result := CalculateBMI(tc.weight, tc.height)
		if !approxEqual(result, tc.expected, 0.01) {
			t.Errorf("CalculateBMI(%.2f, %.2f) = %.2f; expected %.2f", tc.weight, tc.height, result, tc.expected)
		}
	}
}


func approxEqual(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}
