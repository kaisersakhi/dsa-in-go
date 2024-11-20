package dynamic_array

import (
	"errors"
	"fmt"
)

type DynamicArray struct {
	items []int
	isSorted bool
}

func NewDynamicArray() *DynamicArray{
	return &DynamicArray{
		items: make([]int, 0),
		isSorted: false,
	}
}

func (d *DynamicArray) Append(item int) {
	d.items = append(d.items, item)
}

func (d *DynamicArray) PrintAll() {
	fmt.Print("[")
	mLen := len(d.items)

	for index, val := range d.items {
		fmt.Print(val)

		if index < (mLen - 1) {
			fmt.Print(", ")
		}
	}
	fmt.Print("]\n")
}

func (d *DynamicArray) Remove(index int) (int, error) {
	if index < 0 || index >= len(d.items) {
		return 0, errors.New("array index out of bounds")
	}

	item := d.items[index]

	d.items = append(d.items[:index], d.items[index + 1:]...)

	return item, nil
}


func (d *DynamicArray) BubbleSort() []int {
	itemDup := make([]int, len(d.items))
	copy(itemDup, d.items)

	for i := 0; i < len(itemDup); i++ {
		for j := 0; j < len(itemDup) - 1 - i; j++ {
			if itemDup[j + 1] < itemDup[j] {
				temp := itemDup[j]
				itemDup[j] = itemDup[j + 1]
				itemDup[j + 1] = temp
			}
		}
	}

	return itemDup
}

func (d *DynamicArray) BubbleSortHard() []int {
	if d.isSorted {
		return d.items
	}

	d.items = d.BubbleSort()
	d.isSorted = true

	return d.items
}

func (d *DynamicArray) LinearSearch(item int) (int, error) {
	for index, val := range d.items {
		if val == item {
			return index, nil
		}
	}

	return 0, errors.New("item not found")
}
