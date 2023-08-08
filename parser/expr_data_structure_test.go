package parser

import (
	"testing"

	"github.com/indeedhat/monkey-lang/ast"
	"github.com/stretchr/testify/require"
)

func TestParsingArrayLiterals(t *testing.T) {
	input := `[1, 2 * 2, 3 + 3]`

	program := parseProgram(t, input)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	array, ok := stmt.Expression.(*ast.ArrayLiteral)
	if !ok {
		t.Fatalf("bad type: expect(*ast.ArrayLiteral) found(%T)", stmt.Expression)
	}

	require.Len(t, array.Elements, 3, "array.Elements")

	testIntegerLiteral(t, array.Elements[0], 1)
	testInfixExpression(t, array.Elements[1], 2, "*", 2)
	testInfixExpression(t, array.Elements[2], 3, "+", 3)
}

func TestParsingIndexExpressions(t *testing.T) {
	input := `myArray[1 + 1]`

	program := parseProgram(t, input)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	indexExp, ok := stmt.Expression.(*ast.IndexExpression)
	if !ok {
		t.Fatalf("bad type: expect(*ast.IndexExpression) found(%T)", stmt.Expression)
	}

	testIdentifier(t, indexExp.Subject, "myArray")
	testInfixExpression(t, indexExp.Index, 1, "+", 1)
}

func TestParsStringLiteralingIndexExpressions(t *testing.T) {
	input := `"myString"[1 + 1]`

	program := parseProgram(t, input)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	indexExp, ok := stmt.Expression.(*ast.IndexExpression)
	if !ok {
		t.Fatalf("bad type: expect(*ast.IndexExpression) found(%T)", stmt.Expression)
	}

	testStringLiteral(t, indexExp.Subject, "myString")
	testInfixExpression(t, indexExp.Index, 1, "+", 1)
}
