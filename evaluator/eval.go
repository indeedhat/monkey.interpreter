package evaluator

import (
	"log"

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
	case *ast.BooleanLiteral:
		return nativeBool(val.Value)
	case *ast.NullLiteral:
		return Null
	case *ast.PrefixExpression:
		return evalPrefixExpression(val.Operator, Eval(val.Right))
	case *ast.InfixExpression:
		return evalInfixExpression(Eval(val.Left), val.Operator, Eval(val.Right))
	}

	return nil
}

func evalStatements(statements []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range statements {
		result = Eval(statement)
	}

	return result
}

func nativeBool(input bool) *object.Boolean {
	if input {
		return True
	}

	return False
}

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangPrefixOperator(right)
	case "-":
		return evalMinusPrefixOperator(right)
	}
	return nil
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
		return Null
	}

	return &object.Integer{
		Value: -(right.(*object.Integer).Value),
	}
}

func evalInfixExpression(left object.Object, operator string, right object.Object) object.Object {
	switch {
	case left.Type() == object.IntegerObj && right.Type() == object.IntegerObj:
		return evalIntegerInfixExpression(left, operator, right)

	case operator == "==":
		log.Println(left, operator, right)
		return nativeBool(left == right)
	case operator == "!=":
		log.Println(left, operator, right)
		return nativeBool(left != right)
	}
	return Null
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

	return Null
}
