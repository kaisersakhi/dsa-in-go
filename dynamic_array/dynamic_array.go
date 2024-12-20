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

func (d *DynamicArray) Push(item int) {
	d.Append(item)
}

func (d *DynamicArray) Pop() (int, error) {
	if len(d.items) == 0 {
		return 0, errors.New("list is already empty, can't perfom pop operation")
	}

	item := d.items[len(d.items) - 1]

	d.items = d.items[:len(d.items) - 1]

	return item, nil
}

func (d *DynamicArray) Clone() *DynamicArray {
	clonedItems := make([]int, len(d.items))
	copy(clonedItems, d.items)

	return &DynamicArray{
		items: clonedItems,
		isSorted: d.isSorted,
	}
}

func (d *DynamicArray) Get(index int) (int, error) {
	if index < 0 || index >= len(d.items) {
		return 0, errors.New("index out of bound")
	}

	return d.items[index], nil
}

func (d *DynamicArray) IsEmpty() bool {
	if len(d.items) == 0 {
		return true
	}

	return false
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

func (d *DynamicArray) Map(tranformFunc func(int) int) *DynamicArray {
	clone := d.Clone()

	clone.ForEach(func(index, value int) {
		clone.items[index] = tranformFunc(value)
	})

	return clone
}

func (d *DynamicArray) Filter(predicate func(int) bool) *DynamicArray{
	var dupItems []int

	d.ForEach(func(_, value int){
		if predicate(value) {
			dupItems = append(dupItems, value)
		}
	})

	return &DynamicArray{
		items: dupItems,
		isSorted: d.isSorted,
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

func (d *DynamicArray) Contains(item int) bool {
	_, err := d.LinearSearch(item)

	if err != nil {
		return false
	}

	return true
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
func (d *DynamicArray) SetifyHard(){
	arrayMap := make(map[int]bool)
	var uniqItems []int

	for _, val := range d.items {
		if !arrayMap[val] {
			arrayMap[val] = true // Mark item as read.
			uniqItems = append(uniqItems, val)
		}
	}

	d.items = uniqItems
	d.isSorted = false
}

func (d *DynamicArray) Union(d1 *DynamicArray) *DynamicArray {
	d2 := NewDynamicArray()

	d2.MergeHard(d)
	d2.MergeHard(d1)
	d2.SetifyHard()

	return d2
}

func (d *DynamicArray) UnionHard(d1 *DynamicArray) {
	d.items = (d.Union(d1)).items
}

func (d *DynamicArray) Difference(d1 *DynamicArray) *DynamicArray{
	var diff []int

	for _, val := range d.items {
		if !elementIn(d1.items, val) {
			diff = append(diff, val)
		}
	}

	return &DynamicArray{
		items: diff,
		isSorted: false,
	}
}

func (d *DynamicArray) Intersect(d1 *DynamicArray) *DynamicArray {
	var common []int

	for _, val := range d.items {
		if elementIn(d1.items, val) {
			common = append(common, val)
		}
	}

	return &DynamicArray{
		items: common,
		isSorted: false,
	}
}

func elementIn(elements []int, item int) bool {
	for _, element := range elements {
		if element == item {
			return true
		}
	}

	return false
}