package radixtree

import "math"

// Minimum returns the lexicographically smallest word in the tree.
// Returns empty string if the tree is empty.
func (r *MapRadixTree) Minimum() string {
	if r.root == nil {
		return ""
	}

	acc := ""
	cursor := r.root

	for {
		if cursor.isTerminal {
			return acc + cursor.prefix
		}

		minKey := byte(math.MaxUint8)

		for k := range cursor.children {
			if k < minKey {
				minKey = k
			}
		}

		if len(cursor.children) == 0 {
			return ""
		}

		acc += cursor.prefix
		cursor = cursor.children[minKey]
	}
}
