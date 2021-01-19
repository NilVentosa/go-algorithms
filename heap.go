package main

import (
	"fmt"
    "os"
)

func main() {
    input := os.Args[1:]

    result := permutate(input)

    for _, s := range result {
        fmt.Println(s)
    }
}

// Implementation of Heap's algorithm to get all the permutations of
// a slice
func permutate(input []string) [][]string {
	var helper func(int, []string)
	result := [][]string{}


	helper = func(k int, input []string) {
		if k == 1 {
			part := make([]string, len(input))
			copy(part, input)
			result = append(result, part)
		} else {
			helper(k-1, input)
			for i := 0; i < k-1; i++ {
				if k % 2 == 1 {
					input[0], input[k-1] = input[k-1], input[0]
				} else {
					input[i], input[k-1] = input[k-1], input[i]
				}
				helper(k-1, input)
			}
		}
	}
	helper(len(input), input)
	return result
}
