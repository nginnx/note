package main

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	if haystack == "" && needle == "" {
		return 0
	}
	for i := 0; i <= len(haystack)-len(needle); i++ {
		match := true
		if needle == "" {
			return 0
		}
		for n, v := range needle {
			if byte(v) != haystack[i+n] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Print(strStr("aaa", "aaaa"))
}
