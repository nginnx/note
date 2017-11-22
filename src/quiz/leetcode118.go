package main

/*import (
	"fmt"
)*/

func Generate(numRows int) [][]int {

	var arr [][]int
	if numRows == 0 {
		return arr
	}
	arr = append(arr, []int{1})
	for i := 1; i < numRows; i++ {
		var temp []int
		temp = append(temp, 1)
		for j := 1; j < i; j++ {
			temp = append(temp, arr[i-1][j-1]+arr[i-1][j])
		}
		temp = append(temp, 1)
		arr = append(arr, temp)
	}
	return arr
}
