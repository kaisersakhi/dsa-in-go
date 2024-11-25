package linked_list

import "errors"

type node struct {
	previous *node
	data int
	next *node
}

type LinkedList struct {
	head *node
	tail *node
	size int
}

func newNode(data int) *node {
	return &node{
		previous: nil,
		data: data,
		next: nil,
	}
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		head: nil,
		size: 0,
	}
}

// Append add new node to the list given data, inserts in O(1)
func (ll *LinkedList) Append(data int) {

	newNode := newNode(data)

	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
	} else {
		ll.tail.next = newNode
		ll.tail = newNode
	}

	ll.size++
}

func (ll *LinkedList) Pop() (int, error) {
	if ll.size == 0 {
		return 0, errors.New("can't pop from empty list")
	}

	data := ll.tail.data

	ll.tail = ll.tail.previous
	ll.tail.next = nil

	ll.size--

	return data, nil
}