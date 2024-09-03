package _38

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	m := make(map[*Node]*Node)

	var reverse func(n *Node) *Node
	// go func函数变量的使用，用于在单个函数内进行递归
	// go func时，函数体后跟参数用于goroutine执行时传入

	reverse = func(n *Node) *Node {
		if n == nil {
			return nil
		}
		if n1, ok := m[n]; ok {
			return n1
		}

		n1 := &Node{
			Val:    n.Val,
			Next:   nil,
			Random: nil,
		}
		// 递归出口在于复制的某个节点
		m[n] = n1
		n1.Next = reverse(n.Next)
		n1.Random = reverse(n.Random)
		return n1
	}
	return reverse(head)
}
