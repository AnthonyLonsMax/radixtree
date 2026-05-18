package radixtree

// Contains returns true if the word exists in the tree.
func (r *RadixTree) Contains(word string) bool {
	return r.contains(r.root, word)
}

func (r *RadixTree) contains(nodeCursor *node, word string) bool {
	if nodeCursor == nil {
		return false
	}

	commonLen := commonPrefixLength(nodeCursor.prefix, word)

	if commonLen == 0 && nodeCursor.prefix != "" {
		return false
	}

	if commonLen == len(word) {
		return true
	}

	if _, ok := nodeCursor.children[word[commonLen]]; ok {
		return r.contains(nodeCursor.children[word[commonLen]], word[commonLen:])
	}

	return false
}
