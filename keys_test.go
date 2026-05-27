package radixtree_test

import (
	"slices"
	"testing"

	"github.com/AnthonyLonsMax/radixtree"
)

func TestReadAllTheKeys(t *testing.T) {
	t.Parallel()

	type testCase struct {
		source         []string
		shouldContains []string
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
		},
		{
			source: []string{},
		},
		{
			source: []string{"single"},
		},
		{
			source: []string{"a", "b", "c"},
		},
	}

	for _, test := range tc {
		var tree radixtree.RadixTree
		for _, e := range test.source {
			tree.Add(e)
		}
		keys := *tree.Keys()
		slices.Sort(keys)

		for _, e := range test.source {
			if _, ok := slices.BinarySearch(keys, e); !ok {
				t.Fatalf("Word %s should be in the keys slice", e)
			}
		}
	}
}
