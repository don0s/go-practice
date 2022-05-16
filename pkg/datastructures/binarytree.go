package datastructures

import (
	"fmt"
)

type islessthan interface {
	IsLessThan(islessthan) bool
	GetValue() string
}

type BinaryTree[T islessthan] struct {
	Root *TreeNode[T]
}

type TreeNode[T islessthan] struct {
	Value T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func InOrderTraverse[T islessthan](b *BinaryTree[T]) []T {
	if b.Root != nil {
		return inOrderTraverse(b.Root)
	}
	return nil
}

func String[T islessthan](node *TreeNode[T]) string {
	return fmt.Sprintf("[%v]", node.Value)
}

func inOrderTraverse[T islessthan](node *TreeNode[T]) []T {

	var leftList []T
	var rightList []T

	if node.Left != nil {
		leftList = inOrderTraverse(node.Left)
	}

	//var midList []T

	midList := append(leftList, node.Value)

	if node.Right != nil {
		rightList = inOrderTraverse(node.Right)
	}

	return append(midList, rightList...)

}

func AddSorted[T islessthan](b *BinaryTree[T], v T) {

	if b.Root == nil {
		b.Root = &TreeNode[T]{Value: v}
	} else {
		addSorted(b.Root, v)
	}
}

func addSorted[T islessthan](node *TreeNode[T], v T) {
	if v.IsLessThan(node.Value) {

		if node.Left == nil {
			node.Left = &TreeNode[T]{Value: v}
		} else {
			addSorted(node.Left, v)
		}

	} else {

		if node.Right == nil {
			node.Right = &TreeNode[T]{Value: v}
		} else {
			addSorted(node.Right, v)
		}

	}
}
