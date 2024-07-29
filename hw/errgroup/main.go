package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
)

type Name string

var names = []Name{"Anna", "Ivan", "Fedor", "Katya", "Gleb"}

// Здесь напишите метод для Name

type ErrName struct {
	name Name
}

func (en *ErrName) Print() error {
	fmt.Printf("Hello %s\n", en.name)
	return nil
}

func main() {
	g := &errgroup.Group{}

	for _, name := range names {
		en := ErrName{name}
		g.Go(en.Print)

	}

	g.Wait()
}
