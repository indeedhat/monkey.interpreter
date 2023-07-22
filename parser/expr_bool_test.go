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
