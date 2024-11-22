package main

import (
	"dsa-interview-prep/dynamic_array"
)

func main(){
	da := dynamic_array.NewDynamicArray()

//	da.Append(70)
//	da.Append(20)
//	da.Append(80)
//	da.Append(90)
//	da.Append(60)
//	da.Append(60)
//	da.Append(60)
//	da.Append(50)
//	da.Append(30)
//	da.Append(10)


//
//	da.PrintAll()
//
////	fmt.Printf("Removing 20\n")
////
////	item , _ := da.Remove(1)
////
////	fmt.Println("Removed : ", item)
////
////
////	index, _ := da.LinearSearch(80)
////
////	da.PrintAll()
////
////	fmt.Println("80 found at index : ", index)
////
////	fmt.Println("Sorted array is: ", da.BubbleSort())

//	_ = da.BubbleSortHard()
//
//	index, _ = da.BinarySearch(70)
//	fmt.Println("70 found at index : ", index)

//	da.ForEach(func(index, value int) {
//		fmt.Println(value)
//	})


//	fmt.Println("max is : ", da.Max())
//	fmt.Println("min is : ", da.Min())
//	fmt.Println("Reversed : ", da.Reverse())

//	da.BubbleSortHard()
//	da.ShuffleHard()

	da.Append(1)
	da.Append(2)
	da.Append(3)
	da.Append(4)

	d2 := dynamic_array.NewDynamicArray()
	d2.Append(4)
	d2.Append(5)
	d2.Append(6)

//	da.SetifyHard()
	da.UnionHard(d2)

	da.PrintAll()

	filtered := da.Filter(func(value int) bool {
		if value % 2 == 0 {
			return true
		}

		return false
	})

	filtered.PrintAll()
}

