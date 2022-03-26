package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func removeNthFromEnd0(head *ListNode, n int) *ListNode {
	var list []*ListNode
	cur := head
	for cur != nil {
		list = append(list, cur)
		cur = cur.Next
	}
	count := len(list)
	if count == 0 {
		return nil
	}
	if count - n <= 0 {
		head = list[0].Next
	} else {
		node := list[count - n - 1]
		node.Next = node.Next.Next
	}
	return head
}


func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var (
		fast = head
		slow *ListNode = nil
	)
	for fast != nil {
		if n == 0 {
			if slow == nil {
				slow = head
			} else {
				slow = slow.Next
			}
		}
		fast = fast.Next
		if n > 0 {
			n--
		}
	}
	if slow == nil {
		head = head.Next
	} else {
		slow.Next = slow.Next.Next
	}
	return head
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: nil,
					},
				},
			},
		},
	}
	removeNthFromEnd(head, 2)
	println(head)
}

