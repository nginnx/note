package main

import (
	"fmt"
	"runtime"
	"sync"
)

var count int = 0

func main() {
	var lock sync.Mutex = sync.Mutex{}
	for i := 0; i < 10000; i++ {
		go adder(&lock)
	}

	for {
		lock.Lock()
		c := count
		lock.Unlock()
		runtime.Gosched()
		if c == 10000 {
			fmt.Println(c)
			break
		}
	}
}

func adder(llock *sync.Mutex) {
	llock.Lock()
	count++
	llock.Unlock()
}
