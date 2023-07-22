package parser

import (
	"testing"

	"github.com/indeedhat/monkey-lang/ast"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var infixTests = []struct {
	inpuut   string
	left     int64
	operator string
	right    int64
}{
	{"5 + 5;", 5, "+", 5},
	{"5 - 5;", 5, "-", 5},
	{"5 * 5;", 5, "*", 5},
	{"5 / 5;", 5, "/", 5},
	{"5 < 5;", 5, "<", 5},
	{"5 > 5;", 5, ">", 5},
	{"5 == 5;", 5, "==", 5},
	{"5 != 5;", 5, "!=", 5},
}

func TestParsingInfixExpressions(t *testing.T) {
	for _, tCase := range infixTests {
		t.Run(tCase.inpuut, func(t *testing.T) {
			program := parseProgram(t, tCase.inpuut)

			require.Len(t, program.Statements, 1, "program.Statement")

			stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
			if !ok {
				t.Fatalf("program.Statements[0] bad type: expect(*ast.ExpressionStatement) found(%T)",
					program.Statements[0],
				)
			}

			infix, ok := stmt.Expression.(*ast.InfixExpression)
			if !ok {
				t.Fatalf("stmt.Expression bad type: expected(*ast.InfixExpression) found(%T)", stmt.Expression)
			}

			testIntegerLiteral(t, infix.Left, tCase.left)
			assert.Equal(t, tCase.operator, infix.Operator, "infix.Operator")
			testIntegerLiteral(t, infix.Right, tCase.right)
		})
	}
}

var operatorPresedenceTests = []struct {
	input    string
	expected string
}{
	{"-a * b", "((-a) * b)"},
	{"!-a", "(!(-a))"},
	{"a + b + c", "((a + b) + c)"},
	{"a + b - c", "((a + b) - c)"},
	{"a * b * c", "((a * b) * c)"},
	{"a * b / c", "((a * b) / c)"},
	{"a + b / c", "(a + (b / c))"},
	{"a + b * c + d / e - f", "(((a + (b * c)) + (d / e)) - f)"},
	{"3 + 4; -5 * 5", "(3 + 4)((-5) * 5)"},
	{"5 > 4 == 3 < 4", "((5 > 4) == (3 < 4))"},
	{"5 < 4 != 3 > 4", "((5 < 4) != (3 > 4))"},
	{"3 + 4 * 5 == 3 * 1 + 4 * 5", "((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))"},
}

func TestOperatorPresedenceParsing(t *testing.T) {
	for _, tCase := range operatorPresedenceTests {
		t.Run(tCase.input, func(t *testing.T) {
			program := parseProgram(t, tCase.input)

			assert.Equal(t, tCase.expected, program.String())
		})
	}
}
