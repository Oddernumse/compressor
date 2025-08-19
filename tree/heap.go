package tree

type Leaf struct {
	Char  string
	Freq  int
	Left  *Leaf
	Right *Leaf
	Code  int
}

type BST struct {
	Root *Leaf
}

type PriorityQueue []*Leaf

func (pg PriorityQueue) Len() int { return len(pg) }

func (pg PriorityQueue) Less(i, j int) bool {
	return pg[i].Freq < pg[j].Freq
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
