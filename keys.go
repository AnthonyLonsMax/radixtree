package radixtree

const (
	initialBufferSize = 256
	initialStackDepth = 64
)

type walkFrame struct {
	cursor   *node
	startLen int
}

func walk(cursor *node, buf *[]byte, yield func(string)) {
	if cursor == nil {
		return
	}

	stack := make([]walkFrame, 0, initialStackDepth)
	stack = append(stack, walkFrame{cursor: cursor, startLen: len(*buf)})

	for len(stack) > 0 {
		frame := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		*buf = (*buf)[:frame.startLen]
		*buf = append(*buf, frame.cursor.prefix...)

		if frame.cursor.isTerminal {
			yield(string(*buf))
		}

		childStartLen := len(*buf)

		for k := range frame.cursor.children {
			child := frame.cursor.children[k]
			stack = append(stack, walkFrame{cursor: child, startLen: childStartLen})
		}
	}
}

// Keys returns all stored words in the tree.
func (r *RadixTree) Keys() []string {
	result := make([]string, 0)

	if r.root == nil {
		return result
	}

	buf := make([]byte, 0, initialBufferSize)

	walk(r.root, &buf, func(s string) {
		result = append(result, s)
	})

	return result
}
