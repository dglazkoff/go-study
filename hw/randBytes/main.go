package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func randBytes(size int) (string, error) {
	b := make([]byte, size)
	_, err := rand.Read(b)

	if err != nil {
		return "", err
	}

	fmt.Println(b)
	fmt.Println(hex.EncodeToString(b))

	return base64.StdEncoding.EncodeToString(b), nil
}

func main() {
	fmt.Println(randBytes(10))
}
