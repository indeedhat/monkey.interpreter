package parser

import (
	"testing"

	"github.com/indeedhat/monkey-lang/ast"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var infixTests = []struct {
	inpuut   string
	left     any
	operator string
	right    any
}{
	{"5 + 5;", 5, "+", 5},
	{"5 - 5;", 5, "-", 5},
	{"5 * 5;", 5, "*", 5},
	{"5 / 5;", 5, "/", 5},
	{"5 < 5;", 5, "<", 5},
	{"5 > 5;", 5, ">", 5},
	{"5 == 5;", 5, "==", 5},
	{"5 != 5;", 5, "!=", 5},
	{"5 >= 5;", 5, ">=", 5},
	{"5 <= 5;", 5, "<=", 5},
	{"true == true", true, "==", true},
	{"true != false", true, "!=", false},
	{"false == false", false, "==", false},
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

			testInfixExpression(t, stmt.Expression, tCase.left, tCase.operator, tCase.right)
		})
	}
}

var operatorPresedenceTests = []struct {
	input    string
	expected string
}{
	// natural presedence ordering
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
	{"true", "true"},
	{"false", "false"},
	{"3 > 5 == false", "((3 > 5) == false)"},
	{"3 < 5 == true", "((3 < 5) == true)"},

	// grouped expression ordering
	{"1 + (2 + 3) + 4", "((1 + (2 + 3)) + 4)"},
	{"(5 + 5) * 2", "((5 + 5) * 2)"},
	{"2 / (5 + 5)", "(2 / (5 + 5))"},
	{"-(5 + 5)", "(-(5 + 5))"},
	{"!(true == true)", "(!(true == true))"},

	// function calls
	{"a + add(b * c) + d", "((a + add((b * c))) + d)"},
	{"add(a, b, 1, 2 * 3, 4 + 5, add(6, 7 * 8))", "add(a, b, 1, (2 * 3), (4 + 5), add(6, (7 * 8)))"},
	{"add(a + b + c * d / f + g)", "add((((a + b) + ((c * d) / f)) + g))"},
}

func TestOperatorPresedenceParsing(t *testing.T) {
	for _, tCase := range operatorPresedenceTests {
		t.Run(tCase.input, func(t *testing.T) {
			program := parseProgram(t, tCase.input)

			assert.Equal(t, tCase.expected, program.String())
		})
	}
}
