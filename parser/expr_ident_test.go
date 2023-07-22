package parser

import (
	"testing"

	"github.com/indeedhat/monkey-lang/ast"
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

	testIdentifier(t, stmt.Expression, "ook")
}
