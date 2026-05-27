package radixtree_test

import (
	"testing"

	"github.com/AnthonyLonsMax/radixtree"
)

func TestDeleteWord(t *testing.T) {
	t.Parallel()

	type testCase struct {
		source            []string
		shouldNotContains []string
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
			shouldNotContains: []string{
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
			source:            []string{},
			shouldNotContains: []string{"hello"},
		},
		{
			source:            []string{"hello", "world"},
			shouldNotContains: []string{"nonexistent"},
		},
		{
			source:            []string{"hello", "world"},
			shouldNotContains: []string{"hello", "world"},
		},
		{
			source:            []string{"test", "testing", "tester"},
			shouldNotContains: []string{"test", "testing", "tester"},
		},
	}

	for _, test := range tc {
		var tree radixtree.RadixTree
		for _, e := range test.source {
			tree.Add(e)
		}
		for _, e := range test.shouldNotContains {
			tree.Delete(e)
			if tree.Contains(e) {
				t.Fatalf("Word %s should not be in the tree", e)
			}
		}
	}
}
