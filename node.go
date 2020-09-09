package main

type node struct {
	Left  *node `json:"left,omitempty"`
	Right *node `json:"right,omitempty"`
	C     rune  `json:"c,omitempty"`
	Tot   int   `json:"-"`
}

func (n *node) isLeaf() bool {
	if n.Left != nil && n.Right != nil {
		return false
	}
	if n.Left != nil || n.Right != nil {
		panic("not full binary tree")
	}
	return true
}

type nodeHeap []*node

func (n nodeHeap) Len() int {
	return len(n)
}

func (n nodeHeap) Less(i, j int) bool {
	if n[i].Tot != n[j].Tot {
		return n[i].Tot > n[j].Tot
	}
	return n[i].C > n[j].C
}

func (n nodeHeap) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n *nodeHeap) Push(x interface{}) {
	*n = append(*n, x.(*node))
}

func (n *nodeHeap) Pop() interface{} {
	old := *n
	sz := len(old)
	x := old[sz-1]
	*n = old[0 : sz-1]
	return x
}
