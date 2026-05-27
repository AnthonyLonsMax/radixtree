package radixtree

// Remaining returns all stored words that have the given prefix.
func (r *RadixTree) Remaining(prefix string) []string {
	result := make([]string, 0)

	if r.root == nil {
		return result
	}

	if prefix == "" {
		r.keys(r.root, "", &result)

		return result
	}

	r.remaining(r.root, prefix, "", &result)

	return result
}

func (r *RadixTree) remaining(cursor *node, word, acc string, words *[]string) {
	if cursor == nil {
		return
	}

	commonLen := commonPrefixLength(cursor.prefix, word)

	if commonLen == 0 && cursor.prefix != "" {
		return
	}

	if commonLen == len(word) {
		r.keys(cursor, acc, words)

		return
	}

	if commonLen < len(cursor.prefix) {
		return
	}

	if child, ok := cursor.children[word[commonLen]]; ok {
		r.remaining(child, word[commonLen:], acc+cursor.prefix, words)
	}
}
