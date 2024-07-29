package main

import (
	"fmt"
	"time"
)

func main() {
	// допишите код здесь
	birthday := time.Date(2093, time.November, 26, 0, 0, 0, 0, time.Local)
	days := time.Until(birthday).Hours() / 24
	fmt.Println(days)
}
