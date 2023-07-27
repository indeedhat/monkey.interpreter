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
	require.Len(t, expr.IfBlock.Statements, 1, "expr.IfBlock")
	require.Nil(t, expr.ElseBlock, "expr.ElseBlock")

	consiquence, ok := expr.IfBlock.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("expr.IfBlock.Statements[0] bad type: expect(*ast.ExpressionStatement) found(%T)",
			expr.IfBlock.Statements[0],
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

	require.NotNil(t, expr.IfBlock, "expr.IfBlock")
	require.Len(t, expr.IfBlock.Statements, 1, "expr.IfBlock")

	require.NotNil(t, expr.ElseBlock, "expr.ElseBlock")

	elseBlock, ok := expr.ElseBlock.(*ast.BlockStatement)
	if !ok {
		t.Fatalf("bad ElseBlock type: expect(*ast.BlockStatement) found(%T)", expr.ElseBlock)
	}

	require.Len(t, elseBlock.Statements, 1, "expr.ElseBlock")

	consiquence, ok := expr.IfBlock.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("expr.IfBlock.Statements[0] bad type: expect(*ast.ExpressionStatement) found(%T)",
			expr.IfBlock.Statements[0],
		)
	}

	testIdentifier(t, consiquence.Expression, "x")

	alternative, ok := elseBlock.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("expr.ElseBlock.Statements[0] bad type: expect(*ast.ExpressionStatement) found(%T)",
			elseBlock.Statements[0],
		)
	}
	testIdentifier(t, alternative.Expression, "y")
}

func TestIfElseChainExpression(t *testing.T) {
	program := parseProgram(t, `if (x < y) { x } else if (false) { y } else { z }`)

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

	require.NotNil(t, expr.IfBlock, "expr.IfBlock")
	require.Len(t, expr.IfBlock.Statements, 1, "expr.IfBlock")

	require.NotNil(t, expr.ElseBlock, "expr.ElseBlock")

	elseifBlock, ok := expr.ElseBlock.(*ast.IfExpression)
	if !ok {
		t.Fatalf("bad ElseBlock type: expect(*ast.IfExpression) found(%T)", expr.ElseBlock)
	}

	testBooleanLiteral(t, elseifBlock.Condition, false)

	require.NotNil(t, expr.IfBlock, "expr.IfBlock")
	require.Len(t, expr.IfBlock.Statements, 1, "expr.IfBlock")

	require.NotNil(t, expr.ElseBlock, "expr.ElseBlock")

	elseBlock, ok := elseifBlock.ElseBlock.(*ast.BlockStatement)
	require.Len(t, elseBlock.Statements, 1, "expr.ElseBlock")

	ifStmt, ok := expr.IfBlock.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("expr.IfBlock.Statements[0] bad type: expect(*ast.ExpressionStatement) found(%T)",
			expr.IfBlock.Statements[0],
		)
	}

	testIdentifier(t, ifStmt.Expression, "x")

	elseifStmt, ok := elseifBlock.IfBlock.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("expr.ElseBlock.IfBlock.Statements[0] bad type: expect(*ast.ExpressionStatement) found(%T)",
			elseBlock.Statements[0],
		)
	}
	testIdentifier(t, elseifStmt.Expression, "y")

	elseIfElseBlock, ok := elseifBlock.ElseBlock.(*ast.BlockStatement)
	if !ok {
		t.Fatalf("expr.ElseBlock.ElseBlock bad type: expect(*ast.BlockStatement) found(%T)",
			elseBlock.Statements[0],
		)
	}

	elseStmt, ok := elseIfElseBlock.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("expr.ElseBlock.ElseBlock.Statements[0] bad type: expect(*ast.ExpressionStatement) found(%T)",
			elseBlock.Statements[0],
		)
	}
	testIdentifier(t, elseStmt.Expression, "z")
}
