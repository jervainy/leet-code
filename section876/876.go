package section876


type ListNode struct {
	Val int
	Next *ListNode
}


func middleNode0(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var (
		list []*ListNode
		cur = head
	)
	for cur != nil {
		list = append(list, cur)
		cur = cur.Next
	}
	return list[len(list) / 2]
}


func middleNode(head *ListNode) *ListNode {
	fast, slow, n := head, head, 1
	for fast != nil {
		if n % 2 == 0 {
			slow = slow.Next
		}
		fast = fast.Next
		n += 1
	}
	return slow
}

