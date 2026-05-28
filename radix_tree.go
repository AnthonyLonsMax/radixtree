package radixtree

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"
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

// node represents a node in the radix tree.
type node struct {
	prefix     string
	children   map[byte]*node
	isTerminal bool
}

type edge struct {
	prefixes  string
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
		return rune(e.prefixes[i]) > char
	})
	if index < len(e.prefixes) && rune(e.prefixes[index]) > char {
		e.prefixes = e.prefixes[:index] + string(char) + e.prefixes[index:]
		e.childrens = slices.Insert(e.childrens, index, insertEdge)
	} else {
		e.prefixes += string(char)
		e.childrens = append(e.childrens, insertEdge)
	}

}

// newNode creates a new node with the given prefix and terminal status.
func newNode(prefix string, isTerminal bool) *node {
	return &node{
		prefix:     prefix,
		children:   make(map[byte]*node),
		isTerminal: isTerminal,
	}
}

// MapRadixTree represents a radix tree data structure.
type MapRadixTree struct {
	size int64
	root *node
}

func RadixMap() RadixTree {
	return new(MapRadixTree)
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

// Clear removes all words from the tree.
func (r *MapRadixTree) Clear() {
	r.root = nil
	r.size = 0
}

// PrintDebug prints the DFS traversal of the tree prefixes with level indentation.
func (r *MapRadixTree) PrintDebug() {
	if r.root == nil {
		os.Stderr.WriteString("<empty>\n")

		return
	}

	r.dfsPrint(r.root, 0)
	os.Stderr.WriteString("\n")
}

func (r *MapRadixTree) dfsPrint(node *node, level int) {
	formatFn := func(terminal bool) string {
		if terminal {
			return "W"
		}

		return "N"
	}

	indent := strings.Repeat(" ", level)
	fmt.Printf("%s%s %v\n", indent, node.prefix, formatFn(node.isTerminal))

	for _, childNode := range node.children {
		r.dfsPrint(childNode, level+1)
	}
}
