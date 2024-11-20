package main

import "fmt"

func main() {
	memo := make([]int, 8)
	fmt.Println(fib(7, memo))
}

func fib(n int, memo []int) int {
	if n <= 1 {
		return n
	}

	if memo[n - 1] == 0 {
		memo[n - 1] = fib(n - 1, memo)
	}

	if memo[n - 2] == 0 {
		memo[n - 2] = fib(n - 2, memo)
	}

	return memo[n - 1] + memo[n - 2]
}