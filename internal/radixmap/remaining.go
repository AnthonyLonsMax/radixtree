package radixmap

// Remaining returns all stored words that have the given prefix.
func (r *MapRadixTree) Remaining(prefix string) []string {
	result := make([]string, 0)

	if r.root == nil {
		return result
	}

	buf := make([]byte, 0, initialBufferSize)

	if prefix == "" {
		walk(r.root, &buf, func(s string) {
			result = append(result, s)
		})

		return result
	}

	cursor := r.root
	remaining := prefix

	for {
		commonLen := commonPrefixLength(cursor.prefix, remaining)

		if commonLen == 0 && cursor.prefix != "" {
			return result
		}

		if commonLen == len(remaining) {
			walk(cursor, &buf, func(s string) {
				result = append(result, s)
			})

			return result
		}

		if commonLen < len(cursor.prefix) {
			return result
		}

		buf = append(buf, cursor.prefix...)

		child, ok := cursor.children[remaining[commonLen]]

		if !ok {
			return result
		}

		cursor = child
		remaining = remaining[commonLen:]
	}
}
