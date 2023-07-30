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
}

func TestBuiltinFunctions(t *testing.T) {
	for _, tCase := range builtinTests {
		t.Run(tCase.input, func(t *testing.T) {
			evald := testLoggedEval(t, tCase.input)

			switch expected := tCase.expected.(type) {
			case int:
				testIntegerObject(t, evald, int64(expected))
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
