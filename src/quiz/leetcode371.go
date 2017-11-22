package main

import (
	"fmt"
)

func add(a, b int) int {
	v := a
	v1 := b

	for v1 != 0 {
		temp := v ^ v1
		temp1 := (v & v1) << 1
		v = temp
		v1 = temp1
	}

	return v
}

/*011
011

000
0110


110
111


0001
1100

1101
0000





*/
