package radixmap_test

import (
	"math/rand"
	"slices"
	"strings"
	"testing"

	"github.com/AnthonyLonsMax/radixtree/internal/radixmap"
)

func TestClear(t *testing.T) {
	t.Parallel()

	tree := radixmap.MapRadixTree{}
	tree.Add("wordHello")
	tree.Add("world")
	tree.Clear()

	if tree.Size() != 0 {
		t.Fatalf("Expected size 0 after Clear, got %d", tree.Size())
	}

	if tree.Contains("wordHello") {
		t.Fatal("Tree should not contain hello after Clear")
	}

	if len(tree.Keys()) != 0 {
		t.Fatal("Keys should be empty after Clear")
	}
}

func TestAddReturnsBool(t *testing.T) {
	t.Parallel()

	tree := radixmap.MapRadixTree{}

	if !tree.Add("wordHello") {
		t.Fatal("Add should return true for new word")
	}

	if tree.Add("wordHello") {
		t.Fatal("Add should return false for duplicate")
	}

	if !tree.Add("") {
		t.Fatal("Add should return true for new empty string")
	}

	if tree.Add("") {
		t.Fatal("Add should return false for duplicate empty string")
	}
}

func TestDeleteReturnsBool(t *testing.T) {
	t.Parallel()

	tree := radixmap.MapRadixTree{}
	tree.Add("wordHello")

	if !tree.Delete("wordHello") {
		t.Fatal("Delete should return true for existing word")
	}

	if tree.Delete("wordHello") {
		t.Fatal("Delete should return false for already deleted word")
	}

	if tree.Delete("nonexistent") {
		t.Fatal("Delete should return false for nonexistent word")
	}
}

func TestAddDeleteRoundtrip(t *testing.T) {
	t.Parallel()

	tree := radixmap.MapRadixTree{}
	words := []string{"a", "wordAb", "wordAbc", "wordAbcd", "abcde"}

	for _, w := range words {
		tree.Add(w)
	}

	if int(tree.Size()) != len(words) {
		t.Fatalf("Expected size %d, got %d", len(words), tree.Size())
	}

	slices.Reverse(words)

	for _, w := range words {
		tree.Delete(w)
	}

	if tree.Size() != 0 {
		t.Fatalf("Expected size 0, got %d", tree.Size())
	}

	if len(tree.Keys()) != 0 {
		t.Fatal("Keys should be empty after deleting all")
	}
}

func TestLargeDataSet(t *testing.T) {
	t.Parallel()

	const count = 10000
	tree := radixmap.MapRadixTree{}

	for i := range count {
		word := strings.Repeat("a", i+1)
		tree.Add(word)
	}

	if int(tree.Size()) != count {
		t.Fatalf("Expected size %d, got %d", count, tree.Size())
	}

	for i := range count {
		word := strings.Repeat("a", i+1)

		if !tree.Contains(word) {
			t.Fatalf("Should contain word of length %d", i+1)
		}
	}

	for i := range count {
		word := strings.Repeat("a", i+1)
		tree.Delete(word)
	}

	if tree.Size() != 0 {
		t.Fatalf("Expected size 0 after deleting all, got %d", tree.Size())
	}
}

func TestRandomOperations(t *testing.T) {
	t.Parallel()

	rng := rand.New(rand.NewSource(42)) //nolint:gosec
	tree := radixmap.MapRadixTree{}
	reference := make(map[string]bool)

	ops := 5000

	for range ops {
		word := randomWord(rng, 1, 20)

		switch rng.Intn(3) {
		case 0:
			tree.Add(word)
			reference[word] = true
		case 1:
			tree.Delete(word)
			delete(reference, word)
		case 2:
			tree.Contains(word)
		}
	}

	keys := tree.Keys()

	if len(keys) != int(tree.Size()) {
		t.Fatalf("Keys length %d != size %d", len(keys), tree.Size())
	}

	for k := range reference {
		if !tree.Contains(k) {
			t.Fatalf("Reference word %q missing from tree", k)
		}
	}

	for _, k := range keys {
		if !reference[k] {
			t.Fatalf("Tree key %q not in reference", k)
		}
	}
}

func TestForEachEmptyTree(t *testing.T) {
	t.Parallel()

	tree := radixmap.MapRadixTree{}
	count := 0
	tree.ForEach(func(string) { count++ })

	if count != 0 {
		t.Fatalf("Expected 0, got %d", count)
	}
}

