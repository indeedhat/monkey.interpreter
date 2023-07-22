package parser

import (
	"testing"

	"github.com/indeedhat/monkey-lang/ast"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var letStatementTests = []struct {
	expectedIdent string
	value         any
}{
	{"x", 5},
	{"y", 10},
	{"zee", 835628},
	{"t", true},
	{"f", false},
}

func TestLetStatements(t *testing.T) {
	program := parseProgram(t, `
let x = 5;
let y = 10;
let zee = 835628;
let t = true;
let f = false;
`)

	require.Len(t, program.Statements, 5, "program.Statements")

	for i, testCase := range letStatementTests {
		s := program.Statements[i]
		require.Equal(t, "let", s.TokenLiteral(), "program.Stamenets[i]")

		stmt, ok := s.(*ast.LetStatement)
		if !ok {
			t.Fatalf("s is not a *ast.LetStatement: got(%T)", s)
		}

		assert.Equal(t, testCase.expectedIdent, stmt.Name.Value, "stmt.Name.Value")
		assert.Equal(t, testCase.expectedIdent, stmt.Name.TokenLiteral(), "stmt.Name.TokenLIteral()")
	}
}
