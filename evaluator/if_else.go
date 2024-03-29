package evaluator

import (
	"github.com/indeedhat/monkey-lang/ast"
	"github.com/indeedhat/monkey-lang/evaluator/object"
)

func evalIfExpression(expr *ast.IfExpression, env *object.Environment) object.Object {
	condition := Eval(expr.Condition, env)
	if isErr(condition) {
		return condition
	}

	if isTruthful(condition) {
		return Eval(expr.IfBlock, env)
	}

	if expr.ElseBlock != nil {
		return Eval(expr.ElseBlock, env)
	}

	return Null
}

func isTruthful(obj object.Object) bool {
	switch obj {
	case True:
		return true
	case False, Null:
		return false
	}

	switch val := obj.(type) {
	case *object.Integer:
		return val.Value != 0
	case *object.String:
		return val.Value != ""
	}

	return false
}
