package radixtree_test

import (
	"testing"

	"github.com/AnthonyLonsMax/radixtree"
)

func TestForEach(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name   string
		source []string
	}

	tc := []testCase{
		{name: "empty tree", source: []string{}},
		{name: "single word", source: []string{"a"}},
		{name: "multiple unrelated words", source: []string{"hello", "world", "hi"}},
		{name: "multiple words with common prefixes", source: []string{
			"worderland", "word", "worddy", "work", "worry",
			"wor", "worries", "wallet", "love", "lonnly",
			"lovers", "anthony", "ony", "anth",
		}},
	}

	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			var tree radixtree.RadixTree
			for _, e := range test.source {
				tree.Add(e)
			}
			collected := make(map[string]bool)
			tree.ForEach(func(key string) {
				collected[key] = true
			})
			for _, e := range test.source {
				if !collected[e] {
					t.Fatalf("Key %q should be returned by ForEach", e)
				}
			}
		})
	}
}
