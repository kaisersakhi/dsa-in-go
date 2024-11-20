package main

import "fmt"

func main(){
	n := 5
	fmt.Printf("The factorial of 5 is: %d", fact(n))
}

func fact(n int) int{
	if n == 0 {
		return 1
	}

	return n * fact(n - 1)
}