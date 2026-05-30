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

	// Split the prefix word
	newNode.prefix = src.prefix[position:]
	src.prefix = src.prefix[:position]

	// Cero the values
	src.prefixes = make([]rune, 0)
	src.childrens = make([]*edge, 0)

	// Append the split part
	src.childrens = append(src.childrens, newNode)
	src.prefixes = append(src.prefixes, rune(newNode.prefix[0]))
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

func (r *RadixOrdered) add(cursor *edge, word string) (*edge, bool) {
	if cursor == nil {
		return newEdge(word, true), true
	}

	commonLen := commonPrefixLength(cursor.prefix, word)

	var added bool

	minLength := min(len(cursor.prefix), len(word))

	switch {
	case commonLen == 0:
		if cursor.prefix != "" {
			splitEdge(cursor, commonLen)
		}
		cursor.insertOrdered(rune(word[0]), newEdge(word, true))
		return cursor, true

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
		return r.add(rest, word[commonLen:])

	case commonLen == len(word) && commonLen < len(cursor.prefix):
		if cursor.prefix != "" {
			splitEdge(cursor, commonLen)
		}
		cursor.isTerminal = true
		return cursor, true

	case 0 < commonLen && commonLen < minLength: // partial cover
		if cursor.prefix != "" {
			splitEdge(cursor, commonLen)
		}
		rest := cursor.getChildren(rune(word[commonLen]))
		if rest == nil {
			newNode := newEdge(word[commonLen:], true)
			cursor.insertOrdered(rune(newNode.prefix[0]), newNode)
			return cursor, true
		}
		return r.add(rest, word[commonLen:])
	}

	//if !cursor.isTerminal && len(cursor.childrens) == 1 {
	//	for _, child := range cursor.childrens {
	//		cursor.prefix += child.prefix
	//		cursor.childrens = child.childrens
	//		cursor.isTerminal = child.isTerminal
	//	}
	//}

	return cursor, added
}

func (r *RadixOrdered) Contains(word string) bool {
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
	case commonLen == len(cursor.prefix) && commonLen < len(word):
		rest := cursor.getChildren(rune(word[commonLen]))
		if rest == nil {
			return false
		}
		return r.contains(rest, word[commonLen:])
	}
	return false
}
