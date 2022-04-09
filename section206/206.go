package section206

type ListNode struct {
	Val int
	Next *ListNode
}


func reverseList0(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var prev *ListNode = nil
	for head != nil {
		tmp := head.Next
		head.Next = prev
		prev = head
		head = tmp
	}
	return prev
}


func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}