package main

import "fmt"

func main(){
	n := 7

	fmt.Printf("The sume of n natural numbers upt0 7 is: %d", sum(n))
}

func sum(n int) int {
	if n == 0 {
		return 0
	} else {
		return n + sum(n - 1)
	}
}