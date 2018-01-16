package main

import (
	"math"
)

func main() {
	count := 0
	number := 1
	res := 0
	for count < 4263116 {
		if isPrime(number) {
			count++
			res = number
		}
		number++
	}
	print(res)
}

func isPrime(n int) bool {
	if n == 2 {
		return true
	}
	if n <= 1 || n%2 == 0 {
		return false
	}
	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}
