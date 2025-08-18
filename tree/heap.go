package tree

type Leaf struct {
	char  string
	freq  int
	left  *Leaf
	right *Leaf
}

type PriorityQueue []*Leaf

func (pg PriorityQueue) Len() int { return len(pg) }

func (pg PriorityQueue) Less(i, j int) bool {
	return pg[i].freq < pg[j].freq
}

func (pg PriorityQueue) Swap(i, j int) {
	pg[i], pg[j] = pg[j], pg[i]
}

func (pg *PriorityQueue) Push(x interface{}) {
	item := x.(*Leaf)
	*pg = append(*pg, item)
}

func (pg *PriorityQueue) Pop() interface{} {
	old := *pg
	n := len(old)
	item := old[n-1]
	*pg = old[0 : n-1]
	return item
}
