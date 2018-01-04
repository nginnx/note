package main

func main() {
	print(climbStairs(1000))
}

var dummy map[int]int = make(map[int]int)

func climbStairs(n int) int {

	if n == 1 || n == 2 || n == 0 {
		return n
	}
	_, ok := dummy[n]
	if ok {
		return dummy[n]
	} else {
		dummy[n] = climbStairs(n-1) + climbStairs(n-2)
		return dummy[n]
	}

}
