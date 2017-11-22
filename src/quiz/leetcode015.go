package main

import (
	"fmt"
	"sort"
)

/*Given an array S of n integers, are there elements a, b, c in S such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note: The solution set must not contain duplicate triplets.

For example, given array S = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]*/

func threeSum(nums []int) [][]int {

	sort.Ints(nums)
	var data map[int][]int
	data = make(map[int][]int)
	var result [][]int

	//var result [][]int
	for i := 0; i < len(nums); i++ {

		_, ok := data[nums[i]]
		if ok {
			fmt.Printf("%v contains \n", nums[i])
			continue
		} else {

			target := -nums[i]
			//fmt.Printf("target %v \n", target)
			subn := nums[i+1:]
			m := make(map[int]int, len(subn))

			for subi := 0; subi < len(subn); subi++ {

				v, has := m[target-subn[subi]]
				if has {

					v1, has1 := data[nums[i]]
					if has1 {
						if v1[0] != subn[v] && v1[0] != subn[subi] {
							result = append(result, []int{nums[i], subn[v], subn[subi]})
							data[nums[i]] = []int{subn[v], subn[subi]}

						}
					} else {
						result = append(result, []int{nums[i], subn[v], subn[subi]})
						data[nums[i]] = []int{subn[v], subn[subi]}

					}

					/*result = append(result, []int{nums[i], subn[v], subn[subi]})
					data[nums[i]] = []int{subn[v], subn[subi]}
					*/
				}
				m[subn[subi]] = subi
			}

		}
	}

	//fmt.Printf("%v", result)
	return result

}

func twoSum(nums []int, target int) []int {

	m := make(map[int]int, len(nums))

	for i, num := range nums {
		v, has := m[target-nums[i]]
		if has {
			var result []int = []int{v, i}
			return result
		}
		m[num] = i
	}

	return nil
}
