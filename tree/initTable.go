package tree

func BuildCode(nodes *Leaf, currentCode string, codeMap map[string]string) {
	if nodes == nil {
		return
	}

	if nodes.Char != 0 {
		nodes.Code = currentCode
	}

	BuildCode(nodes.Left, currentCode+"0", codeMap)
	codeMap[string(nodes.Char)] = currentCode
	BuildCode(nodes.Right, currentCode+"1", codeMap)
	codeMap[string(nodes.Char)] = currentCode
}

func PrintTree(nodes *Leaf) {
	if nodes != nil {
		println(nodes.Char, " ", nodes.Code)
		PrintTree(nodes.Right)
		PrintTree(nodes.Left)
	}
}
