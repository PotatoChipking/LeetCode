package _3

type ListNode struct {
	Val  int
	Next *ListNode
}

func fuckMe(node1, node2 *ListNode) *ListNode {
	ans := &ListNode{}
	for node1 != nil && node2 != nil {
		if node1.Val < node2.Val {
			ans.Next = node1
			node1 = node1.Next
		} else {
			ans.Next = node2
			node2 = node2.Next
		}
	}
	if node1 != nil {
		ans.Next = node1
	}
	if node2 != nil {
		ans.Next = node2
	}
	return ans.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	l := len(lists)

	if l == 0 {
		return nil
	}
	if l == 1 {
		return lists[0]
	}
	left := mergeKLists(lists[:l/2])
	right := mergeKLists(lists[l/2:])
	fuckMe(left, right)
}
