package _17

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

type queue []*Node

func (q *queue) push(n *Node) {
	if n != nil {
		*q = append(*q, n)
	}
}

func (q *queue) pop() *Node {
	if len(*q) == 0 {
		return nil
	}
	node := (*q)[0]
	*q = (*q)[1:]
	return node
}

func (q *queue) empty() bool {
	return len(*q) == 0
}

func connect(root *Node) *Node {
	var q queue
	q.push(root)
	for !q.empty() {
		var prev *Node
		size := len(q)
		for i := 0; i < size; i++ {
			t := q.pop()
			// q.pop()

			if prev != nil {
				prev.Next = t
			}
			prev = t

			if t.Left != nil {
				q.push(t.Left)
			}
			if t.Right != nil {
				q.push(t.Right)
			}
		}
	}
	return root
}
