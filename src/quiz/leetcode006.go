package leetcode

/*
	1     7
	2   6 8
	3 5   9
	4     10
*/

func convert(s string, numRows int) string {

	if numRows == 1 {
		return s
	}

	var arr []byte
	for i := 0; i < numRows; i++ {
		dis := 2 * (numRows - 1)
		var dis2 int

		if i == 0 {
			dis2 = dis
		} else {
			dis2 = dis - 2*i
		}
		for j := i; j < len(s); j += dis {
			arr = append(arr, s[j])
			if i > 0 && i < numRows-1 {
				temp := j + dis2
				if temp < len(s) {
					arr = append(arr, s[temp])

				}
			}
		}

	}

	return string(arr)
}
