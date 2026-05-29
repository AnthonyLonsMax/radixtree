package radixmap_test

import (
	"testing"

	"github.com/AnthonyLonsMax/radixtree/internal/radixmap"
)

func TestStartsWith(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name     string
		source   []string
		prefix   string
		expected bool
	}

	testCases := []testCase{
		{name: "empty tree", source: []string{}, prefix: "a", expected: false},
		{name: "matching prefix", source: []string{"hello", "world"}, prefix: "he", expected: true},
		{name: "matching another prefix", source: []string{"hello", "world"}, prefix: "wor", expected: true},
		{name: "non-matching prefix", source: []string{"hello", "world"}, prefix: "xyz", expected: false},
		{name: "full word as prefix", source: []string{"hello", "world"}, prefix: "hello", expected: true},
		{name: "prefix longer than any word", source: []string{"hello"}, prefix: "helloworld", expected: false},
		{name: "empty prefix on empty tree", source: []string{""}, prefix: "", expected: true},
		{name: "empty prefix on non-empty tree", source: []string{"test"}, prefix: "", expected: true},
		{name: "single char prefix present", source: []string{"abc", "def"}, prefix: "a", expected: true},
		{name: "single char another prefix present", source: []string{"abc", "def"}, prefix: "d", expected: true},
		{name: "two char prefix", source: []string{"abc", "def"}, prefix: "ab", expected: true},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			var tree radixmap.MapRadixTree

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
		})
	}
}
