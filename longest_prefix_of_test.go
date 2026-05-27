package radixtree_test

import (
	"testing"

	"github.com/AnthonyLonsMax/radixtree"
)

func TestLongestPrefixOf(t *testing.T) {
	t.Parallel()

	type testCase struct {
		source   []string
		word     string
		expected string
	}

	tc := []testCase{
		{source: []string{}, word: "hello", expected: ""},
		{source: []string{"hello"}, word: "hello", expected: "hello"},
		{source: []string{"hello", "helloworld"}, word: "helloworld", expected: "helloworld"},
		{source: []string{"hello", "helloworld"}, word: "hello", expected: "hello"},
		{source: []string{"hello"}, word: "helloworld", expected: "hello"},
		{source: []string{"helloworld"}, word: "hello", expected: ""},
		{source: []string{"a", "ab", "abc"}, word: "abcd", expected: "abc"},
		{source: []string{"a", "ab", "abc"}, word: "ab", expected: "ab"},
		{source: []string{"a", "ab", "abc"}, word: "a", expected: "a"},
		{source: []string{""}, word: "anything", expected: ""},
		{source: []string{"test"}, word: "xyz", expected: ""},
		{source: []string{"pre", "prefix"}, word: "prefixed", expected: "prefix"},
	}

	for _, test := range tc {
		var tree radixtree.RadixTree
		for _, e := range test.source {
			tree.Add(e)
		}
		result := tree.LongestPrefixOf(test.word)
		if result != test.expected {
			t.Fatalf("LongestPrefixOf(%q) = %q, want %q", test.word, result, test.expected)
		}
	}
}
