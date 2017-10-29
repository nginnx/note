package leetcode

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
