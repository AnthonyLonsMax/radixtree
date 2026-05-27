package radixtree_test

import (
	"testing"

	"github.com/AnthonyLonsMax/radixtree"
)

func TestStartsWith(t *testing.T) {
	t.Parallel()

	type testCase struct {
		source   []string
		prefix   string
		expected bool
	}

	tc := []testCase{
		{source: []string{}, prefix: "a", expected: false},
		{source: []string{"hello", "world"}, prefix: "he", expected: true},
		{source: []string{"hello", "world"}, prefix: "wor", expected: true},
		{source: []string{"hello", "world"}, prefix: "xyz", expected: false},
		{source: []string{"hello", "world"}, prefix: "hello", expected: true},
		{source: []string{"hello"}, prefix: "helloworld", expected: false},
		{source: []string{""}, prefix: "", expected: true},
		{source: []string{"test"}, prefix: "", expected: true},
		{source: []string{"abc", "def"}, prefix: "a", expected: true},
		{source: []string{"abc", "def"}, prefix: "d", expected: true},
		{source: []string{"abc", "def"}, prefix: "ab", expected: true},
	}

	for _, test := range tc {
		var tree radixtree.RadixTree
		for _, e := range test.source {
			tree.Add(e)
		}
		if tree.StartsWith(test.prefix) != test.expected {
			if test.expected {
				t.Fatalf("Tree should start with prefix %q", test.prefix)
			} else {
				t.Fatalf("Tree should not start with prefix %q", test.prefix)
			}
		}
	}
}
