// The Object Tree pattern is a design pattern used to manage
// hierarchical data structures such as trees.
// In this pattern, each node in the tree is an object with a pointer to its parent node
// and an array or a slice of pointers to its child nodes.
package structural_patterns

import "fmt"

type Node struct {
	Value    string
	Parent   *Node
	Childern []*Node
}

func (n *Node) AddChildren(child *Node) {
	child.Parent = n
	n.Childern = append(n.Childern, child)
}

func (n *Node) ResetParent(parent *Node) {
	oldParent := n.Parent
	oldParentChildenLength := len(oldParent.Childern)
	if oldParentChildenLength != 0 {
		// rm from old parent's children
		oldParent.Childern = oldParent.Childern[:oldParentChildenLength-1]
	}
	// add to new parent's children
	parent.AddChildren(n)
	// set new parent
	n.Parent = parent
}

func (n *Node) Print(indent string) {
	fmt.Println(indent + n.Value)
	for _, c := range n.Childern {
		c.Print(indent + "  ")
	}
}

func PrintObjectTree() {
	root := &Node{Value: "Root"}

	// level 1 child nodes
	child1 := &Node{Value: "c1"}
	child2 := &Node{Value: "c2"}
	root.AddChildren(child1)
	root.AddChildren(child2)

	// level 2 child nodes
	child11 := &Node{Value: "c11"}
	child12 := &Node{Value: "c12"}
	child21 := &Node{Value: "c21"}

	child1.AddChildren(child11)
	child1.AddChildren(child12)
	child2.AddChildren(child21)

	root.Print(" ")
	// child1.Print(" ")
	// child2.Print(" ")
	child21.ResetParent(child1)
	root.Print(" ")
}
