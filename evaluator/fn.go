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

	var (
		args  = make([]object.Object, 0, len(val.Arguments))
		scope = env.NewScope()
	)

	// evaluate arguments
	for _, arg := range val.Arguments {
		a := Eval(arg, env)
		if isErr(a) {
			return a
		}

		args = append(args, a)
	}

	// apply function
	switch fn := ret.(type) {
	case *object.Function:
		if len(val.Arguments) != len(fn.Parameters) {
			return error("unexpected arg count: expect(%d) found(%d)", len(fn.Parameters), len(val.Arguments))
		}
		// bind arguments to the funciton scope
		for i, arg := range args {
			scope.Set(fn.Parameters[i].String(), arg)
		}

		// unwrap return values
		evald := Eval(fn.Body, scope)
		if ret, ok := evald.(*object.ReturnValue); ok {
			return ret.Value
		}

		return evald

	case *object.Builtin:
		return fn.Fn(args...)
	}

	return error("not a function: %T", val)
}
