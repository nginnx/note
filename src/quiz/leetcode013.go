package main

import "fmt"

func main() {
	fmt.Printf("%v", romanToInt("MDCCCLXXXIV"))
}

/*
相同的数字连写，所表示的数等于这些数字相加得到的数，例如：III = 3
小的数字在大的数字右边，所表示的数等于这些数字相加得到的数，例如：VIII = 8
小的数字，限于（I、X和C）在大的数字左边，所表示的数等于大数减去小数所得的数，例如：IV = 4
正常使用时，连续的数字重复不得超过三次
从前向后遍历罗马数字，如果某个数比前一个数小，则加上该数。反之，减去前一个数的两倍然后加上该数
*/
func romanToInt(s string) int {
	var dic = make(map[rune]int)
	dic['I'] = 1
	dic['V'] = 5
	dic['X'] = 10
	dic['L'] = 50
	dic['C'] = 100
	dic['D'] = 500
	dic['M'] = 1000

	arr := []rune(s)
	var result = dic[arr[0]]
	for i := 1; i < len(arr); i++ {
		cur := dic[arr[i]]
		previous := dic[arr[i-1]]
		if cur <= previous {
			result += cur
		} else {
			result -= previous << 1
			result += cur
		}
	}

	return result
}
