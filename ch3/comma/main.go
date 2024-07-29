// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
//
//	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
//	1
//	12
//	123
//	1,234
//	1,234,567,890
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma2(os.Args[i]))
		// comma(os.Args[i])
	}
}

// !+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma2(s[:n-3]) + "," + s[n-3:]
}

func writeToStart(b *bytes.Buffer, s rune) {
	// fmt.Printf("b %s\n", b.String())
	temp := b.String()
	b.Reset()
	b.WriteRune(s)
	b.WriteString(temp)

	//fmt.Printf("Temp %s\n", temp)
	//fmt.Printf("B %s\n", b.String())
}

func comma2(s string) string {
	var buf bytes.Buffer
	var sign string

	if s[0] == '-' || s[0] == '+' {
		sign = string(s[0])
		s = s[1:]
	}

	parts := strings.Split(s, ".")

	for i := 1; i <= len(parts[0]); i++ {
		writeToStart(&buf, rune(s[len(parts[0])-i]))
		// fmt.Println(i)
		if i%3 == 0 {
			// fmt.Printf("buf %s\n", buf.String())
			writeToStart(&buf, ',')
			// fmt.Printf("buf %s\n", buf.String())
		}
	}

	if len(parts) > 1 {
		buf.WriteString(".")
		buf.WriteString(parts[1])
	}

	return sign + buf.String()
}

//!-
