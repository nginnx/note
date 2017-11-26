package main

import "fmt"

func main() {
	println(maxArea([]int{1,1}))
}

func maxArea(height []int) int {
	var startIndex = 0
	var endIndex = len(height) - 1
	startHeight := height[0]
	endHeight := height[len(height)-1]
	var res = 0
	if startHeight > endHeight {
		res = endHeight * (len(height)-1)
		endIndex--
	} else {
		res = startHeight * (len(height)-1)
		startIndex++
	}

	fmt.Println(res)

	for startIndex < endIndex {
		startHeight := height[startIndex]
		endHeight := height[endIndex]
		var temp = 0
		if startHeight > endHeight {
			temp = endHeight * (endIndex - startIndex)
			fmt.Println(temp, "   ", endHeight, "   ", endIndex, "   ", startIndex)

		} else {
			temp = startHeight * (endIndex - startIndex)
			fmt.Println(temp, "   ", startHeight, "   ", endIndex, "   ", startIndex)

		}

		if temp > res {
			res = temp
		}
		if startHeight < endHeight {
			startIndex++
		} else {
			endIndex--
		}
	}

	return res

}
