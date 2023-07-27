package evaluator

import (
	"github.com/indeedhat/monkey-lang/evaluator/object"
)

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangPrefixOperator(right)
	case "-":
		return evalMinusPrefixOperator(right)
	}
	return error("unknown operator: %s%s", operator, right.Type())
}

func evalBangPrefixOperator(right object.Object) object.Object {
	switch right {
	case True:
		return False
	case False:
		return True
	case Null:
		return True
	}

	switch val := right.(type) {
	case *object.Integer:
		if val.Value == 0 {
			return True
		}
	case *object.String:
		if val.Value == "" {
			return True
		}
	}

	return False
}

func evalMinusPrefixOperator(right object.Object) object.Object {
	if right.Type() != object.IntegerObj {
		return error("unknown operator: -%s", right.Type())
	}

	return &object.Integer{
		Value: -(right.(*object.Integer).Value),
	}
}

func evalInfixExpression(left object.Object, operator string, right object.Object) object.Object {
	switch {
	case left.Type() == object.IntegerObj && right.Type() == object.IntegerObj:
		return evalIntegerInfixExpression(left, operator, right)

	case left.Type() != right.Type():
		return error("type mismatch: %s %s %s", left.Type(), operator, right.Type())

	case operator == "==":
		return nativeBool(left == right)
	case operator == "!=":
		return nativeBool(left != right)
	}

	return error("unknown operator: %s %s %s", left.Type(), operator, right.Type())
}

func evalIntegerInfixExpression(left object.Object, operator string, right object.Object) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value
	switch operator {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "-":
		return &object.Integer{Value: leftVal - rightVal}
	case "*":
		return &object.Integer{Value: leftVal * rightVal}
	case "/":
		return &object.Integer{Value: leftVal / rightVal}

	case ">":
		return nativeBool(leftVal > rightVal)
	case "<":
		return nativeBool(leftVal < rightVal)
	case "==":
		return nativeBool(leftVal == rightVal)
	case "!=":
		return nativeBool(leftVal != rightVal)
	case ">=":
		return nativeBool(leftVal >= rightVal)
	case "<=":
		return nativeBool(leftVal <= rightVal)
	}

	return error("unknown operator %s %s %s", left.Type(), operator, right.Type())
}
