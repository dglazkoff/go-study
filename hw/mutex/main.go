package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	mut := sync.Mutex{}
	m := make(map[int]int)

	for i := 0; i < 100; i++ {
		go func(v int) {
			mut.Lock()
			m[v] = 1
			mut.Unlock()
		}(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println(len(m))
}
