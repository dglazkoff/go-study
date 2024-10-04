package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
)

func main() {
	src := `package main
    
func main() {
     ids := 77
     id := ids + 1
     fmt.Println("id равно:", id/2 )
}`

	// создаём token.FileSet
	fset := token.NewFileSet()
	// получаем дерево разбора
	f, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
	}

	// обходим дерево разбора
	ast.Inspect(f, func(n ast.Node) bool {
		if v, ok := n.(*ast.Ident); ok {
			// и содержит подстроку "Hello"
			if v.Name == "id" {
				v.Name = "Ident"
			}
		}
		return true
	})
	// выливаем изменённый AST
	// обратно в текст исходного кода
	// и выводим в консоль
	printer.Fprint(os.Stdout, fset, f)
}
