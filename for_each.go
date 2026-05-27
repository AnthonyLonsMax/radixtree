package radixtree

// ForEach calls callbackFn for each word in the tree.
func (r *RadixTree) ForEach(callbackFn func(key string)) {
	buf := make([]byte, 0, initialBufferSize)

	r.forEach(r.root, &buf, callbackFn)
}

func (r *RadixTree) forEach(cursor *node, buf *[]byte, callbackFn func(string)) {
	if cursor == nil {
		return
	}

	startLen := len(*buf)

	*buf = append(*buf, cursor.prefix...)

	if cursor.isTerminal {
		callbackFn(string(*buf))
	}

	for _, child := range cursor.children {
		r.forEach(child, buf, callbackFn)
	}

	*buf = (*buf)[:startLen]
}
