package main

// I think this works. Finds a 2d peak with O(n log m)
func peakFinder2d(input [][]int) (xResult int, yResult int) {
	var helper func([][]int, int, int) (int, int)
	xResult = 0
	yResult = 0

	helper = func(input [][]int, min int, max int) (int, int) {
		if len(input[min:max+1]) == 1 {
			maxValue := input[min : max+1][0][0]
			for x, column := range input[min : max+1][0] {
				if column > maxValue {
					maxValue = column
					xResult = x
				}
			}
			return xResult, min
		} else if len(input[0]) == 1 || len(input[min:max+1]) == 1 {
			maxValue := input[min : max+1][0][0]
			for y, row := range input[min : max+1] {
				if row[0] > maxValue {
					maxValue = row[0]
					yResult = y
				}
			}
			return xResult, yResult
		} else {
			middleIndex := (min + max) / 2
			middleRow := input[middleIndex]
			maxValue := middleRow[0]
			for x, column := range middleRow {
				if column > maxValue {
					maxValue = column
					xResult = x
				}
			}
			if input[middleIndex-1][xResult] > input[middleIndex][xResult] {
				return helper(input, min, middleIndex-1)
			} else if input[middleIndex+1][xResult] > input[middleIndex][xResult] {
				return helper(input, middleIndex+1, max)
			} else {
				return xResult, middleIndex
			}
		}
	}
	return helper(input, 0, len(input)-1)
}
