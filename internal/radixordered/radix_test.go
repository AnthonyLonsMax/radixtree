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
		{
			name: "Test empty",
			sources: []string{
				"",
			},
		},
		{
			name: "Unicode words",
			sources: []string{
				"café", "résumé", "jalapeño", "naïve",
				"cœur", "über", "façade",
				"你好", "世界",
				"😀emoji",
			},
		},
	}
	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			tree := radixordered.RadixOrdered{}
			for _, src := range test.sources {
				if !tree.Add(src) {
					if src != "" {
						t.Fatalf("Element %s should be added to the tree", src)
					}
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
		{
			name: "Should contains with empty check",
			sources: []string{
				"worderland", "word", "worddy", "work", "", "worry",
				"wor", "worries", "wallet", "love", "", "lonnly",
				"lovers", "anthony", "ony", "anth",
			},
			shouldContains: []string{
				"worderland", "word", "worddy", "work", "worry",
			},
		},
		{
			name: "Unicode words",
			sources: []string{
				"café", "résumé", "jalapeño",
				"你好", "世界",
				"😀emoji", "🌍🌎🌏",
			},
			shouldContains: []string{
				"café", "résumé", "jalapeño",
				"你好", "世界",
				"😀emoji", "🌍🌎🌏",
			},
		},
		{
			name: "Unicode prefix sharing",
			sources: []string{"é", "ë"},
			shouldContains: []string{"é", "ë"},
		},
	}
	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			tree := radixordered.RadixOrdered{}
			for _, src := range test.sources {
				if !tree.Add(src) {
					if src != "" {
						t.Fatalf("Element %s should be added to the tree", src)
					}
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

func TestContainsWithExpected(t *testing.T) {
	type tt struct {
		name     string
		sources  []string
		find     string
		expected bool
	}
	tc := []tt{
		{
			name:     "Empty tree",
			find:     "word",
			expected: false,
		},
		{
			name:     "Should contains",
			sources:  []string{"word"},
			find:     "word",
			expected: true,
		},
		{
			name:     "No match",
			sources:  []string{"word", "source", "wordy"},
			find:     "worda",
			expected: false,
		},
		{
			name:     "Empty word",
			find:     "",
			expected: false,
		},
		{
			name:     "Splitted root",
			sources:  []string{"word", "match", "human"},
			find:     "fuzzy",
			expected: false,
		},
		{
			name:     "Partial prefix divergence",
			sources:  []string{"hello"},
			find:     "hey",
			expected: false,
		},
		{
			name:     "Unicode contains",
			sources:  []string{"café", "résumé"},
			find:     "café",
			expected: true,
		},
		{
			name:     "Unicode no match",
			sources:  []string{"café"},
			find:     "caféa",
			expected: false,
		},
		{
			name:     "Unicode partial prefix divergence",
			sources:  []string{"naïve"},
			find:     "naiv",
			expected: false,
		},
	}
	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			tree := radixordered.RadixOrdered{}
			for _, src := range test.sources {
				if !tree.Add(src) {
					if src != "" {
						t.Fatalf("Element %s should be added to the tree", src)
					}
				}
			}
			if exp := tree.Contains(test.find); exp != test.expected {
				t.Fatalf("Element find %s expect to be %v but got %v", test.find, test.expected, exp)
			}
		})
	}

}

func TestDelete(t *testing.T) {
	type tt struct {
		name        string
		sources     []string
		deleteItems []string
	}
	tc := []tt{
		{
			name: "Single source",
			sources: []string{
				"worderland", "word", "worddy", "work", "worry",
				"wor", "worries", "wallet", "love", "lonnly",
				"lovers", "anthony", "ony", "anth",
			},
			deleteItems: []string{
				"lovers", "anthony", "ony", "anth",
			},
		},
		{
			name:    "Empty tree",
			sources: []string{},
			deleteItems: []string{
				"lovers", "anthony", "ony", "anth",
			},
		},
		{
			name: "Partial match",
			sources: []string{
				"worderland", "word", "worddy", "work", "worry",
			},
			deleteItems: []string{
				"work", "worry",
			},
		},
		{
			name: "Diferent words",
			sources: []string{
				"worderland", "human", "root",
			},
			deleteItems: []string{
				"work", "factory",
			},
		},
		{
			name: "Common length 0",
			sources: []string{
				"word", "work",
			},
			deleteItems: []string{
				"worddy",
			},
		},
		{
			name: "Delete empty word",
			sources: []string{
				"worderland", "human", "root",
			},
			deleteItems: []string{
				"",
			},
		},
		{
			name: "Exact match",
			sources: []string{
				"worderland",
			},
			deleteItems: []string{
				"worderland",
			},
		},
		{
			name: "Delete terminal with multiple children",
			sources: []string{
				"test", "testing", "tested",
			},
			deleteItems: []string{
				"test",
			},
		},
		{
			name: "Delete already removed word from prefix node",
			sources: []string{
				"test", "testing", "tested",
			},
			deleteItems: []string{
				"test", "test",
			},
		},
		{
			name: "Unicode delete",
			sources: []string{
				"café", "cafeteria", "cafetal",
			},
			deleteItems: []string{
				"café",
			},
		},
		{
			name: "Unicode prefix sharing delete",
			sources: []string{"é", "ë"},
			deleteItems: []string{"é"},
		},
	}
	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			tree := radixordered.RadixOrdered{}
			for _, src := range test.sources {
				if !tree.Add(src) {
					if src != "" {
						t.Fatalf("Element %s should be added to the tree", src)
					}
				}
			}

			for _, src := range test.deleteItems {
				tree.Delete(src)
			}

			for _, src := range test.deleteItems {
				if tree.Contains(src) {
					t.Fatalf("Element %s should be not in the tree", src)
				}
			}
		})
	}

}

func TestStartWith(t *testing.T) {
	type tt struct {
		name     string
		sources  []string
		find     string
		expected bool
	}
	tc := []tt{
		{
			name:     "Empty tree",
			find:     "word",
			expected: false,
		},
		{
			name:     "Should contains",
			sources:  []string{"word"},
			find:     "word",
			expected: true,
		},
		{
			name:     "No match",
			sources:  []string{"word", "source", "wordy"},
			find:     "worda",
			expected: false,
		},
		{
			name:     "Empty word",
			find:     "",
			expected: false,
		},
		{
			name:     "Splitted root",
			sources:  []string{"word", "match", "human"},
			find:     "fuzzy",
			expected: false,
		},
		{
			name:     "Partial prefix divergence",
			sources:  []string{"hello"},
			find:     "hey",
			expected: false,
		},
		{
			name:     "Unicode starts with",
			sources:  []string{"café", "cafeteria", "cafetal"},
			find:     "café",
			expected: true,
		},
		{
			name:     "Unicode no starts with",
			sources:  []string{"café"},
			find:     "caféa",
			expected: false,
		},
		{
			name:     "Unicode prefix sharing starts with",
			sources:  []string{"é", "ë"},
			find:     "é",
			expected: true,
		},
	}
	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			tree := radixordered.RadixOrdered{}
			for _, src := range test.sources {
				if !tree.Add(src) {
					if src != "" {
						t.Fatalf("Element %s should be added to the tree", src)
					}
				}
			}
			if exp := tree.StartsWith(test.find); exp != test.expected {
				t.Fatalf("Element find %s expect to be %v but got %v", test.find, test.expected, exp)
			}
		})
	}

}
