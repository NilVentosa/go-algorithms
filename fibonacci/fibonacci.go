package main

import (
	"fmt"
	"os"
	"strconv"
)

var instructions = "First argument should be a number and second memo for the memoized algorithm bad for the non memoized and bot for bottom-up"

func main() {
	arg, err := strconv.Atoi(os.Args[1])
	if len(os.Args) != 3 || err != nil {
		fmt.Println(instructions)
		os.Exit(1)
	}
	if os.Args[2] == "memo" {
		fmt.Println("Memoized version:", goodFibonacci(arg))
	} else if os.Args[2] == "bot" {
		fmt.Println("Bottom-up version:", bottomUpFibonacci(arg))
	} else if os.Args[2] == "bad" {
		fmt.Println("Bad version:", badFibonacci(arg))
	} else {
		fmt.Println(instructions)
		os.Exit(1)
	}
}

// Basic recurrent implementation of algorith to get the nth Fibonacci number
// This runs at exponential time
func badFibonacci(n int) int {
	if n <= 2 {
		return 1
	}
	return badFibonacci(n-1) + badFibonacci(n-2)
}

// Memoized version of the algorith runs at linear time
func goodFibonacci(n int) int {
	var m map[int]int
	m = make(map[int]int)
	var helper func(int) int

	helper = func(n int) int {
		if _, ok := m[n]; ok {
			return m[n]
		}
		if n <= 2 {
			return 1
		}
		m[n] = helper(n-1) + helper(n-2)
		return m[n]
	}
	return helper(n)
}

// Bottom-up version
func bottomUpFibonacci(n int) int {
	var m map[int]int
	m = make(map[int]int)

	for i := 1; i <= n; i++ {
		if i <= 2 {
			m[i] = 1
		} else {
			m[i] = m[i-1] + m[i-2]
		}
	}
	return m[n]
}
