package parser

import (
	"testing"

	"github.com/indeedhat/monkey-lang/ast"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIdentifierExpression(t *testing.T) {
	program := parseProgram(t, `ook;`)

	require.Len(t, program.Statements, 1, "program.Statement")

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] bad type: expect(*ast.ExpressionStatement) found(%T)",
			program.Statements[0],
		)
	}

	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("stmt.Expression bad type: expected(*ast.Identifier) found(%T)", stmt.Expression)
	}

	assert.Equal(t, "ook", ident.Value, "ident.Value")
	assert.Equal(t, "ook", ident.TokenLiteral(), "ident.TokenLiteral()")
}
