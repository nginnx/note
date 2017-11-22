package main

import (
	"fmt"
)

/*Given a non-negative integer represented as a non-empty array of digits, plus one to the integer.

You may assume the integer do not contain any leading zero, except the number 0 itself.

The digits are stored such that the most significant digit is at the head of the list.*/

func main() {

	//var i uint
	//count := 0
	/*for i = 32; i != 0; i-- {
		if num>>i&1 == 0 {
			continue
		}
		fmt.Printf("%v ", num>>i&1)

	}*/

	//fmt.Printf("%v", plusOne([]int{1, 0, 0, 1, 1}))
	fmt.Printf("%v", plusOne([]int{3, 9, 9}))
}

//2,9,9

func plusOne(digits []int) []int {

	carry := 1
	var res []int
	res = make([]int, len(digits))
	for i := len(digits) - 1; i >= 0; i-- {
		temp := digits[i] + carry
		if temp == 10 {

			res[i] = 0

		} else {
			carry = 0
			res[i] = temp
		}

		if i == 0 && carry == 1 {
			temp := make([]int, len(res)+1)
			temp[0] = 1
			for i := 1; i < len(temp); i++ {
				temp[i] = res[i-1]
			}
			return temp
		}
	}

	return res
}
