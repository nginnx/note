package leetcode

func longestPalindrome(s string) string {

	var str string
	if len(s) == 1 {
		return s
	}
	for i := 0; i < len(s)-1; i++ {

		res := search(i, i, s)
		//fmt.Printf("left %v right %v res  %v \n", i, i, res)
		if len(res) > len(str) {
			str = res
		}

		if i != len(s)-1 {

			res1 := search(i, i+1, s)
			//fmt.Printf("left %v right %v res  %v \n", i, i+1, res1)

			if len(res1) > len(str) {
				str = res1
			}
		}

	}
	return str
}

func search(left, right int, s string) string {

	for left >= 0 && right < len(s) && s[left] == s[right] {
		//fmt.Printf("%v %v %v\n", left, right, s)
		left--
		right++
	}
	//fmt.Printf("left %v right %v \n", left, right)

	return s[left+1 : right]
}
