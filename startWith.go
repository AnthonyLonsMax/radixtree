package radixtree

// StartsWith returns true if any word in the tree has the given prefix.
func (r *RadixTree) StartsWith(word string) bool {
	if word == "" {
		return r.root != nil
	}
	return r.startWith(r.root, word)
}

func (r *RadixTree) startWith(nodeCursor *node, word string) bool {
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
		return r.startWith(nodeCursor.children[word[commonLen]], word[commonLen:])
	}

	return false
}
