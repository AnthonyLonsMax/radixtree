package radixtree

import (
	"fmt"
	"os"
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
		os.Stderr.WriteString("<empty>\n")

		return
	}

	r.dfsPrint(r.root, 0)
	os.Stderr.WriteString("\n")
}

func (r *RadixTree) dfsPrint(node *node, level int) {
	formatFn := func(terminal bool) string {
		if terminal {
			return "W"
		}

		return "N"
	}

	indent := strings.Repeat(" ", level)
	os.Stderr.WriteString(fmt.Sprintf("%s%s %v\n", indent, node.prefix, formatFn(node.isTerminal)))

	for _, childNode := range node.children {
		r.dfsPrint(childNode, level+1)
	}
}
