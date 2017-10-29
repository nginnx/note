package leetcode

import (
	"bytes"
	"fmt"
)

/*Given a string s consists of upper/lower-case alphabets and empty space characters ' ', return the length of last word in the string.

If the last word does not exist, return 0.

Note: A word is defined as a character sequence consists of non-space characters only.

For example,
Given s = "Hello World",
return 5.*/

func lengthOfLastWord(s string) int {
	//fmt.Printf("%v", []byte(s))

	//fmt.Printf("s %v \n", s)
	length := 0
	bt := []byte(s)
	bt = bytes.Trim(bt, " ")
	for i, v := range bt {
		if i == len(bt)-1 {
			if v != 32 {
				length++
			}
		} else if v == 32 {
			length = 0
		} else {
			length++
		}

	}

	return length
}

func main() {
	fmt.Printf("%v", lengthOfLastWord("b  a  "))
}
