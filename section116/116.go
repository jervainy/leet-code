package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}




type Data struct {
	data []*Node
}

func (t *Data) pushQueue(node *Node) {
	t.data = append(t.data, node)
}

func (t *Data) popQueue() *Node {
	length := len(t.data)
	if length == 0 {
		panic(length)
	}
	a := t.data[0]
	t.data = t.data[1:length]
	return a
}

func (t *Data) isQueueEmpty() bool {
	return len(t.data) == 0
}


func connect0(root *Node) *Node {
	if root == nil {
		return root
	}
	bfs(root)
	return root
}


func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	for leftmost := root; leftmost.Left != nil; leftmost = leftmost.Left {
		for node := leftmost; node != nil; node = node.Next {
			node.Left.Next = node.Right
			if node.Next != nil {
				node.Right.Next = node.Next.Left
			}
		}
	}
	return root
}




func bfs(node *Node) {
	var (
		st   = &Data{}
		list []*Node
	)
	st.pushQueue(node)
	for !st.isQueueEmpty() {
		list = append(list, st.popQueue())
		if st.isQueueEmpty() && len(list) > 0 {
			for i, l := 0, len(list); i < l; i++ {
				c := list[i]
				if i < l - 1 {
					c.Next = list[i + 1]
				}
				if c.Left != nil {
					st.pushQueue(c.Left)
				}
				if c.Right != nil {
					st.pushQueue(c.Right)
				}
			}
			list = list[0:0]
		}
	}
}

func main() {
	root := &Node{
		Val: 1,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val: 4,
			},
			Right: &Node{
				Val: 5,
			},
		},
		Right: &Node{
			Val: 3,
			Left: &Node{
				Val: 6,
			},
			Right: &Node{
				Val: 7,
			},
		},
	}

	connect(root)
	println(root)
}

