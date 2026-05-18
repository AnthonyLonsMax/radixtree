// Package radixtree implements a radix tree data structure.
package radixtree

import (
	"fmt"
	"strings"
)

// node represents a node in the radix tree.
type node struct {
	prefix     string
	children   map[byte]*node
	isTerminal bool
}

// newNode creates a new node with the given prefix and terminal status.
func newNode(prefix string, isTerminal bool) *node {
	return &node{
		prefix:     prefix,
		children:   make(map[byte]*node),
		isTerminal: isTerminal,
	}
}

// RadixTree represents a radix tree data structure.
type RadixTree struct {
	root *node
}

// Add inserts a word into the radix tree.
func (r *RadixTree) Add(word string) {
	r.root = r.add(r.root, word)
}

// String returns a string representation of the tree.
func (r *RadixTree) String() string {
	panic("unimplemented")
}

// MarshalText implements encoding.TextMarshaler.
func (r *RadixTree) MarshalText() ([]byte, error) {
	panic("unimplemented")
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (r *RadixTree) UnmarshalText(_ []byte) error {
	panic("unimplemented")
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

// Delete removes a word from the tree. Returns true if the word was found and deleted.
func (r *RadixTree) Delete(_ string) bool {
	panic("unimplemented")
}

// StartsWith returns true if any word in the tree has the given prefix.
func (r *RadixTree) StartsWith(_ string) bool {
	panic("unimplemented")
}

// CommonPrefix returns the longest common prefix among all stored words.
func (r *RadixTree) CommonPrefix() string {
	panic("unimplemented")
}

// Keys returns all stored words in the tree.
func (r *RadixTree) Keys() []string {
	panic("unimplemented")
}

// Size returns the number of words stored in the tree.
func (r *RadixTree) Size() int {
	panic("unimplemented")
}

// Clear removes all words from the tree.
func (r *RadixTree) Clear() {
	panic("unimplemented")
}

// ForEach calls fn for each word in the tree. If fn returns false, iteration stops.
func (r *RadixTree) ForEach(_ func(key string) bool) {
	panic("unimplemented")
}

// LongestPrefixOf returns the longest key that is a prefix of the given word.
func (r *RadixTree) LongestPrefixOf(_ string) string {
	panic("unimplemented")
}

// PrintDebug prints the DFS traversal of the tree prefixes with level indentation.
func (r *RadixTree) PrintDebug() {
	if r.root == nil {
		fmt.Println("<empty>")
		return
	}
	r.dfsPrint(r.root, 0)
	fmt.Println()
}

func (r *RadixTree) dfsPrint(n *node, level int) {
	indent := strings.Repeat(" ", level)
	fmt.Printf("%s%s %v\n", indent, n.prefix, n.isTerminal)

	for _, node := range n.children {
		r.dfsPrint(node, level+1)
	}
}
