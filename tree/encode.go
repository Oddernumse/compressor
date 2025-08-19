package tree

import "fmt"

func (bst *BST) BuildCode(nodes *Leaf) {
	if nodes == nil {
		return
	} else {
		bst.BuildCode(nodes.Left)
		fmt.Println(nodes.Char, " ", nodes.Freq)
		bst.BuildCode(nodes.Right)

	}
}
