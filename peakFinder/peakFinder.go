package main

// Linear approach O(n)
func peakFinderLinear(input []int) int {
	if len(input) == 0 {
		return -1
	} else if len(input) == 1 {
		return 0
	} else {
		if input[0] >= input[1] {
			return 0
		} else if input[len(input)-1] >= input[len(input)-2] {
			return len(input) - 1
		} else {
			for i := 1; i < len(input)-1; i++ {
				if input[i-1] <= input[i] && input[i] >= input[i+1] {
					return i
				}
			}
		}
	}
	return -1
}

// Divide and conquer approach O(log2n)
func peakFinderDAC(input []int, min int, max int) int {
	if len(input) == 0 {
		return -1
	}
	if min-max == -1 || min-max == 1 {
		if input[min] >= input[max] {
			return min
		}
		return max
	} else {
		left := input[min : min+(max-min)/2]
		right := input[min+(max-min)/2 : max]

		if left[len(left)-1] >= right[0] {
			return peakFinderDAC(input, min, min+(max-min)/2)
		} else {
			return peakFinderDAC(input, min+(max-min)/2, max)
		}
	}
}
