package main

import "testing"

func TestPeakFinder2d(t *testing.T) {
	var tests = []struct {
		input     [][]int
		expectedX int
		expectedY int
	}{
		{[][]int{
			{0, 1, 0},
			{1, 9, 1},
			{0, 1, 0}}, 1, 1},
		{[][]int{
			{4, 5, 6, 5},
			{5, 6, 7, 6},
			{6, 7, 8, 7},
			{7, 8, 9, 8}}, 2, 3},
		{[][]int{
			{0},
			{1},
			{0}}, 0, 1},
		{[][]int{
			{1, 0, 0},
			{9, 1, 0},
			{1, 0, 0}}, 0, 1},
		{[][]int{
			{9, 1, 0},
			{1, 0, 0},
			{0, 0, 0}}, 0, 0},
		{[][]int{
			{0, 1, 0}}, 1, 0},
	}
	for _, test := range tests {
		resultX, resultY := peakFinder2d(test.input)
		if resultX != test.expectedX || resultY != test.expectedY {
			t.Errorf("TestPeakFinder failed: input: %v, expected: (%v,%v), received: (%v,%v)", test.input, test.expectedX, test.expectedY, resultX, resultY)
		}
	}
}
