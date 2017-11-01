package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v", myPow(2, -2))
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	half := myPow(x, n/2)

	if n%2 == 0 {
		return half * half
	} else if n > 0 {
		return half * half * x
	} else {
		return half / x * half
	}

}
