package evaluator

import (
	"testing"

	"github.com/indeedhat/monkey-lang/evaluator/object"
	"github.com/stretchr/testify/require"
)

func TestFunctionObject(t *testing.T) {
	input := "fn(x) { x + 2; };"
	evald := testEval(input)

	fn, ok := evald.(*object.Function)
	if !ok {
		t.Fatalf("bad object type. expect(*object.Function) found(%T)", evald)
	}

	require.Len(t, fn.Parameters, 1)
	require.Equal(t, "x", fn.Parameters[0].String(), "fn.Parameters[0]")

	expectedBody := `{
    (x + 2)
}`
	require.Equal(t, expectedBody, fn.Body.String(), "fn.Body.String()")
}

var fnCallTests = []struct {
	input    string
	expected int64
}{
	{"let identity = fn(x) { x; }; identity(5);", 5},
	{"let identity = fn(x) { return x; }; identity(5);", 5},
	{"let double = fn(x) { x * 2; }; double(5);", 10},
	{"let add = fn(x, y) { x + y; }; add(5, 5);", 10},
	{"let add = fn(x, y) { x + y; }; add(5 + 5, add(5, 5));", 20},
	{"fn(x) { x; }(5)", 5},
}

func TestFunctionCalls(t *testing.T) {
	for _, tCase := range fnCallTests {
		t.Run(tCase.input, func(t *testing.T) {
			testIntegerObject(t, testEval(tCase.input), tCase.expected)
		})
	}
}
