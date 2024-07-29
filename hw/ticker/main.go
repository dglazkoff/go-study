package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	ticker := time.NewTicker(time.Second * 2)

	for i := 0; i < 10; i++ {
		t := <-ticker.C
		fmt.Println(t.Sub(start).Seconds())
	}

}
