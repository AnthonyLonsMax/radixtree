package radixtree

import "strings"

// CommonPrefix returns the longest common prefix among all stored words.
func (r *RadixTree) CommonPrefix() string {
	if r.root == nil {
		return ""
	}

	prefix := strings.Builder{}
	cursor := r.root

	for {
		prefix.WriteString(cursor.prefix)

		if cursor.isTerminal || len(cursor.children) > 1 {
			break
		}

		for _, uniqueNode := range cursor.children {
			cursor = uniqueNode
		}
	}

	return prefix.String()
}
