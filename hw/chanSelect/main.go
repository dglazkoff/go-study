package main

import "fmt"

func main() {
	chIn := make(chan int)
	chOut := make(chan int)
	quit := make(chan struct{})

	go func() {
		for i := 0; i < 15; i++ {
			chIn <- i
		}
		close(chIn)
	}()
	go func() {
		var x int
		for x = range chIn {
			chOut <- x * 2
		}
		close(chOut)
		//for {
		//	select {
		//	case x = <-chIn:
		//		chOut <- x * 2
		//	case <-quit:
		//		return
		//	}
		//}
	}()
	go func() {
		for x := range chOut {
			fmt.Printf("%d ", x)
		}
		quit <- struct{}{}
	}()
	<-quit
}
