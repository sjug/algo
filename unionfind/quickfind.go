package unionfind

type quickFind struct {
	id []int
}

// NewQuickFind is the constructor class to create a new struct containing id[i]=i
func NewQuickFind(n int) *quickFind {
	qf := &quickFind{}
	qf.id = make([]int, n)
	for i := 0; i < n; i++ {
		qf.id[i] = i
	}
	return qf
}

// find is not necessary to have in its own function, just for sake of comparison and learning
// find is quick as it only accesses the []id slice once for each call
func (qf *quickFind) find(i int) int {
	return qf.id[i]
}

// connected returns true if the two []id values are in the same component, as they have the same value
// connected uses two array accesses for each call (one for each find)
func (qf *quickFind) connected(p, q int) bool {
	return qf.find(p) == qf.find(q)
}

// union sets the id-value of one component to that of another to join the two components
// union needs to scan through the whole array for each call which causes issues with a large number of components
func (qf *quickFind) union(p, q int) {
	pid := qf.id[p]
	qid := qf.id[q]
	if pid != qid {
		for i := 0; i < len(qf.id); i++ {
			if qf.id[i] == pid {
				qf.id[i] = qid
			}
		}
	}
}
