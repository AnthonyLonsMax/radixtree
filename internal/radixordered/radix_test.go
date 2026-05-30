package radixordered_test

import (
	"testing"

	"github.com/AnthonyLonsMax/radixtree/internal/radixordered"
)

func TestAddFunction(t *testing.T) {
	type tt struct {
		name    string
		sources []string
	}
	tc := []tt{
		{
			name: "Single source",
			sources: []string{
				"worderland", "word", "worddy", "work", "worry",
				"wor", "worries", "wallet", "love", "lonnly",
				"lovers", "anthony", "ony", "anth",
			},
		},
	}
	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			tree := radixordered.RadixOrdered{}
			for _, src := range test.sources {
				if !tree.Add(src) {
					t.Fatalf("Element %s should be added to the tree", src)
				}
			}
		})
	}

}

func TestContains(t *testing.T) {
	type tt struct {
		name           string
		sources        []string
		shouldContains []string
	}
	tc := []tt{
		{
			name: "Should contains all the elements",
			sources: []string{
				"worderland", "word", "worddy", "work", "worry",
				"wor", "worries", "wallet", "love", "lonnly",
				"lovers", "anthony", "ony", "anth",
			},
			shouldContains: []string{
				"worderland", "word", "worddy", "work", "worry",
				"wor", "worries", "wallet", "love", "lonnly",
				"lovers", "anthony", "ony", "anth",
			},
		},
	}
	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			tree := radixordered.RadixOrdered{}
			for _, src := range test.sources {
				if !tree.Add(src) {
					t.Fatalf("Element %s should be added to the tree", src)
				}
			}

			for _, src := range test.shouldContains {
				if !tree.Contains(src) {
					t.Fatalf("Element %s should be contains to the tree", src)
				}
			}
		})
	}

}
