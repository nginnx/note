package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v", mySqrt(2147395599))
}

func mySqrt(x int) int {
	var z float64
	z = 10
	for i := 0; i < 100; i++ {
		z -= ((z*z - float64(x)) / (2 * z))
	}
	return int(z)
}
