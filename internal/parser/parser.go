package parser

import "github.com/pratikdev/the-super-tiny-compiler-go/internal/tokenizer"

// valueLiteral is a struct that represents a value literal (number or string)
type valueLiteral struct {
	Type  string
	Value string
}

// GetType is a method that returns the type of the value literal
func (v valueLiteral) getType() string {
	return v.Type
}

// node is an interface that represents a node in the AST
type node interface {
	getType() string
}

// callExpression is a struct that represents a function call expression
type callExpression struct {
	Type   string
	Name   string
	Params []node
}

// GetType is a method that returns the type of the call expression
func (c callExpression) getType() string {
	return c.Type
}

// walk is a function that takes a slice of tokens and returns a CallExpression
// It also takes a pointer to an int that is used to keep track of the index in the given tokens slice
func walk(tokens []tokenizer.Token, j *int) callExpression {
	callExpression := callExpression{
		Type:   "CallExpression",
		Name:   tokens[1].Value,
		Params: []node{},
	}

	for i := 2; i < len(tokens); i++ {
		token := tokens[i]

		if token.Type == "number" {
			callExpression.Params = append(callExpression.Params, valueLiteral{
				Type:  "NumberLiteral",
				Value: token.Value,
			})
		}

		if token.Type == "string" {
			callExpression.Params = append(callExpression.Params, valueLiteral{
				Type:  "StringLiteral",
				Value: token.Value,
			})
		}

		if token.Type == "paren" && token.Value == "(" {
			// pass the rest of the tokens (without the last paren as we didn't include the first one in the loop)
			// to the `walk` function
			// and get the new call expression
			// and append it to the `Params` array

			callExpression.Params = append(callExpression.Params, walk(tokens[i:len(tokens)-1], &i))
		}

		// j is used to keep track of the index in the original tokens array
		// so we don't process the same token again
		(*j)++
	}

	return callExpression
}

// AST is a struct that represents the Abstract Syntax Tree (AST)
type AST struct {
	Type string
	Body []callExpression
}

// GetType is a method that returns the type of the AST
func Parser(tokens []tokenizer.Token) AST {
	ast := AST{
		Type: "Program",
		Body: []callExpression{},
	}

	for i := 0; i < len(tokens); {
		token := tokens[i]

		if token.Type == "paren" && token.Value == "(" {
			callExpression := walk(tokens[i:], &i)
			ast.Body = append(ast.Body, callExpression)
			i++
			continue
		}

		i++
	}

	return ast
}
