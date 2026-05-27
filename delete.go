package radixtree

// Delete removes a word from the tree. Returns true if the word was found and deleted.
func (r *RadixTree) Delete(word string) bool {
	node, ok := r.delete(r.root, word)

	r.root = node

	if ok {
		r.size--
	}

	return ok
}

func (r *RadixTree) delete(nodeCursor *node, word string) (*node, bool) {
	if nodeCursor == nil {
		return nodeCursor, false
	}

	commonLen := commonPrefixLength(nodeCursor.prefix, word)

	switch {
	case commonLen == 0 && nodeCursor.prefix != "":
		return nodeCursor, false

	case commonLen == len(word):
		if commonLen < len(nodeCursor.prefix) || !nodeCursor.isTerminal {
			return nodeCursor, false
		}

		return r.deleteExactMatch(nodeCursor)

	default:
		return r.deleteRecurse(nodeCursor, word, commonLen)
	}
}

func (r *RadixTree) deleteExactMatch(nodeCursor *node) (*node, bool) {
	nodeCursor.isTerminal = false

	if len(nodeCursor.children) == 0 {
		return nil, true
	}

	if len(nodeCursor.children) == 1 {
		for _, v := range nodeCursor.children {
			nodeCursor.prefix += v.prefix
			nodeCursor.children = v.children

			nodeCursor.isTerminal = v.isTerminal
		}
	}

	return nodeCursor, true
}

func (r *RadixTree) deleteRecurse(nodeCursor *node, word string, commonLen int) (*node, bool) {
	child, ok := r.delete(nodeCursor.children[word[commonLen]], word[commonLen:])

	if child != nil {
		nodeCursor.children[word[commonLen]] = child
	} else {
		delete(nodeCursor.children, word[commonLen])
	}

	if !nodeCursor.isTerminal && len(nodeCursor.children) == 0 {
		return nil, ok
	}

	if !nodeCursor.isTerminal && len(nodeCursor.children) == 1 {
		for _, child := range nodeCursor.children {
			nodeCursor.prefix += child.prefix
			nodeCursor.children = child.children

			nodeCursor.isTerminal = child.isTerminal
		}
	}

	return nodeCursor, ok
}
