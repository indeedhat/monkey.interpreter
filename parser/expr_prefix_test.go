package parser

import (
	"testing"

	"github.com/indeedhat/monkey-lang/ast"
	"github.com/stretchr/testify/require"
)

var prefixTests = []struct {
	input    string
	operator string
	intValue int64
}{
	{"!5", "!", 5},
	{"-15", "-", 15},
}

func TestParsingPrefixExpressions(t *testing.T) {
	for _, tCase := range prefixTests {
		t.Run(tCase.input, func(t *testing.T) {
			program := parseProgram(t, tCase.input)

			require.Len(t, program.Statements, 1, "program")

			stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
			if !ok {
				t.Fatalf("bad stmt type: expect(*ast.ExpressionStatement) found(%T)", program.Statements[0])
			}

			prefix, ok := stmt.Expression.(*ast.PrefixExpression)
			if !ok {
				t.Fatalf("bad expression type: expect(*ast.PrefixExpression) found(%T)", stmt.Expression)
			}

			require.Equal(t, tCase.operator, prefix.Operator)
			testIntegerLiteral(t, prefix.Right, tCase.intValue)
		})
	}
}
