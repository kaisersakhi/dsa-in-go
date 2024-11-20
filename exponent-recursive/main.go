package main

import "fmt"

func main(){
	n := 2
	k := 4

	fmt.Printf("2 to the power 4 is : %d", expo(n, k))
}

func expo(n, k int) int {
	if k == 0 {
		return 1
	}
	return n * expo(n, k - 1)
}