package main

import (
	"fmt"
)

/*
Find the contiguous subarray within an array (containing at least one number) which has the largest sum.

For example, given the array [-2,1,-3,4,-1,2,1,-5,4],
the contiguous subarray [4,-1,2,1] has the largest sum = 6.
*/

func main() {
	fmt.Printf("%v", maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

func maxSubArray(nums []int) int {
	const MaxUint = ^uint(0)
	const MinUint = 0
	const MaxInt = int(MaxUint >> 1)
	const MinInt = -MaxInt - 1
	var curSum int
	var maxSum int = MinInt
	for i := 0; i < len(nums); i++ {
		if curSum <= 0 {
			curSum = nums[i]
		} else {
			curSum += nums[i]
		}
		fmt.Printf("%v \n", curSum)
		//fmt.Printf("cur %v  i  %v \n", curSum, nums[i])
		if curSum > maxSum {
			maxSum = curSum
		}
	}
	return maxSum
}
