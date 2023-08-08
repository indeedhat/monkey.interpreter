package evaluator

import (
	"github.com/indeedhat/monkey-lang/ast"
	"github.com/indeedhat/monkey-lang/evaluator/object"
)

func evalArrayLiteral(array *ast.ArrayLiteral, env *object.Environment) object.Object {
	elems := evalExpressions(array.Elements, env)
	if isErr(elems[0]) {
		return elems[0]
	}

	return &object.Array{
		Elements: elems,
	}
}

func evalIndexExpression(node *ast.IndexExpression, env *object.Environment) object.Object {
	subject := Eval(node.Subject, env)
	if isErr(subject) {
		return subject
	}

	index := Eval(node.Index, env)
	if isErr(index) {
		return subject
	}

	switch {
	case subject.Type() == object.ArrayObj && index.Type() == object.IntegerObj:
		return evalArrayIndexExpression(subject, index)
	case subject.Type() == object.StringObj && index.Type() == object.IntegerObj:
		return evalStringIndexExpression(subject, index)
	}

	return error("index operator not supported: <%s[%s]>", subject.Type(), index.Type())
}

func evalStringIndexExpression(subject, index object.Object) object.Object {
	str := subject.(*object.String)
	i := index.(*object.Integer)

	if i.Value < 0 || i.Value >= int64(len(str.Value)) {
		return Null
	}

	c := str.Value[i.Value]
	return &object.String{Value: string(c)}
}

func evalArrayIndexExpression(subject, index object.Object) object.Object {
	array := subject.(*object.Array)
	i := index.(*object.Integer)

	if i.Value < 0 || i.Value >= int64(len(array.Elements)) {
		return Null
	}

	return array.Elements[i.Value]
}
