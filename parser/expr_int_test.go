package parser

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/indeedhat/monkey-lang/ast"
	"github.com/stretchr/testify/require"
)

func TestIntLiteralExpression(t *testing.T) {
	program := parseProgram(t, `5;`)

	spew.Dump(program.Statements)
	require.Len(t, program.Statements, 1, "program.Statements")

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("bat statement type: expect(*ast.ExpressionStatement) found(%T)", program.Statements[0])
	}

	testIntegerLiteral(t, stmt.Expression, 5)
}
