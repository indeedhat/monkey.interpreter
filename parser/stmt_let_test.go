package parser

import (
	"testing"

	"github.com/indeedhat/monkey-lang/ast"
	"github.com/stretchr/testify/require"
)

var letStatementTests = []struct {
	input string
	ident string
	value any
}{
	{"let x = 5;", "x", 5},
	{"let y = 10;", "y", 10},
	{"let zee = 835628;", "zee", 835628},
	{"let t = true;", "t", true},
	{"let f = false;", "f", false},
	{"let x = y;", "x", "y"},
}

func TestLetStatements(t *testing.T) {
	for _, tCase := range letStatementTests {
		t.Run(tCase.input, func(t *testing.T) {
			program := parseProgram(t, tCase.input)

			require.Len(t, program.Statements, 1, "program.Statements")

			testLetStatement(t, program.Statements[0], tCase.ident)
			testLiteralExpression(t, program.Statements[0].(*ast.LetStatement).Value, tCase.value)
		})
	}
}

func testLetStatement(t *testing.T, stmt ast.Statement, name string) {
	require.Equal(t, "let", stmt.TokenLiteral(), "stmt.TokenLiteral()")

	letStmt, ok := stmt.(*ast.LetStatement)
	if !ok {
		t.Fatalf("stmt not *ast.LetStatement. got=%T", stmt)
	}

	require.Equal(t, name, letStmt.Name.Value, "letStmt.Name.Value")
	require.Equal(t, name, letStmt.Name.TokenLiteral(), "letStmt.Name.TokenLiteral()")
}
