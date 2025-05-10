package linkedlist

import "fmt"

type LinkedList struct {
	Data int
	Next *LinkedList
}

func NewLinkedList(data int) *LinkedList {
	return &LinkedList{Data: data}
}

func (ll *LinkedList) Insert(n int, newData *LinkedList) *LinkedList {
	if ll.Length() < n {
		fmt.Println("LinkedList is not long enough.")
		return ll
	}
	count := 1
	for ll != nil {
		if count == n-1 {
			temp := ll.Next
			ll.Next = newData
			ll = ll.Next
			ll.Next = temp
			return ll
		}
		count++
		ll = ll.Next
	}
	return ll
}

func (ll *LinkedList) Length() int {
	count := 0
	for ll != nil {
		count++
		ll = ll.Next
	}
	return count
}

func (ll *LinkedList) Push(data int) *LinkedList {
	newNode := &LinkedList{Data: data}
	for ll != nil {
		if ll.Next == nil {
			ll.Next = newNode
			return ll
		}
		ll = ll.Next
	}
	return ll
}

func (ll *LinkedList) Pop() (int, *LinkedList) {
	res := 0
	for ll != nil {
		if ll.Next == nil {
			res = ll.Data
			ll = nil
			return res, ll
		}
	}
	return res, ll
}

func (ll *LinkedList) Max() int {
	res := ll.Data
	for ll != nil {
		if res < ll.Data {
			res = ll.Data
		}
		ll = ll.Next
	}
	return res
}
