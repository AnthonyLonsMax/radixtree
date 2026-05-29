package radixmap

// Contains returns true if the word exists in the tree.
func (r *MapRadixTree) Contains(word string) bool {
	return r.contains(r.root, word)
}

func (r *MapRadixTree) contains(nodeCursor *node, word string) bool {
	if nodeCursor == nil {
		return false
	}

	commonLen := commonPrefixLength(nodeCursor.prefix, word)

	if commonLen == 0 && nodeCursor.prefix != "" {
		return false
	}

	if commonLen == len(word) {
		if commonLen == len(nodeCursor.prefix) {
			return nodeCursor.isTerminal
		}
		return false
	}

	if _, ok := nodeCursor.children[word[commonLen]]; ok {
		return r.contains(nodeCursor.children[word[commonLen]], word[commonLen:])
	}

	return false
}
