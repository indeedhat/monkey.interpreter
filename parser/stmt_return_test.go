package parser

import (
	"testing"

	"github.com/indeedhat/monkey-lang/ast"
	"github.com/stretchr/testify/require"
)

var returnStatemnetTests = []struct {
	input string
	value any
}{
	{"return 1", 1},
	{"return 100;", 100},
	{"return true;", true},
	{"return datas;", "datas"},
}

func TestReturnStatements(t *testing.T) {
	for _, tCase := range returnStatemnetTests {
		program := parseProgram(t, tCase.input)

		require.Len(t, program.Statements, 1, "prorgam.Statements")

		require.Equal(t, "return", program.Statements[0].TokenLiteral(), "stmt")
		require.IsType(t, program.Statements[0], &ast.ReturnStatement{}, "stmt")
		testLiteralExpression(t, program.Statements[0].(*ast.ReturnStatement).Vaule, tCase.value)
	}
}
