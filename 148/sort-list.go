package _48

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	l, r := head, head
	for l != nil {
		r = l.Next
		for r != nil {
			if r.Val < l.Val {
				swap(l, r)
			}
			r = r.Next
		}
		l = l.Next
	}
	return head
}

func swap(l, r *ListNode) {
	i := l.Val
	l.Val = r.Val
	r.Val = i
}
