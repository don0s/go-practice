package datastructures

import "fmt"

type LinkedList[T any] struct {
	Head *ListNode[T]
	Tail *ListNode[T]
}

type ListNode[T any] struct {
	Value T
	Next  *ListNode[T]
}

func CreateLinkedList[T any](i ...T) *LinkedList[T] {

	var list LinkedList[T]
	for _, v := range i {
		AddToTail(&list, v)
	}

	return &list
}

func AddToHead[T any](l *LinkedList[T], v T) {

	next := ListNode[T]{Value: v}
	if l.Head == nil {
		l.Head = &next
		l.Tail = &next
	}

	next.Next = l.Head
	l.Head = &next

}

func AddToTail[T any](l *LinkedList[T], v T) {

	next := ListNode[T]{Value: v}
	if l.Head == nil {
		l.Head = &next
		l.Tail = &next
	} else {
		l.Tail.Next = &next
		l.Tail = &next
	}

}

func (l *LinkedList[T]) String() string {
	copy := l.Head

	response := fmt.Sprintf("[value: %v]", copy.Value)
	for copy.Next != nil {
		copy = copy.Next
		response = response + fmt.Sprintf(" [value: %v]", copy.Value)
	}

	return response

}
