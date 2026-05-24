package radixtree

import "maps"

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

	if commonLen == 0 && nodeCursor.prefix != "" {
		return nodeCursor, false
	}

	if commonLen == len(word) {
		if nodeCursor.isTerminal {
			nodeCursor.isTerminal = false
			if len(nodeCursor.children) == 0 {
				return nil, true
			}
			if len(nodeCursor.children) == 1 {
				// Merge with it's childrens
				for k, v := range nodeCursor.children {
					if v.isTerminal == false {
						nodeCursor.prefix += v.prefix
						delete(nodeCursor.children, k)
						maps.Copy(nodeCursor.children, v.children)
						nodeCursor.isTerminal = v.isTerminal
					}
				}
			}
			return nodeCursor, true
		}
		return nodeCursor, false
	}

	if _, ok := nodeCursor.children[word[commonLen]]; ok {
		node, ok := r.delete(nodeCursor.children[word[commonLen]], word[commonLen:])
		nodeCursor.children[word[commonLen]] = node
		return nodeCursor, ok
	}

	return nodeCursor, false
}
