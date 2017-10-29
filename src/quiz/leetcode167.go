package leetcode

func twoSum(numbers []int, target int) []int {
	m := make(map[int]int)

	for i, v := range numbers {
		_, ok := m[target-v]
		if ok {
			res := []int{m[target-v] + 1, i + 1}
			return res
		} else {
			m[v] = i
		}

	}

	return nil

}
