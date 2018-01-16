package main

import (
	"fmt"
)

/*Given a matrix of m x n elements (m rows, n columns), return all elements of the matrix in spiral order.

For example,
Given the following matrix:

[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]*/
func spiralOrder(matrix [][]int) []int {

	var res []int
	if len(matrix) == 0 {
		return res
	}

	x := 0
	y := 0
	carryX := 0
	carryY := 0
	endX := len(matrix[0]) - 1
	endY := len(matrix) - 1
	plus := true
	res = append(res, matrix[x][y])
	for i := 1; len(res) < len(matrix)*len(matrix[0]); i++ {

		if plus {
			if x < endX {
				x++
			} else if y < endY {
				y++
			} else {
				plus = false
				carryY++
				endX = carryX
				endY = carryY
				x--

			}
		} else {
			if x > endX {
				x--
			} else if y > endY {
				y--
			} else {
				plus = true
				carryX++
				endX = len(matrix[0]) - carryX - 1
				endY = len(matrix) - carryY - 1
				x++
			}
		}

		res = append(res, matrix[y][x])

	}
	return res
}

/*
1 2
3

*/

func main() {
	var quiz [][]int
	//quiz = append(quiz, []int{3})
	//quiz = append(quiz, []int{2})

	quiz = append(quiz, []int{12, 13, 14, 44, 55})
	quiz = append(quiz, []int{11, 3, 5, 2, 66})
	quiz = append(quiz, []int{11, 1, 6, 7, 66})
	quiz = append(quiz, []int{11, 3, 9, 1, 66})
	quiz = append(quiz, []int{11, 5, 15, 77, 66})

	res := spiralOrder(quiz)
	for _, v := range res {
		fmt.Printf("%v ", v)
	}

}
