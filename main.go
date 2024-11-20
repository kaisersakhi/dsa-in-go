package main

import (
	"dsa-interview-prep/dynamic_array"
	"fmt"
)

func main(){
	da := dynamic_array.NewDynamicArray()

	da.Append(70)
	da.Append(20)
	da.Append(80)
	da.Append(90)
	da.Append(60)
	da.Append(50)
	da.Append(30)
	da.Append(10)



	da.PrintAll()

	fmt.Printf("Removing 20\n")

	item , _ := da.Remove(1)

	fmt.Println("Removed : ", item)


	index, _ := da.LinearSearch(80)

	da.PrintAll()

	fmt.Println("80 found at index : ", index)

	fmt.Println("Sorted array is: ", da.BubbleSort())

	_ = da.BubbleSortHard()

	index, _ = da.BinarySearch(70)
	fmt.Println("70 found at index : ", index)

}

