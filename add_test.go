package radixtree_test

import (
	"testing"

	"github.com/AnthonyLonsMax/radixtree"
)

func TestAddNodes(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name           string
		source         []string
		shouldContains []string
	}

	testCases := []testCase{
		{
			name: "multiple words with common prefixes",
			source: []string{
				"worderland", "word", "worddy", "work", "worry",
				"wor", "worries", "wallet", "love", "lonnly",
				"lovers", "anthony", "ony", "anth",
			},
			shouldContains: []string{
				"worddy", "work", "worry", "wor", "worries", "wallet", "love",
			},
		},
		{
			name:           "empty string",
			source:         []string{""},
			shouldContains: []string{""},
		},
		{
			name:           "single characters",
			source:         []string{"a", "b", "c"},
			shouldContains: []string{"a", "b", "c"},
		},
		{
			name:           "duplicate words",
			source:         []string{"hello", "hello", "hello"},
			shouldContains: []string{"hello"},
		},
		{
			name:           "accented characters",
			source:         []string{"café", "cafè"},
			shouldContains: []string{"café", "cafè"},
		},
		{
			name:           "numeric strings",
			source:         []string{"123", "456", "789"},
			shouldContains: []string{"123", "789"},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			var tree radixtree.RadixTree

			for _, e := range test.source {
				tree.Add(e)
			}

			for _, e := range test.shouldContains {
				if !tree.Contains(e) {
					t.Fatalf("Word %s should be in the tree", e)
				}
			}
		})
	}
}
