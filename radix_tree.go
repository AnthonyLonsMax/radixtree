package radixtree

import (
	"slices"
	"sort"

	"github.com/AnthonyLonsMax/radixtree/internal/radixmap"
)

type RadixTree interface {
	Add(word string) bool
	Delete(word string) bool
	Contains(word string) bool
	StartsWith(word string) bool
	LongestPrefixOf(word string) string
	ForEach(fn func(string))
	Keys() []string
	Size() int64
	Remaining(prefix string) []string
}

type edge struct {
	prefixes  []rune
	prefix    string
	childrens []*edge
}

// Avoid the use of `""` as root node
type edgeRoot struct {
	childrens [512]*edge
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

func (e *edge) insert(char rune, insertEdge *edge) {
	e.prefixes = append(e.prefixes, char)
	e.childrens = append(e.childrens, insertEdge)
}

func (e *edge) splitEdge(src *edge, position int) {
	newNode := new(edge)
	newNode.childrens = e.childrens
	newNode.prefixes = e.prefixes
	newNode.prefix = src.prefix[position:]
	src.prefix = src.prefix[:position]
	src.prefixes = make([]rune, 0)
	src.childrens = make([]*edge, 0)
	e.insert(rune(newNode.prefix[0]), newNode)
}

func RadixMap() RadixTree {
	return new(radixmap.MapRadixTree)
}
