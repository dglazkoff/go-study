package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	firstString := os.Args[1]
	secondString := os.Args[2]

	fmt.Println(isAnagram(firstString, secondString))
}

func isAnagram(firstString, secondString string) bool {
	for i := 0; i < len(firstString); i++ {
		index := strings.Index(secondString, string(firstString[i]))

		if index == -1 {
			return false
		}
	}

	return true
}
