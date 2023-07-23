package parser

import (
	"testing"

	"github.com/indeedhat/monkey-lang/ast"
	"github.com/stretchr/testify/require"
)

func TestFunctionExpression(t *testing.T) {
	program := parseProgram(t, `fn(a, b) { a < b; }`)

	require.Len(t, program.Statements, 1, "program.Statements")

	expr, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] bad type: expect(*ast.ExpressionStatement) found(%T)",
			program.Statements[0],
		)
	}

	fn, ok := expr.Expression.(*ast.FunctionLiteral)
	if !ok {
		t.Fatalf("expr.Expression bad type: expect(*ast.FunctionLiteral) found(%T)", expr.Expression)
	}

	require.Len(t, fn.Parameters, 2, "fn.Parameters")
	testLiteralExpression(t, fn.Parameters[0], "a")
	testLiteralExpression(t, fn.Parameters[1], "b")

	require.NotNil(t, fn.Body, "fn.Body")
	require.Len(t, fn.Body.Statements, 1, "fn.Body.Statements")

	body, ok := fn.Body.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("fn.Body.Statements[0] bad type: expect(*ast.ExpressionStatement) found(%T)", expr.Expression)
	}

	testInfixExpression(t, body.Expression, "a", "<", "b")
}

var functionArgTests = []struct {
	input string
	args  []string
}{
	{"fn() {}", []string{}},
	{"fn(a) {}", []string{"a"}},
	{"fn(a, b, c) {}", []string{"a", "b", "c"}},
}

func TestFunctionArgumentParsing(t *testing.T) {
	for _, tCase := range functionArgTests {
		t.Run(tCase.input, func(t *testing.T) {
			program := parseProgram(t, tCase.input)

			stmt := program.Statements[0].(*ast.ExpressionStatement)
			fn := stmt.Expression.(*ast.FunctionLiteral)

			require.Len(t, fn.Parameters, len(tCase.args), "fn.Parameters")
			for i, arg := range tCase.args {
				testIdentifier(t, fn.Parameters[i], arg)
			}
		})
	}
}
