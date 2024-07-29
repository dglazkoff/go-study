// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 83.

// The sha256 command computes the SHA256 hash (an array) of a string.
package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func countDifferentBits(hash1, hash2 [32]byte) int {
	count := 0
	for i := 0; i < len(hash1); i++ {
		// XOR the bytes and count the number of set bits
		diff := hash1[i] ^ hash2[i]

		fmt.Printf("%x\n", diff)

		count += popCount(diff)
	}
	return count
}

// popCount returns the number of set bits (1s) in a byte.
func popCount(x byte) int {
	count := 0
	for x != 0 {
		count += int(x & 1)
		x >>= 1
	}
	return count
}

func main() {
	//args := os.Args[1:]
	//var cod string
	//var hash hash.Hash
	//reader := bufio.NewReader(os.Stdin)
	//
	//for _, v := range args {
	//	if strings.HasPrefix(v, "-") {
	//		cod = v[1:]
	//		args = args[1:]
	//	}
	//}
	//
	//switch cod {
	//case "sha384":
	//	hash = sha512.New384()
	//case "sha512":
	//	hash = sha512.New()
	//default:
	//	hash = sha256.New()
	//}
	//
	//input, _ := reader.ReadString('\n')
	//
	//// только проблема в том что не используются стримы
	//fmt.Printf("%x\n", hash.Sum([]byte(input)))

	c1 := sha256.Sum256([]byte(os.Args[1]))
	c2 := sha256.Sum256([]byte(os.Args[2]))
	fmt.Printf("%x\n", c1)
	fmt.Printf("%x\n", c2)
	fmt.Printf("%v\n", countDifferentBits(c1, c2))
	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8
}

//!-
