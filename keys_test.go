package radixtree_test

import (
	"slices"
	"testing"

	"github.com/AnthonyLonsMax/radixtree"
)

func TestReadAllTheKeys(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name   string
		source []string
	}

	tc := []testCase{
		{
			name:   "multiple words with common prefixes",
			source: []string{"worderland", "word", "worddy", "work", "worry", "wor", "worries", "wallet", "love", "lonnly", "lovers", "anthony", "ony", "anth"},
		},
		{
			name:   "empty tree",
			source: []string{},
		},
		{
			name:   "single word",
			source: []string{"single"},
		},
		{
			name:   "single characters",
			source: []string{"a", "b", "c"},
		},
	}

	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			var tree radixtree.RadixTree
			for _, e := range test.source {
				tree.Add(e)
			}
			keys := tree.Keys()
			slices.Sort(keys)

			for _, e := range test.source {
				if _, ok := slices.BinarySearch(keys, e); !ok {
					t.Fatalf("Word %s should be in the keys slice", e)
				}
			}
		})
	}
}
