package radixtree

import "maps"

// Add inserts a word into the radix tree. Returns true if the word was newly added.
func (r *RadixTree) Add(word string) bool {
	if word == "" {
		if r.root == nil {
			r.root = newNode("", true)
			r.size++

			return true
		}

		if r.root.prefix != "" {
			r.nodeSplit(r.root, 0)
		}

		if !r.root.isTerminal {
			r.root.isTerminal = true
			r.size++

			return true
		}

		return false
	}

	node, added := r.add(r.root, word)

	r.root = node

	if added {
		r.size++
	}

	return added
}

func (r *RadixTree) nodeSplit(current *node, pos int) {
	child := newNode(current.prefix[pos:], current.isTerminal)

	maps.Copy(child.children, current.children)

	current.prefix = current.prefix[:pos]
	current.children = make(map[byte]*node)

	current.children[child.prefix[0]] = child
	current.isTerminal = false
}

func (r *RadixTree) add(cursor *node, word string) (*node, bool) {
	if cursor == nil {
		return newNode(word, true), true
	}

	commonLen := commonPrefixLength(cursor.prefix, word)

	var added bool

	switch {
	case commonLen == 0:
		added = r.addCommonLenZero(cursor, word, commonLen)

	case commonLen == len(cursor.prefix) && commonLen == len(word):
		added = !cursor.isTerminal
		cursor.isTerminal = true

	case commonLen == len(cursor.prefix) && commonLen < len(word):
		added = r.addRecurseOrCreate(cursor, word, commonLen)

	case commonLen == len(word) && commonLen < len(cursor.prefix):
		r.nodeSplit(cursor, commonLen)
		added = true

		cursor.isTerminal = true

	default:
		r.nodeSplit(cursor, commonLen)
		added = r.addRecurseOrCreate(cursor, word, commonLen)
	}

	if !cursor.isTerminal && len(cursor.children) == 1 {
		for _, child := range cursor.children {
			cursor.prefix += child.prefix
			cursor.children = child.children

			cursor.isTerminal = child.isTerminal
		}
	}

	return cursor, added
}

func (r *RadixTree) addCommonLenZero(cursor *node, word string, commonLen int) bool {
	if cursor.prefix != "" {
		r.nodeSplit(cursor, 0)
	}

	return r.addRecurseOrCreate(cursor, word, commonLen)
}

func (r *RadixTree) addRecurseOrCreate(cursor *node, word string, commonLen int) bool {
	if child, ok := cursor.children[word[commonLen]]; ok {
		child, added := r.add(child, word[commonLen:])
		cursor.children[word[commonLen]] = child

		return added
	}

	cursor.children[word[commonLen]] = newNode(word[commonLen:], true)

	return true
}
