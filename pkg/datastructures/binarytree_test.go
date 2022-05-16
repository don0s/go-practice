package datastructures

import (
	"fmt"
	"reflect"
	"testing"
)

type item struct {
	value string
}

func TestInOrderTraverse(t *testing.T) {

	tree := BinaryTree[item]{}

	AddSorted(&tree, item{value: "E"})
	AddSorted(&tree, item{value: "C"})
	AddSorted(&tree, item{value: "A"})
	AddSorted(&tree, item{value: "D"})
	AddSorted(&tree, item{value: "B"})

	values := InOrderTraverse(&tree)
	fmt.Println(values)

	expected := []item{{value: "A"}, {value: "B"}, {value: "C"}, {value: "D"}, {value: "E"}}

	if !reflect.DeepEqual(expected, values) {
		t.Errorf("Checking Output, Received %v; Wanted %v", expected, values)
	}

}

func (n item) IsLessThan(v islessthan) bool {
	return n.value < v.GetValue()
}

func (n item) GetValue() string {
	return n.value
}

func (n item) String() string {
	return fmt.Sprintf("[%v]", n.value)
}
