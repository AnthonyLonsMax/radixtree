package radixordered

import (
	"slices"
	"sort"
)

type RadixOrdered struct {
	root  *edge
	count int
}

func (r *RadixOrdered) Add(word string) bool {
	if word == "" {
		return false
	}
	added := false
	r.root, added = r.add(r.root, word)
	if added {
		r.count++
	}
	return added
}

type edge struct {
	prefixes   []rune
	prefix     string
	childrens  []*edge
	isTerminal bool
}

func newEdge(prefix string, isTerminal bool) *edge {
	return &edge{
		prefixes:   make([]rune, 0),
		childrens:  make([]*edge, 0),
		prefix:     prefix,
		isTerminal: isTerminal,
	}
}

func commonPrefixLength(word1, word2 string) int {
	minLength := min(len(word1), len(word2))
	count := 0

	for i := range minLength {
		if word1[i] == word2[i] {
			count++
		} else {
			break
		}
	}

	return count
}

func (e *edge) getChildren(char rune) *edge {
	for index, prefixChild := range e.prefixes {
		if prefixChild == char {
			return e.childrens[index]
		}
	}
	return nil
}

func (e *edge) insertOrdered(char rune, insertEdge *edge) {
	index := sort.Search(len(e.prefixes), func(i int) bool {
		return e.prefixes[i] > char
	})
	if index < len(e.prefixes) && e.prefixes[index] > char {
		e.prefixes = slices.Insert(e.prefixes, index, char)
		e.childrens = slices.Insert(e.childrens, index, insertEdge)
	} else {
		e.prefixes = append(e.prefixes, char)
		e.childrens = append(e.childrens, insertEdge)
	}
}

func splitEdge(src *edge, position int) {
	// Create a copy the node data
	newNode := new(edge)
	newNode.childrens = src.childrens
	newNode.prefixes = src.prefixes
	newNode.isTerminal = src.isTerminal

	// Split the prefix word
	newNode.prefix = src.prefix[position:]
	src.prefix = src.prefix[:position]

	// Cero the values
	src.prefixes = make([]rune, 0)
	src.childrens = make([]*edge, 0)
	src.isTerminal = false

	// Append the split part
	src.insertOrdered(rune(newNode.prefix[0]), newNode)
}

func (r *RadixOrdered) add(cursor *edge, word string) (*edge, bool) {
	if cursor == nil {
		return newEdge(word, true), true
	}

	commonLen := commonPrefixLength(cursor.prefix, word)

	var added bool

	switch {
	case commonLen == 0:
		if cursor.prefix != "" {
			splitEdge(cursor, commonLen)
		}
		rest := cursor.getChildren(rune(word[commonLen]))
		if rest == nil {
			newNode := newEdge(word[commonLen:], true)
			cursor.insertOrdered(rune(newNode.prefix[0]), newNode)
			return cursor, true
		}
		rest, added = r.add(rest, word[commonLen:])
		return cursor, added

	case commonLen == len(cursor.prefix) && commonLen == len(word):
		cursor.isTerminal = true
		return cursor, true

	case commonLen == len(cursor.prefix) && commonLen < len(word):
		rest := cursor.getChildren(rune(word[commonLen]))
		if rest == nil {
			newNode := newEdge(word[commonLen:], true)
			cursor.insertOrdered(rune(newNode.prefix[0]), newNode)
			return cursor, true
		}
		rest, added = r.add(rest, word[commonLen:])
		return cursor, added

	case commonLen == len(word) && commonLen < len(cursor.prefix):
		if cursor.prefix != "" {
			splitEdge(cursor, commonLen)
		}
		cursor.isTerminal = true
		return cursor, true

	default: // partial cover
		if cursor.prefix != "" {
			splitEdge(cursor, commonLen)
		}
		newNode := newEdge(word[commonLen:], true)
		cursor.insertOrdered(rune(newNode.prefix[0]), newNode)
		return cursor, true
	}
}

func (r *RadixOrdered) Contains(word string) bool {
	if word == "" {
		return false
	}
	return r.contains(r.root, word)
}

func (r *RadixOrdered) contains(cursor *edge, word string) bool {
	if cursor == nil {
		return false
	}

	commonLen := commonPrefixLength(cursor.prefix, word)
	if commonLen == 0 && cursor.prefix == "" {
		rest := cursor.getChildren(rune(word[0]))
		if rest == nil {
			return false
		}
		return r.contains(rest, word)
	}

	switch {
	case commonLen == len(word) && commonLen == len(cursor.prefix):
		return cursor.isTerminal
	default:
		rest := cursor.getChildren(rune(word[commonLen]))
		return r.contains(rest, word[commonLen:])
	}
}

func (r *RadixOrdered) Delete(word string) bool {
	if word == "" {
		return false
	}
	var deleted bool
	r.root, deleted = r.delete(r.root, word)
	return deleted
}

func (r *RadixOrdered) delete(cursor *edge, word string) (*edge, bool) {
	if cursor == nil {
		return nil, false
	}

	var deleted bool

	commonLen := commonPrefixLength(cursor.prefix, word)
	if commonLen == 0 && cursor.prefix == "" {
		rest := cursor.getChildren(rune(word[0]))
		if rest == nil {
			return cursor, false
		}
		rest, deleted = r.delete(rest, word)
		return cursor, deleted
	}

	switch {
	case commonLen == len(word) && commonLen == len(cursor.prefix):
		if cursor.isTerminal {
			cursor.isTerminal = false
			childCount := len(cursor.childrens)
			switch childCount {
			case 0:
				return nil, true
			case 1:
				for _, child := range cursor.childrens {
					cursor.prefix += child.prefix
					cursor.childrens = child.childrens
					cursor.isTerminal = child.isTerminal
				}
				return cursor, true
			default:
				return cursor, true
			}
		}
		return cursor, false
	default:
		rest := cursor.getChildren(rune(word[commonLen]))
		rest, deleted = r.delete(rest, word)
		return cursor, deleted
	}
}
