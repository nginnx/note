package main

import (
	"fmt"
	"sort"
)

/*There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

Example 1:
nums1 = [1, 3]
nums2 = [2]

The median is 2.0
Example 2:
nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var temp []int
	res := float64(0)
	for _, v := range nums1 {
		temp = append(temp, v)
	}

	for _, v := range nums2 {
		temp = append(temp, v)
	}
	sort.Ints(temp)
	if len(temp)%2 == 0 {
		left := float64(temp[len(temp)/2])
		right := float64(temp[len(temp)/2-1])
		res = (left + right) / 2
	} else {
		res = float64(temp[len(temp)/2])
	}
	return res

}

func main() {
	fmt.Printf("%v", findMedianSortedArrays([]int{1, 3}, []int{2}))
}
