package unionfind

import "fmt"

type quickUnion struct {
	id []int
}

// NewQuickUnion is the constructor class to create a new struct containing id[i]=i
func NewQuickUnion(n int) *quickUnion {
	qu := &quickUnion{}
	qu.id = make([]int, n)
	for i := 0; i < n; i++ {
		qu.id[i] = i
	}
	return qu
}

// find starts at a given node and follows the links up the tree to reach the root
// find needs to access the array one plus twice the depth of the node at the corresponding site
func (qu *quickUnion) Find(i int) int {
	for i != qu.id[i] {
		i = qu.id[i]
	}
	return i
}

// connected returns true if the two []id values are equal, which means that both nodes have the same root
// connected has two times the array access of find
func (qu *quickUnion) Connected(p, q int) bool {
	return qu.Find(p) == qu.Find(q)
}

// union combines two trees by making the root of one, the parent of the other by changing only one value
// union has two times the array access of find plus one if the sites are in different trees
func (qu *quickUnion) Union(p, q int) {
	i := qu.Find(p)
	j := qu.Find(q)
	if i != j {
		qu.id[i] = j
	}
}

func (qu *quickUnion) PrintIds() {
	fmt.Printf("%v\n", qu.id)
}
