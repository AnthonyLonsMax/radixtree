package radixmap_test

import (
	"testing"

	"github.com/AnthonyLonsMax/radixtree/internal/radixmap"
)

func TestDeleteWord(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name              string
		source            []string
		shouldNotContains []string
	}

	testCases := []testCase{
		{
			name: "delete multiple words with common prefixes",
			source: []string{
				"worderland", "word", "worddy", "work", "worry",
				"wor", "worries", "wallet", "love", "lonnly",
				"lovers", "anthony", "ony", "anth",
			},
			shouldNotContains: []string{
				"worddy", "work", "worry", "wor", "worries", "wallet", "love",
			},
		},
		{
			name:              "delete from empty tree",
			source:            []string{},
			shouldNotContains: []string{"hello"},
		},
		{
			name:              "delete non-existent word",
			source:            []string{"hello", "world"},
			shouldNotContains: []string{"nonexistent"},
		},
		{
			name:              "delete all words",
			source:            []string{"hello", "world"},
			shouldNotContains: []string{"hello", "world"},
		},
		{
			name:              "delete words with shared prefixes",
			source:            []string{"test", "testing", "tester"},
			shouldNotContains: []string{"test", "testing", "tester"},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			var tree radixmap.MapRadixTree

			for _, e := range test.source {
				tree.Add(e)
			}

			for _, e := range test.shouldNotContains {
				tree.Delete(e)

				if tree.Contains(e) {
					t.Fatalf("Word %s should not be in the tree", e)
				}
			}
		})
	}
}
