package radixtree_test

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"

	"github.com/AnthonyLonsMax/radixtree"
)

func BenchmarkAdd(b *testing.B) {
	for _, size := range []int{10, 100, 1000, 10000} {
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			words := generateWords(size, 10)
			b.ResetTimer()

			for idx := range b.N {
				b.StopTimer()
				tree := radixtree.MapRadixTree{}
				start := (idx * size) % len(words)
				b.StartTimer()

				for j := range size {
					tree.Add(words[(start+j)%len(words)])
				}
			}
		})
	}
}

func BenchmarkContains(b *testing.B) {
	for _, size := range []int{10, 100, 1000, 10000} {
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			words := generateWords(size, 10)
			tree := radixtree.MapRadixTree{}

			for _, w := range words {
				tree.Add(w)
			}

			b.ResetTimer()

			for idx := range b.N {
				tree.Contains(words[idx%len(words)])
			}
		})
	}
}

func BenchmarkContainsNegative(b *testing.B) {
	for _, size := range []int{10, 100, 1000, 10000} {
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			words := generateWords(size, 10)
			tree := radixtree.MapRadixTree{}

			for _, w := range words {
				tree.Add(w)
			}

			b.ResetTimer()

			for range b.N {
				tree.Contains("nonexistent")
			}
		})
	}
}

func BenchmarkDelete(b *testing.B) {
	for _, size := range []int{10, 100, 1000, 10000} {
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			words := generateWords(size, 10)
			b.ResetTimer()

			for idx := range b.N {
				b.StopTimer()
				tree := radixtree.MapRadixTree{}

				for _, w := range words {
					tree.Add(w)
				}

				b.StartTimer()
				tree.Delete(words[idx%len(words)])
			}
		})
	}
}

func BenchmarkForEach(b *testing.B) {
	for _, size := range []int{10, 100, 1000, 10000} {
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			words := generateWords(size, 10)
			tree := radixtree.MapRadixTree{}

			for _, w := range words {
				tree.Add(w)
			}

			b.ResetTimer()

			for range b.N {
				tree.ForEach(func(string) {})
			}
		})
	}
}

func BenchmarkKeys(b *testing.B) {
	for _, size := range []int{10, 100, 1000, 10000} {
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			words := generateWords(size, 10)
			tree := radixtree.MapRadixTree{}

			for _, w := range words {
				tree.Add(w)
			}

			b.ResetTimer()

			for range b.N {
				tree.Keys()
			}
		})
	}
}

func BenchmarkLongestPrefixOf(b *testing.B) {
	for _, size := range []int{10, 100, 1000, 10000} {
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			words := generateWords(size, 10)
			tree := radixtree.MapRadixTree{}

			for _, w := range words {
				tree.Add(w)
			}

			query := strings.Repeat("a", 15)
			b.ResetTimer()

			for range b.N {
				tree.LongestPrefixOf(query)
			}
		})
	}
}

func BenchmarkStartsWith(b *testing.B) {
	for _, size := range []int{10, 100, 1000, 10000} {
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			words := generateWords(size, 10)
			tree := radixtree.MapRadixTree{}

			for _, w := range words {
				tree.Add(w)
			}

			b.ResetTimer()

			for range b.N {
				tree.StartsWith("a")
			}
		})
	}
}

func BenchmarkCommonPrefix(b *testing.B) {
	for _, size := range []int{10, 100, 1000, 10000} {
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			words := make([]string, 0, size)
			base := "commonprefix"

			for i := range size {
				words = append(words, fmt.Sprintf("%s_%d", base, i))
			}

			tree := radixtree.MapRadixTree{}

			for _, w := range words {
				tree.Add(w)
			}

			b.ResetTimer()

			for range b.N {
				tree.CommonPrefix()
			}
		})
	}
}

func BenchmarkAddSequentialPrefixes(b *testing.B) {
	b.Run("size=1000", func(b *testing.B) {
		for range b.N {
			tree := radixtree.MapRadixTree{}

			for i := range 1000 {
				tree.Add(strings.Repeat("a", i+1))
			}
		}
	})
}

func BenchmarkMixedOperations(b *testing.B) {
	words := generateWords(1000, 10)
	b.ResetTimer()

	for range b.N {
		tree := radixtree.MapRadixTree{}

		for _, w := range words {
			tree.Add(w)
		}

		for _, w := range words[:100] {
			tree.Contains(w)
		}

		for _, w := range words[:50] {
			tree.Delete(w)
		}
	}
}

func generateWords(n, maxLen int) []string { //nolint:unparam
	rng := rand.New(rand.NewSource(42)) //nolint:gosec
	words := make([]string, 0, n)
	seen := make(map[string]bool)

	for len(words) < n {
		w := randomWord(rng, 1, maxLen)

		if !seen[w] {
			seen[w] = true
			words = append(words, w)
		}
	}

	return words
}
