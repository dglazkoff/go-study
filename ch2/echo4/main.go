// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 33.
//!+

// Echo4 prints its command-line arguments.
package main

import (
	"flag"
	"fmt"
	"strings"
)

// зачем возвращать указатель? потому что n будет менятся со временем. если запишем сразу примитив, то он и будет хранится как значение переменной
// но если передавать указатель, то мы можем менять значение переменной неявно и потом взять его по указателю
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

// тут мы создаем новую переменную и соответственно у нее будет другой адрес в памяти.
// если бы мы написали var a = n, то адреса были бы одинаковые
var a = *n

func main() {
	flag.Parse()

	fmt.Println(a)
	fmt.Println(*n)

	fmt.Println(&a)
	fmt.Println(n)

	fmt.Print(strings.Join(flag.Args(), *sep))

	if !*n {
		fmt.Println()
	}
}

//!-
