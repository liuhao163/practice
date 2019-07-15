package tree

import "fmt"

//import "fmt"

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func InsertNode(tree *TreeNode, v int) {
	node := tree

	if node == nil {
		tree = &TreeNode{}
		tree.Value = v
		return
	}

	for node != nil {
		if v < node.Value {
			if node.Left == nil {
				node.Left = &TreeNode{}
				node.Left.Value = v
				return
			}
			node = node.Left
		} else if v > node.Value {
			if node.Right == nil {
				node.Right = &TreeNode{}
				node.Right.Value = v
				return
			}
			node = node.Right
		} else {
			fmt.Printf("error duplicate Value %d", node)
			return
		}
	}

}

func (t *TreeNode) PreIterate() {
	if t != nil {
		fmt.Printf("		Value=%v\n", t.Value)
		t.Left.PreIterate()
		t.Right.PreIterate()
	}
}
