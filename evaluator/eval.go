package evaluator

import (
	"fmt"

	"github.com/indeedhat/monkey-lang/ast"
	"github.com/indeedhat/monkey-lang/evaluator/object"
)

var (
	True  = &object.Boolean{Value: true}
	False = &object.Boolean{Value: false}
	Null  = &object.Null{}
)

func Eval(node ast.Node) object.Object {

	switch val := node.(type) {
	case *ast.Program:
		return evalStatements(val.Statements)
	case *ast.ExpressionStatement:
		return Eval(val.Expression)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: val.Value}
	case *ast.StringLiteral:
		return &object.String{Value: val.Value}
	case *ast.BooleanLiteral:
		return nativeBool(val.Value)
	case *ast.NullLiteral:
		return Null
	case *ast.PrefixExpression:
		ret := Eval(val.Right)
		if isErr(ret) {
			return ret
		}
		return evalPrefixExpression(val.Operator, ret)
	case *ast.InfixExpression:
		left := Eval(val.Right)
		if isErr(left) {
			return left
		}
		right := Eval(val.Right)
		if isErr(right) {
			return right
		}
		return evalInfixExpression(left, val.Operator, right)
	case *ast.IfExpression:
		return evalIfExpression(val)
	case *ast.BlockStatement:
		return evalBlockStatement(val)
	case *ast.ReturnStatement:
		return evalReturnStatement(val)
	}

	return error("unknown node: %T", node)
}

func evalStatements(statements []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range statements {
		result = Eval(statement)

		switch ret := result.(type) {
		case *object.Error:
			// break exec on error
			return ret
		case *object.ReturnValue:
			// break exec on early return
			return ret.Value
		}
	}

	return result
}

func evalBlockStatement(block *ast.BlockStatement) object.Object {
	var result object.Object

	for _, statement := range block.Statements {
		result = Eval(statement)

		// lets us return early from a block
		if result != nil &&
			(result.Type() == object.ReturnObj || result.Type() == object.ErrObj) {

			return result
		}
	}

	return result
}

func evalReturnStatement(ret *ast.ReturnStatement) object.Object {
	val := Eval(ret.Vaule)
	if isErr(val) {
		return val
	}
	return &object.ReturnValue{Value: val}
}

func nativeBool(input bool) *object.Boolean {
	if input {
		return True
	}

	return False
}

func error(format string, args ...any) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, args...)}
}

func isErr(obj object.Object) bool {
	return obj != nil && obj.Type() == object.ErrObj
}
