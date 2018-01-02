/*Given a string containing just the characters '(', ')', '{', '}', '[' and ']',
determine if the input string is valid.

The brackets must close in the correct order,
"()" and "()[]{}" are all valid but "(]" and "([)]" are not.*/
package main

import (
	"container/list"
)

func main() {
	print(isValid("[()]"))
}

func isValid(s string) bool {
	operatorStack := list.New()
	for _, v := range s {
		if v == '(' || v == '{' || v == '[' {
			operatorStack.PushFront(v)
		} else if v == ')' {
			front := operatorStack.Front()
			if front != nil {
				if front.Value.(int32) == '(' {
					operatorStack.Remove(operatorStack.Front())
				} else {
					return false
				}
			} else {
				return false
			}
		} else if v == ']' {
			front := operatorStack.Front()
			if front != nil {
				if front.Value.(int32) == '[' {
					operatorStack.Remove(operatorStack.Front())
				} else {
					return false
				}
			} else {
				return false
			}
		} else if v == '}' {
			front := operatorStack.Front()
			if front != nil {
				if front.Value.(int32) == '{' {
					operatorStack.Remove(operatorStack.Front())
				} else {
					return false
				}
			} else {
				return false
			}
		}
	}
	if operatorStack.Front() != nil {
		return false
	}
	return true
}
