# radixtree

A radix tree (compressed prefix tree) implementation in Go. Zero external dependencies.

## Features

- Insert, delete, and search words in O(k) time (k = key length)
- Longest prefix matching
- Prefix existence checks (StartsWith)
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
| `Keys() *[]string` | Returns a pointer to a slice of all words. |
| `Size() int64` | Returns the number of stored words. |
| `Clear()` | Removes all words from the tree. |

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
