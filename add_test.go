package radixtree_test

import (
	"testing"

	"github.com/AnthonyLonsMax/radixtree"
)

func TestAddNodes(t *testing.T) {
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
			shouldContains: []string{
				"worddy",
				"work",
				"worry",
				"wor",
				"worries",
				"wallet",
				"love",
			},
		},
		{
			source:         []string{""},
			shouldContains: []string{""},
		},
		{
			source:         []string{"a", "b", "c"},
			shouldContains: []string{"a", "b", "c"},
		},
		{
			source:         []string{"hello", "hello", "hello"},
			shouldContains: []string{"hello"},
		},
		{
			source:         []string{"café", "cafè"},
			shouldContains: []string{"café", "cafè"},
		},
		{
			source:         []string{"123", "456", "789"},
			shouldContains: []string{"123", "789"},
		},
	}

	for _, test := range tc {
		var tree radixtree.RadixTree
		for _, e := range test.source {
			tree.Add(e)
		}
		for _, e := range test.shouldContains {
			if !tree.Contains(e) {
				t.Fatalf("Word %s should be in the tree", e)
			}
		}
	}
}
