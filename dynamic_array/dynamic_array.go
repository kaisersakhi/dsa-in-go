package dynamic_array

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
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

func (d *DynamicArray) Get(index int) (int, error) {
	if index < 0 || index >= len(d.items) {
		return 0, errors.New("index out of bound")
	}

	return d.items[index], nil
}

func (d *DynamicArray) Sum() int {
	var sum int

	for _, val := range d.items {
		sum += val
	}

	return sum
}

func (d *DynamicArray) Average() float64{
	return float64(d.Sum()) / float64(len(d.items))
}

func (d *DynamicArray) Max() int {
	if len(d.items) < 1 {
		return 0
	}

	var maxx = d.items[0]

	d.ForEach(func(_, val int) {
		if val > maxx {
			maxx = val
		}
	})

	return maxx
}

func (d *DynamicArray) Min() int {
	if len(d.items) < 1 {
		return 0
	}

	var minn = d.items[0]

	d.ForEach(func(_, val int) {
		if val < minn {
			minn = val
		}
	})

	return minn
}

func (d *DynamicArray) ForEach(callback func(int, int)) {
	for index, val := range d.items {
		callback(index, val)
	}
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

func (d *DynamicArray) BinarySearch(item int) (int, error) {
	if !d.isSorted {
		return 0, errors.New("items are not sorted, hard sort the items list first")
	}

	var newMid = func (l, r int) int {
		return int(math.Floor(float64(l + r)/2))
	}

	var mLen = len(d.items)
	var left = 0
	var right = mLen
	var mid = newMid(left, right)

	for mid >= 0 && mid < mLen {
		if d.items[mid] == item {
			return mid, nil
		} else if d.items[mid] > item {
			// Go left.
			right = mid
		} else if d.items[mid] < item {
			// Go right.
			left = mid
		}

		mid = newMid(left, right)
	}

	return -1, errors.New("item not found")
}

func (d *DynamicArray) Reverse() []int {
	itemsDup := make([]int, len(d.items))
	copy(itemsDup, d.items)

	var left = 0
	var right = len(itemsDup) - 1

	for left < right {
		temp := itemsDup[left]
		itemsDup[left] = itemsDup[right]
		itemsDup[right] = temp

		left++
		right--
	}

	return itemsDup
}

func (d *DynamicArray) ReverseHard() []int {
	d.items = d.Reverse()
	d.isSorted = false

	return d.items
}

func (d *DynamicArray) Shuffle() []int {
	itemsDup := make([]int, len(d.items))
	copy(itemsDup, d.items)

	var mLen = len(itemsDup)

	for i := 0; i < mLen; i++ {
		randomIndex := rand.Intn(mLen)
		temp := itemsDup[0]
		itemsDup[0] = itemsDup[randomIndex]
		itemsDup[randomIndex] = temp
	}

	return itemsDup
}

func (d *DynamicArray) ShuffleHard() []int {
	d.items = d.Shuffle()
	d.isSorted = false

	return d.items
}

func (d *DynamicArray) MergeHard(od *DynamicArray) {
	combinedItems := append(d.items, od.items...)
	d.items = combinedItems
	d.isSorted = false
}

// SetifyHard Removes duplicates.
func (d *DynamicArray) SetifyHard() {
	
}