package main

import (
	"encoding/json"
	"fmt"

	"github.com/pratikdev/the-super-tiny-compiler-go/internal/parser"
	"github.com/pratikdev/the-super-tiny-compiler-go/internal/tokenizer"
)

func main() {
	input := `(add 52 (sub 23 12))`
	tokens := tokenizer.Tokenizer(input)
	ast := parser.Parser(tokens)

	// print the ast in json format with indentation
	astJSON, err := json.MarshalIndent(ast, "", "  ")
	if err != nil {
		fmt.Println("Error converting AST to JSON:", err)
		return
	}
	fmt.Println(string(astJSON))
}
