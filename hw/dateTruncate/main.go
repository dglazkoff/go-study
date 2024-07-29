package main

import (
	"fmt"
	"time"
)

func main() {
	var today time.Time = time.Now()
	today = today.Truncate(24 * time.Hour)

	fmt.Println(today)
}
