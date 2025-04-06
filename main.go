package main

import (
	"encoding/json"
	"fmt"

	"github.com/pratikdev/the-super-tiny-compiler-go/internal/parser"
	"github.com/pratikdev/the-super-tiny-compiler-go/internal/tokenizer"
	"github.com/pratikdev/the-super-tiny-compiler-go/internal/transformer"
)

func main() {
	input := `(add 52 (sub 23 12))`
	tokens := tokenizer.Tokenizer(input)
	ast := parser.Parser(tokens)
	newAst := transformer.Transformer(ast)

	// print in json format with indentation
	astJSON, err := json.MarshalIndent(newAst, "", "  ")
	if err != nil {
		fmt.Println("Error converting AST to JSON:", err)
		return
	}
	fmt.Println(string(astJSON))
}
