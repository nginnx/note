package leetcode

func SearchInsert(nums []int, target int) int {
	if len(nums) == 1 {
		if target > nums[0] {
			return 1
		} else {
			return 0
		}
	}

	for i := 1; i < len(nums); i++ {
		if target == nums[i] {
			return i
		} else if target < nums[i] {
			if target > nums[i-1] {
				return i
			} else {
				return i - 1
			}
		}
	}

	return len(nums)
}
