package radixtree_test

import (
	"testing"

	"github.com/AnthonyLonsMax/radixtree"
)

func TestForEach(t *testing.T) {
	t.Parallel()

	type testCase struct {
		source []string
	}

	tc := []testCase{
		{source: []string{}},
		{source: []string{"a"}},
		{source: []string{"hello", "world", "hi"}},
		{source: []string{
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
		}},
	}

	for _, test := range tc {
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
	}
}
