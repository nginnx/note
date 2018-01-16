package main

/*import (
	"fmt"
)
*/
func main() {
	head := ListNode{1, nil}
	generate(&head, 1)
	generate(&head, 1)
	deleteDuplicates(&head)
	printLink(&head)

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	previous := head
	var current *ListNode = previous.Next
	for current != nil {
		curVal := current.Val
		preVal := previous.Val

		if preVal == curVal {

			if current.Next != nil {
				previous.Next = current.Next
				current = current.Next
				continue
			} else {

				previous.Next = nil
			}
		}
		previous = current
		current = current.Next
	}
	return head
}

func generate(head *ListNode, val int) {
	temp := head
	for temp.Next != nil {
		temp = temp.Next
	}

	temp.Next = &ListNode{val, nil}

}

func printLink(head *ListNode) *ListNode {
	temp := head
	for temp != nil {
		print(temp.Val)
		temp = temp.Next
	}
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
