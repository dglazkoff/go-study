package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"strings"
)

func main() {
	src := `/* Тестовый пакет */
package main

// Double умножает значение на 2.
func Double(i int) int {
    return i*2
}

func main() {
   // умножаем в цикле
   for i := 1; i < 5; i++ {
      fmt.Println(Double(i))
   }
}`
	// создаём token.FileSet
	fset := token.NewFileSet()
	// получаем дерево разбора
	f, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
	}

	for _, gr := range f.Comments {
		for _, c := range gr.List {
			if strings.HasPrefix(c.Text, "//") || strings.HasPrefix(c.Text, "/*") {
				fmt.Println(fset.Position(c.Slash).String(), c.Text)
			}
		}
	}
}
