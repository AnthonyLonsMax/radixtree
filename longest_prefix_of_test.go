package radixtree_test

import (
	"testing"

	"github.com/AnthonyLonsMax/radixtree"
)

func TestLongestPrefixOf(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name     string
		source   []string
		word     string
		expected string
	}

	tc := []testCase{
		{name: "empty tree", source: []string{}, word: "hello", expected: ""},
		{name: "exact match", source: []string{"hello"}, word: "hello", expected: "hello"},
		{name: "prefix word is shorter than input", source: []string{"hello", "helloworld"}, word: "helloworld", expected: "helloworld"},
		{name: "exact match with multiple options", source: []string{"hello", "helloworld"}, word: "hello", expected: "hello"},
		{name: "partial prefix match", source: []string{"hello"}, word: "helloworld", expected: "hello"},
		{name: "no prefix match", source: []string{"helloworld"}, word: "hello", expected: ""},
		{name: "deepest prefix match", source: []string{"a", "ab", "abc"}, word: "abcd", expected: "abc"},
		{name: "match intermediate prefix", source: []string{"a", "ab", "abc"}, word: "ab", expected: "ab"},
		{name: "match shortest prefix", source: []string{"a", "ab", "abc"}, word: "a", expected: "a"},
		{name: "empty string in tree", source: []string{""}, word: "anything", expected: ""},
		{name: "no match", source: []string{"test"}, word: "xyz", expected: ""},
		{name: "prefix nested within longer word", source: []string{"pre", "prefix"}, word: "prefixed", expected: "prefix"},
	}

	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			var tree radixtree.RadixTree
			for _, e := range test.source {
				tree.Add(e)
			}
			result := tree.LongestPrefixOf(test.word)
			if result != test.expected {
				t.Fatalf("LongestPrefixOf(%q) = %q, want %q", test.word, result, test.expected)
			}
		})
	}
}
