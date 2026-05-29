package radixmap

// ForEach calls callbackFn for each word in the tree.
func (r *MapRadixTree) ForEach(callbackFn func(key string)) {
	if r.root == nil {
		return
	}

	buf := make([]byte, 0, initialBufferSize)

	walk(r.root, &buf, callbackFn)
}
