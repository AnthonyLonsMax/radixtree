package radixtree

import "maps"

func (r *RadixTree) nodeSplit(current *node, pos int) *node {
	child := newNode(current.prefix[pos:], current.isTerminal)

	maps.Copy(child.children, current.children)

	current.prefix = current.prefix[:pos]
	current.children = make(map[byte]*node)

	current.children[child.prefix[0]] = child
	current.isTerminal = false

	return child
}

func (r *RadixTree) add(cursor *node, word string) *node {
	if cursor == nil {
		return newNode(word, true)
	}

	commonLen := commonPrefixLength(cursor.prefix, word)

	switch {
	case commonLen == 0:
		if cursor.prefix != "" {
			r.nodeSplit(cursor, 0)
		}
		if _, ok := cursor.children[word[commonLen]]; ok {
			cursor.children[word[commonLen]] = r.add(cursor.children[word[commonLen]], word[commonLen:])
		} else {
			cursor.children[word[commonLen]] = newNode(word[commonLen:], true)
		}

	// Avoid duplicates
	case commonLen == len(cursor.prefix) && commonLen == len(word):
		cursor.isTerminal = true

	case commonLen == len(cursor.prefix) && commonLen < len(word):
		if _, ok := cursor.children[word[commonLen]]; ok {
			cursor.children[word[commonLen]] = r.add(cursor.children[word[commonLen]], word[commonLen:])
		} else {
			cursor.children[word[commonLen]] = newNode(word[commonLen:], true)
		}

	case commonLen == len(word) && commonLen < len(cursor.prefix):
		r.nodeSplit(cursor, commonLen)
		cursor.isTerminal = true

	default:
		r.nodeSplit(cursor, commonLen)

		if _, ok := cursor.children[word[commonLen]]; ok {
			cursor.children[word[commonLen]] = r.add(cursor.children[word[commonLen]], word[commonLen:])
		} else {
			cursor.children[word[commonLen]] = newNode(word[commonLen:], true)
		}

	}

	// If there is one child and it's no terminal -> merge it
	if !cursor.isTerminal && len(cursor.children) == 1 {
		cursor.prefix += cursor.children[0].prefix
		children := cursor.children[0].children
		terminal := cursor.children[0].isTerminal
		cursor.children = make(map[byte]*node)
		maps.Copy(cursor.children, children)
		cursor.isTerminal = terminal
	}

	return cursor
}
