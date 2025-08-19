package tree

func BuildCode(nodes *Leaf, currentCode string) {
	if nodes == nil {
		return
	}

	if nodes.Char != 0 {
		nodes.Code = currentCode

		BuildCode(nodes.Left, nodes.Code+"0")
		BuildCode(nodes.Right, nodes.Code+"1")
	}

	println(nodes.Char, " ", nodes.Code)
}

func PrintTree(nodes *Leaf) {
	if nodes != nil {
		println(nodes.Char, " ", nodes.Code)
		PrintTree(nodes.Right)
		PrintTree(nodes.Left)
	}
}
