package parser

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/indeedhat/monkey-lang/ast"
	"github.com/stretchr/testify/assert"
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

func testIntegerLiteral(t *testing.T, expr ast.Expression, expected int64) {
	lit, ok := expr.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("bad expression type: expect(*ast.IntegerLiteral) found(%T)", expr)
	}

	assert.Equal(t, expected, lit.Value, "lit.Value")
	assert.Equal(t, fmt.Sprint(expected), lit.TokenLiteral(), "lit.TokenLiteral()")
}
