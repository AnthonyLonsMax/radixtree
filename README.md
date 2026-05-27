# radixtree

A radix tree (compressed prefix tree) implementation in Go. Zero external dependencies.

## Features

- Insert, delete, and search words in O(k) time (k = key length)
- Longest prefix matching
- Prefix existence checks (StartsWith)
- Autocomplete / prefix-based word retrieval (Remaining)
- Minimum / Maximum lexicographic word retrieval
- Common prefix computation across all stored words
- Iteration over all keys (ForEach, Keys)
- No external dependencies (stdlib only)

## Install

```
go get github.com/AnthonyLonsMax/radixtree
```

## Usage

```go
import "github.com/AnthonyLonsMax/radixtree"

tree := radixtree.RadixTree{}

tree.Add("hello")
tree.Add("world")
tree.Add("hell")

fmt.Println(tree.Contains("hello")) // true
fmt.Println(tree.Contains("world")) // true
fmt.Println(tree.Contains("xyz"))   // false

fmt.Println(tree.StartsWith("he"))  // true

fmt.Println(tree.LongestPrefixOf("helloworld")) // "hello"

fmt.Println(tree.CommonPrefix()) // ""

fmt.Println(tree.Minimum()) // "hell"
fmt.Println(tree.Maximum()) // "world"

fmt.Println(tree.Remaining("hel")) // ["hell", "hello"]

tree.Delete("world")

tree.ForEach(func(key string) {
    fmt.Println(key)
})
```

## API

| Method | Description |
|---|---|
| `Add(word string) bool` | Inserts a word. Returns true if newly added. |
| `Delete(word string) bool` | Removes a word. Returns true if it existed. |
| `Contains(word string) bool` | Returns true if the word is in the tree. |
| `StartsWith(prefix string) bool` | Returns true if any word has the given prefix. |
| `LongestPrefixOf(word string) string` | Returns the longest stored word that is a prefix of the input. |
| `CommonPrefix() string` | Returns the longest common prefix among all stored words. |
| `ForEach(fn func(key string))` | Calls fn for every word in the tree (DFS order). |
| `Keys() []string` | Returns a slice of all words. |
| `Size() int64` | Returns the number of stored words. |
| `Minimum() string` | Returns the lexicographically smallest word (empty string if tree is empty). |
| `Maximum() string` | Returns the lexicographically largest word (empty string if tree is empty). |
| `Remaining(prefix string) []string` | Returns all words that start with the given prefix. |
| `Clear()` | Removes all words from the tree. |

## Considerations

### Thread safety

This library is **not thread-safe**. Concurrent reads and writes from multiple goroutines
without external synchronization will cause a data race. If you need concurrent access,
protect the tree with a `sync.RWMutex` or use a dedicated concurrent data structure.

### Set semantics

The tree tracks word *presence* only — it has no key-value storage. Use it when you
need to know *whether* a word exists, not when you need to associate data with it.

### When to use a radix tree vs a `map[string]bool`

| Use case | Recommended |
|---|---|
| Fast exact lookup, no prefix ops | `map[string]bool` |
| Prefix matching / autocomplete | Radix tree |
| Longest common prefix | Radix tree |
| Lexicographic ordering | Radix tree |
| Simplicity | `map[string]bool` |

### Compared to other radix tree libraries

- **hashicorp/go-immutable-radix**: Immutable, concurrent-safe via copy-on-write,
  supports key-value storage, but has external dependencies and higher allocation
  overhead per write.
- **This library**: Simple mutable API, zero dependencies, minimal allocation
  overhead, set-only semantics. Best when you need a lightweight prefix tree.

### Memory

Each node carries a `map[byte]*node` for children, so per-node overhead is higher
than a raw trie, but fewer nodes exist due to prefix compression. The tree is
optimized for string keys of moderate length.

### Character encoding

Keys are treated as raw byte sequences. Any valid Go string is accepted, including
arbitrary binary data and multi-byte UTF-8. Comparisons are lexicographic by byte
value — "café" and "cafè" are distinct keys.

### Iterative traversal (no recursion limit)

`ForEach`, `Keys`, and `Remaining` use an **iterative DFS** with an explicit stack
instead of recursion. This means trees with thousands of nested levels won't cause
a stack overflow. The shared `[]byte` buffer with save/truncate avoids O(n²) string
allocations — each node appends its prefix to the buffer and truncates after its
subtree is processed, reusing the underlying array across the entire traversal.

## Benchmarks

```
go test -bench=. -benchmem
```

Results on a typical machine:

| Operation | Size | Time |
|---|---|---|
| Add | 10 | ~XXX ns/op |
| Add | 100 | ~XXX ns/op |
| Add | 1000 | ~XXX ns/op |
| Contains | 1000 | ~XXX ns/op |
| Delete | 1000 | ~XXX ns/op |
| ForEach | 1000 | ~XXX ns/op |

## Tests

```
# Unit + extended tests
go test -v -count 1 ./...

# Fuzz test (run for 10 seconds)
go test -fuzz FuzzRadixTree -fuzztime 10s

# Benchmarks
go test -bench=. -benchmem ./...
```

## License

MIT
