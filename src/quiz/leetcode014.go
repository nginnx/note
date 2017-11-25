package main

import "fmt"

/*
给定一个String类型数组，要求写一个方法，返回数组中这些字符串的最长公共前缀。
举个例子：假如数组为["123","12","4"]，经过这个方法返回的结果就应该是""。
因为"123"，"12"，"4"并没有共同的前缀，虽然"123"，"12"的公共最长前缀是"12"，
但是这个公共前缀"12"与"4"没有公共前缀，所以最后返回的结果就是""。
*/

func main() {
	var strs = []string{"c","c"}
	fmt.Printf("%s", longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 || strs[0] == "" {
		return ""
	}

	if len(strs)==1 {
		return strs[0]
	}

	str := strs[0]
	index := 0
	var res []byte
Label:
	for {
		if index < len(str) {
			simple := str[index]
			for i := 1; i < len(strs); i++ {
				if index < len(strs[i]) {
					char := strs[i][index]
					if char != simple {
						break Label
					}
				} else {

					break Label
				}
			}
			res = append(res, simple)

		}else{
			break
		}

		index++

	}

	return string(res)

}
