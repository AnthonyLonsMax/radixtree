package radixtree_test

import (
	"testing"

	"github.com/AnthonyLonsMax/radixtree"
)

func TestContains(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name       string
		source     []string
		search     string
		shouldFind bool
	}

	testCases := []testCase{
		{name: "empty tree", source: []string{}, search: "a", shouldFind: false},
		{name: "find existing word", source: []string{hello, world}, search: hello, shouldFind: true},
		{name: "find another existing word", source: []string{hello, world}, search: world, shouldFind: true},
		{name: "non-existent word", source: []string{hello, world}, search: "xyz", shouldFind: false},
		{name: "partial match should not be found", source: []string{hello, world}, search: "hell", shouldFind: false},
		{name: "prefix that is also a word", source: []string{"hell", hello}, search: "hell", shouldFind: true},
		{name: "empty string in tree", source: []string{""}, search: "", shouldFind: true},
		{name: "search empty string in non-empty tree", source: []string{"test"}, search: "", shouldFind: false},
		{name: "nested prefixes", source: []string{"a", "ab", "abc"}, search: "ab", shouldFind: true},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			var tree radixtree.RadixTree

			for _, e := range test.source {
				tree.Add(e)
			}

			if tree.Contains(test.search) != test.shouldFind {
				if test.shouldFind {
					t.Fatalf("Word %q should be in the tree", test.search)
				} else {
					t.Fatalf("Word %q should not be in the tree", test.search)
				}
			}
		})
	}
}
