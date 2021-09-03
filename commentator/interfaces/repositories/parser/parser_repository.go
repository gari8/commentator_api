package parser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

type ParserRepository struct{}

func NewParserRepository() ParserRepository {
	return ParserRepository{}
}

func (r ParserRepository) ParseFile(body []byte) (string, error) {
	fSet := token.NewFileSet()
	pf, err := parser.ParseFile(fSet, "commentator.go", string(body), 0)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	err = ast.Print(fSet, pf)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
