package main

import "fmt"

type Node struct {
	value string
	left  *Node
	right *Node
}

/**
		a
	b		c
   d e     f g
*/

func InitTree() Node {
	a := Node{value: "A"}
	b := Node{value: "B"}
	c := Node{value: "C"}
	d := Node{value: "D"}
	e := Node{value: "E"}
	f := Node{value: "F"}
	g := Node{value: "G"}

	a.left = &b
	a.right = &c

	b.left = &d
	b.right = &e

	c.left = &f
	c.right = &g

	return a
}

// a-b-d-e-c-f-g
func PreOrder(n *Node) {
	if n != nil {
		fmt.Print("Value=" + n.value + " ")
		PreOrder(n.left)
		PreOrder(n.right)
	}
}

//d-b-e-a-f-c-g
func MidOrder(n *Node) {
	if n != nil {
		MidOrder(n.left)
		fmt.Print("Value=" + n.value + " ")
		MidOrder(n.right)
	}
}

//d-e-b-f-g-c-a
func PostOrder(n *Node) {
	if n != nil {
		PostOrder(n.left)
		PostOrder(n.right)
		fmt.Print("Value=" + n.value + " ")
	}
}
