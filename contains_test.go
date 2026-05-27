package radixtree_test

import (
	"testing"

	"github.com/AnthonyLonsMax/radixtree"
)

func TestContains(t *testing.T) {
	t.Parallel()

	type testCase struct {
		source     []string
		search     string
		shouldFind bool
	}

	tc := []testCase{
		{source: []string{}, search: "a", shouldFind: false},
		{source: []string{"hello", "world"}, search: "hello", shouldFind: true},
		{source: []string{"hello", "world"}, search: "world", shouldFind: true},
		{source: []string{"hello", "world"}, search: "xyz", shouldFind: false},
		{source: []string{"hello", "world"}, search: "hell", shouldFind: false},
		{source: []string{"hell", "hello"}, search: "hell", shouldFind: true},
		{source: []string{""}, search: "", shouldFind: true},
		{source: []string{"test"}, search: "", shouldFind: false},
		{source: []string{"a", "ab", "abc"}, search: "ab", shouldFind: true},
	}

	for _, test := range tc {
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
	}
}
