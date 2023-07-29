package evaluator

import (
	"github.com/indeedhat/monkey-lang/ast"
	"github.com/indeedhat/monkey-lang/evaluator/object"
)

func evalFunctionLiteral(val *ast.FunctionLiteral, env *object.Environment) object.Object {
	return &object.Function{
		Parameters: val.Parameters,
		Body:       val.Body,
	}
}

func evalFuncionCall(val *ast.FunctionCallExpression, env *object.Environment) object.Object {
	ret := Eval(val.Function, env)
	if isErr(ret) {
		return ret
	}

	fn, ok := ret.(*object.Function)
	if !ok {
		return error("ret is not a function: %T", ret)
	}

	if len(val.Arguments) != len(fn.Parameters) {
		return error("unexpected arg count: expect(%d) found(%d)", len(fn.Parameters), len(val.Arguments))
	}

	scope := env.NewScope()

	// bind arguments to the funciton scope
	for i, arg := range val.Arguments {
		a := Eval(arg, env)
		if isErr(a) {
			return a
		}

		scope.Set(fn.Parameters[i].String(), a)
	}

	// unwrap return values
	evald := Eval(fn.Body, scope)
	if ret, ok := evald.(*object.ReturnValue); ok {
		return ret.Value
	}

	return evald
}
