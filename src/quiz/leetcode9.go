package leetcode

func isPalindrome(x int) bool {
	//carry := 1
	//var arr []int
	if x < 0 {
		return false
	}
	lens := 1
	for x/(lens*10) != 0 {
		lens *= 10
	}

	left := 1
	right := lens
	for left < right {
		//fmt.Printf("%v %v \n", (x / left), (x / lens))
		if x/left%10 != x/right%10 {
			return false
		}
		left *= 10
		right /= 10

	}
	/*for x/carry != 0 {
		arr = append(arr, x/carry%10)

		carry *= 10
	}

	i := 0
	j := len(arr) - 1

	for i < len(arr)-1 && j >= 0 {
		//fmt.Printf("%v  %v \n", arr[i], arr[j])
		if arr[i] != arr[j] {
			return false
		}
		i++
		j--
	}*/

	return true
}
