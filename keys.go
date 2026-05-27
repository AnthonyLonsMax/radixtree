package radixtree

const initialBufferSize = 256

// Keys returns all stored words in the tree.
func (r *RadixTree) Keys() []string {
	result := make([]string, 0)

	buf := make([]byte, 0, initialBufferSize)

	r.keys(r.root, &buf, &result)

	return result
}

func (r *RadixTree) keys(cursor *node, buf *[]byte, words *[]string) {
	if cursor == nil {
		return
	}

	startLen := len(*buf)

	*buf = append(*buf, cursor.prefix...)

	if cursor.isTerminal {
		*words = append(*words, string(*buf))
	}

	for _, child := range cursor.children {
		r.keys(child, buf, words)
	}

	*buf = (*buf)[:startLen]
}