func TestKeysEmptyTree(t *testing.T) {
	t.Parallel()

	tree := radixmap.MapRadixTree{}
	keys := tree.Keys()

	if len(keys) != 0 {
		t.Fatal("Keys of empty tree should be empty slice")
	}
}

func TestLongestPrefixOfOnEmptyTree(t *testing.T) {
	t.Parallel()

	tree := radixmap.MapRadixTree{}

	if r := tree.LongestPrefixOf("anything"); r != "" {
		t.Fatalf("Expected empty, got %q", r)
	}
}

func TestStartsWithOnEmptyTree(t *testing.T) {
	t.Parallel()

	tree := radixmap.MapRadixTree{}

	if tree.StartsWith("a") {
		t.Fatal("Empty tree should not start with any prefix")
	}
}

func TestCommonPrefixOnSingleChar(t *testing.T) {
	t.Parallel()

	tree := radixmap.MapRadixTree{}
	tree.Add("x")
	tree.Add("y")

	if p := tree.CommonPrefix(); p != "" {
		t.Fatalf("Expected empty common prefix, got %q", p)
	}
}

func TestAddDeleteInterleaved(t *testing.T) {
	t.Parallel()

	tree := radixmap.MapRadixTree{}
	words := []string{"wordTest", "wordTesting", "tester", "tested"}

	for _, w := range words {
		tree.Add(w)
	}

	tree.Delete("wordTesting")
	tree.Add("wordTesting")

	if !tree.Contains("wordTesting") {
		t.Fatal("testing should exist after re-add")
	}

	if int(tree.Size()) != len(words) {
		t.Fatalf("Expected size %d, got %d", len(words), tree.Size())
	}

	tree.Delete("wordTest")
	tree.Delete("tester")
	tree.Add("wordTest")

	if int(tree.Size()) != 3 {
		t.Fatalf("Expected size 3, got %d", tree.Size())
	}
}

func TestUnicodeWords(t *testing.T) {
	t.Parallel()

	tree := radixmap.MapRadixTree{}
	words := []string{"日本語", "日本", "にほんご", "日本語テスト", "café", "cafè", "cafeteria"}

	for _, w := range words {
		tree.Add(w)
	}

	for _, w := range words {
		if !tree.Contains(w) {
			t.Fatalf("Should contain %q", w)
		}
	}

	if int(tree.Size()) != len(words) {
		t.Fatalf("Expected size %d, got %d", len(words), tree.Size())
	}

	keys := tree.Keys()

	if len(keys) != len(words) {
		t.Fatalf("Expected %d keys, got %d", len(words), len(keys))
	}

	for _, w := range words {
		if !slices.Contains(keys, w) {
			t.Fatalf("Key %q missing from Keys result", w)
		}
	}
}

func TestCommonPrefixEdgeCases(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		words    []string
		expected string
	}{
		{name: "all same", words: []string{"wordAbc", "wordAbc", "wordAbc"}, expected: "wordAbc"},
		{name: "nested prefixes", words: []string{"a", "wordAb", "wordAbc"}, expected: "a"},
		{name: "single char differ", words: []string{"a", "b"}, expected: ""},
		{name: "long common", words: []string{"prefix", "prefixed", "prefixes"}, expected: "prefix"},
		{name: "empty tree", words: []string{}, expected: ""},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			tree := radixmap.MapRadixTree{}

			for _, w := range test.words {
				tree.Add(w)
			}

			if p := tree.CommonPrefix(); p != test.expected {
				t.Fatalf("Expected %q, got %q", test.expected, p)
			}
		})
	}
}

func TestStartsWithEdgeCases(t *testing.T) {
	t.Parallel()

	tree := radixmap.MapRadixTree{}
	tree.Add("wordAbc")

	cases := []struct {
		prefix   string
		expected bool
	}{
		{"a", true},
		{"wordAb", true},
		{"wordAbc", true},
		{"wordAbcd", false},
		{"", true},
		{"xyz", false},
		{"abC", false},
	}

	for _, c := range cases {
		t.Run(c.prefix, func(t *testing.T) {
			t.Parallel()

			if r := tree.StartsWith(c.prefix); r != c.expected {
				t.Fatalf("StartsWith(%q) = %v, want %v", c.prefix, r, c.expected)
			}
		})
	}
}

