package evaluator

import (
	"testing"

	"github.com/indeedhat/monkey-lang/evaluator/object"
	"github.com/indeedhat/monkey-lang/lexer"
	"github.com/indeedhat/monkey-lang/parser"
	"github.com/stretchr/testify/require"
)

var integerTests = []struct {
	input string
	value int64
}{
	{"5", 5},
	{"10", 10},

	// negatives
	{"-5", -5},
	{"-10", -10},

	// expressions
	{"5 + 5 + 5 + 5 - 10", 10},
	{"2 * 2 * 2 * 2 * 2", 32},
	{"-50 + 100 + -50", 0},
	{"5 * 2 + 10", 20},
	{"5 + 2 * 10", 25},
	{"20 + 2 * -10", 0},
	{"50 / 2 * 2 + 10", 60},
	{"2 * (5 + 10)", 30},
	{"3 * 3 * 3 + 10", 37},
	{"3 * (3 * 3) + 10", 37},
	{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
}

func TestEvalInteger(t *testing.T) {
	for _, tCase := range integerTests {
		t.Run(tCase.input, func(t *testing.T) {
			evald := testEval(tCase.input)
			testIntegerObject(t, evald, tCase.value)
		})
	}
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Fatalf("bad object type: expect(Integer) found(%T)", obj)
	}

	require.Equal(t, expected, result.Value)
}

var boolTests = []struct {
	input string
	value bool
}{
	{"true", true},
	{"false", false},
}

func TestEvalBool(t *testing.T) {
	for _, tCase := range boolTests {
		t.Run(tCase.input, func(t *testing.T) {
			evald := testEval(tCase.input)
			testBoolObject(t, evald, tCase.value)
		})
	}
}

func testBoolObject(t *testing.T, obj object.Object, expected bool) {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Fatalf("bad object type: expected(Boolean) found(%T)", obj)
	}

	require.Equal(t, expected, result.Value)
}

func TestEvalNull(t *testing.T) {
	evald := testEval("null")
	_, ok := evald.(*object.Null)
	if !ok {
		t.Fatalf("bad object type: expected(Null) found(%T)", evald)
	}
}

var bangTests = []struct {
	input string
	value bool
}{
	{"!true", false},
	{"!false", true},
	{"!!true", true},
	{"!!false", false},
	{"!!!true", false},
	{"!!!false", true},
	{"!0", true},
	{"!!0", false},
	{"!!!0", true},
}

func TestEvalBang(t *testing.T) {
	for _, tCase := range bangTests {
		t.Run(tCase.input, func(t *testing.T) {
			evald := testEval(tCase.input)
			testBoolObject(t, evald, tCase.value)
		})
	}
}

var booleanExpressionTests = []struct {
	input string
	value bool
}{
	{"5 == 5", true},
	{"5 != 5", false},
	{"5 > 5", false},
	{"5 < 5", false},
	{"5 > 10", false},
	{"5 < 10", true},
	{"5 <= 10", true},
	{"5 >= 10", false},
	{"10 <= 5", false},
	{"10 >= 5", true},
	{"5 <= 5", true},
	{"10 >= 10", true},
	{"true == true", true},
	{"true == false", false},
	{"false == false", true},
	{"false == true", false},
	{"(1 < 2) == true", true},
	{"(1 < 2) == false", false},
	{"(1 > 2) == true", false},
	{"(1 > 2) == false", true},
	{"(5 > 5) == true", false},
	{"(10 > 5) == true", true},
}

func TestEvalBoolExpression(t *testing.T) {
	for _, tCase := range booleanExpressionTests {
		t.Run(tCase.input, func(t *testing.T) {
			evald := testEval(tCase.input)
			testBoolObject(t, evald, tCase.value)
		})
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	prog := p.ParseProgram()

	return Eval(prog)
}
