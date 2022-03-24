package evaluator

import (
	"github.com/shoma3571/go_interpreter/ast"
	"github.com/shoma3571/go_interpreter/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	}

	return nil
}