func TestLongestPrefixOfEdgeCases(t *testing.T) {
	t.Parallel()

	tree := radixmap.MapRadixTree{}
	tree.Add("a")
	tree.Add("wordAb")
	tree.Add("wordAbc")

	cases := []struct {
		word     string
		expected string
	}{
		{"a", "a"},
		{"wordAb", "wordAb"},
		{"wordAbc", "wordAbc"},
		{"wordAbcd", "wordAbc"},
		{"b", ""},
		{"", ""},
	}

	for _, c := range cases {
		t.Run(c.word, func(t *testing.T) {
			t.Parallel()

			if r := tree.LongestPrefixOf(c.word); r != c.expected {
				t.Fatalf("LongestPrefixOf(%q) = %q, want %q", c.word, r, c.expected)
			}
		})
	}
}

func TestForEachOrder(t *testing.T) {
	t.Parallel()

	tree := radixmap.MapRadixTree{}
	words := []string{"b", "a", "c", "wordAb", "ac"}

	for _, w := range words {
		tree.Add(w)
	}

	visited := make([]string, 0)
	tree.ForEach(func(key string) {
		visited = append(visited, key)
	})

	slices.Sort(words)
	slices.Sort(visited)

	if !slices.Equal(words, visited) {
		t.Fatalf("ForEach visited %v, expected %v", visited, words)
	}
}

func TestMinimumEmptyTree(t *testing.T) {
	t.Parallel()

	tree := radixmap.MapRadixTree{}

	if m := tree.Minimum(); m != "" {
		t.Fatalf("Expected empty, got %q", m)
	}
}

func TestMinimumSingleWord(t *testing.T) {
	t.Parallel()

	tree := radixmap.MapRadixTree{}
	tree.Add("wordHello")

	if m := tree.Minimum(); m != "wordHello" {
		t.Fatalf("Expected %q, got %q", "wordHello", m)
	}
}

func TestMinimumMultipleWords(t *testing.T) {
	t.Parallel()

	tree := radixmap.MapRadixTree{}
	tree.Add("zebra")
	tree.Add("apple")
	tree.Add("mango")

	if m := tree.Minimum(); m != "apple" {
		t.Fatalf("Expected \"apple\", got %q", m)
	}
}

func TestMinimumWithCommonPrefix(t *testing.T) {
	t.Parallel()

	tree := radixmap.MapRadixTree{}
	tree.Add("wordAbcd")
	tree.Add("wordAbc")
	tree.Add("wordAb")

	if m := tree.Minimum(); m != "wordAb" {
		t.Fatalf("Expected %q, got %q", "wordAb", m)
	}
}

func TestMinimumWithEmptyString(t *testing.T) {
	t.Parallel()

	tree := radixmap.MapRadixTree{}
	tree.Add("")
	tree.Add("wordHello")

	if m := tree.Minimum(); m != "" {
		t.Fatalf("Expected \"\", got %q", m)
	}
}

func TestMaximumEmptyTree(t *testing.T) {
	t.Parallel()

	tree := radixmap.MapRadixTree{}

	if m := tree.Maximum(); m != "" {
		t.Fatalf("Expected empty, got %q", m)
	}
}

func TestMaximumSingleWord(t *testing.T) {
	t.Parallel()

	tree := radixmap.MapRadixTree{}
	tree.Add("wordHello")

	if m := tree.Maximum(); m != "wordHello" {
		t.Fatalf("Expected %q, got %q", "wordHello", m)
	}
}

func TestMaximumMultipleWords(t *testing.T) {
	t.Parallel()

	tree := radixmap.MapRadixTree{}
	tree.Add("apple")
	tree.Add("zebra")
	tree.Add("mango")

	if m := tree.Maximum(); m != "zebra" {
		t.Fatalf("Expected \"zebra\", got %q", m)
	}
}

func TestMaximumWithCommonPrefix(t *testing.T) {
	t.Parallel()

	tree := radixmap.MapRadixTree{}
	tree.Add("wordAb")
	tree.Add("wordAbc")
	tree.Add("wordAbcd")

	if m := tree.Maximum(); m != "wordAbcd" {
		t.Fatalf("Expected %q, got %q", "wordAbcd", m)
	}
}

func TestMaximumWithEmptyString(t *testing.T) {
	t.Parallel()

	tree := radixmap.MapRadixTree{}
	tree.Add("wordHello")
	tree.Add("")

	if m := tree.Maximum(); m != "wordHello" {
		t.Fatalf("Expected %q, got %q", "wordHello", m)
	}
}

