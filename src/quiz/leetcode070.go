package main

func main() {
	print(climbStairs(100))
}

var test map[int]int = make(map[int]int)

func climbStairs(n int) int {

	if n == 1 || n == 2 || n == 0 {
		return n
	}
	_, ok := test[n]
	if ok {
		return test[n]
	} else {
		test[n] = climbStairs(n-1) + climbStairs(n-2)
		return test[n]
	}

}
