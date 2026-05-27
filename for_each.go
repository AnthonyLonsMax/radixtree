package radixtree

// ForEach calls callbackFn for each word in the tree.
func (r *RadixTree) ForEach(callbackFn func(key string)) {
	r.forEach(r.root, "", callbackFn)
}

func (r *RadixTree) forEach(cursor *node, currentWord string, callbackFn func(string)) {
	if cursor == nil {
		return
	}

	if cursor.isTerminal {
		callbackFn(currentWord + cursor.prefix)
	}

	currentWord += cursor.prefix

	for _, child := range cursor.children {
		r.forEach(child, currentWord, callbackFn)
	}
}
