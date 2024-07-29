package main

import (
	"fmt"
	"unicode/utf8"
)

func reverseRune(s []byte) {
	// Reverse the entire byte slice.
	for i, j := 0, len(s)-1; i < j; {
		_, size1 := utf8.DecodeRune(s[i:])
		_, size2 := utf8.DecodeRune(s[j:])

		// Move runes in place
		copy(s[i:], s[j+1-size2:j+1])
		copy(s[j+1-size2:], s[i:i+size1])

		// Adjust indices
		i += size1
		j -= size2
	}
}

func main() {
	testCases := [][]byte{
		[]byte("Hello, 世界"),
		[]byte("Привет, мир"),
		[]byte("こんにちは世界"),
		[]byte("안녕하세요 세계"),
	}

	for _, s := range testCases {
		fmt.Printf("Original: %s\n", s)
		reverseRune(s)
		fmt.Printf("Reversed: %s\n\n", s)
	}
}
