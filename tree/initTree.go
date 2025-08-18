package tree

import "container/heap"

func BuildTree(freqMap map[string]int) *Leaf {
	pg := make(PriorityQueue, 0)
	heap.Init(&pg)

	for char, freq := range freqMap {
		heap.Push(&pg, &Leaf{char: char, freq: freq})
	}

	for pg.Len() > 1 {
		left := heap.Pop(&pg).(*Leaf)
		right := heap.Pop(&pg).(*Leaf)

		combined := &Leaf{
			freq:  left.freq + right.freq,
			left:  left,
			right: right,
		}

		heap.Push(&pg, combined)
	}

	return heap.Pop(&pg).(*Leaf)
}
