package quiz

/*Given two binary strings, return their sum (also a binary string).

For example,
a = "11"
b = "1"
Return "100".*/

func AddBinary(a string, b string) string {
	var arr_a []byte
	var arr_b []byte
	if len(a) > len(b) {
		arr_a = []byte(a)
		arr_b = []byte(b)
	} else {
		arr_a = []byte(b)
		arr_b = []byte(a)
	}

	var result []byte
	i := len(arr_a) - 1
	j := len(arr_b) - 1
	carry := 0
	for i >= 0 && j >= 0 {
		temp := arr_a[i] + arr_b[j] + byte(carry)

		if temp >= byte(98) {
			result = append(result, 48+temp%98)
			if carry == 0 {
				carry++

			}

		} else {
			result = append(result, temp-48)
			if carry > 0 {
				carry--
			}
		}
		i--
		j--
	}

	for i := len(arr_a) - len(arr_b) - 1; i >= 0; i-- {
		temp := arr_a[i]
		sum := temp + byte(carry)

		if sum >= 50 {
			result = append(result, 48+(50-sum))
			if carry == 0 {
				carry++
			}
		} else {
			result = append(result, sum)
			if carry > 0 {
				carry--
			}
		}
	}

	for carry > 0 {
		result = append(result, 49)
		carry--
	}

	return reverse(string(result))
}

func reverse(str string) string {
	var res []byte
	temp := []byte(str)

	for i := len(temp) - 1; i >= 0; i-- {
		res = append(res, temp[i])
	}

	return string(res)
}
