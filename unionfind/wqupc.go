package unionfind

type wQuickUnionPC struct {
	id []int
	sz []int
}

// NewWeightedQuickUnionPC is the constructor class to create a new struct containing id[i]=i and sz[i]=1
func NewWeightedQuickUnionPC(n int) *wQuickUnionPC {
	wqu := &wQuickUnionPC{}
	wqu.id = make([]int, n)
	wqu.sz = make([]int, n)
	for i := 0; i < n; i++ {
		wqu.id[i] = i
		wqu.sz[i] = 1
	}
	return wqu
}

// find starts at a given node and follows the links up the tree to reach the root
// find needs to access the array one plus twice the depth of the node at the corresponding site
// find modified with simple one-pass path compression variant to make every other node point to its grandparent
func (wqu *wQuickUnionPC) find(i int) int {
	for i != wqu.id[i] {
		wqu.id[i] = wqu.id[wqu.id[i]]
		i = wqu.id[i]
	}
	return i
}

// connected returns true if the two []id values are equal, which means that both nodes have the same root
// connected has two times the array access of find
func (wqu *wQuickUnionPC) connected(p, q int) bool {
	return wqu.find(p) == wqu.find(q)
}

// union in the weighted union version determines which tree is the smaller of the two then dynamically
// links the root of the smaller tree to the root of the larger tree, we also increment the size slice
func (wqu *wQuickUnionPC) union(p, q int) {
	i := wqu.find(p)
	j := wqu.find(q)
	if i != j {
		if wqu.sz[i] < wqu.sz[j] {
			wqu.id[i] = j
			wqu.sz[j] += wqu.sz[i]
		} else {
			wqu.id[j] = i
			wqu.sz[i] += wqu.sz[j]
		}
	}
}
