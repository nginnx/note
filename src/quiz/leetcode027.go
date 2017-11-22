package main

//3,2,2,3
//2,2,3,3
func removeElement(nums []int, val int) int {
	//index := 1
	lastindex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[lastindex] = nums[i]
			lastindex++
		}

	}
	return lastindex
}
