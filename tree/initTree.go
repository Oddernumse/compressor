package tree

import "container/heap"

func BuildTree(freqMap map[rune]int) *Leaf {
	pg := make(PriorityQueue, 0)
	heap.Init(&pg)

	for char, freq := range freqMap {
		heap.Push(&pg, &Leaf{Char: char, Freq: freq})
	}

	for pg.Len() > 1 {
		left := heap.Pop(&pg).(*Leaf)
		right := heap.Pop(&pg).(*Leaf)

		combined := &Leaf{
			Freq:  left.Freq + right.Freq,
			Left:  left,
			Right: right,
		}

		heap.Push(&pg, combined)
	}

	return heap.Pop(&pg).(*Leaf)
}
