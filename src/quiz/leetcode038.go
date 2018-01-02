/*
The count-and-say sequence is the sequence of integers with the first five terms as following:

1.     1
2.     11
3.     21
4.     1211
5.     111221
1 is read off as "one 1" or 11.
11 is read off as "two 1s" or 21.
21 is read off as "one 2, then one 1" or 1211.
Given an integer n, generate the nth term of the count-and-say sequence.

Note: Each term of the sequence of integers will be represented as a string.

Example 1:

Input: 1
Output: "1"
Example 2:

Input: 4
Output: "1211"*/
/*
1
11
21
1211
111221
*/
package main

import (
	"fmt"
)

func main() {
	print(countAndSay(6))
}

func countAndSay(count int) string {
	var res string = "1"
	for count > 1 {
		res = say(res)
		count--
		print(res + "\n")
	}

	return res
}

func say(str string) string {

	var arr = []byte(str)
	var res string
	count := 1
	for i := 0; i < len(arr); i++ {
		previous := arr[i]
		if i+1 < len(arr) {

			current := arr[i+1]
			if current == previous {
				count++
			} else {
				res = res + fmt.Sprint(count)
				res = res + string(previous)
				count = 1
			}
		} else {
			res = res + fmt.Sprint(count)
			res = res + string(previous)
		}

	}

	return res
}
