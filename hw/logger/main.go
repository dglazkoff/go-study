package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	var buf bytes.Buffer

	mylogger := log.New(&buf, "mylog: ", 0)

	mylogger.Println("Hello, world!")
	mylogger.Println("Goodbye")
	// допишите код
	// 1) создайте переменную типа *log.Logger
	// 2) запишите в неё нужные строки

	// ...

	fmt.Print(&buf)
	// должна вывести
	// mylog: Hello, world!
	// mylog: Goodbye
}
