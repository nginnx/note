package main

import (
	"fmt"
)

type ListNode struct {
	Next *ListNode
	Val  interface{}
}

func main() {
	head := generateLinkedList()
	temp := reverse(&head)
	for temp.Next != nil {
		fmt.Printf("%v", temp.Val)
		temp = temp.Next
	}
}

func generateLinkedList() ListNode {
	head := ListNode{nil, 0}
	for i := 9; i > 0; i-- {
		addLast(&head,i)
	}

	return head
}

func addLast(head *ListNode,val int){
	for head.Next != nil {
		head = head.Next
	}
	head.Next = &ListNode{nil, val}

}

func reverse(head *ListNode) *ListNode {
	current := head
	next := current.Next
	for next != nil {
		temp := next.Next
		next.Next = current
		current = next
		next = temp
	}
	head.Next = nil
	return current
}
