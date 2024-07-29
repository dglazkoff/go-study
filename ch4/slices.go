package main

import "fmt"

func main() {
	months := [...]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"}

	summer := months[5:8]
	summer = summer[:9]

	fmt.Println(summer)
}
