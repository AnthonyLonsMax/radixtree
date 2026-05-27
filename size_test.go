package radixtree_test

import (
	"testing"

	"github.com/AnthonyLonsMax/radixtree"
)

func TestSize(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name          string
		source        []string
		wordsToDelete []string
	}

	tc := []testCase{
		{
			name:   "delete subset of words with common prefixes",
			source: []string{"worderland", "word", "worddy", "work", "worry", "wor", "worries", "wallet", "love", "lonnly", "lovers", "anthony", "ony", "anth"},
			wordsToDelete: []string{
				"anthony", "ony", "anth",
			},
		},
		{
			name:          "empty tree",
			source:        []string{},
			wordsToDelete: []string{},
		},
		{
			name:          "single word no deletion",
			source:        []string{"a"},
			wordsToDelete: []string{},
		},
		{
			name:          "delete some words",
			source:        []string{"a", "b", "c"},
			wordsToDelete: []string{"a", "b"},
		},
		{
			name:          "delete non-existent word",
			source:        []string{"hello", "world"},
			wordsToDelete: []string{"nonexistent"},
		},
	}

	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			var tree radixtree.RadixTree
			for _, e := range test.source {
				tree.Add(e)
			}
			deleteCount := 0
			for _, e := range test.wordsToDelete {
				if tree.Contains(e) {
					deleteCount++
				}
				tree.Delete(e)
			}
			expected := len(test.source) - deleteCount
			if expected != int(tree.Size()) {
				t.Fatalf("Expected size %d got %d", expected, tree.Size())
			}
		})
	}
}
