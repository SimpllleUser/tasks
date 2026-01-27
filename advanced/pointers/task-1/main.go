package main

import (
	"fmt"
)

// type TreeNode struct {
// Value int
// Left  *TreeNode
// Right *TreeNode
// }

// Реалізуй: Insert(value int), Search(value int) bool, InOrderTraversal() []int

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) Insert(value int) bool {
	if t == nil {
		return false
	}

	if t.Value == value {
		return false
	}

	if t.Value > value {
		if t.Left == nil {
			t.Left = &TreeNode{Value: value}
			return true
		}
		return t.Left.Insert(value)
	}

	if t.Right == nil {
		t.Right = &TreeNode{Value: value}
		return true
	}
	return t.Right.Insert(value)
}

func (t *TreeNode) InOrderTraversal() []int {
	if t == nil {
		return nil
	}

	result := []int{}

	if t.Left != nil {
		result = append(result, t.Left.InOrderTraversal()...)
	}

	result = append(result, t.Value)

	if t.Right != nil {
		result = append(result, t.Right.InOrderTraversal()...)
	}

	return result
}

func (t *TreeNode) Search(value int) bool {
	if t == nil {
		return false
	}

	if t.Value == value {
		return true
	}

	if value < t.Value {
		return t.Left != nil && t.Left.Search(value)
	}

	return t.Right != nil && t.Right.Search(value)

}

func main() {

	treeNode := TreeNode{
		Value: 1,
	}

	treeNode.Insert(2)
	treeNode.Insert(-1)
	treeNode.Insert(3)
	treeNode.Insert(4)

	fmt.Printf("Result %v", treeNode.InOrderTraversal())
	fmt.Println("Result ", treeNode.Search(5))
	fmt.Println("Result ", treeNode.Search(2))
}
