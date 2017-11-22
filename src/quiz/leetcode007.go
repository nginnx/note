package main

func reverse(x int) int {

	res := 0
	for x != 0 {

		res = res*10 + x%10
		x /= 10
	}

	return res
}
