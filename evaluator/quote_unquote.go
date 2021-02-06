package evaluator

import (
	"monkey/m/v2/ast"
	"monkey/m/v2/object"
)

func quote(exp ast.Expression) object.Object {
	return &object.Quote{Node: exp}
}
