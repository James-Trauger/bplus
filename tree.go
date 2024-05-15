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
	insert(val V) *node[K, V]
	/*
		searches for the key and returns a pointer to V on success or an error if it's not found
	*/
	search(key K) (error, *V)
}
