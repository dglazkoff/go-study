package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	freq := make(map[string]int)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	// На MacOS (и других Unix-подобных системах) стандартный ввод (stdin) буферизуется на уровне терминала, и ввод обрабатывается построчно.
	//Это означает, что данные передаются программе только после того, как вы нажмете Enter.
	for scanner.Scan() {
		text := scanner.Text()

		freq[text] += 1
		// fmt.Println(text)
	}

	for text, count := range freq {
		fmt.Printf("%s\t%d\n", text, count)
	}
}
