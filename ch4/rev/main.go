// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 86.

// Rev reverses a slice.
package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	//!+array
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(&a)
	fmt.Println(a) // "[5 4 3 2 1 0]"
	//!-array

	//!+slice
	s := []int{0, 1, 2, 3, 4, 5}
	// Rotate s left by two positions.
	s = rotate(s, 2)
	fmt.Println(s) // "[2 3 4 5 0 1]"
	//!-slice
	fmt.Printf("%s\n", reverseRune([]byte("Hello, 世界")))
	fmt.Println(removeDuplicates([]string{"a", "b", "b", "c", "c", "c"}))
	fmt.Printf("Squashed: %q\n", removeDuplicatesSpaces([]byte("This  \u00A0is   a \t\t test \n\n string  .")))

	// Interactive test of reverse.
	//	input := bufio.NewScanner(os.Stdin)
	//outer:
	//	for input.Scan() {
	//		var ints []int
	//		for _, s := range strings.Fields(input.Text()) {
	//			x, err := strconv.ParseInt(s, 10, 64)
	//			if err != nil {
	//				fmt.Fprintln(os.Stderr, err)
	//				continue outer
	//			}
	//			ints = append(ints, int(x))
	//		}
	//		reverse(ints)
	//		fmt.Printf("%v\n", ints)
	//	}
	// NOTE: ignoring potential errors from input.Err()
}

// !+rev
// reverse reverses a slice of ints in place.
func reverse(s *[6]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverseBytes(b []byte) {
	for i, j := 0, 0; i < len(b); j++ {
		r, size := utf8.DecodeRune(b)
		b[j] = r
		i += size
	}

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

//!-rev
/*
	Вы можете заранее выделить память для нового среза z с помощью make, чтобы избежать возможных повторных выделений памяти при вызовах append:
	z := make([]int, 0, len(s))
*/
func rotate(s []int, n int) []int {
	var z []int
	// fmt.Println(s)
	z = append(z, s[n:]...)
	z = append(z, s[:n]...)
	// fmt.Println(z)

	return z
}

func removeDuplicates(s []string) []string {
	fmt.Println(s)
	for i := 0; i <= len(s)-1; i++ {
		if s[i+1] == s[i] {
			lenS := len(s)
			s = s[:i+1]
			if i+1 != lenS-1 {
				s = append(s, s[i+2:lenS]...)
				i--
			}
		}
	}

	return s
}

// используется []byte чтобы делать изменения in place. иначе string нельзя изменять и так не получится
// мы должны проверять последовательность bytes, потому что некоторые Unicode символы (как \u00a0 например) занимают несколько байт но при этом это один символ (rune)
func removeDuplicatesSpaces(s []byte) []byte {
	fmt.Printf("Original: %q\n", s)
	fmt.Printf("Original: %s\n", s)
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCount(s))
	inSpace := false
	j := 0

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRune(s[i:])
		if unicode.IsSpace(r) {
			if !inSpace {
				s[j] = ' '
				j++
				inSpace = true
			}
		} else {
			inSpace = false
			copy(s[j:], s[i:i+size])
			j += size
		}

		i += size
	}

	return s[:j]
}
