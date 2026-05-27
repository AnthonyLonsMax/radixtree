package radixtree

// Keys returns all stored words in the tree.
func (r *RadixTree) Keys() []string {
	result := make([]string, 0)
	r.keys(r.root, "", &result)
	return result
}

func (r *RadixTree) keys(cursor *node, currentWord string, words *[]string) {
	if cursor == nil {
		return
	}
	if cursor.isTerminal {
		*words = append(*words, currentWord+cursor.prefix)
	}
	currentWord += cursor.prefix
	for _, child := range cursor.children {
		r.keys(child, currentWord, words)
	}
}
