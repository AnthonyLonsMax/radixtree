package radixtree_test

import (
	"testing"

	"github.com/AnthonyLonsMax/radixtree"
)

func TestSize(t *testing.T) {
	t.Parallel()

	type testCase struct {
		source        []string
		wordsToDelete []string
	}

	tc := []testCase{
		{
			source: []string{
				"worderland",
				"word",
				"worddy",
				"work",
				"worry",
				"wor",
				"worries",
				"wallet",
				"love",
				"lonnly",
				"lovers",
				"anthony",
				"ony",
				"anth",
			},
			wordsToDelete: []string{
				"anthony",
				"ony",
				"anth",
			},
		},
		{
			source:        []string{},
			wordsToDelete: []string{},
		},
		{
			source:        []string{"a"},
			wordsToDelete: []string{},
		},
		{
			source:        []string{"a", "b", "c"},
			wordsToDelete: []string{"a", "b"},
		},
		{
			source:        []string{"hello", "world"},
			wordsToDelete: []string{"nonexistent"},
		},
	}

	for _, test := range tc {
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
	}
}
