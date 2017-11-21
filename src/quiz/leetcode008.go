package main

import (
	"fmt"
)

func myAtoi(str string) int {
	if str == "" {
		return 0
	}
	bytes := []byte(str)
	var res int64 = 0
	carry := 1

	positive := true
	start := true
	startindex := 0
	for i, v := range bytes {
		if !start && v == 32 {
			return 0
		}

		if start {
			if v==32 {
				continue
			}
			if v == '-' {
				positive = false
			} else if v > 48 && v <= 57 {
				startindex = i
				break
			}else if v!='+'&&v<48||v>57{
				return 0
			}
			start = false
		} else if !start {
			fmt.Printf("begin %v", v)
			if v == 48 {

				continue
			}

			if v > 48 && v <= 57 {
				if i+1 < len(bytes) && bytes[i+1] == 32 {
					return 0
				} else {
					startindex = i
					break
				}

			} else {
				return 0
			}
		}
	}
	//fmt.Printf("start %v %v \n",startindex,bytes[startindex+1])
	for i := len(bytes) - 1; i >= startindex; i-- {

		v := bytes[i]

		if v >= 48 && v <= 57 {
			var num int
			if positive {
				num = int(v - 48)
			} else {
				num = -(int(v - 48))
			}
			res += int64(num * carry)

			if res > 2147483647 {
				return 2147483647
			}
			if res < -2147483648 {
				return -2147483648
			}

			carry *= 10

		} else if v == 32 {
			res = 0

			carry = 1
		} else {
			res = 0

			carry = 1
		}
	}

	return int(res)
}

func main() {

	print(myAtoi("9223372036854775809"))
}
