package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}



type Item struct {
	node1 *TreeNode
	node2 *TreeNode
	parent *TreeNode
	direction int
}

type Data struct {
	data []*Item
}




func (t *Data) pushQueue(node1, node2, parent *TreeNode, direction int) {
	t.data = append(t.data, &Item{node1: node1, node2: node2, parent: parent, direction: direction})
}

func (t *Data) popQueue() (*TreeNode, *TreeNode, *TreeNode, int) {
	length := len(t.data)
	if length == 0 {
		panic(length)
	}
	a := t.data[0]
	t.data = t.data[1:length]
	return a.node1, a.node2, a.parent, a.direction
}

func (t *Data) isQueueEmpty() bool {
	return len(t.data) == 0
}



func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	bfs(root1, root2)
	return root1
}

func bfs(node1 *TreeNode, node2 *TreeNode) {
	st := &Data{}
	var (
		parent    *TreeNode = nil
		direction           = 0
	)
	for {
		var (
			n1l *TreeNode = nil
			n1r *TreeNode = nil
			n2l *TreeNode = nil
			n2r *TreeNode = nil
			sum = 0
		)
		if node1 != nil {
			n1l = node1.Left
			n1r = node1.Right
			sum += node1.Val
		}
		if node2 != nil {
			n2l = node2.Left
			n2r = node2.Right
			sum += node2.Val
		}
		if node1 == nil {
			node1 = &TreeNode{}
		}
		node1.Val = sum
		if parent != nil {
			if direction == 0 {
				parent.Left = node1
			} else {
				parent.Right = node1
			}
		}
		if n1l != nil || n2l != nil {
			st.pushQueue(n1l, n2l, node1, 0)
		}
		if n1r != nil || n2r != nil {
			st.pushQueue(n1r, n2r, node1, 1)
		}
		if !st.isQueueEmpty() {
			node1, node2, parent, direction = st.popQueue()
		} else {
			break
		}
	}
}

func main() {
	root1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 2,
		},
	}
	root2 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	mergeTrees(root1, root2)
	println(root1)
}
