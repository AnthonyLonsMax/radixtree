package radixmap

// Maximum returns the lexicographically largest word in the tree.
// Returns empty string if the tree is empty.
func (r *MapRadixTree) Maximum() string {
	if r.root == nil {
		return ""
	}

	result := ""
	acc := ""
	cursor := r.root

	for {
		if cursor.isTerminal {
			result = acc + cursor.prefix
		}

		var (
			maxKey   byte
			maxChild *node
		)

		for k := range cursor.children {
			if maxChild == nil || k > maxKey {
				maxKey = k
				maxChild = cursor.children[k]
			}
		}

		if maxChild == nil {
			return result
		}

		acc += cursor.prefix
		cursor = maxChild
	}
}
