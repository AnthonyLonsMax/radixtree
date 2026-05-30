package radixtree

import (
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

func RadixMap() RadixTree {
	return new(radixmap.MapRadixTree)
}
