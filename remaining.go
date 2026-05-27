package radixtree

// Remaining returns all stored words that have the given prefix.
func (r *RadixTree) Remaining(prefix string) []string {
	result := make([]string, 0)

	if r.root == nil {
		return result
	}

	buf := make([]byte, 0, initialBufferSize)

	if prefix == "" {
		r.keys(r.root, &buf, &result)

		return result
	}

	r.remaining(r.root, prefix, &buf, &result)

	return result
}

func (r *RadixTree) remaining(cursor *node, word string, buf *[]byte, words *[]string) {
	if cursor == nil {
		return
	}

	commonLen := commonPrefixLength(cursor.prefix, word)

	if commonLen == 0 && cursor.prefix != "" {
		return
	}

	if commonLen == len(word) {
		r.keys(cursor, buf, words)

		return
	}

	if commonLen < len(cursor.prefix) {
		return
	}

	if child, ok := cursor.children[word[commonLen]]; ok {
		startLen := len(*buf)

		*buf = append(*buf, cursor.prefix...)

		r.remaining(child, word[commonLen:], buf, words)

		*buf = (*buf)[:startLen]
	}
}
