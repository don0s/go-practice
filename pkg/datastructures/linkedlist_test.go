package datastructures

import (
	"testing"
)

func TestCreateSingleItemValidHead(t *testing.T) {

	var want int32 = 5
	list := CreateLinkedList(want)

	if list.Head.Value != want {
		t.Errorf("1 = %d; want 1", list.Head.Value)
	}

}

func TestCreateMultiItemValidHead(t *testing.T) {

	want := []string{"a", "b", "c"}
	list := CreateLinkedList(want...)

	if list.Head.Value != want[0] {
		t.Errorf("Received %v; Wanted %v", list.Head.Value, want[0])
	}

}

func TestCreateMultiItemValidTail(t *testing.T) {

	want := []int32{1, 2, 3}
	list := CreateLinkedList(want...)

	if list.Tail.Value != want[len(want)-1] {
		t.Errorf("1 = %d; want 1", list.Tail.Value)
	}

}

func TestCreateSingleItemValidTail(t *testing.T) {

	var want int32 = 5
	list := CreateLinkedList(want)

	if list.Tail.Value != want {
		t.Errorf("1 = %d; want 1", list.Tail.Value)
	}

}
