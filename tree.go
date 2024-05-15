package tree

/*
All keys must be ordered somehow
*/
type Comparer[K any] interface {
	/*
		returns -1 if y is greater, 0 if they're equal, and 1 if y is less
	*/
	Compare(y K) int
}

/*
two different types of nodes: interior and leaves.
Use this node interface to group their similarities
*/

type node[K Comparer[K], V any] interface {
	/*
		inserts the value v into the subtree rooted at this node
	*/
	insert(key K, val V) node[K, V]
	/*
		searches for the key and returns a pointer to V on success or an error if it's not found
	*/
	search(key K) (error, *V)

	height() int

	rangeSearch(min, max K) (*leaf[K, V], *leaf[K, V])
}

type interior[K Comparer[K], V any] struct {
	depth    int
	keys     []K          // size = order-1
	children []node[K, V] // size = order
	order    uint
	occupied uint
}

func (n interior[K, V]) insert(key K, val V) node[K, V] {
	var i uint
	for i = 0; key.Compare(n.keys[i]) >= 0 && i < n.occupied; {
		i += 1
	}
	// the new key is larger than all the current keys
	n.children[i] = n.children[i].insert(key, val)
	return n
}

func (n interior[K, V]) search(key K) (error, *V) {
	return nil, nil
}

func (n interior[K, V]) height() int {
	return 0
}

func (n interior[K, V]) rangeSearch(min, max K) (*leaf[K, V], *leaf[K, V]) {
	return nil, nil
}

type leaf[K Comparer[K], V any] struct {
	keys   []K
	bucket []V
	next   *leaf[K, V]
	order  uint
}
