package main

import "fmt"

func main() {
	ch := generator("Hello")
	for msg := range ch {
		fmt.Println(msg)
	}
}

// generator — генератор, который создает канал и сразу возвращает его
func generator(msg string) chan string {
	c := make(chan string)

	// через отдельную горутину генератор отправляет данные в канал
	go func() {
		// закрываем канал по завершению горутины — это отправитель
		defer close(c)
		for i := 0; i < 5; i++ {
			// отправляем данные в канал
			c <- msg + fmt.Sprintf(" %d", i)
		}
	}()

	return c
}
