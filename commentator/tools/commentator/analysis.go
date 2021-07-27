package commentator

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

type (
	Analysis struct {
		Src string
	}
)

func (a *Analysis) Exec() {
	fmt.Println(a.Src)
	src := `package moke
import ( "fmt")

func main() {
    fmt.Printf("OK!")
}
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		fmt.Printf("Failed to parse file\n")
		return
	}
	ast.Print(fset, f)
}