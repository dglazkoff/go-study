package main

import (
	"fmt"
	"time"
)

func main() {
	currentTimeStr := "2021-09-19T15:59:41+03:00"
	currentTime, err := time.Parse(time.RFC3339, currentTimeStr)
	// или
	// currentTime, err := time.Parse("2006-01-02T15:04:05Z07:00", currentTimeStr)
	if err != nil {
		panic(err)
	}
	fmt.Println(currentTime)
}
