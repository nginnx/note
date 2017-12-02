package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var res = &ListNode{-1, nil}
	var dummy = res
	for l1 != nil && l2 != nil {
		l1Value := l1.Val
		l2Value := l2.Val
		if l1Value < l2Value {
			dummy.Next = l1
			l1 = l1.Next
		} else {
			dummy.Next = l2
			l2 = l2.Next
		}
		dummy = dummy.Next
	}
	if l1 != nil {
		dummy.Next = l1
	} else {
		dummy.Next = l2
	}

	return res.Next

}

func main() {

}
