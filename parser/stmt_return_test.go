package parser

import (
	"testing"

	"github.com/indeedhat/monkey-lang/ast"
	"github.com/stretchr/testify/require"
)

func TestReturnStatements(t *testing.T) {
	program := parseProgram(t, `
return 1;
return 11;
return 69420;
`)

	require.Len(t, program.Statements, 3, "program.Statements")

	for _, stmt := range program.Statements {
		require.Equal(t, "return", stmt.TokenLiteral(), "stmt")
		require.IsType(t, stmt, &ast.ReturnStatement{}, "stmt")
	}
}
