package BinarySearchTree

import "fmt"

type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

func (tree *TreeNode) Insert(data int) *TreeNode {

	if tree == nil {
		return &TreeNode{data: data, left: nil, right: nil}

	} else {
		if data > tree.data {
			tree.right = tree.right.Insert(data)
		}
		if data < tree.data {
			tree.left = tree.left.Insert(data)
		}
	}
	return tree

}

func (tree *TreeNode) Invert() *TreeNode {
	if tree != nil {
		aux := tree.left
		tree.left = tree.right
		tree.right = aux
	}
	return tree
}

func (tree *TreeNode) Print() {
	if tree != nil {
		fmt.Printf("%v ", tree.data)
		tree.left.Print()
		tree.right.Print()
	}
}
