package evaluator

import (
	"testing"

	"github.com/indeedhat/monkey-lang/evaluator/object"
	"github.com/stretchr/testify/require"
)

func TestArrayLiterals(t *testing.T) {
	evaluated := testEval(t, `[1, 2 * 2, 3 + 3]`)

	result, ok := evaluated.(*object.Array)
	if !ok {
		t.Fatalf("bad type: expect(Array) found(%T)", evaluated)
	}

	require.Len(t, result.Elements, 3)

	testIntegerObject(t, result.Elements[0], 1)
	testIntegerObject(t, result.Elements[1], 4)
	testIntegerObject(t, result.Elements[2], 6)
}

var arrayIndexTests = []struct {
	input    string
	expected any
}{
	{`[1, 2, 3][0]`, 1},
	{`[1, 2, 3][1]`, 2},
	{`[1, 2, 3][2]`, 3},
	{`let i = 0; [1][i];`, 1},
	{`[1, 2, 3][1 + 1];`, 3},
	{`let myArray = [1, 2, 3]; myArray[2];`, 3},
	{`let myArray = [1, 2, 3]; myArray[0] + myArray[1] + myArray[2];`, 6},
	{`let myArray = [1, 2, 3]; let i = myArray[0]; myArray[i]`, 2},
	{`[1, 2, 3][3]`, nil},
	{`[1, 2, 3][-1]`, nil},
	{`let myString = "my String"; myString[3];`, "S"},
	{`let myString = "my String"; myString[20];`, nil},
	{`let myString = "my String"; myString[-1];`, nil},
	{`"myString"[1]`, "y"},
}

func TestArrayIndexExpressions(t *testing.T) {
	for _, tCase := range arrayIndexTests {
		t.Run(tCase.input, func(t *testing.T) {
			evaluated := testLoggedEval(t, tCase.input)

			switch val := tCase.expected.(type) {
			case int:
				testIntegerObject(t, evaluated, int64(val))
			case string:
				t.Log(evaluated)
				t.Log(val)
				testStringObject(t, evaluated, val)
			default:
				t.Log(evaluated)
				t.Log(val)
				testNullObject(t, evaluated)
			}
		})
	}
}
