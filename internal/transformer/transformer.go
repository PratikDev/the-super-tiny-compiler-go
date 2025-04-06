package transformer

import (
	"github.com/pratikdev/the-super-tiny-compiler-go/internal/parser"
)

type callee struct {
	Type string
	Name string
}

type expression struct {
	Type       string
	Callee     callee
	Arguements []parser.Node
}

// GetType is a method that returns the type of the call expression
func (c expression) GetType() string {
	return c.Type
}

type expressionStatement struct {
	Type       string
	Expression parser.Node
}

type NewAST struct {
	Type string
	Body []expressionStatement
}

func Transformer(ast parser.AST) NewAST {
	newAst := NewAST{
		Type: "Program",
		Body: []expressionStatement{{
			Type: "ExpressionStatement",
		}},
	}

	for i, body := range ast.Body {
		newAst.Body[i].Expression = walk(body)
	}

	return newAst
}

func walk(astBody parser.Node) parser.Node {
	switch body := astBody.(type) {
	case parser.CallExpression:
		callExpression := expression{
			Type: body.Type,
			Callee: callee{
				Type: "Identifier",
				Name: body.Name,
			},
		}

		for _, param := range body.Params {
			switch param := param.(type) {
			case parser.CallExpression:
				callExpression.Arguements = append(callExpression.Arguements, walk(param))
			case parser.ValueLiteral:
				callExpression.Arguements = append(callExpression.Arguements, walk(param))
			}
		}

		return callExpression

	case parser.ValueLiteral:
		return parser.ValueLiteral{
			Type:  body.Type,
			Value: body.Value,
		}

	default:
		return parser.ValueLiteral{
			Type:  "Unknown",
			Value: "",
		}
	}
}
