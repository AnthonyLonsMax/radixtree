package radixtree

// Size returns the number of words stored in the tree.
func (r *MapRadixTree) Size() int64 {
	return r.size
}
