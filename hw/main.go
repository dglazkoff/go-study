package main

import "fmt"

type User struct {
	Balance int
}

func main() {
	user := User{Balance: 100}
	AddBalance(user)

	fmt.Println(user.Balance)
}

func AddBalance(user User) {
	user.Balance += 100
}
