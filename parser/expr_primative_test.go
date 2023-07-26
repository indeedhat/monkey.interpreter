package parser

import (
	"testing"

	"github.com/indeedhat/monkey-lang/ast"
	"github.com/stretchr/testify/require"
)

var boolTests = []struct {
	name  string
	value bool
}{
	{"true", true},
	{"false", false},
}

func TestParseBool(t *testing.T) {
	program := parseProgram(t, `true;false;`)

	require.Len(t, program.Statements, 2, "program.Statement")

	for i, tCase := range boolTests {
		t.Run(tCase.name, func(t *testing.T) {
			stmt, ok := program.Statements[i].(*ast.ExpressionStatement)
			if !ok {
				t.Fatalf("program.Statements[i] bad type: expect(*ast.ExpressionStatement) found(%T)",
					program.Statements[i],
				)
			}

			testBooleanLiteral(t, stmt.Expression, tCase.value)
		})
	}
}

var stringTests = []struct {
	input string
	value string
}{
	{`"this is a string"`, "this is a string"},
	{`"this is a string with \"quotes\""`, `this is a string with "quotes"`},
	{`"this is a string with a \ (backslash)"`, `this is a string with a \ (backslash)`},
}

func TestParseString(t *testing.T) {
	for _, tCase := range stringTests {
		t.Run(tCase.input, func(t *testing.T) {
			program := parseProgram(t, tCase.input)
			require.Len(t, program.Statements, 1)

			stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
			if !ok {
				t.Fatalf("program.Statements[0] bad type: expect(*ast.ExpressionStatement) found(%T)",
					program.Statements[0],
				)
			}

			testStringLiteral(t, stmt.Expression, tCase.value)
		})
	}
}

func TestIntLiteralExpression(t *testing.T) {
	program := parseProgram(t, `5;`)

	require.Len(t, program.Statements, 1, "program.Statements")

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("bat statement type: expect(*ast.ExpressionStatement) found(%T)", program.Statements[0])
	}

	testIntegerLiteral(t, stmt.Expression, 5)
}
