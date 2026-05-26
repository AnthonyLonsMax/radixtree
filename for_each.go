package radixtree

// ForEach calls fn for each word in the tree.
func (r *RadixTree) ForEach(fn func(key string)) {
	r.forEach(r.root, "", fn)
}

func (r *RadixTree) forEach(cursor *node, currentWord string, fn func(string)) {
	if cursor == nil {
		return
	}
	if cursor.isTerminal {
		fn(currentWord + cursor.prefix)
	}
	currentWord += cursor.prefix
	for _, child := range cursor.children {
		r.forEach(child, currentWord, fn)
	}
}
