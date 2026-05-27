package radixtree_test

import (
	"testing"

	"github.com/AnthonyLonsMax/radixtree"
)

func TestCommonPrefix(t *testing.T) {
	t.Parallel()

	type testCase struct {
		source       []string
		commonPrefix string
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
			},
			commonPrefix: "wor",
		},
		{
			source:       []string{},
			commonPrefix: "",
		},
		{
			source:       []string{"single"},
			commonPrefix: "single",
		},
		{
			source:       []string{"abc", "def"},
			commonPrefix: "",
		},
		{
			source:       []string{"flower", "flow", "flight"},
			commonPrefix: "fl",
		},
		{
			source:       []string{"prefix", "pre", "prefixed"},
			commonPrefix: "pre",
		},
	}

	for _, test := range tc {
		var tree radixtree.RadixTree
		for _, e := range test.source {
			tree.Add(e)
		}
		if tree.CommonPrefix() != test.commonPrefix {
			t.Fatalf("Word %s should be the common prefix", test.commonPrefix)
		}
	}
}
