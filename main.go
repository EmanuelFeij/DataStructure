package main

import (
	"fmt"

	"github.com/EmanuelFeij/DataStructures/BinarySearchTree"
)

func main() {
	var t *BinarySearchTree.TreeNode

	t = t.Insert(20)
	t = t.Insert(31)
	t = t.Insert(32)
	t = t.Insert(30)
	t = t.Insert(4)
	t = t.Insert(5)
	t = t.Insert(3)
	t.Print()
	t2 := t.Invert()
	fmt.Println()
	t2.Print()

}