func TestRemainingEmptyTree(t *testing.T) {
	t.Parallel()

	tree := radixmap.MapRadixTree{}

	if r := tree.Remaining("abc"); len(r) != 0 {
		t.Fatalf("Expected empty, got %v", r)
	}
}

func TestRemainingEmptyPrefix(t *testing.T) {
	t.Parallel()

	tree := radixmap.MapRadixTree{}
	words := []string{"wordAb", "wordAbc", "wordAbcd", "wordHello"}

	for _, w := range words {
		tree.Add(w)
	}

	r := tree.Remaining("")
	slices.Sort(r)
	slices.Sort(words)

	if !slices.Equal(r, words) {
		t.Fatalf("Expected %v, got %v", words, r)
	}
}

func TestRemainingMatchingPrefix(t *testing.T) {
	t.Parallel()

	tree := radixmap.MapRadixTree{}
	tree.Add("wordAb")
	tree.Add("wordAbc")
	tree.Add("wordAbcd")
	tree.Add("wordHello")

	got := tree.Remaining("wordAb")
	slices.Sort(got)

	expected := []string{"wordAb", "wordAbc", "wordAbcd"}

	if !slices.Equal(got, expected) {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
}

func TestRemainingNoMatch(t *testing.T) {
	t.Parallel()

	tree := radixmap.MapRadixTree{}
	tree.Add("wordAbc")

	got := tree.Remaining("xyz")

	if len(got) != 0 {
		t.Fatalf("Expected empty, got %v", got)
	}
}

func TestRemainingExactMatch(t *testing.T) {
	t.Parallel()

	tree := radixmap.MapRadixTree{}
	tree.Add("wordAbc")
	tree.Add("wordAbcd")

	result := tree.Remaining("wordAbc")

	if len(result) != 2 || !slices.Contains(result, "wordAbc") || !slices.Contains(result, "wordAbcd") {
		t.Fatalf("Expected [%q %q], got %v", "wordAbc", "wordAbcd", result)
	}
}

func TestRemainingDeepPrefix(t *testing.T) {
	t.Parallel()

	tree := radixmap.MapRadixTree{}
	tree.Add("abcdef")
	tree.Add("abcxyz")
	tree.Add("wordHello")

	got := tree.Remaining("abcd")
	slices.Sort(got)

	expected := []string{"abcdef"}

	if !slices.Equal(got, expected) {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
}

func TestRemainingWithEmptyString(t *testing.T) {
	t.Parallel()

	tree := radixmap.MapRadixTree{}
	tree.Add("")
	tree.Add("wordHello")

	got := tree.Remaining("")

	if len(got) != 2 || !slices.Contains(got, "") || !slices.Contains(got, "wordHello") {
		t.Fatalf("Expected [\"\" %q], got %v", "wordHello", got)
	}
}

func randomWord(rng *rand.Rand, minLen, maxLen int) string {
	n := rng.Intn(maxLen-minLen+1) + minLen

	var b strings.Builder

	for range n {
		b.WriteByte(byte('a' + rng.Intn(26))) //nolint:gosec
	}

	return b.String()
}

func FuzzRadixTree(f *testing.F) {
	seeds := []string{"a", "wordAb", "wordAbc", "wordHello", "world", "wordTest", "", "x"}

	for _, s := range seeds {
		f.Add(s, s, int32(0))
	}

	f.Fuzz(func(t *testing.T, addWord, delWord string, _ int32) {
		tree := radixmap.MapRadixTree{}

		tree.Add(addWord)

		if !tree.Contains(addWord) {
			t.Skip()
		}

		tree.Delete(delWord)

		if tree.Contains(delWord) {
			t.Skip()
		}

		tree.Add(addWord)

		if tree.Size() < 1 {
			t.Skip()
		}

		tree.ForEach(func(string) {})
		_ = tree.Keys()
		_ = tree.CommonPrefix()
		_ = tree.LongestPrefixOf(addWord)
		_ = tree.StartsWith(addWord[:minInt(1, len(addWord))])
		_ = tree.Minimum()
		_ = tree.Maximum()
		_ = tree.Remaining(addWord[:minInt(1, len(addWord))])
	})
}

func minInt(a, b int) int {
	if a < b {
		return a
	}

	return b
}
