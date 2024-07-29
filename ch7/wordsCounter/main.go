package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordsCounter int

func (c *WordsCounter) Write(p []byte) (int, error) {
	currentCount := 0
	scanner := bufio.NewScanner(bytes.NewBuffer(p))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		currentCount += 1
		*c += 1
	}

	return currentCount, nil
}

func main() {
	const input = "Now is the winter of our discontent,\nMade glorious summer by this sun of York.\n"

	var c WordsCounter
	fmt.Fprintf(&c, input)

	fmt.Println(c)
}
