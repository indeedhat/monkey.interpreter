package parser

import (
	"testing"

	"github.com/indeedhat/monkey-lang/ast"
	"github.com/stretchr/testify/require"
)

func TestIfExpression(t *testing.T) {
	program := parseProgram(t, `if (x < y) { x }`)

	require.Len(t, program.Statements, 1, "program.Statements")

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] bad type: expect(*ast.ExpressionStatement) found(%T)",
			program.Statements[0],
		)
	}

	expr, ok := stmt.Expression.(*ast.IfExpression)
	if !ok {
		t.Fatalf("stmt.Expression bad type: expect(*ast.IfExpression) found(%T)",
			program.Statements[0],
		)
	}

	testInfixExpression(t, expr.Condition, "x", "<", "y")
	require.Len(t, expr.Consiquence.Statements, 1, "expr.Consiquence")
	require.Nil(t, expr.Alternative, "expr.Alternative")

	consiquence, ok := expr.Consiquence.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("expr.Consiquence.Statements[0] bad type: expect(*ast.ExpressionStatement) found(%T)",
			expr.Consiquence.Statements[0],
		)
	}

	testIdentifier(t, consiquence.Expression, "x")
}

func TestIfElseExpression(t *testing.T) {
	program := parseProgram(t, `if (x < y) { x } else { y }`)

	require.Len(t, program.Statements, 1, "program.Statements")

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] bad type: expect(*ast.ExpressionStatement) found(%T)",
			program.Statements[0],
		)
	}

	expr, ok := stmt.Expression.(*ast.IfExpression)
	if !ok {
		t.Fatalf("stmt.Expression bad type: expect(*ast.IfExpression) found(%T)",
			program.Statements[0],
		)
	}

	testInfixExpression(t, expr.Condition, "x", "<", "y")

	require.NotNil(t, expr.Consiquence, "expr.Consiquence")
	require.Len(t, expr.Consiquence.Statements, 1, "expr.Consiquence")

	require.NotNil(t, expr.Alternative, "expr.Alternative")
	require.Len(t, expr.Alternative.Statements, 1, "expr.Alternative")

	consiquence, ok := expr.Consiquence.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("expr.Consiquence.Statements[0] bad type: expect(*ast.ExpressionStatement) found(%T)",
			expr.Consiquence.Statements[0],
		)
	}

	testIdentifier(t, consiquence.Expression, "x")

	alternative, ok := expr.Alternative.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("expr.Alternative.Statements[0] bad type: expect(*ast.ExpressionStatement) found(%T)",
			expr.Alternative.Statements[0],
		)
	}
	testIdentifier(t, alternative.Expression, "y")
}