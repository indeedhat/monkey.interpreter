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

func Eval(node ast.Node, env *object.Environment) object.Object {
	switch val := node.(type) {
	case *ast.Program:
		return evalStatements(val.Statements, env)
	case *ast.ExpressionStatement:
		return Eval(val.Expression, env)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: val.Value}
	case *ast.StringLiteral:
		return &object.String{Value: val.Value}
	case *ast.BooleanLiteral:
		return nativeBool(val.Value)
	case *ast.NullLiteral:
		return Null
	case *ast.PrefixExpression:
		return evalPrefixExpression(val, env)
	case *ast.InfixExpression:
		return evalInfixExpression(val, env)
	case *ast.IfExpression:
		return evalIfExpression(val, env)
	case *ast.BlockStatement:
		return evalBlockStatement(val, env)
	case *ast.ReturnStatement:
		return evalReturnStatement(val, env)
	case *ast.Identifier:
		return evalIdentifier(val, env)
	case *ast.LetStatement:
		return evalLetStatement(val, env)
	case *ast.FunctionLiteral:
		return evalFunctionLiteral(val, env)
	case *ast.FunctionCallExpression:
		return evalFuncionCall(val, env)
	}

	return error("unknown node: %T", node)
}

func evalIdentifier(val *ast.Identifier, env *object.Environment) object.Object {
	ret, ok := env.Get(val.Value)
	if ok {
		return ret
	}

	if b, ok := builtins[val.Value]; ok {
		return b
	}

	return error("undefined identifier: %s", val.Value)
}

func evalLetStatement(val *ast.LetStatement, env *object.Environment) object.Object {
	ret := Eval(val.Value, env)
	if isErr(ret) {
		return ret
	}

	return env.Set(val.Name.String(), ret)
}

func evalStatements(statements []ast.Statement, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range statements {
		result = Eval(statement, env)

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

func evalBlockStatement(block *ast.BlockStatement, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range block.Statements {
		result = Eval(statement, env)

		// lets us return early from a block
		if result != nil &&
			(result.Type() == object.ReturnObj || result.Type() == object.ErrObj) {

			return result
		}
	}

	return result
}

func evalReturnStatement(ret *ast.ReturnStatement, env *object.Environment) object.Object {
	val := Eval(ret.Vaule, env)
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
