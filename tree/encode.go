package tree

func BuildCode(nodes *Leaf, currentCode string) {
	if nodes == nil {
		return
	}

	if nodes.Char != 0 {
		nodes.Code = currentCode

	}

	BuildCode(nodes.Left, currentCode+"0")
	BuildCode(nodes.Right, currentCode+"1")
}

func PrintTree(nodes *Leaf) {
	if nodes != nil {
		println(nodes.Char, " ", nodes.Code)
		PrintTree(nodes.Right)
		PrintTree(nodes.Left)
	}
}
