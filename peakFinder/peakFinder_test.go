package main

import "testing"

func TestPeakFinderLinear(t *testing.T) {
	var tests = []struct {
		input    []int
		min      int
		max      int
		expected int
	}{
		{[]int{}, 0, 0, -1},
		{[]int{0, 0, 0}, 0, 2, 0},
		{[]int{1, 1, 1}, 0, 2, 0},
		{[]int{1, 2, 3}, 0, 2, 2},
		{[]int{1, 5, 1}, 0, 2, 1},
		{[]int{10, 5, 1}, 0, 2, 0},
		{[]int{0, 0, 0, 0, 1, 2, 3, 4, 5}, 0, 8, 0},
		{[]int{0, 5, 9, 0}, 0, 3, 2},
	}
	for _, test := range tests {
		result := peakFinderLinear(test.input)
		if result != test.expected {
			t.Errorf("TestPeakFinder failed: input: %v, expected: %v, received: %v", test.input, test.expected, result)
		}
	}
}

func TestPeakFinderDAC(t *testing.T) {
	var tests = []struct {
		input    []int
		min      int
		max      int
		expected int
	}{
		{[]int{}, 0, 0, -1},
		{[]int{0, 0, 0}, 0, 2, 0},
		{[]int{1, 1, 1}, 0, 2, 0},
		{[]int{1, 2, 3}, 0, 2, 2},
		{[]int{1, 5, 1}, 0, 2, 1},
		{[]int{10, 5, 1}, 0, 2, 0},
		{[]int{0, 0, 0, 0, 1, 2, 3, 4, 5}, 0, 8, 8},
	}
	for _, test := range tests {
		result := peakFinderDAC(test.input, test.min, test.max)
		if result != test.expected {
			t.Errorf("TestPeakFinder failed: input: %v, expected: %v, received: %v", test.input, test.expected, result)
		}
	}
}
