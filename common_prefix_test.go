package radixtree_test

import (
	"testing"

	"github.com/AnthonyLonsMax/radixtree"
)

func TestCommonPrefix(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name         string
		source       []string
		commonPrefix string
	}

	tc := []testCase{
		{
			name:         "multiple words with common prefix",
			source:       []string{"worderland", "word", "worddy", "work", "worry", "wor"},
			commonPrefix: "wor",
		},
		{
			name:         "empty tree",
			source:       []string{},
			commonPrefix: "",
		},
		{
			name:         "single word",
			source:       []string{"single"},
			commonPrefix: "single",
		},
		{
			name:         "no common prefix",
			source:       []string{"abc", "def"},
			commonPrefix: "",
		},
		{
			name:         "words sharing partial prefix",
			source:       []string{"flower", "flow", "flight"},
			commonPrefix: "fl",
		},
		{
			name:         "words with longer common prefix",
			source:       []string{"prefix", "pre", "prefixed"},
			commonPrefix: "pre",
		},
	}

	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			var tree radixtree.RadixTree
			for _, e := range test.source {
				tree.Add(e)
			}
			if tree.CommonPrefix() != test.commonPrefix {
				t.Fatalf("Word %s should be the common prefix", test.commonPrefix)
			}
		})
	}
}
