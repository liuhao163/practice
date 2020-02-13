package linked

import "fmt"

type Node struct {
	Next *Node
	Val  int
}

func (n *Node) Print() {
	haed := n
	for haed != nil {
		fmt.Printf("%v ", haed.Val)
		haed = haed.Next
	}
}

func ReverseAdjacentNode(head *Node) {
	if head == nil || head.Next == nil {
		return
	}

	tmpV := head.Val
	head.Val = head.Next.Val
	head.Next.Val = tmpV

	if head.Next.Next != nil {
		ReverseAdjacentNode(head.Next.Next)
	}
}
