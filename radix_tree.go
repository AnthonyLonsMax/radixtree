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
	size int64
	root *node
}

// Add inserts a word into the radix tree.
func (r *RadixTree) Add(word string) {
	r.size++
	r.root = r.add(r.root, word)
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
func (r *RadixTree) Clear() {
	r.root = nil
	r.size = 0
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
