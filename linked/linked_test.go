package linked

import "testing"

func TestNode_ReverseAdjacentNode(t *testing.T) {
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}
	node5 := &Node{Val: 5}
	node6 := &Node{Val: 6}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6

	ReverseAdjacentNode(node1)

	node1.Print()
}
