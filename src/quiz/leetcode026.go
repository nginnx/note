package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	index := 1
	for i := 1; i < len(nums); i++ {
		if nums[index] != nums[i-1] {
			nums[index] = nums[i]
			index++
		}
	}
	fmt.Printf("%v", nums)
	return index
}
