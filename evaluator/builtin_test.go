package evaluator

import (
	"testing"

	"github.com/indeedhat/monkey-lang/evaluator/object"
	"github.com/stretchr/testify/require"
)

var builtinTests = []struct {
	input    string
	expected interface{}
}{
	{`len("")`, 0},
	{`len("four")`, 4},
	{`len("hello world")`, 11},
	{`len(1)`, "invalid arg type `int`"},
	{`len("one", "two"); 1`, "bad arg count: expect(1) found(2)"},

	{`append([1], 2)`, []int64{1, 2}},
	{`append([1], 2, 3)`, []int64{1, 2, 3}},
	{`append([1])`, "append requires at least 2 arguments"},
	{`append("not an array", 2)`, "bad subject type: expected(*object.Array) found(*object.String)"},
}

func TestBuiltinFunctions(t *testing.T) {
	for _, tCase := range builtinTests {
		t.Run(tCase.input, func(t *testing.T) {
			evald := testEval(t, tCase.input)

			switch expected := tCase.expected.(type) {
			case int:
				testIntegerObject(t, evald, int64(expected))
			case []int64:
				testIntegerArrayObject(t, evald, expected)
			case string:
				err, ok := evald.(*object.Error)
				if !ok {
					t.Errorf("unexpected type. expect(*object.Error) found(%T)", evald)
					return
				}

				require.Equal(t, expected, err.Message)
			}
		})
	}
}

func testIntegerArrayObject(t *testing.T, obj object.Object, expected []int64) {
	result, ok := obj.(*object.Array)
	if !ok {
		t.Fatalf("bad object type: expect(object.Array) found(%T) -> %v", obj, obj)
	}

	require.Len(t, result.Elements, len(expected))
	for i, elem := range result.Elements {
		i64, ok := elem.(*object.Integer)
		if !ok {
			t.Fatalf("bad type: expected(*object.Integer) found(%T)", elem)
		}

		require.Equal(t, expected[i], i64.Value)
	}
}
