package main

import "fmt"

func main() {
	c := gen(2, 3)
	out := square(c)

	for res := range out {
		fmt.Println(res)
	}
}

// реализация генератора gen здесь
func gen(nums ...int) chan int {
	outCh := make(chan int)

	go func() {
		defer close(outCh)

		for _, data := range nums {
			outCh <- data
		}
	}()

	return outCh
}

// реализация square здесь
func square(in chan int) chan int {
	outCh := make(chan int)

	go func() {
		defer close(outCh)

		for data := range in {
			outCh <- data * data
		}
	}()

	return outCh
}
