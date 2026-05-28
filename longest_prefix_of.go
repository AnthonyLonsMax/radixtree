package radixtree

// LongestPrefixOf returns the longest key that is a prefix of the given word.
func (r *MapRadixTree) LongestPrefixOf(word string) string {
	return r.longestPrefixOf(r.root, word)
}

func (r *MapRadixTree) longestPrefixOf(cursor *node, word string) string {
	if cursor == nil {
		return ""
	}
	commonLength := commonPrefixLength(word, cursor.prefix)
	if commonLength < len(cursor.prefix) {
		return ""
	}

	var result string
	if cursor.isTerminal {
		result = cursor.prefix
	}

	if commonLength == len(word) {
		return result
	}

	if childResult := r.longestPrefixOf(cursor.children[word[commonLength]], word[commonLength:]); childResult != "" {
		return cursor.prefix + childResult
	}

	return result
}
