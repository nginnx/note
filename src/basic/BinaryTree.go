package main

/*import (
	"fmt"
)
*/
func main() {
	printNode(generateTree("1234567"))
}

type Node struct {
	LeftChild  *Node
	RightChild *Node
	Val        interface{}
}

func generateTree(raw string) *Node {
	head := Node{nil, nil, nil}
	for _, v := range raw {
		str := string(v)
		if head.LeftChild == nil {
			head.LeftChild = &Node{nil, nil, str}
		} else if head.RightChild == nil {
			head.RightChild = &Node{nil, nil, str}
		} else {
			head = *head.LeftChild
			head.LeftChild = &Node{nil, nil, str}
		}
	}
	return &head
}

func printNode(head *Node) {
	if head == nil {
		return
	}
	print(head.Val.(string) + "\n")
	printNode(head.LeftChild)
	printNode(head.RightChild)
}
